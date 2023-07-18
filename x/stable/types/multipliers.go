package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	Multiplier        sdk.Int = sdk.NewInt(int64(10000))
	MintUsqMultiplier sdk.Int = sdk.NewInt(int64(1000))
	BurnUsqMultiplier sdk.Int = sdk.NewInt(int64(1000000000))
)
