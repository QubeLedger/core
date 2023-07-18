package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgMintUsq = "mint_usq"

var _ sdk.Msg = &MsgMintUsq{}

func NewMsgMintUsq(creator string, amount string) *MsgMintUsq {
	return &MsgMintUsq{
		Creator: creator,
		Amount:  amount,
	}
}

func (msg *MsgMintUsq) Route() string {
	return RouterKey
}

func (msg *MsgMintUsq) Type() string {
	return TypeMsgMintUsq
}

func (msg *MsgMintUsq) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMintUsq) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMintUsq) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

const TypeMsgBurnUsq = "burn_usq"

var _ sdk.Msg = &MsgBurnUsq{}

func NewMsgBurnUsq(creator string, amount string) *MsgBurnUsq {
	return &MsgBurnUsq{
		Creator: creator,
		Amount:  amount,
	}
}

func (msg *MsgBurnUsq) Route() string {
	return RouterKey
}

func (msg *MsgBurnUsq) Type() string {
	return TypeMsgBurnUsq
}

func (msg *MsgBurnUsq) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBurnUsq) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBurnUsq) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
