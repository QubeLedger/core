package types_test

import (
	"testing"

	"github.com/QuadrateOrg/core/x/perpetual/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestMsgPerpetualDeposit(t *testing.T) {
	addrs := []sdk.AccAddress{
		sdk.AccAddress([]byte("addr1_______________")),
	}
	tests := []struct {
		addr        sdk.AccAddress
		amountInt   sdk.Coin
		expectedErr string
	}{
		{
			addrs[0],
			sdk.NewCoin("uatom", sdk.OneInt()),
			"",
		},
		{
			sdk.AccAddress{},
			sdk.NewCoin("uatom", sdk.OneInt()),
			"invalid creator address (empty address string is not allowed): invalid address",
		},
	}

	for _, tc := range tests {
		msg := types.NewMsgPerpetualDeposit(tc.addr.String(), tc.amountInt.String())
		if tc.expectedErr == "" {
			require.Nil(t, msg.ValidateBasic())
		} else {
			require.EqualError(t, msg.ValidateBasic(), tc.expectedErr)
		}
	}
}

func TestMsgPerpetualWithdraw(t *testing.T) {
	addrs := []sdk.AccAddress{
		sdk.AccAddress([]byte("addr1_______________")),
	}
	tests := []struct {
		addr        sdk.AccAddress
		deposit_id  string
		expectedErr string
	}{
		{
			addrs[0],
			"test",
			"",
		},
		{
			sdk.AccAddress{},
			"test",
			"invalid creator address (empty address string is not allowed): invalid address",
		},
		{
			addrs[0],
			"",
			"deposit id empty (ErrDepositIdEmpty err): ErrDepositIdEmpty err",
		},
	}

	for _, tc := range tests {
		msg := types.NewMsgPerpetualWithdraw(tc.addr.String(), tc.deposit_id)
		if tc.expectedErr == "" {
			require.Nil(t, msg.ValidateBasic())
		} else {
			require.EqualError(t, msg.ValidateBasic(), tc.expectedErr)
		}
	}
}
