package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/qredo/fusionchain/policy"
	"github.com/qredo/fusionchain/x/blackbird/types"
)

func (k msgServer) NewPolicy(goCtx context.Context, msg *types.MsgNewPolicy) (*types.MsgNewPolicyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var policy policy.Policy
	if err := k.cdc.UnpackAny(msg.Policy, policy); err != nil {
		return nil, err
	}
	if err := policy.Validate(); err != nil {
		return nil, err
	}

	p := &types.Policy{
		Name:   msg.Name,
		Policy: msg.Policy,
	}
	id := k.PolicyRepo().Append(ctx, p)

	return &types.MsgNewPolicyResponse{
		Id: id,
	}, nil
}
