package kms

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/qredo/fusionchain/keyring/pkg/api"
	"github.com/qredo/fusionchain/keyring/pkg/common"
	"github.com/qredo/fusionchain/keyring/pkg/database"
	"github.com/qredo/fusionchain/keyring/pkg/services/mpcrelayer"
	"github.com/qredo/fusionchain/x/treasury/types"
)

type keyController struct {
	KeyringAddr        string
	queue              chan *keyRequestQueueItem
	keyRequestsHandler KeyRequestsHandler
	log                *logrus.Entry
	threads            chan struct{}
	stop               chan struct{}
	wait               chan struct{}

	retrySleep time.Duration
}

func newFusionKeyController(logger *logrus.Entry, prefixDB database.Database, q chan *keyRequestQueueItem, keyringClient Keyring, txc mpcrelayer.TxClient) *keyController {
	k := &FusionKeyRequestHandler{
		KeyDB:         prefixDB,
		keyringClient: keyringClient,
		TxClient:      txc,
		Logger:        logger,
	}

	return &keyController{
		queue:              q,
		keyRequestsHandler: k,
		log:                logger,
		threads:            makeThreads(defaultThreads),
		stop:               make(chan struct{}, 1),
		wait:               make(chan struct{}, 1),
		retrySleep:         defaultRetryTimeout,
	}
}

// Start implements Module.Start()
func (k *keyController) Start() error {
	if k.queue == nil || k.stop == nil {
		return fmt.Errorf("empty work channels")
	}
	k.log.WithField("threads", len(k.threads)).Info("starting keyRequestHandler")
	go k.startExecutor()
	return nil
}

func (k *keyController) startExecutor() {
	for {
		select {
		case <-k.stop:
			k.log.Debug("keyController received shutdown signal")
			for i := 0; i < defaultThreads; i++ {
				<-k.threads // empty thread chan
			}
			k.log.Debug("terminated keyController")
			k.wait <- struct{}{}
			return
		case item := <-k.queue:
			go func() {
				i := item
				<-k.threads
				defer func() { k.threads <- struct{}{} }()
				if err := k.executeRequest(i); err != nil {
					k.log.WithFields(logrus.Fields{
						"retries": i.retries,
						"error":   err.Error(),
					}).Error("keyRequestErr")
				}
			}()
		}
	}
}

// Stop implements Module.Stop()
func (k *keyController) Stop() error {
	k.stop <- struct{}{}
	<-k.wait
	return nil
}

func (k *keyController) executeRequest(item *keyRequestQueueItem) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), defaultHandlerTimeout)
	defer cancelFunc()
	if err := k.keyRequestsHandler.HandleKeyRequests(ctx, item); err != nil {
		if item.retries <= item.maxTries {
			requeueKeyItemWithTimeout(k.queue, item, k.retrySleep) // Requeue items until maxTries limit has been reached
		}
		return err
	}
	return nil
}

func (k *keyController) healthcheck() *api.HealthResponse {
	return k.keyRequestsHandler.healthcheck()
}

type keyRequestQueueItem struct {
	retries  int
	maxTries int
	request  *types.KeyRequest
}

type KeyRequestsHandler interface {
	HandleKeyRequests(ctx context.Context, item *keyRequestQueueItem) error
	healthcheck() *api.HealthResponse
}

// FusionKeyRequestHandler implements KeyRequestsHandler.
type FusionKeyRequestHandler struct {
	KeyDB         database.Database
	keyringClient Keyring
	TxClient      mpcrelayer.TxClient
	Logger        *logrus.Entry
}

var _ KeyRequestsHandler = &FusionKeyRequestHandler{}

// HandleKeyRequests processes the pending key request supplied by fusiond, requesting a public key
// via the MPC client and fulfilling the request via the TxClient.
func (h *FusionKeyRequestHandler) HandleKeyRequests(ctx context.Context, item *keyRequestQueueItem) error {
	if item == nil || item.request == nil {
		return fmt.Errorf("malformed keyRequest item")
	}
	start := time.Now()

	// make 64 character keyID from the ID supplied for the keys request
	keyIDStr := fmt.Sprintf("%0*x", keyIDLength, item.request.Id)

	keyID, err := hex.DecodeString(keyIDStr)
	if err != nil {
		return err
	}

	// Request an ECDSA/EdDSA public key from the MPC service
	pk, err := h.keyringClient.PublicKey(keyID)
	if err != nil {
		return err
	}
	h.Logger.WithFields(logrus.Fields{
		"keyID":     keyIDStr,
		"publicKey": fmt.Sprintf("%x", pk),
	}).Debug("pubKeyReturned")

	// Verify that a signature can be generated for the supplied public key.
	// The response is validated inside the keyringClient.
	if _, err = h.keyringClient.PubkeySignature(keyID); err != nil {
		return err
	}

	// Approve the user item.request, write the generated public key to fusiond.
	if err = h.TxClient.FulfilKeyRequest(ctx, item.request.Id, pk); err != nil {
		return err
	}

	// Store the generated secret key in our database, will be used when user requests signatures.
	if err = makePkEntry(h.KeyDB, keyIDStr, fmt.Sprintf("%x", pk)); err != nil {
		return err
	}
	h.Logger.WithFields(logrus.Fields{
		"timeTaken": common.RoundFloat(time.Since(start).Seconds(), 2),
	}).Info("keyRequestFulfilled")
	return nil
}

func (*FusionKeyRequestHandler) healthcheck() *api.HealthResponse {
	return &api.HealthResponse{}
}

func makePkEntry(db database.Database, keyIDStr, pkStr string) error {
	k := makeDBKey(keyIDStr)
	v, err := json.Marshal(&api.PkData{
		PublicKey: pkStr,
		Created:   time.Now().Format(time.RFC3339),
	})
	if err != nil {
		return err
	}
	if err := db.Persist(k, v); err != nil {
		return err
	}
	return nil
}

func makeDBKey(keyID string) string {
	return fmt.Sprintf("%s%s", pkPrefix, keyID)
}
