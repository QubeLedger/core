package types_test

import (
	"testing"

	apptesting "github.com/QuadrateOrg/core/app/apptesting"
	. "github.com/QuadrateOrg/core/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgDexDeposit_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDexDeposit
		err  error
	}{
		{
			name: "invalid creator",
			msg: MsgDexDeposit{
				Creator:         "invalid_address",
				Receiver:        apptesting.CreateRandomAccounts(1)[0].String(),
				Fees:            []uint64{0},
				TickIndexesAToB: []int64{0},
				AmountsA:        []sdk.Int{sdk.OneInt()},
				AmountsB:        []sdk.Int{sdk.OneInt()},
				Options:         []*DepositOptions{{false}},
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid receiver",
			msg: MsgDexDeposit{
				Creator:         apptesting.CreateRandomAccounts(1)[0].String(),
				Receiver:        "invalid address",
				Fees:            []uint64{0},
				TickIndexesAToB: []int64{0},
				AmountsA:        []sdk.Int{sdk.OneInt()},
				AmountsB:        []sdk.Int{sdk.OneInt()},
				Options:         []*DepositOptions{{false}},
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid fee indexes length",
			msg: MsgDexDeposit{
				Creator:         apptesting.CreateRandomAccounts(1)[0].String(),
				Receiver:        apptesting.CreateRandomAccounts(1)[0].String(),
				Fees:            []uint64{0},
				TickIndexesAToB: []int64{},
				AmountsA:        []sdk.Int{},
				AmountsB:        []sdk.Int{},
				Options:         []*DepositOptions{{false}},
			},
			err: ErrUnbalancedTxArray,
		},
		{
			name: "invalid tick indexes length",
			msg: MsgDexDeposit{
				Creator:         apptesting.CreateRandomAccounts(1)[0].String(),
				Receiver:        apptesting.CreateRandomAccounts(1)[0].String(),
				Fees:            []uint64{},
				TickIndexesAToB: []int64{0},
				AmountsA:        []sdk.Int{},
				AmountsB:        []sdk.Int{},
			},
			err: ErrUnbalancedTxArray,
		},
		{
			name: "invalid amounts A length",
			msg: MsgDexDeposit{
				Creator:         apptesting.CreateRandomAccounts(1)[0].String(),
				Receiver:        apptesting.CreateRandomAccounts(1)[0].String(),
				Fees:            []uint64{},
				TickIndexesAToB: []int64{},
				AmountsA:        []sdk.Int{sdk.OneInt()},
				AmountsB:        []sdk.Int{},
			},
			err: ErrUnbalancedTxArray,
		},
		{
			name: "invalid amounts B length",
			msg: MsgDexDeposit{
				Creator:         apptesting.CreateRandomAccounts(1)[0].String(),
				Receiver:        apptesting.CreateRandomAccounts(1)[0].String(),
				Fees:            []uint64{},
				TickIndexesAToB: []int64{},
				AmountsA:        []sdk.Int{},
				AmountsB:        []sdk.Int{sdk.OneInt()},
			},
			err: ErrUnbalancedTxArray,
		},
		{
			name: "invalid no deposit",
			msg: MsgDexDeposit{
				Creator:         apptesting.CreateRandomAccounts(1)[0].String(),
				Receiver:        apptesting.CreateRandomAccounts(1)[0].String(),
				Fees:            []uint64{},
				TickIndexesAToB: []int64{},
				AmountsA:        []sdk.Int{},
				AmountsB:        []sdk.Int{},
				Options:         []*DepositOptions{},
			},
			err: ErrZeroDeposit,
		},
		{
			name: "invalid no deposit",
			msg: MsgDexDeposit{
				Creator:         apptesting.CreateRandomAccounts(1)[0].String(),
				Receiver:        apptesting.CreateRandomAccounts(1)[0].String(),
				Fees:            []uint64{},
				TickIndexesAToB: []int64{},
				AmountsA:        []sdk.Int{},
				AmountsB:        []sdk.Int{},
				Options:         []*DepositOptions{{false}, {false}, {false}},
			},
			err: ErrUnbalancedTxArray,
		},
		{
			name: "invalid no deposit",
			msg: MsgDexDeposit{
				Creator:         apptesting.CreateRandomAccounts(1)[0].String(),
				Receiver:        apptesting.CreateRandomAccounts(1)[0].String(),
				Fees:            []uint64{0},
				TickIndexesAToB: []int64{0},
				AmountsA:        []sdk.Int{sdk.ZeroInt()},
				AmountsB:        []sdk.Int{sdk.ZeroInt()},
				Options:         []*DepositOptions{{false}},
			},
			err: ErrZeroDeposit,
		},
		{
			name: "valid msg",
			msg: MsgDexDeposit{
				Creator:         apptesting.CreateRandomAccounts(1)[0].String(),
				Receiver:        apptesting.CreateRandomAccounts(1)[0].String(),
				Fees:            []uint64{0},
				TickIndexesAToB: []int64{0},
				AmountsA:        []sdk.Int{sdk.OneInt()},
				AmountsB:        []sdk.Int{sdk.OneInt()},
				Options:         []*DepositOptions{{false}},
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
