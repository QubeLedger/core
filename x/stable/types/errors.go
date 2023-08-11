package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/stable module sentinel errors
var (
	ErrSample                          = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrSdkIntError                     = sdkerrors.Register(ModuleName, 1101, "sdk.Int error")
	ErrInvalidLength                   = sdkerrors.Register(ModuleName, 1102, "invalid length")
	ErrInvalidPacketTimeout            = sdkerrors.Register(ModuleName, 1500, "invalid packet timeout")
	ErrInvalidVersion                  = sdkerrors.Register(ModuleName, 1501, "invalid version")
	ErrAfpNegative                     = sdkerrors.Register(ModuleName, 1200, "AFP negative")
	ErrArNegative                      = sdkerrors.Register(ModuleName, 1201, "AR negative")
	ErrQmNegative                      = sdkerrors.Register(ModuleName, 1202, "QM negative")
	ErrARAlreadyInit                   = sdkerrors.Register(ModuleName, 1203, "ErrARAlreadyInit")
	ErrBackingRatioNegative            = sdkerrors.Register(ModuleName, 1204, "BackingRatio negative")
	ErrBackingRatioNil                 = sdkerrors.Register(ModuleName, 1205, "BackingRatio nil")
	ErrQMAlreadyInit                   = sdkerrors.Register(ModuleName, 1206, "ErrQMAlreadyInit")
	ErrCalculateMintingFee             = sdkerrors.Register(ModuleName, 1301, "CalculateMintingFee err")
	ErrCalculateBurningFee             = sdkerrors.Register(ModuleName, 1302, "CalculateBurningFee err")
	ErrInvalidCoins                    = sdkerrors.Register(ModuleName, 500, "ErrInvalidCoins err")
	ErrAtomPriceNil                    = sdkerrors.Register(ModuleName, 1000, "ErrAtomPriceNil err")
	ErrMultipleCoinsLockupNotSupported = sdkerrors.Register(ModuleName, 1600, "ErrMultipleCoinsLockupNotSupported multiple coins lockup is not supported")
	ErrSendBaseTokenDenom              = sdkerrors.Register(ModuleName, 1601, "ErrSendBaseTokenDenom err")
	ErrSendSendTokenDenom              = sdkerrors.Register(ModuleName, 1602, "ErrSendSendTokenDenom err")
	ErrPairNotFound                    = sdkerrors.Register(ModuleName, 1701, "ErrPairNotFound err")
	ErrAmountInGTEminAmountIn          = sdkerrors.Register(ModuleName, 1801, "ErrAmountInGTEminAmountIn err")
	ErrAmountOutGTEminAmountOut        = sdkerrors.Register(ModuleName, 1802, "ErrAmountOutGTEminAmountOut err")
)

var (
	ErrUnrecognizedPacketType = sdkerrors.Register(ModuleName, 15, "unrecognized packet type")
)

// GMD errors
var (
	ErrMintBlocked = sdkerrors.Register(ModuleName, 100, "Backing Ration >= 120%")
	ErrBurnBlocked = sdkerrors.Register(ModuleName, 101, "Backing Ration < 85%")
)
