package types

import (
	"testing"

	"github.com/QuadrateOrg/core/app/apptesting"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateLend_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateLend
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateLend{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateLend{
				Creator: apptesting.CreateRandomAccounts(1)[0].String(),
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

func TestMsgDeleteLend_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteLend
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteLend{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteLend{
				Creator: apptesting.CreateRandomAccounts(1)[0].String(),
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

func TestMsgDeposit_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeposit
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeposit{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeposit{
				Creator: apptesting.CreateRandomAccounts(1)[0].String(),
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

func TestMsgWithdrawal_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgWithdrawal
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgWithdrawal{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgWithdrawal{
				Creator: apptesting.CreateRandomAccounts(1)[0].String(),
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
