package types_test

import (
	"testing"

	"github.com/QuadrateOrg/core/app/apptesting"
	"github.com/QuadrateOrg/core/x/grow/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateBorrow_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgCreateBorrow
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgCreateBorrow{
				Borrower: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: types.MsgCreateBorrow{
				Borrower: apptesting.CreateRandomAccounts(1)[0].String(),
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

func TestMsgDeleteBorrow_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgDeleteBorrow
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgDeleteBorrow{
				Borrower: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: types.MsgDeleteBorrow{
				Borrower: apptesting.CreateRandomAccounts(1)[0].String(),
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
		msg  types.MsgGrowDeposit
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgGrowDeposit{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: types.MsgGrowDeposit{
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
		msg  types.MsgGrowWithdrawal
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgGrowWithdrawal{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: types.MsgGrowWithdrawal{
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

func TestMsgCreateLend_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgCreateLend
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgCreateLend{
				Depositor: "invalid_address",
				AmountIn:  "100uosmo",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid",
			msg: types.MsgCreateLend{
				Depositor: apptesting.CreateRandomAccounts(1)[0].String(),
				AmountIn:  "100uosmo",
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

func TestMsgWithdrawalLend_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgWithdrawalLend
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgWithdrawalLend{
				Depositor: "invalid_address",
				DenomOut:  "uosmo",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid",
			msg: types.MsgWithdrawalLend{
				Depositor: apptesting.CreateRandomAccounts(1)[0].String(),
				DenomOut:  "uosmo",
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

func TestMsgOpenLiquidationPosition_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgOpenLiquidationPosition
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgOpenLiquidationPosition{
				Creator:  "invalid_address",
				AmountIn: "20uosmo",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "valid",
			msg: types.MsgOpenLiquidationPosition{
				Creator:  apptesting.CreateRandomAccounts(1)[0].String(),
				AmountIn: "20uosmo",
				Asset:    "OSMO",
				Premium:  "3",
			},
		},
		{
			name: "invalid amountIn",
			msg: types.MsgOpenLiquidationPosition{
				Creator:  apptesting.CreateRandomAccounts(1)[0].String(),
				AmountIn: "",
				Asset:    "OSMO",
				Premium:  "3",
			},
			err: sdkerrors.ErrInvalidCoins,
		},
		{
			name: "invalid asset",
			msg: types.MsgOpenLiquidationPosition{
				Creator:  apptesting.CreateRandomAccounts(1)[0].String(),
				AmountIn: "20uosmo",
				Asset:    "",
				Premium:  "3",
			},
			err: types.ErrInvalidLength,
		},
		{
			name: "invalid Premium",
			msg: types.MsgOpenLiquidationPosition{
				Creator:  apptesting.CreateRandomAccounts(1)[0].String(),
				AmountIn: "20uosmo",
				Asset:    "OSMO",
				Premium:  "",
			},
			err: types.ErrInvalidLength,
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

func TestMsgCloseLiquidationPosition_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgCloseLiquidationPosition
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgCloseLiquidationPosition{
				Creator:              "invalid_address",
				LiquidatorPositionId: "testid",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "valid address",
			msg: types.MsgCloseLiquidationPosition{
				Creator:              apptesting.CreateRandomAccounts(1)[0].String(),
				LiquidatorPositionId: "testid",
			},
		},
		{
			name: "invalid LiquidatorPositionId",
			msg: types.MsgCloseLiquidationPosition{
				Creator:              apptesting.CreateRandomAccounts(1)[0].String(),
				LiquidatorPositionId: "",
			},
			err: types.ErrInvalidLength,
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
