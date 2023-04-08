package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/printer module sentinel errors
var (
	ErrSample               = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrABIPack              = sdkerrors.Register(ModuleName, 305, "contract ABI pack failed")
	ErrInvalidPacketTimeout = sdkerrors.Register(ModuleName, 1500, "invalid packet timeout")
	ErrInvalidVersion       = sdkerrors.Register(ModuleName, 1501, "invalid version")
	ErrUnknownRequest       = sdkerrors.Register(ModuleName, 4, "unknown request")
	ErrWrongDenom           = sdkerrors.Register(ModuleName, 321, "wrong denom")
	ErrValNotFound          = sdkerrors.Register(ModuleName, 322, "validator not found")
)
