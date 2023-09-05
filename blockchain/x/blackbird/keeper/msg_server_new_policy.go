package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/qredo/fusionchain/x/blackbird/types"
)

func (k msgServer) NewPolicy(goCtx context.Context, msg *types.MsgNewPolicy) (*types.MsgNewPolicyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgNewPolicyResponse{}, nil
}
