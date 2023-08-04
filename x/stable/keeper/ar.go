package keeper

import (
	"github.com/QuadrateOrg/core/x/stable/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	AtomReserve sdk.Int
)

func (k Keeper) GetAtomReserve(ctx sdk.Context) sdk.Int {
	return AtomReserve
}

func (k Keeper) IncreaseAtomReserve(ctx sdk.Context, amount sdk.Int) error {
	if amount.IsNil() {
		return types.ErrArNegative
	}
	if amount.IsNegative() {
		return types.ErrArNegative
	}
	AtomReserve = AtomReserve.Add(amount)
	return nil
}

func (k Keeper) ReduceAtomReserve(ctx sdk.Context, amount sdk.Int) error {
	if amount.IsNil() {
		return types.ErrArNegative
	}
	if amount.IsNegative() {
		return types.ErrArNegative
	}
	if amount.GTE(AtomReserve) {
		return types.ErrArNegative
	}
	AtomReserve = AtomReserve.Sub(amount)
	return nil
}

func (k Keeper) InitAtomReserve(ctx sdk.Context) error {
	if !AtomReserve.IsNil() {
		return types.ErrARAlreadyInit
	}
	AtomReserve = sdk.NewInt(0)
	if !AtomReserve.IsZero() {
		return types.ErrARAlreadyInit
	}
	return nil
}
