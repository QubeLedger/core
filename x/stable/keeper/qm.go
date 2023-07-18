package keeper

import (
	"github.com/QuadrateOrg/core/x/stable/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	StablecoinSupply sdk.Int
)

func (k Keeper) GetStablecoinSupply(ctx sdk.Context) sdk.Int {
	return StablecoinSupply
}

func (k Keeper) IncreaseStablecoinSupply(ctx sdk.Context, amount sdk.Int) error {
	if amount.IsNil() {
		return types.ErrQmNegative
	}
	if amount.IsNegative() {
		return types.ErrQmNegative
	}
	StablecoinSupply = StablecoinSupply.Add(amount)
	return nil
}

func (k Keeper) ReduceStablecoinSupply(ctx sdk.Context, amount sdk.Int) error {
	if amount.IsNil() {
		return types.ErrQmNegative
	}
	if amount.IsNegative() {
		return types.ErrQmNegative
	}
	if amount.GTE(StablecoinSupply) {
		return types.ErrQmNegative
	}
	StablecoinSupply = StablecoinSupply.Sub(amount)
	return nil
}

func (k Keeper) InitStablecoinSupply(ctx sdk.Context) error {
	if !StablecoinSupply.IsNil() {
		return types.ErrQMAlreadyInit
	}
	StablecoinSupply = sdk.NewInt(0)
	if !StablecoinSupply.IsZero() {
		return types.ErrQMAlreadyInit
	}
	return nil
}
