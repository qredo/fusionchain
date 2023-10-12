package service

import (
	"fmt"
	"net/http"

	"github.com/qredo/fusionchain/mpc-relayer/pkg/common"
	"github.com/qredo/fusionchain/mpc-relayer/pkg/rpc"
)

const (
	statusEndPnt  = "/status"
	healthEndPnt  = "/healthcheck"
	pubKeysEndPnt = "/pubkeys"
)

// Response represents the superset of all API responses
type Response struct {
	Message  string    `json:"message,omitempty"`
	Version  string    `json:"version,omitempty"`
	Service  string    `json:"service,omitempty"`
	Failures []string  `json:"failures,omitempty"`
	PubKeys  []*PubKey `json:"pubkeys,omitempty"`
}

type PubKey struct {
	KeyID     string `json:"key_id"`
	PublicKey string `json:"pubkey"`
}

func makeAPIHandlers(s *Service) *rpc.Api {
	r := &rpc.Api{}
	r.AddEndpoint(rpc.NewEndpoint(statusEndPnt, http.MethodGet, s.status))
	r.AddEndpoint(rpc.NewEndpoint(healthEndPnt, http.MethodGet, s.healthcheck))
	r.AddEndpoint(rpc.NewEndpoint(pubKeysEndPnt, http.MethodGet, s.pubKeys))
	return r
}

// Status handles the /api/status query and will always respond OK
func (s *Service) status(w http.ResponseWriter, req *http.Request) {
	resp := Response{Message: "OK", Version: common.FullVersion, Service: serviceName}
	if err := rpc.RespondWithJSON(w, http.StatusOK, resp); err != nil {
		s.log.Error(err)
	}
}

// Healthcheck handles the the /api/healthcheck query.
func (s *Service) healthcheck(w http.ResponseWriter, req *http.Request) {
	health := &Response{
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

func (s *Service) pubKeys(w http.ResponseWriter, req *http.Request) {
	pKeyReponse := &Response{
		Service: serviceName,
		Version: common.FullVersion,
		Message: "OK",
	}

	keyMap, err := s.keyDB.Read("")
	if err != nil {
		pKeyReponse.Message = err.Error()
		if err := rpc.RespondWithJSON(w, http.StatusInternalServerError, pKeyReponse); err != nil {
			s.log.Error(err)
		}
		return
	}

	var pubKeyList []*PubKey
	for keyID, pK := range keyMap {
		pubKeyList = append(pubKeyList, &PubKey{KeyID: keyID, PublicKey: fmt.Sprintf("%x", pK)})
	}
	pKeyReponse.PubKeys = pubKeyList

	if err := rpc.RespondWithJSON(w, http.StatusOK, pKeyReponse); err != nil {
		s.log.Error(err)
	}
}
