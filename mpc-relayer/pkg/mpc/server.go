package mpc

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type LocalMPCServer struct {
	initVersion int
	keyType     CryptoSystem
}

func NewLocalMPCServer(initVersion int) LocalMPCServer {
	return LocalMPCServer{
		initVersion: initVersion,
	}
}

func (service LocalMPCServer) Routes() http.Handler {
	router := mux.NewRouter()
	router.Handle(Status, http.HandlerFunc(service.Check)).Methods(http.MethodGet)
	router.Handle(ECDSAKeys, http.HandlerFunc(service.KeysHandler)).Methods(http.MethodPost)
	router.Handle(ECDSASig, http.HandlerFunc(service.Sign)).Methods(http.MethodPost)
	router.Handle(EdDSAKeys, http.HandlerFunc(service.KeysHandler)).Methods(http.MethodPost)
	router.Handle(EdDSASig, http.HandlerFunc(service.Sign)).Methods(http.MethodPost)

	return router
}

func (service LocalMPCServer) KeysHandler(w http.ResponseWriter, req *http.Request) {
	request := KeysRequest{}
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	response, err := localMPCKeys(&request, service.initVersion, service.keyType)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	b, _ := json.Marshal(response)
	_, _ = w.Write(b)
}

func (service LocalMPCServer) Sign(w http.ResponseWriter, req *http.Request) {
	request := SigRequest{}
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	response, err := localMPCSign(&request, service.initVersion, service.keyType)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	b, _ := json.Marshal(response)
	_, _ = w.Write(b)
}

func (service LocalMPCServer) Check(w http.ResponseWriter, _ *http.Request) {
	response := KeysResponse{
		Message: "OK",
	}
	b, _ := json.Marshal(response)
	_, _ = w.Write(b)
}
