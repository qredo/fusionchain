package kms

import (
	"bytes"

	"gopkg.in/yaml.v3"
)

// ServiceConfig represents the main application configuration struct.
// Example YAML config can be founs in github.com/qredo/fusionchain/mpc-relayer/docker-compose/config-example.yaml
type ServiceConfig struct {
	Port          int    `yaml:"port"`
	Path          string `yaml:"path"`
	KeyringAddr   string `yaml:"keyring_addr"`
	ChainID       string `yaml:"chain_id"`
	FusionURL     string `yaml:"fusion_url"`
	Password      string `yaml:"password"` // User supplied passphrase. Must be supplied in the http header
	Mnemonic      string `yaml:"mnemonic"` // (Optional) The user can supply a mnemonic
	LogLevel      string `yaml:"loglevel"`
	LogFormat     string `yaml:"logformat"`
	LogToFile     bool   `yaml:"logtofile"`
	QueryInterval int64  `yaml:"query_interval"`
	RetrySleep    int64  `yaml:"retrySleep"`
	MaxTries      int64  `yaml:"maxTries"`
}

var emptyConfig = ServiceConfig{}

var defaultConfig = ServiceConfig{
	Port:        9000,
	Path:        "db",
	KeyringAddr: "qredokeyring1ph63us46lyw56vrzgaq",
}

func isEmpty(c ServiceConfig) bool {
	b, _ := yaml.Marshal(c)
	e, _ := yaml.Marshal(emptyConfig)
	return bytes.Equal(b, e)
}

// sanitizeConfig Partially empty configs will be sanitized with default values.
func sanitizeConfig(config ServiceConfig) (cfg ServiceConfig, defaultUsed bool) {
	if isEmpty(config) {
		defaultUsed = true
		cfg = defaultConfig
		return
	}
	cfg = config

	if config.MaxTries == 0 {
		cfg.MaxTries = defaultMaxRetries
	}

	if config.QueryInterval == 0 {
		cfg.QueryInterval = defaultQueryInterval
	}

	if config.Port == 0 {
		cfg.Port = defaultPort
	}

	if config.FusionURL == "" {
		cfg.FusionURL = defaultFusionURL
	}

	if config.ChainID == "" {
		cfg.FusionURL = defaultFusionChainID
	}
	return
}
