package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDexDeposit = "dex_deposit"

var _ sdk.Msg = &MsgDexDeposit{}

func NewMsgDexDeposit(
	creator,
	receiver,
	tokenA,
	tokenB string,
	amountsA,
	amountsB []sdk.Int,
	tickIndexes []int64,
	fees []uint64,
	depositOptions []*DepositOptions,
) *MsgDexDeposit {
	return &MsgDexDeposit{
		Creator:         creator,
		Receiver:        receiver,
		TokenA:          tokenA,
		TokenB:          tokenB,
		AmountsA:        amountsA,
		AmountsB:        amountsB,
		TickIndexesAToB: tickIndexes,
		Fees:            fees,
		Options:         depositOptions,
	}
}

func (msg *MsgDexDeposit) Route() string {
	return RouterKey
}

func (msg *MsgDexDeposit) Type() string {
	return TypeMsgDexDeposit
}

func (msg *MsgDexDeposit) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{creator}
}

func (msg *MsgDexDeposit) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDexDeposit) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid receiver address (%s)", err)
	}

	// Verify that the lengths of TickIndexes, Fees, AmountsA, AmountsB are all equal
	numDeposits := len(msg.AmountsA)
	if numDeposits != len(msg.Fees) ||
		numDeposits != len(msg.TickIndexesAToB) ||
		numDeposits != len(msg.AmountsB) ||
		numDeposits != len(msg.Options) {
		return ErrUnbalancedTxArray
	}
	if numDeposits == 0 {
		return ErrZeroDeposit
	}

	for i := 0; i < numDeposits; i++ {
		if msg.AmountsA[i].LT(sdk.ZeroInt()) || msg.AmountsB[i].LT(sdk.ZeroInt()) {
			return ErrZeroDeposit
		}
		if msg.AmountsA[i].Equal(sdk.ZeroInt()) && msg.AmountsB[i].Equal(sdk.ZeroInt()) {
			return ErrZeroDeposit
		}
	}

	return nil
}
