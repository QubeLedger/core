package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	Multiplier        sdk.Int = sdk.NewInt(int64(10000))
	SdkMultiplier     sdk.Int = sdk.NewInt(int64(1000000))
	FeeMultiplier     sdk.Int = sdk.NewInt(int64(100000))
	MintUsqMultiplier sdk.Int = sdk.NewInt(int64(1000))
	BurnUsqMultiplier sdk.Int = sdk.NewInt(int64(1000000000))
)
