package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgErc20Cw20 = "erc_20_cw_20"

var _ sdk.Msg = &MsgErc20Cw20{}

func NewMsgErc20Cw20(creator string, amount string, token string) *MsgErc20Cw20 {
	return &MsgErc20Cw20{
		Creator: creator,
		Amount:  amount,
		Token:   token,
	}
}

func (msg *MsgErc20Cw20) Route() string {
	return RouterKey
}

func (msg *MsgErc20Cw20) Type() string {
	return TypeMsgErc20Cw20
}

func (msg *MsgErc20Cw20) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgErc20Cw20) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgErc20Cw20) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
