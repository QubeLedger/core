package types_test

import (
	"testing"

	"github.com/QubeLedger/core/x/stable/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestMsgMintUsq(t *testing.T) {
	addrs := []sdk.AccAddress{
		sdk.AccAddress([]byte("addr1_______________")),
	}
	tests := []struct {
		addr        sdk.AccAddress
		amountInt   sdk.Coin
		denomOut    string
		expectedErr string
	}{
		{
			addrs[0],
			sdk.NewCoin("uatom", sdk.OneInt()),
			"uusd",
			"",
		},
		{
			sdk.AccAddress{},
			sdk.NewCoin("uatom", sdk.OneInt()),
			"uusd",
			"invalid creator address (empty address string is not allowed): invalid address",
		},
	}

	for _, tc := range tests {
		msg := types.NewMsgMint(tc.addr.String(), tc.amountInt.String(), tc.denomOut)
		if tc.expectedErr == "" {
			require.Nil(t, msg.ValidateBasic())
		} else {
			require.EqualError(t, msg.ValidateBasic(), tc.expectedErr)
		}
	}
}

func TestMsgBurnUsq(t *testing.T) {
	addrs := []sdk.AccAddress{
		sdk.AccAddress([]byte("addr1_______________")),
	}
	tests := []struct {
		addr        sdk.AccAddress
		amount      sdk.Coin
		denomOut    string
		expectedErr string
	}{
		{
			addrs[0],
			sdk.NewCoin("uusd", sdk.OneInt()),
			"uatom",
			"",
		},
		{
			sdk.AccAddress{},
			sdk.NewCoin("uusd", sdk.OneInt()),
			"uatom",
			"invalid creator address (empty address string is not allowed): invalid address",
		},
	}

	for _, tc := range tests {
		msg := types.NewMsgBurn(tc.addr.String(), tc.amount.String(), tc.denomOut)
		if tc.expectedErr == "" {
			require.Nil(t, msg.ValidateBasic())
		} else {
			require.EqualError(t, msg.ValidateBasic(), tc.expectedErr)
		}
	}
}
