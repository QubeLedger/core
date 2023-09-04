package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateLend = "create_lend"
const TypeMsgDeleteLend = "delete_lend"
const TypeMsgDeposit = "deposit"
const TypeMsgWithdrawal = "withdrawal"
const TypeMsgDepositCollateral = "deposit_collateral"
const TypeMsgWithdrawalCollateral = "withdrawal_collateral"
const TypeMsgCreateLiquidationPosition = "create_liquidation_position"
const TypeMsgCloseLiquidationPosition = "close_liquidation_position"

var _ sdk.Msg = &MsgCreateLend{}
var _ sdk.Msg = &MsgDeleteLend{}
var _ sdk.Msg = &MsgDeposit{}
var _ sdk.Msg = &MsgWithdrawal{}
var _ sdk.Msg = &MsgCreateLiquidationPosition{}
var _ sdk.Msg = &MsgCloseLiquidationPosition{}
var _ sdk.Msg = &MsgDepositCollateral{}
var _ sdk.Msg = &MsgWithdrawalCollateral{}

/*
create_lend
*/
func NewMsgCreateLend(creator string, denomIn string, desiredAmount string) *MsgCreateLend {
	return &MsgCreateLend{
		Borrower:      creator,
		DenomIn:       denomIn,
		DesiredAmount: desiredAmount,
	}
}

func (msg *MsgCreateLend) Route() string {
	return RouterKey
}

func (msg *MsgCreateLend) Type() string {
	return TypeMsgCreateLend
}

func (msg *MsgCreateLend) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Borrower)
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
	_, err := sdk.AccAddressFromBech32(msg.Borrower)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

/*
delete_lend
*/

func NewMsgDeleteLend(creator string, amount string, loadId string, denomOut string) *MsgDeleteLend {
	return &MsgDeleteLend{
		Borrower: creator,
		AmountIn: amount,
		LoanId:   loadId,
		DenomOut: denomOut,
	}
}

func (msg *MsgDeleteLend) Route() string {
	return RouterKey
}

func (msg *MsgDeleteLend) Type() string {
	return TypeMsgDeleteLend
}

func (msg *MsgDeleteLend) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Borrower)
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
	_, err := sdk.AccAddressFromBech32(msg.Borrower)
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

/*
deposit collateral
*/
func NewMsgDepositCollateral(creator string, amountIn string) *MsgDepositCollateral {
	return &MsgDepositCollateral{
		Depositor: creator,
		AmountIn:  amountIn,
	}
}

func (msg *MsgDepositCollateral) Route() string {
	return RouterKey
}

func (msg *MsgDepositCollateral) Type() string {
	return TypeMsgDepositCollateral
}

func (msg *MsgDepositCollateral) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Depositor)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDepositCollateral) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDepositCollateral) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Depositor)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid depositor address (%s)", err)
	}
	return nil
}

/*
withdrawal collateral
*/

func NewMsgWithdrawalCollateral(creator string, denom string) *MsgWithdrawalCollateral {
	return &MsgWithdrawalCollateral{
		Depositor: creator,
		Denom:     denom,
	}
}

func (msg *MsgWithdrawalCollateral) Route() string {
	return RouterKey
}

func (msg *MsgWithdrawalCollateral) Type() string {
	return TypeMsgWithdrawalCollateral
}

func (msg *MsgWithdrawalCollateral) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Depositor)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgWithdrawalCollateral) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgWithdrawalCollateral) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Depositor)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid depositor address (%s)", err)
	}
	return nil
}

/*
create liq position
*/

func NewMsgCreateLiquidationPosition(creator string, amountIn string, asset string, premium string) *MsgCreateLiquidationPosition {
	return &MsgCreateLiquidationPosition{
		Creator:  creator,
		AmountIn: amountIn,
		Asset:    asset,
		Premium:  premium,
	}
}

func (msg *MsgCreateLiquidationPosition) Route() string {
	return RouterKey
}

func (msg *MsgCreateLiquidationPosition) Type() string {
	return TypeMsgCreateLiquidationPosition
}

func (msg *MsgCreateLiquidationPosition) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateLiquidationPosition) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateLiquidationPosition) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid depositor address (%s)", err)
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
	return nil
}
