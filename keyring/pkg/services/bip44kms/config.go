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
	Keyring       string `yaml:"keyring"`
	ChainID       string `yaml:"chainid"`
	FusionURL     string `yaml:"fusionurl"`
	Password      string `yaml:"password"` // User supplied passphrase. Must be supplied in the http header
	Mnemonic      string `yaml:"mnemonic"` // (Optional) The user can supply a mnemonic
	LogLevel      string `yaml:"loglevel"`
	LogFormat     string `yaml:"logformat"`
	LogToFile     bool   `yaml:"logtofile"`
	QueryInterval int64  `yaml:"queryinterval"`
	RetrySleep    int64  `yaml:"retrySleep"`
	MaxTries      int64  `yaml:"maxTries"`
}

var emptyConfig = ServiceConfig{}

var defaultConfig = ServiceConfig{
	Port:      9000,
	Path:      "db",
	LogLevel:  "info",
	LogFormat: "plain",
	LogToFile: false,
	Keyring:   "qredokeyring1ph63us46lyw56vrzgaq",
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
