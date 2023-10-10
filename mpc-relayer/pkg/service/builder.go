package service

/*
import (
	"time"

	"github.com/qredo/assets/libs/nodeconnector"
	"github.com/qredo/assets/libs/watcher/healthcheck"
	"github.com/qredo/assets/libs/watcher/qredochain"
	"github.com/qredo/assets/libs/watcher/services/writer"
	"github.com/qredo/fusionchain/mpc-relayer/pkg/database"
	"github.com/qredo/fusionchain/mpc-relayer/pkg/logger"
	"github.com/qredo/fusionchain/mpc-relayer/pkg/mpc"
)

type ServiceConfig struct {
	Port        int        `yaml:"port"`
	Path        string     `yaml:"path"`
	FusionURL   string     `yaml:"fusion_url"`
	Loglevel    string     `yaml:"loglevel"`
	LogFormat   string     `yaml:"logformat"`
	LogToFile   bool       `yaml:"logtofile"`
	MPC         mpc.Config `yaml:"mpc"`
	InitVersion int        `yaml:"initVersion"`
	RetrySleep  int64      `yaml:"retrySleep"`
	MaxTries    int64      `yaml:"maxTries"`
	IndexStart  int64      `yaml:"indexStart"`
}

func BuildService(config ServiceConfig) (*Service, error) {
	log, err := logger.NewLogger(logger.Level(config.Loglevel), logger.Format(config.LogFormat), config.LogToFile, "signer")
	if err != nil {
		return nil, err
	}

	kv, err := database.NewBadger(config.Path, false)
	if err != nil {
		return nil, err
	}

	mpcClient := mpc.NewClient(config.MPC, log, config.InitVersion)

	catchUp := qredochain.NewCatchUpFunc(nodeconnector.TendermintKeySearcher{Client: tmClient}, log)
	processor := NewProcessor(nodeconnector.TendermintKeySearcher{Client: tmClient}, mpcClient, writer.NewClient(config.Writer))
	retriever := NewRetriever(tmClient)
	httpService := healthcheck.NewHTTPService(healthcheck.Config{Service: "signer"}, config.Qredochain)
	service := New(kv, processor, catchUp, retriever, Config{
		RetrySleep: time.Duration(config.RetrySleep) * time.Millisecond,
		MaxTries:   config.MaxTries,
		IndexStart: config.IndexStart,
	}, httpService, tmClient)

	return service, nil
}
