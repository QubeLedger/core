package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) ChangeBurningFundAddress(ctx sdk.Context, address sdk.AccAddress) error {
	k.SetBurningFundAddress(ctx, address)
	return nil
}

func (k Keeper) SetBurningFundAddress(ctx sdk.Context, newBurningFundAddress sdk.AccAddress) {
	params := k.GetParams(ctx)
	params.BurningFundAddress = newBurningFundAddress.String()
	k.SetParams(ctx, params)
}

func (k Keeper) GetBurningFundAddress(ctx sdk.Context) sdk.AccAddress {
	params := k.GetParams(ctx)
	addr, _ := sdk.AccAddressFromBech32(params.BurningFundAddress)
	return addr
}

func (k Keeper) ChangeReserveFundAddress(ctx sdk.Context, address sdk.AccAddress) error {
	k.SetReserveFundAddress(ctx, address)
	return nil
}

func (k Keeper) SetReserveFundAddress(ctx sdk.Context, newReserveFundAddress sdk.AccAddress) {
	params := k.GetParams(ctx)
	params.ReserveFundAddress = newReserveFundAddress.String()
	k.SetParams(ctx, params)
}

func (k Keeper) GetReserveFundAddress(ctx sdk.Context) sdk.AccAddress {
	params := k.GetParams(ctx)
	addr, _ := sdk.AccAddressFromBech32(params.ReserveFundAddress)
	return addr
}

/*
Check Address for empty
*/

func (k Keeper) AddressEmptyCheck(ctx sdk.Context) bool {
	params := k.GetParams(ctx)
	ReserveFundAddress, _ := sdk.AccAddressFromBech32(params.ReserveFundAddress)
	BurningFundAddress, _ := sdk.AccAddressFromBech32(params.BurningFundAddress)
	if ReserveFundAddress.Empty() {
		return true
	}
	if BurningFundAddress.Empty() {
		return true
	}

	return false
}
