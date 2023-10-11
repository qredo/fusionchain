package service

import "time"

var (
	defaultHandlerTimeout = 60 * time.Second
	defaultQueryTimeout   = 5 * time.Second

	defaultMaxRetries    int64 = 100
	defaultQueryInterval int64 = 5

	defaultRetryTimeout = 30 * time.Second
)

func requeueKeyItemWithTimeout(c chan *keyRequestQueueItem, item *keyRequestQueueItem, timeout time.Duration) {
	time.Sleep(timeout)
	item.retries++
	c <- item
}

func requeueSigItemWithTimeout(c chan *signatureRequestQueueItem, item *signatureRequestQueueItem, timeout time.Duration) {
	time.Sleep(timeout)
	item.retries++
	c <- item
}
