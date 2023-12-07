package kms

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/qredo/fusionchain/keyring/pkg/api"
	"github.com/qredo/fusionchain/keyring/pkg/common"
	"github.com/qredo/fusionchain/keyring/pkg/rpc"
)

const (
	pwdHeaderKey = "password"
)

var (
	errInvalidPswd = errors.New("invalid password")
)

// Status handles the /api/status query and will always respond OK
func (s *Service) Status(w http.ResponseWriter, _ *http.Request) {
	resp := api.Response{Message: "OK", Version: common.FullVersion, Service: serviceName}
	if err := rpc.RespondWithJSON(w, http.StatusOK, resp); err != nil {
		s.log.Error(err)
	}
}

// Healthcheck handles the the /healthcheck query.
func (s *Service) HealthCheck(w http.ResponseWriter, _ *http.Request) {
	health := &api.HealthResponse{
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
func (s *Service) Keyring(w http.ResponseWriter, h *http.Request) {
	pwd := h.Header.Get(pwdHeaderKey)
	if s.secrets.password != pwd {
		rpc.RespondWithError(w, http.StatusBadRequest, errInvalidPswd)
		return
	}
	if err := rpc.RespondWithJSON(w, http.StatusOK, &api.Response{
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
func (s *Service) PubKeys(w http.ResponseWriter, h *http.Request) {
	pwd := h.Header.Get(pwdHeaderKey)
	if s.secrets.password != pwd {
		rpc.RespondWithError(w, http.StatusBadRequest, errInvalidPswd)
		return
	}
	pKeyResponse := &api.Response{
		Service: serviceName,
		Version: common.FullVersion,
		Message: "OK",
	}

	keyMap, err := s.dB.Read(pkPrefix)
	if err != nil {
		pKeyResponse.Message = err.Error()
		if err := rpc.RespondWithJSON(w, http.StatusInternalServerError, pKeyResponse); err != nil {
			s.log.Error(err)
		}
		return
	}

	var pubKeyList []*api.PubKey
	for keyID, pKDat := range keyMap {
		dt := api.PkData{}
		if err := json.Unmarshal(pKDat, &dt); err != nil {
			continue // TODO - handle error properly
		}
		pubKeyList = append(pubKeyList, &api.PubKey{KeyID: keyID, PubKeyData: dt})
	}
	pKeyResponse.PubKeys = pubKeyList

	if err := rpc.RespondWithJSON(w, http.StatusOK, pKeyResponse); err != nil {
		s.log.Error(err)
	}
}

// Mnemonic implements the /mnemonic endpoint, returning the BIP39 seed phrase used to derive the keyring's master seed.
// PASSWORD PROTECTION is used, the http header must contain the correct password for the service.
func (s *Service) Mnemonic(w http.ResponseWriter, h *http.Request) {
	pwd := h.Header.Get(pwdHeaderKey)
	if s.secrets.password != pwd {
		rpc.RespondWithError(w, http.StatusBadRequest, errInvalidPswd)
		return
	}
	if err := rpc.RespondWithJSON(w, http.StatusOK, &api.Response{
		Service:      serviceName,
		Version:      common.FullVersion,
		Message:      "OK",
		Mnemonic:     s.secrets.mnemonic,
		PasswordUsed: (s.secrets.password != ""),
	}); err != nil {
		s.log.Error(err)
	}
}
