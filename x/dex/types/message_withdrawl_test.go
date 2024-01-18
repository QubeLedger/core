package types_test

import (
	"testing"

	apptesting "github.com/QuadrateOrg/core/app/apptesting"
	. "github.com/QuadrateOrg/core/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgDexWithdrawal_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDexWithdrawal
		err  error
	}{
		{
			name: "invalid creator",
			msg: MsgDexWithdrawal{
				Creator:         "invalid_address",
				Receiver:        apptesting.CreateRandomAccounts(1)[0].String(),
				Fees:            []uint64{0},
				TickIndexesAToB: []int64{0},
				SharesToRemove:  []sdk.Int{sdk.OneInt()},
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid receiver",
			msg: MsgDexWithdrawal{
				Creator:         apptesting.CreateRandomAccounts(1)[0].String(),
				Receiver:        "invalid_address",
				Fees:            []uint64{0},
				TickIndexesAToB: []int64{0},
				SharesToRemove:  []sdk.Int{sdk.OneInt()},
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid fee indexes length",
			msg: MsgDexWithdrawal{
				Creator:         apptesting.CreateRandomAccounts(1)[0].String(),
				Receiver:        apptesting.CreateRandomAccounts(1)[0].String(),
				Fees:            []uint64{},
				TickIndexesAToB: []int64{0},
				SharesToRemove:  []sdk.Int{sdk.OneInt()},
			},
			err: ErrUnbalancedTxArray,
		},
		{
			name: "invalid tick indexes length",
			msg: MsgDexWithdrawal{
				Creator:         apptesting.CreateRandomAccounts(1)[0].String(),
				Receiver:        apptesting.CreateRandomAccounts(1)[0].String(),
				Fees:            []uint64{0},
				TickIndexesAToB: []int64{},
				SharesToRemove:  []sdk.Int{sdk.OneInt()},
			},
			err: ErrUnbalancedTxArray,
		},
		{
			name: "invalid shares to remove length",
			msg: MsgDexWithdrawal{
				Creator:         apptesting.CreateRandomAccounts(1)[0].String(),
				Receiver:        apptesting.CreateRandomAccounts(1)[0].String(),
				Fees:            []uint64{0},
				TickIndexesAToB: []int64{0},
				SharesToRemove:  []sdk.Int{},
			},
			err: ErrUnbalancedTxArray,
		},
		{
			name: "no withdraw specs",
			msg: MsgDexWithdrawal{
				Creator:         apptesting.CreateRandomAccounts(1)[0].String(),
				Receiver:        apptesting.CreateRandomAccounts(1)[0].String(),
				Fees:            []uint64{},
				TickIndexesAToB: []int64{},
				SharesToRemove:  []sdk.Int{},
			},
			err: ErrZeroWithdraw,
		},
		{
			name: "no withdraw specs",
			msg: MsgDexWithdrawal{
				Creator:         apptesting.CreateRandomAccounts(1)[0].String(),
				Receiver:        apptesting.CreateRandomAccounts(1)[0].String(),
				Fees:            []uint64{0},
				TickIndexesAToB: []int64{0},
				SharesToRemove:  []sdk.Int{sdk.ZeroInt()},
			},
			err: ErrZeroWithdraw,
		},
		{
			name: "valid msg",
			msg: MsgDexWithdrawal{
				Creator:         apptesting.CreateRandomAccounts(1)[0].String(),
				Receiver:        apptesting.CreateRandomAccounts(1)[0].String(),
				Fees:            []uint64{0},
				TickIndexesAToB: []int64{0},
				SharesToRemove:  []sdk.Int{sdk.OneInt()},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
