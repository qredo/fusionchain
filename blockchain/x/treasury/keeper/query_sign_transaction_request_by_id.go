package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/qredo/fusionchain/x/treasury/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) SignTransactionRequestById(goCtx context.Context, req *types.QuerySignTransactionRequestByIdRequest) (*types.QuerySignTransactionRequestByIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QuerySignTransactionRequestByIdResponse{}, nil
}
