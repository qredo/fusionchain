package kms

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/qredo/fusionchain/keyring/pkg/api"
	"github.com/qredo/fusionchain/keyring/pkg/common"
	"github.com/qredo/fusionchain/keyring/pkg/logger"
)

var testConfig = ServiceConfig{
	Port:      8080,
	Keyring:   "qredokeyring1ph63us46lyw56vrzgaq",
	LogLevel:  "fatal",
	LogFormat: "plain",
	LogToFile: false,
	Mnemonic:  "exclude try nephew main caught favorite tone degree lottery device tissue tent ugly mouse pelican gasp lava flush pen river noise remind balcony emerge",
}

var (
	tests = []struct {
		name                        string
		config                      ServiceConfig
		modules                     []Module
		buildErr, startErr, stopErr bool
	}{
		{
			"empty config",
			ServiceConfig{},
			nil,
			false,
			false,
			false,
		},
		{
			"no mnemonic",
			ServiceConfig{
				Port:      8080,
				Keyring:   "qredokeyring1ph63us46lyw56vrzgaq",
				LogLevel:  "fatal",
				LogFormat: "plain",
				LogToFile: false,
			},
			nil,
			false,
			false,
			false,
		},
		{
			"with mnemonic",
			ServiceConfig{
				Port:      8080,
				Keyring:   "qredokeyring1ph63us46lyw56vrzgaq",
				LogLevel:  "fatal",
				LogFormat: "plain",
				LogToFile: false,
				Mnemonic:  testMnemonic,
			},
			nil,
			false,
			false,
			false,
		},
		{
			"no modules",
			testConfig,
			nil,
			false,
			false,
			false,
		},
		{
			"single module",
			testConfig,
			[]Module{mockModule{}},
			false,
			false,
			false,
		},
		{
			"multiple module",
			testConfig,
			[]Module{mockModule{}, mockModule{}},
			false,
			false,
			false,
		},
		{
			"module with error",
			testConfig,
			[]Module{mockModuleErr{}},
			false,
			true,
			true,
		},
	}

	configTests = []struct {
		name     string
		config   ServiceConfig
		buildErr bool
	}{
		{
			"empty config",
			emptyConfig,
			false,
		},
		{
			"default config",
			defaultConfig,
			false,
		},
		{
			"invalid log parameters",
			ServiceConfig{
				LogLevel: "fatall",
			},
			true,
		},
	}
)

type mockModule struct{}

func (m mockModule) Start() error {
	return nil
}

func (m mockModule) Stop() error {
	return nil
}

func (m mockModule) healthcheck() *api.HealthResponse {
	return &api.HealthResponse{}
}

type mockModuleErr struct{}

func (m mockModuleErr) Start() error {
	return errors.New("error")
}

func (m mockModuleErr) Stop() error {
	return errors.New("error")
}

func (m mockModuleErr) healthcheck() *api.HealthResponse {
	return &api.HealthResponse{Failures: []string{"some failure"}}
}

func Test_ServiceStartStop(t *testing.T) {
	// build service with different 'module' combinations
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := buildTestService(t, tt.config, tt.modules...)
			if (err != nil) != tt.buildErr {
				t.Fatalf("unexpected build error %v", err)
			}
			if err != nil {
				return
			}
			if err := s.Start(); (err != nil) != tt.startErr {
				t.Fatalf("unexpected start error %v", err)
			}
			if err := s.Stop(os.Interrupt); (err != nil) != tt.stopErr {
				t.Fatalf("unexpected stop error %v", err)
			}
		})
	}

}

func Test_ConfigTypes(t *testing.T) {
	// build service with different config types
	for _, tt := range configTests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := BuildService(tt.config)
			if (err != nil) != tt.buildErr {
				t.Fatalf("unexpected build error %v", err)
			}
		})
	}

}

func Test_ServiceAPI(t *testing.T) {
	s, err := buildTestService(t, tests[2].config) // Use test config with fixed mnemonic
	if err != nil {
		t.Fatal(err)
	}

	apiTests := []struct {
		name             string
		endpoint         string
		method           func(w http.ResponseWriter, req *http.Request)
		expectedResponse any
		expectedCode     int
	}{
		{
			"status",
			api.StatusEndPnt,
			s.Status,
			&api.Response{Message: "OK", Version: common.FullVersion, Service: serviceName},
			http.StatusOK,
		},
		{
			"healthcheck",
			api.HealthEndPnt,
			s.HealthCheck,
			&api.HealthResponse{Version: common.FullVersion, Service: serviceName, Failures: []string{}},
			http.StatusOK,
		},
		{
			"keyring",
			api.KeyringEndPnt,
			s.Keyring,
			&api.Response{Message: "OK", Version: common.FullVersion, Service: serviceName, KeyRing: defaultConfig.Keyring, KeyringSigner: "qredo1r7dhrn6ljwj72akjhpgslvqwx6kq2xzypz8sm6"},
			http.StatusOK,
		},
		{
			"pubkeys",
			api.PubKeysEndPnt,
			s.PubKeys,
			&api.Response{Message: "OK", Version: common.FullVersion, Service: serviceName},
			http.StatusOK,
		},
		{
			"mnemonic",
			api.MnemonicEndPnt,
			s.Mnemonic,
			&api.Response{Message: "OK", Version: common.FullVersion, Service: serviceName, Mnemonic: testMnemonic},
			http.StatusOK,
		},
	}

	for _, tt := range apiTests {
		t.Run(tt.name, func(t *testing.T) {
			httpReq := httptest.NewRequest(http.MethodGet, tt.endpoint, nil)
			respRecorder := httptest.NewRecorder()
			tt.method(respRecorder, httpReq)
			if g, w := respRecorder.Code, tt.expectedCode; g != w {
				t.Errorf("unexpected response code, want %v got %v", w, g)
			}
			expectedJSON, _ := json.Marshal(tt.expectedResponse)

			if g, w := respRecorder.Body.Bytes(), expectedJSON; !bytes.Equal(g, w) {
				t.Fatalf("unexpected reponse, want %s, got %s", w, g)
			}
		})
	}

}

func buildTestService(t *testing.T, config ServiceConfig, modules ...Module) (*Service, error) {

	config, _ = sanitizeConfig(config)

	log, err := logger.NewLogger(logger.Level(config.LogLevel), logger.Format(config.LogFormat), config.LogToFile, "test")
	if err != nil {
		return nil, err
	}
	memoryKeyDB, err := makeDB("", true)
	if err != nil {
		t.Fatal(err)
	}
	keyringAddr, mn, p, i, _, err := makeKeyringClient(&config, log, memoryKeyDB)
	if err != nil {
		return nil, err
	}

	return New(keyringAddr, i.Address.String(), mn, p, config.Port, log, memoryKeyDB, modules...), nil
}
