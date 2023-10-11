package service

import (
	"context"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/qredo/fusionchain/mpc-relayer/pkg/database"
	"github.com/qredo/fusionchain/mpc-relayer/pkg/mpc"
	"github.com/qredo/fusionchain/x/treasury/types"
	"github.com/sirupsen/logrus"
)

type signatureController struct {
	KeyringID                uint64
	queue                    chan *signatureRequestQueueItem
	signatureRequestsHandler SignatureRequestsHandler
	log                      *logrus.Entry

	stop       chan struct{}
	retrySleep time.Duration
}

type signatureRequestQueueItem struct {
	retries  int
	maxTries int
	request  *types.SignRequest
}

func newFusionSignatureController(logger *logrus.Entry, db database.Database, q chan *signatureRequestQueueItem, keyringClient mpc.Client, txc TxClient) *signatureController {
	s := &FusionSignatureRequestHandler{
		KeyDB:         db,
		keyringClient: keyringClient,
		TxClient:      txc,
		Logger:        logger,
	}
	return &signatureController{
		queue:                    q,
		signatureRequestsHandler: s,
		stop:                     make(chan struct{}),
		retrySleep:               defaultRetryTimeout,
	}
}

func (s *signatureController) Start() error {
	if s.queue == nil || s.stop == nil {
		return fmt.Errorf("empty work channels")
	}
	go s.startExecutor()
	return nil

}

func (s *signatureController) startExecutor() {
	for {
		select {
		case <-s.stop:
			return
		case item := <-s.queue:
			go func() {
				if err := s.executeRequest(item); err != nil {
					s.log.WithField("error", err.Error()).Error("signRequestErr")
				}
			}()
		}
	}
}

func (s signatureController) Stop() error {
	s.stop <- struct{}{}
	return nil
}
func (s signatureController) executeRequest(item *signatureRequestQueueItem) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), defaultHandlerTimeout)
	defer cancelFunc()
	if err := s.signatureRequestsHandler.HandleSignatureRequest(ctx, item); err != nil {
		if item.retries <= item.maxTries {
			requeueSigItemWithTimeout(s.queue, item, s.retrySleep)
		}
		return err
	}
	return nil
}

type SignatureRequestsHandler interface {
	HandleSignatureRequest(ctx context.Context, item *signatureRequestQueueItem) error
}

// FusionSignatureRequestHandler implements SignatureRequestsHandler.
type FusionSignatureRequestHandler struct {
	KeyDB         database.Database
	keyringClient mpc.Client
	TxClient      TxClient
	Logger        *logrus.Entry
}

var _ SignatureRequestsHandler = &FusionSignatureRequestHandler{}

func (h *FusionSignatureRequestHandler) HandleSignatureRequest(ctx context.Context, item *signatureRequestQueueItem) error {
	l := h.Logger.WithField("request_id", item.request.Id)

	keyID, err := hex.DecodeString(fmt.Sprintf("%0*x", mpcRequestKeyLength, item.request.KeyId))
	if err != nil {
		return err
	}
	requestID, err := hex.DecodeString(fmt.Sprintf("%0*x", mpcRequestKeyLength, item.request.Id))
	if err != nil {
		return err
	}

	sigResponse, traceID, err := h.keyringClient.Signature(&mpc.SigRequestData{
		KeyID:   keyID,
		ID:      requestID,
		SigHash: item.request.DataForSigning,
	}, mpc.EcDSA)
	l = l.WithField("traceID", traceID)
	if err != nil {
		if item.retries >= item.maxTries {
			if rejectErr := h.TxClient.RejectSignatureRequest(ctx, item.request.Id, err.Error()); rejectErr != nil {
				return rejectErr
			}
		}
		return nil
	}

	signature, err := mpc.ExtractSerializedSigECDSA(sigResponse)
	if err != nil {
		return err
	}

	if err = h.TxClient.FulfilSignatureRequest(ctx, item.request.Id, signature); err != nil {
		return err
	}

	l.Info("fulfilled")
	return nil
}
