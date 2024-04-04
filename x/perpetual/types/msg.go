package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgOpen = "perpetual_open"
const TypeMsgClose = "perpetual_close"

var _ sdk.Msg = &MsgOpen{}

func NewMsgOpen(Creator string, TradeType PerpetualTradeType, Leverage sdk.Dec, TradingAsset string, Collateral string) *MsgOpen {
	return &MsgOpen{
		Creator:      Creator,
		TradeType:    TradeType,
		Leverage:     Leverage,
		TradingAsset: TradingAsset,
		Collateral:   Collateral,
	}
}

func (msg *MsgOpen) Route() string {
	return RouterKey
}

func (msg *MsgOpen) Type() string {
	return TypeMsgOpen
}

func (msg *MsgOpen) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgOpen) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgOpen) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	_, err = sdk.ParseCoinsNormalized(msg.Collateral)
	if err != nil {
		return err
	}

	if msg.TradingAsset == "" {
		return sdkerrors.Wrapf(ErrTradingAssetEmpty, "TradingAsset emprty (%s)", err)
	}

	return nil
}

var _ sdk.Msg = &MsgClose{}

func NewMsgClose(Creator string, id uint64, amount sdk.Int) *MsgClose {
	return &MsgClose{
		Creator: Creator,
		Id:      id,
		Amount:  amount,
	}
}

func (msg *MsgClose) Route() string {
	return RouterKey
}

func (msg *MsgClose) Type() string {
	return TypeMsgClose
}

func (msg *MsgClose) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgClose) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgClose) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}
