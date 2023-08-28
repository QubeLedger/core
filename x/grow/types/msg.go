package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateLend = "create_lend"
const TypeMsgDeleteLend = "delete_lend"
const TypeMsgDeposit = "deposit"
const TypeMsgWithdrawal = "withdrawal"

var _ sdk.Msg = &MsgCreateLend{}
var _ sdk.Msg = &MsgDeleteLend{}
var _ sdk.Msg = &MsgDeposit{}
var _ sdk.Msg = &MsgWithdrawal{}

/*
create_lend
*/
func NewMsgCreateLend(creator string, amount string, pairId string) *MsgCreateLend {
	return &MsgCreateLend{
		Creator: creator,
		Amount:  amount,
		PairId:  pairId,
	}
}

func (msg *MsgCreateLend) Route() string {
	return RouterKey
}

func (msg *MsgCreateLend) Type() string {
	return TypeMsgCreateLend
}

func (msg *MsgCreateLend) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateLend) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateLend) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

/*

delete_lend

*/

func NewMsgDeleteLend(creator string, amount string, pairId string) *MsgDeleteLend {
	return &MsgDeleteLend{
		Creator: creator,
		Amount:  amount,
		PairId:  pairId,
	}
}

func (msg *MsgDeleteLend) Route() string {
	return RouterKey
}

func (msg *MsgDeleteLend) Type() string {
	return TypeMsgDeleteLend
}

func (msg *MsgDeleteLend) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteLend) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteLend) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

/*
deposit
*/

func NewMsgDeposit(creator string, amountIn string, denomOut string) *MsgDeposit {
	return &MsgDeposit{
		Creator:  creator,
		AmountIn: amountIn,
		DenomOut: denomOut,
	}
}

func (msg *MsgDeposit) Route() string {
	return RouterKey
}

func (msg *MsgDeposit) Type() string {
	return TypeMsgDeposit
}

func (msg *MsgDeposit) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeposit) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeposit) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

/*
withdrawal
*/

func NewMsgWithdrawal(creator string, amountIn string, denomOut string) *MsgWithdrawal {
	return &MsgWithdrawal{
		Creator:  creator,
		AmountIn: amountIn,
		DenomOut: denomOut,
	}
}

func (msg *MsgWithdrawal) Route() string {
	return RouterKey
}

func (msg *MsgWithdrawal) Type() string {
	return TypeMsgWithdrawal
}

func (msg *MsgWithdrawal) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgWithdrawal) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgWithdrawal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
