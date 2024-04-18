package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/perpetual module sentinel errors
var (
	ErrSample                        = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrTradingAssetEmpty             = sdkerrors.Register(ModuleName, 1101, "ErrTradingAssetEmpty err")
	ErrVaultNotFound                 = sdkerrors.Register(ModuleName, 1102, "ErrVaultNotFound err")
	ErrInCalculationUpdateVault      = sdkerrors.Register(ModuleName, 1103, "ErrInCalculationUpdateVault err")
	ErrPositionNotFound              = sdkerrors.Register(ModuleName, 1104, "ErrPositionNotFound err")
	ErrAmountGreaterThanPositionSize = sdkerrors.Register(ModuleName, 1105, "ErrAmountGreaterThanPositionSize err")
	ErrLeverageEqualZero             = sdkerrors.Register(ModuleName, 1106, "ErrLeverageEqualZero err")
	ErrNotSdkInt                     = sdkerrors.Register(ModuleName, 1107, "ErrNotSdkInt err")
)
