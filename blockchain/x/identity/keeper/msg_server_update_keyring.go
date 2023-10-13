package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/qredo/fusionchain/x/identity/types"
)

func (k msgServer) UpdateKeyring(goCtx context.Context, msg *types.MsgUpdateKeyring) (*types.MsgUpdateKeyringResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	kr, found := k.KeyringsRepo().Get(ctx, msg.Id)
	if !found {
		return nil, fmt.Errorf("keyring not found")
	}
	// No Policy is defined for keyring, we have just to check if the request is realized by one of the admins.
	isAdmin := false
	for _, admin := range kr.Admins {
		if msg.Creator == admin {
			isAdmin = true
			break
		}
	}
	if !isAdmin {
		return nil, fmt.Errorf("keyring updates should be request by admins")
	}
	kr.SetStatus(msg.IsActive)
	kr.SetDescription(msg.Description) //mmmmmmmmmmmmmmm
	k.KeyringsRepo().Set(ctx, kr)
	return &types.MsgUpdateKeyringResponse{}, nil
}
