package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	BurningFundAddress sdk.AccAddress
)

func (k Keeper) ChangeBurningFundAddress(ctx sdk.Context, address sdk.AccAddress) error {
	k.SetBurningFundAddress(ctx, address)
	return nil
}

func (k Keeper) SetBurningFundAddress(ctx sdk.Context, newBurningFundAddress sdk.AccAddress) {
	BurningFundAddress = newBurningFundAddress
}

func (k Keeper) GetBurningFundAddress(ctx sdk.Context) sdk.AccAddress {
	return BurningFundAddress
}
