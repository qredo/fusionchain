package service

import (
	"fmt"
	"strconv"
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

func BuildService(config ServiceConfig) (*Service, error) {
	if isEmpty(config) {
		return nil, fmt.Errorf("no config file supplied")
	}
	log, err := logger.NewLogger(logger.Level(config.Loglevel), logger.Format(config.LogFormat), config.LogToFile, "mpc-relayer")
	if err != nil {
		return nil, err
	}

	kv, err := database.NewBadger(config.Path, false)
	if err != nil {
		return nil, err
	}
	keyDB := database.NewPrefixDB("pk", kv)

	keyringID, identity, mpcClient, err := makeKeyringClient(&config, log)
	if err != nil {
		return nil, err
	}

	queryClient, txClient, err := makeFusionGRPCClient(&config, identity)
	if err != nil {
		return nil, err
	}

	maxRetries := config.MaxTries
	if maxRetries == 0 {
		maxRetries = defaultMaxRetries
	}
	queryInterval := config.QueryInterval
	if queryInterval == 0 {
		queryInterval = defaultQueryInterval
	}
	port := config.Port
	if port == 0 {
		port = defaultPort
	}
	// make modules
	keyChan := make(chan *keyRequestQueueItem, defaultChanSize)
	sigchan := make(chan *signatureRequestQueueItem, defaultChanSize)
	return New(keyringID, port, log, keyDB,
		newKeyQueryProcessor(keyringID, queryClient, keyChan, log, time.Duration(queryInterval)*time.Second, int(maxRetries)),
		newSigQueryProcessor(keyringID, queryClient, sigchan, log, time.Duration(queryInterval)*time.Second, int(maxRetries)),
		newFusionKeyController(log, keyDB, keyChan, mpcClient, txClient),
		newFusionSignatureController(log, keyDB, sigchan, mpcClient, txClient),
	), nil

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
