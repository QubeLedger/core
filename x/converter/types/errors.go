package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/converter module sentinel errors
var (
	ErrSample  = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrABIPack = sdkerrors.Register(ModuleName, 305, "contract ABI pack failed")
)
