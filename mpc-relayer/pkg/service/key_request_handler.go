package service

import (
	"context"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/qredo/fusionchain/mpc-relayer/pkg/database"
	"github.com/qredo/fusionchain/mpc-relayer/pkg/mpc"
	"github.com/qredo/fusionchain/x/treasury/types"
)

type keyController struct {
	KeyringID          uint64
	queue              chan *keyRequestQueueItem
	keyRequestsHandler KeyRequestsHandler
	log                *logrus.Entry

	stop chan struct{}

	retrySleep time.Duration
}

func newFusionKeyController(logger *logrus.Entry, db database.Database, q chan *keyRequestQueueItem, keyringClient mpc.Client, txc TxClient) *keyController {
	k := &FusionKeyRequestHandler{
		KeyDB:         db,
		keyringClient: keyringClient,
		TxClient:      txc,
		Logger:        logger,
	}
	return &keyController{
		queue:              q,
		keyRequestsHandler: k,
		stop:               make(chan struct{}),
		retrySleep:         defaultRetryTimeout,
	}
}

func (k *keyController) Start() error {
	if k.queue == nil || k.stop == nil {
		return fmt.Errorf("empty work channels")
	}
	go k.startExecutor()
	return nil
}

func (k *keyController) startExecutor() {
	for {
		select {
		case <-k.stop:
			return
		case item := <-k.queue:
			// TODO
			go func() {
				if err := k.executeRequest(item); err != nil {
					k.log.WithField("error", err.Error()).Error("keyRequestErr")
				}
			}()
		}
	}
}

// TODO
func (k keyController) Stop() error {
	k.stop <- struct{}{}
	return nil
}

// TODO
func (k keyController) executeRequest(item *keyRequestQueueItem) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), defaultHandlerTimeout)
	defer cancelFunc()
	if err := k.keyRequestsHandler.HandleKeyRequests(ctx, item); err != nil {
		if item.retries <= item.maxTries {
			requeueKeyItemWithTimeout(k.queue, item, k.retrySleep)
		}
		return err
	}
	return nil
}

type keyRequestQueueItem struct {
	retries  int
	maxTries int
	request  *types.KeyRequest
}

type KeyRequestsHandler interface {
	HandleKeyRequests(ctx context.Context, item *keyRequestQueueItem) error
}

// FusionKeyRequestHandler implements KeyRequestsHandler.
type FusionKeyRequestHandler struct {
	KeyDB         database.Database
	keyringClient mpc.Client
	TxClient      TxClient
	Logger        *logrus.Entry
}

// HandleKeyRequests - TODO
func (h *FusionKeyRequestHandler) HandleKeyRequests(ctx context.Context, item *keyRequestQueueItem) error {

	//
	//
	// TODO
	//
	//
	l := h.Logger.WithField("request_id", item.request.Id)

	// generate new key
	keyIDStr := fmt.Sprintf("%0*x", mpcRequestKeyLength, item.request.Id)

	keyID, err := hex.DecodeString(keyIDStr)
	if err != nil {
		return err
	}
	pk, traceID, err := h.keyringClient.PublicKey(keyID, mpc.EcDSA)
	l = l.WithField("trace_id", traceID)
	if err != nil {
		return err
	}

	// approve the user item.request, provide the generated public key
	if err = h.TxClient.FulfilKeyRequest(ctx, item.request.Id, pk); err != nil {
		return err
	}

	// store the generated secret key in our database, will be used when user requests signatures
	err = h.KeyDB.Persist(keyIDStr, pk)
	if err != nil {
		return err
	}

	l.Info("fulfilled")
	return nil
}
