package service

/*
import (
	"encoding/hex"
	"fmt"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/qredo/assets/libs/assets"
	"github.com/qredo/assets/libs/common"
	"github.com/qredo/assets/libs/constants"
	"github.com/qredo/assets/libs/logger"
	"github.com/qredo/assets/libs/store"
	"github.com/qredo/assets/libs/watcher/database"
	"github.com/qredo/assets/libs/watcher/qredochain"
	"github.com/qredo/assets/libs/watcher/services"
)

var log = logger.Empty

type retryableQueueItem struct {
	*QueueItem
	tries int64
	err   error
}

const (
	defaultRetrySleep = 4 * time.Second
	defaultMaxTries   = 5
)

type Service struct {
	processor   Processor
	catchUpFunc qredochain.CatchUpFunc
	retriever   Retriever
	mu          sync.Mutex
	waiting     chan *retryableQueueItem
	modules     []services.Module
	db          *database.Index

	stop    chan struct{}
	stopped atomic.Bool

	retrySleep time.Duration
	maxTries   int64
	indexStart int64
}

type Config struct {
	RetrySleep time.Duration
	MaxTries   int64
	IndexStart int64
}

func defaultIfZero[T int64 | time.Duration](v T, d T) T {
	if v == 0 {
		return d
	}
	return v
}

func New(kv store.KV, processor *QredoProcessor, catchUpFunc qredochain.CatchUpFunc, retriever Retriever, config Config, modules ...services.Module) *Service {
	return &Service{
		db:          database.NewIndex(kv),
		processor:   processor,
		catchUpFunc: catchUpFunc,
		retriever:   retriever,
		waiting:     make(chan *retryableQueueItem),
		stop:        make(chan struct{}),
		modules:     modules,
		retrySleep:  defaultIfZero(config.RetrySleep, defaultRetrySleep),
		maxTries:    defaultIfZero(config.MaxTries, defaultMaxTries),
		indexStart:  config.IndexStart,
	}
}

func (s *Service) process() {
	for {
		select {
		case <-s.stop:
			return
		case item, ok := <-s.waiting:
			if !ok {
				continue
			}
			log := log.WithFields(logrus.Fields{
				"walletID": fmt.Sprintf("%x", item.TxAsset.GetAssetID()),
				"tries":    item.tries,
				"index":    item.Index,
			})
			if item.tries > s.maxTries {
				log.WithFields(logrus.Fields{
					"error": item.err.Error(),
				}).Error("" +
					"signature discarded")
				continue
			}
			go func(item *retryableQueueItem) {
				if item.tries > 1 {
					time.Sleep(s.retrySleep)
				}
				item.tries++
				walletID, mutableIndex, description, err := getItemDetails(item.TxAsset)
				if err != nil {
					log.WithFields(logrus.Fields{
						"error": err.Error(),
					}).Warn("getItemDetailsErr")
					item.err = err
					if s.stopped.Load() {
						return
					}
					s.mu.Lock()
					s.waiting <- item
					s.mu.Unlock()
					return
				}
				result, err := s.processor.SignOnlyProcess(walletID, mutableIndex, description)
				if err != nil {
					log.WithFields(logrus.Fields{
						"error": err.Error(),
					}).Warn("signature failed")
					item.err = err
					if s.stopped.Load() {
						return
					}
					s.mu.Lock()
					s.waiting <- item
					s.mu.Unlock()
					return
				}
				go func() {
					if err := s.db.LastIndex(item.Index); err != nil {
						log.WithFields(logrus.Fields{
							"error": err.Error(),
						}).Warn("persisting completed index failed")
					}
				}()
				log.WithFields(prettyPrintResult(result)).Info("externalWalletSignature")
			}(item)
		}
	}
}

func getItemDetails(tx assets.TXAsset) (walletID []byte, mutableIndex int64, description string, err error) {
	if w, ok := tx.(*assets.Wallet); ok {
		return w.GetAssetID(), w.CurrentAsset.MutableIndex, w.CurrentAsset.Asset.Description, nil
	} else if c, ok := tx.(*assets.Control); ok {
		pl, err := c.Payload()
		if err != nil {
			return nil, 0, "", err
		}
		return pl.LOneAsset.WalletID, 0, "", nil
	}
	return nil, 0, "", fmt.Errorf("unknown asset type")
}

func prettyPrintResult(result signResult) logrus.Fields {
	lf := logrus.Fields{}
	lf["currency"] = result.currency.String()
	// TODO - in theory the signer supports multiple signatures per request
	// but this is not used and/or supported by the qredochain. We need to
	// either add support for a slice of signatures displayed in the logs or refactor the
	// SignOnlyProcess function to not return a slice of externalTxMsg structs - Alex
	if len(result.transactions) > 0 {
		tx := result.transactions[0]
		lf["chainID"] = tx.chainID
		lf["sender"] = tx.sender
		lf["txID"] = tx.txid
		lf["signedMessage"] = fmt.Sprintf("%x", tx.rawSigPayload)
	} else {
		lf["error"] = "no signature result"
	}
	return lf
}

func (s *Service) receive(out <-chan *QueueItem) {
	for {
		select {
		case <-s.stop:
			return
		case item, ok := <-out:
			if !ok {
				continue
			}
			s.mu.Lock()
			s.waiting <- &retryableQueueItem{
				QueueItem: item,
			}
			s.mu.Unlock()
		}
	}
}

func (s *Service) Start() error {
	log.WithFields(logrus.Fields{
		"version":   constants.FullVersion,
		"buildDate": constants.Date,
	}).Info("starting signer service")

	for _, module := range s.modules {
		if err := module.Start(); err != nil {
			log.WithError(err).Error("cannot start module")
		}
	}

	lastIndex, err := s.lastIndex()
	if err != nil {
		return err
	}

	out, err := s.retriever.Start(lastIndex)
	if err != nil {
		return err
	}

	go s.CatchUp()
	go s.process()
	go s.receive(out)

	return nil
}

func (s *Service) Stop(sig os.Signal) {
	log.WithFields(logrus.Fields{"signal": sig}).Info("received shutdown signal")
	if s.stopped.Load() {
		return
	}
	s.stopped.Store(true)
	close(s.stop)
	s.mu.Lock()
	close(s.waiting)
	s.mu.Unlock()
	for _, module := range s.modules {
		if err := module.Stop(); err != nil {
			log.WithError(err).Error("cannot stop module")
		}
	}
}

func (s *Service) CatchUp() {
	walletIDs, err := s.catchUp()
	if err != nil {
		log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("catch up error")
		return
	}
	log.WithFields(logrus.Fields{
		"wallets": walletIDs,
	}).Info("catch up")
}

type catchUpResult struct {
	ID    string
	Index int64
}

func (s *Service) lastIndex() (int64, error) {
	lastIndex, err := s.db.GetLastIndex()
	if err != nil && err != store.ErrNotFound {
		return 0, err
	}
	if lastIndex != 0 {
		lastIndex++ // add one to skip actual last index processed
	}
	return common.MaxInt64(lastIndex, s.indexStart), nil
}

func (s *Service) catchUp() ([]catchUpResult, error) {
	index, err := s.lastIndex()
	if err != nil {
		return nil, err
	}

	items, err := s.catchUpFunc(index)
	if err != nil {
		return nil, err
	}
	var walletIDs []catchUpResult
	for _, item := range items {
		ignore, err := filterSignature(item)
		if err != nil {
			return nil, err
		}
		if ignore {
			continue
		}
		walletIDs = append(walletIDs, catchUpResult{
			ID:    hex.EncodeToString(item.TxAsset.GetAssetID()),
			Index: item.Index,
		})
		s.waiting <- &retryableQueueItem{QueueItem: item}
	}
	return walletIDs, nil
}
