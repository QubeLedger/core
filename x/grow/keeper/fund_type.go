package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	USQReserveAddress         sdk.AccAddress
	GrowYieldReserveAddress   sdk.AccAddress
	GrowStakingReserveAddress sdk.AccAddress
)

/*
USQReserveAddress
*/

func (k Keeper) ChangeUSQReserveAddress(ctx sdk.Context, address sdk.AccAddress) error {
	k.SetUSQReserveAddress(ctx, address)
	return nil
}

func (k Keeper) SetUSQReserveAddress(ctx sdk.Context, newUSQReserveAddress sdk.AccAddress) {
	USQReserveAddress = newUSQReserveAddress
}

func (k Keeper) GetUSQReserveAddress(ctx sdk.Context) sdk.AccAddress {
	return USQReserveAddress
}

/*
GrowYieldReserveAddress
*/

func (k Keeper) ChangeGrowYieldReserveAddress(ctx sdk.Context, address sdk.AccAddress) error {
	k.SetGrowYieldReserveAddress(ctx, address)
	return nil
}

func (k Keeper) SetGrowYieldReserveAddress(ctx sdk.Context, newGrowYieldReserveAddress sdk.AccAddress) {
	GrowYieldReserveAddress = newGrowYieldReserveAddress
}

func (k Keeper) GetGrowYieldReserveAddress(ctx sdk.Context) sdk.AccAddress {
	return GrowYieldReserveAddress
}

/*
GrowStakingReserveAddress
*/

func (k Keeper) ChangeGrowStakingReserveAddress(ctx sdk.Context, address sdk.AccAddress) error {
	k.SetGrowStakingReserveAddress(ctx, address)
	return nil
}

func (k Keeper) SetGrowStakingReserveAddress(ctx sdk.Context, newUSQStakingReserveAddress sdk.AccAddress) {
	GrowStakingReserveAddress = newUSQStakingReserveAddress
}

func (k Keeper) GetGrowStakingReserveAddress(ctx sdk.Context) sdk.AccAddress {
	return GrowStakingReserveAddress
}

/*
Check Address for empty
*/

func (k Keeper) AddressEmptyCheck(ctx sdk.Context) bool {
	if USQReserveAddress.Empty() {
		return true
	}
	if GrowYieldReserveAddress.Empty() {
		return true
	}
	if GrowStakingReserveAddress.Empty() {
		return true
	}

	return false
}
