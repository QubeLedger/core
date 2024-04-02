package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgPerpetualDeposit = "perpetual_deposit"
const TypeMsgPerpetualWithdraw = "perpetual_withdraw"
const TypeMsgCreatePosition = "create_position"
const TypeMsgClosePosition = "close_position"

var _ sdk.Msg = &MsgPerpetualDeposit{}

func NewMsgPerpetualDeposit(trader string, amountIn string) *MsgPerpetualDeposit {
	return &MsgPerpetualDeposit{
		Trader:   trader,
		AmountIn: amountIn,
	}
}

func (msg *MsgPerpetualDeposit) Route() string {
	return RouterKey
}

func (msg *MsgPerpetualDeposit) Type() string {
	return TypeMsgPerpetualDeposit
}

func (msg *MsgPerpetualDeposit) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Trader)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgPerpetualDeposit) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgPerpetualDeposit) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Trader)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgPerpetualWithdraw{}

func NewMsgPerpetualWithdraw(trader string, deposit_id string) *MsgPerpetualWithdraw {
	return &MsgPerpetualWithdraw{
		Trader:    trader,
		DepositId: deposit_id,
	}
}

func (msg *MsgPerpetualWithdraw) Route() string {
	return RouterKey
}

func (msg *MsgPerpetualWithdraw) Type() string {
	return TypeMsgPerpetualWithdraw
}

func (msg *MsgPerpetualWithdraw) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Trader)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgPerpetualWithdraw) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgPerpetualWithdraw) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Trader)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.DepositId == "" {
		return sdkerrors.Wrapf(ErrDepositIdEmpty, "deposit id empty (%s)", ErrDepositIdEmpty)
	}
	return nil
}

var _ sdk.Msg = &MsgCreatePosition{}

func NewMsgCreatePosition(trader string, amountIn string, leverage uint64, tradeType PerpetualTradeType) *MsgCreatePosition {
	return &MsgCreatePosition{
		Trader:    trader,
		AmountIn:  amountIn,
		Leverage:  leverage,
		TradeType: tradeType,
	}
}

func (msg *MsgCreatePosition) Route() string {
	return RouterKey
}

func (msg *MsgCreatePosition) Type() string {
	return TypeMsgCreatePosition
}

func (msg *MsgCreatePosition) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Trader)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreatePosition) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreatePosition) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Trader)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgClosePosition{}

func NewMsgClosePosition(trader string, pos_id string) *MsgClosePosition {
	return &MsgClosePosition{
		Trader:     trader,
		PositionId: pos_id,
	}
}

func (msg *MsgClosePosition) Route() string {
	return RouterKey
}

func (msg *MsgClosePosition) Type() string {
	return TypeMsgClosePosition
}

func (msg *MsgClosePosition) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Trader)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgClosePosition) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgClosePosition) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Trader)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
