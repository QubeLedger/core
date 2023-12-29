package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

/*
USQReserveAddress
*/

func (k Keeper) ChangeUSQReserveAddress(ctx sdk.Context, address sdk.AccAddress) error {
	k.SetUSQReserveAddress(ctx, address)
	return nil
}

func (k Keeper) SetUSQReserveAddress(ctx sdk.Context, newUSQReserveAddress sdk.AccAddress) {
	params := k.GetParams(ctx)
	params.USQReserveAddress = newUSQReserveAddress.String()
	k.SetParams(ctx, params)
}

func (k Keeper) GetUSQReserveAddress(ctx sdk.Context) sdk.AccAddress {
	params := k.GetParams(ctx)
	addr, _ := sdk.AccAddressFromBech32(params.USQReserveAddress)
	return addr
}

/*
GrowYieldReserveAddress
*/

func (k Keeper) ChangeGrowYieldReserveAddress(ctx sdk.Context, address sdk.AccAddress) error {
	k.SetGrowYieldReserveAddress(ctx, address)
	return nil
}

func (k Keeper) SetGrowYieldReserveAddress(ctx sdk.Context, newGrowYieldReserveAddress sdk.AccAddress) {
	params := k.GetParams(ctx)
	params.GrowYieldReserveAddress = newGrowYieldReserveAddress.String()
	k.SetParams(ctx, params)
}

func (k Keeper) GetGrowYieldReserveAddress(ctx sdk.Context) sdk.AccAddress {
	params := k.GetParams(ctx)
	addr, _ := sdk.AccAddressFromBech32(params.GrowYieldReserveAddress)
	return addr
}

/*
GrowStakingReserveAddress
*/

func (k Keeper) ChangeGrowStakingReserveAddress(ctx sdk.Context, address sdk.AccAddress) error {
	k.SetGrowStakingReserveAddress(ctx, address)
	return nil
}

func (k Keeper) SetGrowStakingReserveAddress(ctx sdk.Context, newGrowStakingReserveAddress sdk.AccAddress) {
	params := k.GetParams(ctx)
	params.GrowStakingReserveAddress = newGrowStakingReserveAddress.String()
	k.SetParams(ctx, params)
}

func (k Keeper) GetGrowStakingReserveAddress(ctx sdk.Context) sdk.AccAddress {
	params := k.GetParams(ctx)
	addr, _ := sdk.AccAddressFromBech32(params.GrowStakingReserveAddress)
	return addr
}

/*
Check Address for empty
*/

func (k Keeper) AddressEmptyCheck(ctx sdk.Context) bool {
	params := k.GetParams(ctx)
	USQReserveAddress, _ := sdk.AccAddressFromBech32(params.USQReserveAddress)
	GrowYieldReserveAddress, _ := sdk.AccAddressFromBech32(params.GrowYieldReserveAddress)
	GrowStakingReserveAddress, _ := sdk.AccAddressFromBech32(params.GrowStakingReserveAddress)

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
