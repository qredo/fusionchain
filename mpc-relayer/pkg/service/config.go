package service

import (
	"bytes"

	"github.com/qredo/fusionchain/mpc-relayer/pkg/mpc"
	"gopkg.in/yaml.v3"
)

type ServiceConfig struct {
	Port          int        `yaml:"port"`
	Path          string     `yaml:"path"`
	KeyRingID     string     `yaml:"keyring_id"`
	FusionURL     string     `yaml:"fusion_url"`
	Mnemonic      string     `yaml:"mnemonic"`
	LogLevel      string     `yaml:"loglevel"`
	LogFormat     string     `yaml:"logformat"`
	LogToFile     bool       `yaml:"logtofile"`
	MPC           mpc.Config `yaml:"mpc"`
	QueryInterval int64      `yaml:"query_interval"`
	RetrySleep    int64      `yaml:"retrySleep"`
	MaxTries      int64      `yaml:"maxTries"`
}

var emptyConfig = ServiceConfig{}

func isEmpty(c ServiceConfig) bool {
	b, _ := yaml.Marshal(c)
	e, _ := yaml.Marshal(emptyConfig)
	return bytes.Equal(b, e)
}
