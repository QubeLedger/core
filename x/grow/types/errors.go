package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrSample                        = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrSdkIntError                   = sdkerrors.Register(ModuleName, 1101, "sdk.Int error")
	ErrIntNegativeOrZero             = sdkerrors.Register(ModuleName, 1102, "ErrIntNegativeOrZero error")
	ErrInvalidLength                 = sdkerrors.Register(ModuleName, 1103, "invalid length")
	ErrOracleAssetIdNotFound         = sdkerrors.Register(ModuleName, 1200, "ErrOracleAssetIdNotFound err")
	ErrCoinsLen                      = sdkerrors.Register(ModuleName, 1400, "ErrCoinsLen err")
	ErrDenomsNotEqual                = sdkerrors.Register(ModuleName, 1401, "ErrDenomsNotEqual err")
	ErrLiquidatorAddresesNotEqual    = sdkerrors.Register(ModuleName, 1402, "ErrLiquidatorAddresesNotEqual err")
	ErrLiquidatorPositionIdNotEqual  = sdkerrors.Register(ModuleName, 1403, "ErrLiquidatorPositionIdNotEqual err")
	ErrWrongPremium                  = sdkerrors.Register(ModuleName, 1404, "ErrWrongPremium err")
	ErrLiquidationMechanismError     = sdkerrors.Register(ModuleName, 1405, "Error in Liquidation Mechanism")
	ErrCalculateGrowRate             = sdkerrors.Register(ModuleName, 1500, "ErrCalculateGrowRate err")
	ErrRiskRatioMustBeZero           = sdkerrors.Register(ModuleName, 1501, "ErrRiskRatioMustBeZero err")
	ErrRiskRateIsGreaterThenShouldBe = sdkerrors.Register(ModuleName, 1502, "ErrRiskRateIsGreaterThenShouldBe err")
	ErrPairNotFound                  = sdkerrors.Register(ModuleName, 1701, "ErrPairNotFound err")
	ErrPositionNotFound              = sdkerrors.Register(ModuleName, 1702, "ErrPositionNotFound err")
	ErrNeedSendUSQ                   = sdkerrors.Register(ModuleName, 1703, "ErrNeedSendUSQ err")
	ErrLoanNotFound                  = sdkerrors.Register(ModuleName, 1704, "ErrLoanNotFound err")
	ErrNotEnoughAmountIn             = sdkerrors.Register(ModuleName, 1705, "ErrNotEnoughAmountIn err")
	ErrLoanNotFoundInPosition        = sdkerrors.Register(ModuleName, 1706, "ErrLoanNotFoundInPosition err")
	ErrLiqPositionNotFound           = sdkerrors.Register(ModuleName, 1707, "ErrLiqPositionNotFound err")
	ErrAssetNotFound                 = sdkerrors.Register(ModuleName, 1708, "ErrAssetNotFound err")
	ErrAmountInGTEminAmountIn        = sdkerrors.Register(ModuleName, 1801, "ErrAmountInGTEminAmountIn err")
	ErrAmountOutGTEminAmountOut      = sdkerrors.Register(ModuleName, 1802, "ErrAmountOutGTEminAmountOut err")
	ErrPriceNil                      = sdkerrors.Register(ModuleName, 1900, "ErrPriceNil err")
	ErrUserAlredyDepositCollateral   = sdkerrors.Register(ModuleName, 2000, "ErrUserAlredyDepositCollateral err")
	ErrReserveAddressEmpty           = sdkerrors.Register(ModuleName, 4000, "ErrReserveAddressEmpty err")
	ErrCalculateBIROrSIR             = sdkerrors.Register(ModuleName, 5000, "ErrCalculateBIROrSIR err")
	ErrNeedSendRigthDenom            = sdkerrors.Register(ModuleName, 6000, "ErrNeedSendRigthDenom err")
	ErrLendNotFound                  = sdkerrors.Register(ModuleName, 7000, "ErrLendNotFound err")
	ErrNotEnoughProvideValue         = sdkerrors.Register(ModuleName, 8000, "ErrNotEnoughProvideValue err")
	ErrDepositNotActivated           = sdkerrors.Register(ModuleName, 3000, "Grow Deposit is off. Wait until it is turned on to start using it. To enable the module, make a proposal in the governance.")
	ErrCollateralNotActivated        = sdkerrors.Register(ModuleName, 3001, "Grow Collateral is off. Wait until it is turned on to start using it. To enable the module, make a proposal in the governance.")
	ErrBorrowNotActivated            = sdkerrors.Register(ModuleName, 3002, "Grow Borrow is off. Wait until it is turned on to start using it. To enable the module, make a proposal in the governance.")
)
