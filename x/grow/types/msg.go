package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateBorrow = "create_borrow"
const TypeMsgDeleteBorrow = "delete_borrow"
const TypeMsgGrowDeposit = "grow_deposit"
const TypeMsgGrowWithdrawal = "grow_withdrawal"
const TypeMsgCreateLend = "create_lend"
const TypeMsgWithdrawalLend = "withdrawal_lend"
const TypeMsgOpenLiquidationPosition = "create_liquidation_position"
const TypeMsgCloseLiquidationPosition = "close_liquidation_position"

var _ sdk.Msg = &MsgCreateLend{}
var _ sdk.Msg = &MsgWithdrawalLend{}
var _ sdk.Msg = &MsgGrowDeposit{}
var _ sdk.Msg = &MsgGrowWithdrawal{}
var _ sdk.Msg = &MsgOpenLiquidationPosition{}
var _ sdk.Msg = &MsgCloseLiquidationPosition{}
var _ sdk.Msg = &MsgCreateBorrow{}
var _ sdk.Msg = &MsgDeleteBorrow{}

/*
create_borrow
*/
func NewMsgCreateBorrow(creator string, denomIn string, desiredAmount string) *MsgCreateBorrow {
	return &MsgCreateBorrow{
		Borrower:      creator,
		DenomIn:       denomIn,
		DesiredAmount: desiredAmount,
	}
}

func (msg *MsgCreateBorrow) Route() string {
	return RouterKey
}

func (msg *MsgCreateBorrow) Type() string {
	return TypeMsgCreateBorrow
}

func (msg *MsgCreateBorrow) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Borrower)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateBorrow) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateBorrow) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Borrower)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

/*
delete_borrow
*/

func NewMsgDeleteBorrow(creator string, denomOut string, amount string) *MsgDeleteBorrow {
	return &MsgDeleteBorrow{
		Borrower: creator,
		DenomOut: denomOut,
		AmountIn: amount,
	}
}

func (msg *MsgDeleteBorrow) Route() string {
	return RouterKey
}

func (msg *MsgDeleteBorrow) Type() string {
	return TypeMsgDeleteBorrow
}

func (msg *MsgDeleteBorrow) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Borrower)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteBorrow) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteBorrow) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Borrower)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

/*
deposit
*/

func NewMsgGrowDeposit(creator string, amountIn string, denomOut string) *MsgGrowDeposit {
	return &MsgGrowDeposit{
		Creator:  creator,
		AmountIn: amountIn,
		DenomOut: denomOut,
	}
}

func (msg *MsgGrowDeposit) Route() string {
	return RouterKey
}

func (msg *MsgGrowDeposit) Type() string {
	return TypeMsgGrowDeposit
}

func (msg *MsgGrowDeposit) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgGrowDeposit) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgGrowDeposit) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

/*
withdrawal
*/

func NewMsgGrowWithdrawal(creator string, amountIn string) *MsgGrowWithdrawal {
	return &MsgGrowWithdrawal{
		Creator:  creator,
		AmountIn: amountIn,
	}
}

func (msg *MsgGrowWithdrawal) Route() string {
	return RouterKey
}

func (msg *MsgGrowWithdrawal) Type() string {
	return TypeMsgGrowWithdrawal
}

func (msg *MsgGrowWithdrawal) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgGrowWithdrawal) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgGrowWithdrawal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

/*
deposit collateral
*/
func NewMsgCreateLend(creator string, amountIn string) *MsgCreateLend {
	return &MsgCreateLend{
		Depositor: creator,
		AmountIn:  amountIn,
	}
}

func (msg *MsgCreateLend) Route() string {
	return RouterKey
}

func (msg *MsgCreateLend) Type() string {
	return TypeMsgCreateLend
}

func (msg *MsgCreateLend) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Depositor)
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
	_, err := sdk.AccAddressFromBech32(msg.Depositor)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid depositor address (%s)", err)
	}
	return nil
}

/*
withdrawal collateral
*/

func NewMsgWithdrawalLend(creator string, amountIn string, denom string) *MsgWithdrawalLend {
	return &MsgWithdrawalLend{
		Depositor: creator,
		AmountIn:  amountIn,
		DenomOut:  denom,
	}
}

func (msg *MsgWithdrawalLend) Route() string {
	return RouterKey
}

func (msg *MsgWithdrawalLend) Type() string {
	return TypeMsgWithdrawalLend
}

func (msg *MsgWithdrawalLend) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Depositor)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgWithdrawalLend) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgWithdrawalLend) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Depositor)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid depositor address (%s)", err)
	}
	return nil
}

/*
Open liq position
*/

func NewMsgOpenLiquidationPosition(creator string, amountIn string, asset string, premium string) *MsgOpenLiquidationPosition {
	return &MsgOpenLiquidationPosition{
		Creator:  creator,
		AmountIn: amountIn,
		Asset:    asset,
		Premium:  premium,
	}
}

func (msg *MsgOpenLiquidationPosition) Route() string {
	return RouterKey
}

func (msg *MsgOpenLiquidationPosition) Type() string {
	return TypeMsgOpenLiquidationPosition
}

func (msg *MsgOpenLiquidationPosition) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgOpenLiquidationPosition) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgOpenLiquidationPosition) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid depositor address (%s)", err)
	}

	value, err := sdk.ParseCoinsNormalized(msg.AmountIn)
	if err != nil || value.String() == "" {
		return sdkerrors.ErrInvalidCoins
	}

	if len(msg.Asset) == 0 {
		return ErrInvalidLength
	}

	if len(msg.Premium) == 0 {
		return ErrInvalidLength
	}

	return nil
}

/*
delete liq position
*/

func NewMsgCloseLiquidationPosition(creator string, id string) *MsgCloseLiquidationPosition {
	return &MsgCloseLiquidationPosition{
		Creator:              creator,
		LiquidatorPositionId: id,
	}
}

func (msg *MsgCloseLiquidationPosition) Route() string {
	return RouterKey
}

func (msg *MsgCloseLiquidationPosition) Type() string {
	return TypeMsgCloseLiquidationPosition
}

func (msg *MsgCloseLiquidationPosition) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCloseLiquidationPosition) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCloseLiquidationPosition) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)

	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid depositor address (%s)", err)
	}

	if len(msg.LiquidatorPositionId) == 0 {
		return ErrInvalidLength
	}
	return nil
}
