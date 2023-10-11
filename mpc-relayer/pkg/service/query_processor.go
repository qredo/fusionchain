package service

import (
	"context"
	"sync"
	"time"

	"github.com/qredo/fusionchain/go-client"
	"github.com/sirupsen/logrus"
)

type queryProcessor struct {
	keyRingID      uint64
	queryClient    QueryClient
	keyRequestChan chan *keyRequestQueueItem
	sigRequestChan chan *signatureRequestQueueItem
	stop           chan struct{}
	tickDuration   time.Duration

	log      *logrus.Entry
	maxTries int
}

func newQueryProcessor(keyringId uint64, q QueryClient, k chan *keyRequestQueueItem, s chan *signatureRequestQueueItem, log *logrus.Entry, t time.Duration, maxTries int) *queryProcessor {
	return &queryProcessor{
		keyRingID:      keyringId,
		queryClient:    q,
		keyRequestChan: k,
		sigRequestChan: s,
		stop:           make(chan struct{}),
		tickDuration:   t,
		log:            log,
		maxTries:       maxTries,
	}
}

func (q queryProcessor) Start() error {
	go q.startTicker()
	return nil
}

func (q queryProcessor) startTicker() {
	mu := sync.Mutex{}
	ticker := time.NewTicker(q.tickDuration)
	defer ticker.Stop()
	for {
		select {
		case <-q.stop:
			return
		case <-ticker.C:

			// Process Key request queries
			go func() {
				ctx, cancelFunc := context.WithTimeout(context.Background(), defaultQueryTimeout)
				defer cancelFunc()
				mu.Lock()
				pendingKeyRequests, err := q.queryClient.PendingKeyRequests(ctx, &client.PageRequest{
					Limit: 10,
				}, q.keyRingID)
				mu.Unlock()
				if err != nil {
					q.log.WithField("error", err.Error()).Error("keyQueryErr")
				}
				for _, r := range pendingKeyRequests {
					newItem := &keyRequestQueueItem{
						request:  r,
						maxTries: q.maxTries,
					}
					q.keyRequestChan <- newItem
				}
			}()
			// Process Signature request queries
			go func() {
				ctx, cancelFunc := context.WithTimeout(context.Background(), defaultQueryTimeout)
				defer cancelFunc()
				mu.Lock()
				pendingSigRequests, err := q.queryClient.PendingSignatureRequests(ctx, &client.PageRequest{
					Limit: 10,
				}, q.keyRingID)
				mu.Unlock()
				if err != nil {
					q.log.WithField("error", err.Error()).Error("keyQueryErr")
				}
				for _, r := range pendingSigRequests {
					newItem := &signatureRequestQueueItem{
						request:  r,
						maxTries: q.maxTries,
					}
					q.sigRequestChan <- newItem
				}
			}()

		}
	}
}

func (q queryProcessor) Stop() error {
	q.stop <- struct{}{}
	return nil
}
