package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrSample                   = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrPairNotFound             = sdkerrors.Register(ModuleName, 1701, "ErrPairNotFound err")
	ErrAmountInGTEminAmountIn   = sdkerrors.Register(ModuleName, 1801, "ErrAmountInGTEminAmountIn err")
	ErrAmountOutGTEminAmountOut = sdkerrors.Register(ModuleName, 1802, "ErrAmountOutGTEminAmountOut err")
	ErrCalculatGrowRate         = sdkerrors.Register(ModuleName, 1500, "ErrCalculatGrowRate err")
)
