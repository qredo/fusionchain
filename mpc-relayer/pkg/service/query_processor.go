package service

import (
	"context"
	"time"

	"github.com/qredo/fusionchain/go-client"
	"github.com/sirupsen/logrus"
)

type keyQueryProcessor struct {
	keyRingID      uint64
	queryClient    QueryClient
	keyRequestChan chan *keyRequestQueueItem
	stop           chan struct{}
	wait           chan struct{}
	tickDuration   time.Duration

	log      *logrus.Entry
	maxTries int
}

func newKeyQueryProcessor(keyringId uint64, q QueryClient, k chan *keyRequestQueueItem, log *logrus.Entry, t time.Duration, maxTries int) *keyQueryProcessor {
	return &keyQueryProcessor{
		keyRingID:      keyringId,
		queryClient:    q,
		keyRequestChan: k,
		stop:           make(chan struct{}, 1),
		wait:           make(chan struct{}, 1),
		tickDuration:   t,
		log:            log,
		maxTries:       maxTries,
	}
}

func (q *keyQueryProcessor) Start() error {
	go q.startTicker()
	return nil
}

func (q *keyQueryProcessor) startTicker() {
	ticker := time.NewTicker(q.tickDuration)
	var processing bool
	defer ticker.Stop()
	for {
		select {
		case <-q.stop:
			q.log.Info("keyQueryProcessor received shutdown signal")
			for {
				if !processing {
					break
				}
			}
			q.log.Info("terminated keyQueryProcessor")
			q.wait <- struct{}{}
			return
		case <-ticker.C:
			// Process Key request queries
			go func() {
				processing = true
				defer func() { processing = false }()
				if err := q.executeKeyQuery(); err != nil {
					q.log.WithField("error", err.Error()).Error("pendingKeyQueryErr")
				}
			}()
		}
	}
}

func (q *keyQueryProcessor) executeKeyQuery() error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), defaultQueryTimeout)
	defer cancelFunc()
	pendingKeyRequests, err := q.queryClient.PendingKeyRequests(ctx, &client.PageRequest{}, q.keyRingID)
	if err != nil {
		return err
	}
	for _, r := range pendingKeyRequests {
		newItem := &keyRequestQueueItem{
			request:  r,
			maxTries: q.maxTries,
		}
		q.keyRequestChan <- newItem
	}
	return nil
}

func (q *keyQueryProcessor) Stop() error {
	q.stop <- struct{}{}
	<-q.wait
	return nil
}

func (q *keyQueryProcessor) healthcheck() *Response {
	return &Response{}
}

type sigQueryProcessor struct {
	keyRingID      uint64
	queryClient    QueryClient
	sigRequestChan chan *signatureRequestQueueItem
	stop           chan struct{}
	wait           chan struct{}
	tickDuration   time.Duration

	log      *logrus.Entry
	maxTries int
}

func newSigQueryProcessor(keyringId uint64, q QueryClient, s chan *signatureRequestQueueItem, log *logrus.Entry, t time.Duration, maxTries int) *sigQueryProcessor {
	return &sigQueryProcessor{
		keyRingID:      keyringId,
		queryClient:    q,
		sigRequestChan: s,
		stop:           make(chan struct{}, 1),
		wait:           make(chan struct{}, 1),
		tickDuration:   t,
		log:            log,
		maxTries:       maxTries,
	}
}

func (q sigQueryProcessor) Start() error {
	go q.startTicker()
	return nil
}

func (q sigQueryProcessor) startTicker() {
	ticker := time.NewTicker(q.tickDuration)
	var processing bool
	defer ticker.Stop()
	for {
		select {
		case <-q.stop:
			q.log.Info("sigQueryProcessor received shutdown signal")
			for {
				if !processing {
					break
				}
			}
			q.log.Info("terminated sigQueryProcessor")
			q.wait <- struct{}{}
			return
		case <-ticker.C:
			// Process Signature request queries
			go func() {
				processing = true
				defer func() { processing = false }()
				if err := q.executeSignatureQuery(); err != nil {
					q.log.WithField("error", err.Error()).Error("pendingSigQueryErr")
				}
			}()

		}
	}
}

func (q *sigQueryProcessor) executeSignatureQuery() error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), defaultQueryTimeout)
	defer cancelFunc()
	pendingSigRequests, err := q.queryClient.PendingSignatureRequests(ctx, &client.PageRequest{}, q.keyRingID)
	if err != nil {
		return err
	}
	for _, r := range pendingSigRequests {
		newItem := &signatureRequestQueueItem{
			request:  r,
			maxTries: q.maxTries,
		}
		q.sigRequestChan <- newItem
	}
	return nil
}

func (q *sigQueryProcessor) Stop() error {
	q.stop <- struct{}{}
	<-q.wait
	return nil
}

func (q *sigQueryProcessor) healthcheck() *Response {
	return &Response{}
}
