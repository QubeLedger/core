package types_test

import (
	"testing"

	"github.com/QuadrateOrg/core/app/apptesting"
	"github.com/QuadrateOrg/core/x/grow/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateLend_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgCreateLend
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgCreateLend{
				Borrower: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: types.MsgCreateLend{
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

func TestMsgDeleteLend_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgDeleteLend
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgDeleteLend{
				Borrower: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: types.MsgDeleteLend{
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
		msg  types.MsgDeposit
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgDeposit{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: types.MsgDeposit{
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
		msg  types.MsgWithdrawal
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgWithdrawal{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: types.MsgWithdrawal{
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

func TestMsgDepositCollateral_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgDepositCollateral
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgDepositCollateral{
				Depositor: "invalid_address",
				AmountIn:  "100uosmo",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid",
			msg: types.MsgDepositCollateral{
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

func TestMsgWithdrawalCollateral_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgWithdrawalCollateral
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgWithdrawalCollateral{
				Depositor: "invalid_address",
				Denom:     "uosmo",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid",
			msg: types.MsgWithdrawalCollateral{
				Depositor: apptesting.CreateRandomAccounts(1)[0].String(),
				Denom:     "uosmo",
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
