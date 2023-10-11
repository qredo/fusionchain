package service

import (
	"context"

	"github.com/qredo/fusionchain/go-client"
	"github.com/qredo/fusionchain/x/treasury/types"
)

type QueryClient interface {
	PendingKeyRequests(ctx context.Context, page *client.PageRequest, keyringID uint64) ([]*types.KeyRequest, error)
	PendingSignatureRequests(ctx context.Context, page *client.PageRequest, keyringID uint64) ([]*types.SignRequest, error)
}

type TxClient interface {
	FulfilKeyRequest(ctx context.Context, requestID uint64, publicKey []byte) error

	FulfilSignatureRequest(ctx context.Context, requestID uint64, publicKey []byte) error
	RejectSignatureRequest(ctx context.Context, requestID uint64, reason string) error
}
