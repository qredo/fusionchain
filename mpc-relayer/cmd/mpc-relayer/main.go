package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/qredo/fusionchain/mpc-relayer/pkg/common"
	"github.com/qredo/fusionchain/mpc-relayer/pkg/service"
)

const envPrefix = "MPCRELAYER"

var (
	configFilePath string
	configFilePtr  = flag.String("config", "config.yml", "path to config file")
)

//go run main.go --config ./config.yml
//go run main.go --config {path_to_config_file}

func init() {
	// Parse flag containing path to config file
	flag.Parse()
	if configFilePtr != nil {
		configFilePath = *configFilePtr
	}
}

func main() {
	var config service.ServiceConfig

	if err := common.ParseJSONConfig(configFilePath, &config, envPrefix); err != nil {
		log.Fatal(err)
	}
	mpcRelayer, err := service.BuildService(config)
	if err != nil {
		log.Fatal(err)
	}

	if err := mpcRelayer.Start(); err != nil {
		panic(err)
	}
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	sig := <-sigChan
	mpcRelayer.Stop(sig)
}
