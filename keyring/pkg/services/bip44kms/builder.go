package kms

import (
	"time"

	"github.com/qredo/fusionchain/go-client"
	"github.com/qredo/fusionchain/keyring/pkg/database"
	"github.com/qredo/fusionchain/keyring/pkg/logger"
	"github.com/qredo/fusionchain/keyring/pkg/mpc"
	"github.com/qredo/fusionchain/keyring/pkg/services/mpcrelayer"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultFusionURL     = "localhost:9090"
	defaultFusionChainID = "qredofusiontestnet_257-1"
)

// BuildService constructs the main application based on supplied config parameters
func BuildService(config ServiceConfig) (*Service, error) {

	cfg, useDefault := sanitizeConfig(config) // set default values is none supplied

	log, err := logger.NewLogger(logger.Level(config.LogLevel), logger.Format(config.LogFormat), config.LogToFile, serviceName)
	if err != nil {
		return nil, err
	}
	if useDefault {
		log.Info("no config file supplied, using default values")
	}

	keyDB, err := makeKeyDB(config.Path, false)
	if err != nil {
		return nil, err
	}

	keyringAddr, mnemonic, password, identity, keyRingClient, err := makeKeyringClient(&cfg, log)
	if err != nil {
		return nil, err
	}

	// Warn users is no password has been supplied
	if password == "" {
		log.Warn("WARNING! NO PASSWORD HAS BEEN SUPPLIED. YOUR PRIVATE KEY DATA IS NOT SECURE.")
	}

	queryClient, txClient, err := makeFusionGRPCClient(&cfg, identity)
	if err != nil {
		return nil, err
	}

	// make modules
	keyChan := make(chan *keyRequestQueueItem, defaultChanSize)
	sigchan := make(chan *signatureRequestQueueItem, defaultChanSize)
	return New(keyringAddr, identity.Address.String(), mnemonic, password, cfg.Port, log, keyDB,
		newKeyQueryProcessor(keyringAddr, queryClient, keyChan, log, time.Duration(cfg.QueryInterval)*time.Second, int(cfg.MaxTries)),
		newSigQueryProcessor(keyringAddr, queryClient, sigchan, log, time.Duration(cfg.QueryInterval)*time.Second, int(cfg.MaxTries)),
		newFusionKeyController(log, keyDB, keyChan, keyRingClient, txClient),
		newFusionSignatureController(log, keyDB, sigchan, keyRingClient, txClient),
	), nil
}

func makeKeyDB(path string, inMemory bool) (database.Database, error) {
	kv, err := database.NewBadger(path, inMemory)
	if err != nil {
		return nil, err
	}
	return database.NewPrefixDB("pk", kv), nil
}

func makeKeyringClient(config *ServiceConfig, log *logrus.Entry) (keyringAddr, mnemomic, password string, identity client.Identity, keyRing Keyring, err error) {
	keyringAddr = config.KeyringAddr
	mnemomic = config.Mnemonic
	password = config.Password

	if mnemomic == "" {
		// If no mnemomic ENV VAR is supplied, create a new one
		// Note that the mnemomic is NOT persistently stored
		// Once created it can only be accessed via the '/mnemonic' endpoint
		mnemomic, err = GenerateMnemonic()
		if err != nil {
			return
		}
	}

	keyRing, err = NewBip44KeyRing(mnemomic, password, mpc.EcDSA)
	if err != nil {
		return
	}

	identity, err = client.NewIdentityFromSeed(fusionIdentityDerivationPath, config.Mnemonic)
	if err != nil {
		return
	}
	return
}

func makeFusionGRPCClient(config *ServiceConfig, identity client.Identity) (mpcrelayer.QueryClient, mpcrelayer.TxClient, error) {
	fusionGRPCClient, err := grpc.Dial(
		config.FusionURL,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, nil, err
	}
	queryClient := client.NewQueryClientWithConn(fusionGRPCClient)
	txClient := client.NewTxClient(identity, config.ChainID, fusionGRPCClient, queryClient)
	return queryClient, txClient, nil
}
