package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrSample                        = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrSdkIntError                   = sdkerrors.Register(ModuleName, 1101, "sdk.Int error")
	ErrIntNegativeOrZero             = sdkerrors.Register(ModuleName, 1102, "ErrIntNegativeOrZero error")
	ErrOracleAssetIdNotFound         = sdkerrors.Register(ModuleName, 1200, "ErrOracleAssetIdNotFound err")
	ErrCoinsLen                      = sdkerrors.Register(ModuleName, 1400, "ErrCoinsLen err")
	ErrDenomsNotEqual                = sdkerrors.Register(ModuleName, 1401, "ErrDenomsNotEqual err")
	ErrLiquidatorAddresesNotEqual    = sdkerrors.Register(ModuleName, 1402, "ErrLiquidatorAddresesNotEqual err")
	ErrLiquidatorPositionIdNotEqual  = sdkerrors.Register(ModuleName, 1403, "ErrLiquidatorPositionIdNotEqual err")
	ErrCalculatGrowRate              = sdkerrors.Register(ModuleName, 1500, "ErrCalculatGrowRate err")
	ErrRiskRatioMustBeZero           = sdkerrors.Register(ModuleName, 1501, "ErrRiskRatioMustBeZero err")
	ErrRiskRateIsGreaterThenShouldBe = sdkerrors.Register(ModuleName, 1502, "ErrRiskRateIsGreaterThenShouldBe err")
	ErrPairNotFound                  = sdkerrors.Register(ModuleName, 1701, "ErrPairNotFound err")
	ErrPositionNotFound              = sdkerrors.Register(ModuleName, 1702, "ErrPositionNotFound err")
	ErrNeedSendUSQ                   = sdkerrors.Register(ModuleName, 1703, "ErrNeedSendUSQ err")
	ErrLoanNotFound                  = sdkerrors.Register(ModuleName, 1704, "ErrLoanNotFound err")
	ErrNotEnoughAmountIn             = sdkerrors.Register(ModuleName, 1705, "ErrNotEnoughAmountIn err")
	ErrLoanNotFoundInPosition        = sdkerrors.Register(ModuleName, 1706, "ErrLoanNotFoundInPosition err")
	ErrLiqPositionNotFound           = sdkerrors.Register(ModuleName, 1707, "ErrLiqPositionNotFound err")
	ErrLendAssetNotFound             = sdkerrors.Register(ModuleName, 1708, "ErrLendAssetNotFound err")
	ErrAmountInGTEminAmountIn        = sdkerrors.Register(ModuleName, 1801, "ErrAmountInGTEminAmountIn err")
	ErrAmountOutGTEminAmountOut      = sdkerrors.Register(ModuleName, 1802, "ErrAmountOutGTEminAmountOut err")
	ErrPriceNil                      = sdkerrors.Register(ModuleName, 1900, "ErrPriceNil err")
	ErrUserAlredyDepositCollateral   = sdkerrors.Register(ModuleName, 2000, "ErrUserAlredyDepositCollateral err")
)
