package main

import (
	"net/http"

	"github.com/qredo/fusionchain/mpc-relayer/pkg/rpc"
)

const (
	sigRequest = "/fusionchain/treasury/signature_request_by_id"
)

// Response represents the superset of Status and PubKey API responses.
type Response struct {
	ID             string `json:"id,omitempty"`
	Creator        string `json:"creator,omitempty"`
	KeyID          string `json:"key_id,omitempty"`
	DataForSigning string `json:"data_for_signing,omitempty"`
	Status         string `json:"status,omitempty"`
}

func makeAPIHandlers(s *Service) *rpc.API {
	r := &rpc.API{}
	r.AddEndpoint(rpc.NewEndpoint(sigRequest, http.MethodGet, s.signatureRequestByIDHandler))
	return r
}

func (s *Service) signatureRequestByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Extracting the query parameter 'id'
	id := r.URL.Query().Get("id")

	// Simulating data based on the received ID (replace this logic with your actual data retrieval logic)
	responseData := &Response{
		ID:             id,
		Creator:        s.config.Creator,
		KeyID:          s.config.KeyId,
		DataForSigning: s.config.DataForSigning,
		Status:         "SIGN_REQUEST_STATUS_PENDING",
	}

	if err := rpc.RespondWithJSON(w, http.StatusOK, responseData); err != nil {
		rpc.RespondWithError(w, http.StatusInternalServerError, err)
		s.log.Error(err)
		return
	}
}
