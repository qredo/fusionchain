package kms

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/qredo/fusionchain/keyring/pkg/common"
	"github.com/qredo/fusionchain/keyring/pkg/rpc"
)

const (
	pwdHeaderKey = "password"

	// API
	statusEndPnt   = "/status"
	healthEndPnt   = "/healthcheck"
	keyringEndPnt  = "/keyring"  // Password protected
	pubKeysEndPnt  = "/pubkeys"  // Password protected
	mnemonicEndPnt = "/mnemonic" // Password protected
)

var (
	errInvalidPswd = errors.New("invalid password")
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
	KeyID     string `json:"key_id"`
	PublicKey string `json:"pubkey"`
}

func makeAPIHandlers(s *Service) *rpc.API {
	r := &rpc.API{}
	r.AddEndpoint(rpc.NewEndpoint(statusEndPnt, http.MethodGet, s.status))
	r.AddEndpoint(rpc.NewEndpoint(healthEndPnt, http.MethodGet, s.healthcheck))
	r.AddEndpoint(rpc.NewEndpoint(keyringEndPnt, http.MethodGet, s.keyring))
	r.AddEndpoint(rpc.NewEndpoint(pubKeysEndPnt, http.MethodGet, s.pubKeys))
	r.AddEndpoint(rpc.NewEndpoint(mnemonicEndPnt, http.MethodGet, s.mnemonic))
	return r
}

// Status handles the /api/status query and will always respond OK
func (s *Service) status(w http.ResponseWriter, _ *http.Request) {
	resp := Response{Message: "OK", Version: common.FullVersion, Service: serviceName}
	if err := rpc.RespondWithJSON(w, http.StatusOK, resp); err != nil {
		s.log.Error(err)
	}
}

// Healthcheck handles the the /healthcheck query.
func (s *Service) healthcheck(w http.ResponseWriter, _ *http.Request) {
	health := &HealthResponse{
		Service: serviceName,
		Version: common.FullVersion,
	}
	var failures = []string{}

	for _, sub := range s.modules {
		// verify all subprocesses are healthy
		r := sub.healthcheck()
		failures = append(failures, r.Failures...)
	}

	health.Failures = failures
	if len(failures) > 0 {
		if err := rpc.RespondWithJSON(w, http.StatusServiceUnavailable, health); err != nil {
			s.log.Error(err)
		}
		return
	}
	if err := rpc.RespondWithJSON(w, http.StatusOK, health); err != nil {
		s.log.Error(err)
	}
}

// Keyring implements the /keyring endpoint, keyring address registered for the service.
// PASSWORD PROTECTION is used, the http header must contain the correct password for the service.
func (s *Service) keyring(w http.ResponseWriter, h *http.Request) {
	pwd := h.Header.Get(pwdHeaderKey)
	if s.secrets.password != pwd {
		rpc.RespondWithError(w, http.StatusBadRequest, errInvalidPswd)
		return
	}
	if err := rpc.RespondWithJSON(w, http.StatusOK, &Response{
		Service:       serviceName,
		Version:       common.FullVersion,
		Message:       "OK",
		KeyRing:       s.keyringAddr,
		KeyringSigner: s.keyringSigner,
	}); err != nil {
		s.log.Error(err)
	}
}

// PubKeys implements the /pubkeys endpoint, returning a list of registered keyID and public keys
// stored in the local database. PASSWORD PROTECTION is used, the http header must contain the correct password for
// the service.
func (s *Service) pubKeys(w http.ResponseWriter, h *http.Request) {
	pwd := h.Header.Get(pwdHeaderKey)
	if s.secrets.password != pwd {
		rpc.RespondWithError(w, http.StatusBadRequest, errInvalidPswd)
		return
	}
	pKeyResponse := &Response{
		Service: serviceName,
		Version: common.FullVersion,
		Message: "OK",
	}

	keyMap, err := s.keyDB.Read("")
	if err != nil {
		pKeyResponse.Message = err.Error()
		if err := rpc.RespondWithJSON(w, http.StatusInternalServerError, pKeyResponse); err != nil {
			s.log.Error(err)
		}
		return
	}

	var pubKeyList []*PubKey
	for keyID, pK := range keyMap {
		pubKeyList = append(pubKeyList, &PubKey{KeyID: keyID, PublicKey: fmt.Sprintf("%x", pK)})
	}
	pKeyResponse.PubKeys = pubKeyList

	if err := rpc.RespondWithJSON(w, http.StatusOK, pKeyResponse); err != nil {
		s.log.Error(err)
	}
}

// Mnemonic implements the /mnemonic endpoint, returning the BIP39 seed phrase used to derive the keyring's master seed.
// PASSWORD PROTECTION is used, the http header must contain the correct password for the service.
func (s *Service) mnemonic(w http.ResponseWriter, h *http.Request) {
	pwd := h.Header.Get(pwdHeaderKey)
	if s.secrets.password != pwd {
		rpc.RespondWithError(w, http.StatusBadRequest, errInvalidPswd)
		return
	}
	if err := rpc.RespondWithJSON(w, http.StatusOK, &Response{
		Service:      serviceName,
		Version:      common.FullVersion,
		Message:      "OK",
		Mnemonic:     s.secrets.mnemonic,
		PasswordUsed: (s.secrets.password != ""),
	}); err != nil {
		s.log.Error(err)
	}
}
