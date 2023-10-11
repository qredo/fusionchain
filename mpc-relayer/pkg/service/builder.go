package service

import (
	"fmt"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/qredo/fusionchain/go-client"
	"github.com/qredo/fusionchain/mpc-relayer/pkg/database"
	"github.com/qredo/fusionchain/mpc-relayer/pkg/logger"
	"github.com/qredo/fusionchain/mpc-relayer/pkg/mpc"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	derivationPath = "m/44'/60'/0'/0/0"
	fusionChainID  = "fusion_420-1"
)

type Module interface {
	Start() error
	Stop() error
}

func BuildService(config ServiceConfig) (*Service, error) {
	if isEmpty(config) {
		return nil, fmt.Errorf("no config file supplied")
	}
	log, err := logger.NewLogger(logger.Level(config.Loglevel), logger.Format(config.LogFormat), config.LogToFile, "signer")
	if err != nil {
		return nil, err
	}

	kv, err := database.NewBadger(config.Path, false)
	if err != nil {
		return nil, err
	}

	keyringID, identity, mpcClient, err := makeKeyringClient(&config, log)
	if err != nil {
		return nil, err
	}

	queryClient, txClient, err := makeFusionGRPCClient(&config, identity)
	if err != nil {
		return nil, err
	}
	// make modules

	keyChan := make(chan *keyRequestQueueItem)
	sigchan := make(chan *signatureRequestQueueItem)

	maxRetries := config.MaxTries
	if maxRetries == 0 {
		maxRetries = defaultMaxRetries
	}
	queryInterval := config.QueryInterval
	if queryInterval == 0 {
		queryInterval = defaultQueryInterval
	}

	return &Service{
		keyringID: keyringID,
		modules: []Module{
			newQueryProcessor(keyringID, queryClient, keyChan, sigchan, log, time.Duration(queryInterval)*time.Second, int(maxRetries)),
			newFusionKeyController(log, kv, keyChan, mpcClient, txClient),
			newFusionSignatureController(log, kv, sigchan, mpcClient, txClient),
		},
		stop:    make(chan struct{}),
		stopped: atomic.Bool{},
	}, nil
}

func makeKeyringClient(config *ServiceConfig, log *logrus.Entry) (keyringID uint64, identity client.Identity, mpcClient mpc.Client, err error) {
	mpcClient = mpc.NewClient(config.MPC, log)

	keyringID, err = strconv.ParseUint(config.KeyRingID, 10, 64)
	if err != nil {
		return
	}

	identity, err = client.NewIdentityFromSeed(derivationPath, config.Mnemonic)
	if err != nil {
		return
	}
	return
}

func makeFusionGRPCClient(config *ServiceConfig, identity client.Identity) (QueryClient, TxClient, error) {
	fusionGRPCClient, err := grpc.Dial(
		config.FusionURL,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, nil, err
	}
	queryClient := client.NewQueryClientWithConn(fusionGRPCClient)
	txClient := client.NewTxClient(identity, fusionChainID, fusionGRPCClient, queryClient)
	return queryClient, txClient, nil

}
