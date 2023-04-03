package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCw20Erc20 = "cw_20_erc_20"

var _ sdk.Msg = &MsgCw20Erc20{}

func NewMsgCw20Erc20(creator string, amount string, token string) *MsgCw20Erc20 {
	return &MsgCw20Erc20{
		Creator: creator,
		Amount:  amount,
		Token:   token,
	}
}

func (msg *MsgCw20Erc20) Route() string {
	return RouterKey
}

func (msg *MsgCw20Erc20) Type() string {
	return TypeMsgCw20Erc20
}

func (msg *MsgCw20Erc20) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCw20Erc20) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCw20Erc20) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
