package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/qredo/fusionchain/x/blackbird/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// nolint:stylecheck,st1003
// revive:disable-next-line var-naming
func (k Keeper) PolicyById(goCtx context.Context, req *types.QueryPolicyByIdRequest) (*types.QueryPolicyByIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryPolicyByIdResponse{}, nil
}
