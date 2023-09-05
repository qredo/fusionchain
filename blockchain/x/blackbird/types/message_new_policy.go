package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgNewPolicy = "new_policy"

var _ sdk.Msg = &MsgNewPolicy{}

func NewMsgNewPolicy(creator string) *MsgNewPolicy {
	return &MsgNewPolicy{
		Creator: creator,
	}
}

func (msg *MsgNewPolicy) Route() string {
	return RouterKey
}

func (msg *MsgNewPolicy) Type() string {
	return TypeMsgNewPolicy
}

func (msg *MsgNewPolicy) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgNewPolicy) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgNewPolicy) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
