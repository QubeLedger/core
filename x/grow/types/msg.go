package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateLend = "create_lend"
const TypeMsgDeleteLend = "delete_lend"
const TypeMsgGrowDeposit = "grow_deposit"
const TypeMsgGrowWithdrawal = "grow_withdrawal"
const TypeMsgDepositCollateral = "deposit_collateral"
const TypeMsgWithdrawalCollateral = "withdrawal_collateral"
const TypeMsgOpenLiquidationPosition = "create_liquidation_position"
const TypeMsgCloseLiquidationPosition = "close_liquidation_position"

var _ sdk.Msg = &MsgCreateLend{}
var _ sdk.Msg = &MsgDeleteLend{}
var _ sdk.Msg = &MsgGrowDeposit{}
var _ sdk.Msg = &MsgGrowWithdrawal{}
var _ sdk.Msg = &MsgOpenLiquidationPosition{}
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
