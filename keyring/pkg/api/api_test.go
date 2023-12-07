package api

import (
	"net/http"
	"testing"
)

// MockKeyringService implements the KeyringService interface for testing purposes
type mockKeyringService struct{}

func (m *mockKeyringService) Status(w http.ResponseWriter, r *http.Request) {
}

func (m *mockKeyringService) HealthCheck(w http.ResponseWriter, r *http.Request) {
}

func (m *mockKeyringService) Keyring(w http.ResponseWriter, r *http.Request) {
}

func (m *mockKeyringService) PubKeys(w http.ResponseWriter, r *http.Request) {
}

func (m *mockKeyringService) Mnemonic(w http.ResponseWriter, r *http.Request) {
}

func Test_MakeKeyRingAPI(t *testing.T) {
	mockService := &mockKeyringService{}
	_ = MakeKeyRingAPI(mockService)
}
