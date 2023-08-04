package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	StabilityFundAddress sdk.AccAddress
)

func (k Keeper) ChangeStabilityFundAddress(ctx sdk.Context, address sdk.AccAddress) error {
	k.SetStabilityFundAddress(ctx, address)
	return nil
}

func (k Keeper) SetStabilityFundAddress(ctx sdk.Context, newStabilityFundAddress sdk.AccAddress) {
	StabilityFundAddress = newStabilityFundAddress
}

func (k Keeper) GetStabilityFundAddress(ctx sdk.Context) sdk.AccAddress {
	return StabilityFundAddress
}
