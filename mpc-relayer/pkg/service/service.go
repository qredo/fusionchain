package service

import (
	"os"
	"sync/atomic"

	"github.com/sirupsen/logrus"

	"github.com/qredo/fusionchain/mpc-relayer/pkg/common"
)

type Service struct {
	keyringID uint64
	modules   []Module
	log       *logrus.Entry

	stop    chan struct{}
	stopped atomic.Bool
}

func (s *Service) Start() error {
	s.log.WithFields(logrus.Fields{
		"version":   common.FullVersion,
		"buildDate": common.Date,
	}).Info("starting mpc-relayer service")

	for _, module := range s.modules {
		if err := module.Start(); err != nil {
			s.log.WithError(err).Error("cannot start module")
		}
	}

	return nil
}

func (s *Service) Stop(sig os.Signal) {
	s.log.WithFields(logrus.Fields{"signal": sig}).Info("received shutdown signal")
	if s.stopped.Load() {
		s.log.WithFields(logrus.Fields{"signal": sig}).Warn("already shutting down")
		return
	}
	s.stopped.Store(true)
	close(s.stop)
	for _, module := range s.modules {
		if err := module.Stop(); err != nil {
			s.log.WithError(err).Error("cannot stop module")
		}
	}
	s.log.Info("mpc-relayer stopped")
}
