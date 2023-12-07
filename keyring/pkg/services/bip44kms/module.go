package kms

import "github.com/qredo/fusionchain/keyring/pkg/api"

// Module represents a simple interface for sub-processes within
// a service.
type Module interface {
	Start() error
	Stop() error
	healthcheck() *api.HealthResponse
}
