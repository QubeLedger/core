package types_test

import (
	"testing"

	apptesting "github.com/QuadrateOrg/core/app/apptesting"
	. "github.com/QuadrateOrg/core/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgMultiHopSwap_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgMultiHopSwap
		err  error
	}{
		{
			name: "invalid creator address",
			msg: MsgMultiHopSwap{
				Creator:  "invalid_address",
				Receiver: apptesting.CreateRandomAccounts(1)[0].String(),
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid receiver address",
			msg: MsgMultiHopSwap{
				Creator:  apptesting.CreateRandomAccounts(1)[0].String(),
				Receiver: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "missing route",
			msg: MsgMultiHopSwap{
				Creator:  apptesting.CreateRandomAccounts(1)[0].String(),
				Receiver: apptesting.CreateRandomAccounts(1)[0].String(),
			},
			err: ErrMissingMultihopRoute,
		},
		{
			name: "invalid exit tokens",
			msg: MsgMultiHopSwap{
				Creator:  apptesting.CreateRandomAccounts(1)[0].String(),
				Receiver: apptesting.CreateRandomAccounts(1)[0].String(),
				Routes: []*MultiHopRoute{
					{Hops: []string{"A", "B", "C"}},
					{Hops: []string{"A", "B", "Z"}},
				},
			},
			err: ErrMultihopExitTokensMismatch,
		},
		{
			name: "invalid amountIn",
			msg: MsgMultiHopSwap{
				Creator:  apptesting.CreateRandomAccounts(1)[0].String(),
				Receiver: apptesting.CreateRandomAccounts(1)[0].String(),
				Routes:   []*MultiHopRoute{{Hops: []string{"A", "B", "C"}}},
				AmountIn: sdk.NewInt(-1),
			},
			err: ErrZeroSwap,
		},
		{
			name: "valid",
			msg: MsgMultiHopSwap{
				Routes:   []*MultiHopRoute{{Hops: []string{"A", "B", "C"}}},
				Creator:  apptesting.CreateRandomAccounts(1)[0].String(),
				Receiver: apptesting.CreateRandomAccounts(1)[0].String(),
				AmountIn: sdk.OneInt(),
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
