package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSendShares = "send_shares"

var _ sdk.Msg = &MsgSendShares{}

func NewMsgSendShares(creator string, shares string, denomA string, denomB string, address string) *MsgSendShares {
	return &MsgSendShares{
		Creator: creator,
		Shares:  shares,
		DenomA:  denomA,
		DenomB:  denomB,
		Address: address,
	}
}

func (msg *MsgSendShares) Route() string {
	return RouterKey
}

func (msg *MsgSendShares) Type() string {
	return TypeMsgSendShares
}

func (msg *MsgSendShares) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendShares) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendShares) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
