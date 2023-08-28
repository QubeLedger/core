package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	USQReserveAddress       sdk.AccAddress
	GrowYieldReserveAddress sdk.AccAddress
)

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
