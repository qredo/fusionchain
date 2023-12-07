package api

import (
	"net/http"

	"github.com/qredo/fusionchain/keyring/pkg/rpc"
)

const (

	// API - TODO create single package for multiple services
	StatusEndPnt   = "/status"
	HealthEndPnt   = "/healthcheck"
	KeyringEndPnt  = "/keyring"  // Password protected
	PubKeysEndPnt  = "/pubkeys"  // Password protected
	MnemonicEndPnt = "/mnemonic" // Password protected
)

// Response represents the superset of Status and PubKey API responses.
type Response struct {
	Message       string    `json:"message,omitempty"`
	Version       string    `json:"version,omitempty"`
	Service       string    `json:"service,omitempty"`
	KeyRing       string    `json:"keyring,omitempty"`
	KeyringSigner string    `json:"keyring_signer,omitempty"`
	PubKeys       []*PubKey `json:"pubkeys,omitempty"`
	Mnemonic      string    `json:"mnemonic,omitempty"`
	PasswordUsed  bool      `json:"password_protected,omitempty"`
}

// HealthResponse represents the healthcheck API with no omitted fields.
type HealthResponse struct {
	Version  string   `json:"version"`
	Service  string   `json:"service"`
	Failures []string `json:"failures"`
}

type PubKey struct {
	KeyID      string `json:"key_id"`
	PubKeyData PkData `json:"pubkey_data"`
}

type PkData struct {
	PublicKey string `json:"pubkey"`
	Created   string `json:"created"`
	LastUsed  string `json:"last_used"`
}

type KeyringService interface { // Keyring service APIs
	Status(w http.ResponseWriter, r *http.Request)
	HealthCheck(w http.ResponseWriter, r *http.Request)
	Keyring(w http.ResponseWriter, r *http.Request)
	PubKeys(w http.ResponseWriter, r *http.Request)
	Mnemonic(w http.ResponseWriter, r *http.Request)
}

func MakeKeyRingAPI(k KeyringService) *rpc.API {
	r := &rpc.API{}
	r.AddEndpoint(rpc.NewEndpoint(StatusEndPnt, http.MethodGet, k.Status))
	r.AddEndpoint(rpc.NewEndpoint(HealthEndPnt, http.MethodGet, k.HealthCheck))
	r.AddEndpoint(rpc.NewEndpoint(KeyringEndPnt, http.MethodGet, k.Keyring))
	r.AddEndpoint(rpc.NewEndpoint(PubKeysEndPnt, http.MethodGet, k.PubKeys))
	r.AddEndpoint(rpc.NewEndpoint(MnemonicEndPnt, http.MethodGet, k.Mnemonic))
	return r
}
