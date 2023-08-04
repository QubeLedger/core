package keeper

import sdk "github.com/cosmos/cosmos-sdk/types"

func (k Keeper) GetReserve(ctx sdk.Context) (sdk.Int, sdk.Int) {
	qm := k.GetStablecoinSupply(ctx)
	ar := k.GetAtomReserve(ctx)

	if qm.IsNil() && ar.IsNil() {
		err := k.InitAtomReserve(ctx) // #nosec
		if err != nil {
			panic(err)
		}
		err = k.InitStablecoinSupply(ctx) // #nosec
		if err != nil {
			panic(err)
		}
	}

	qm = k.GetStablecoinSupply(ctx)
	ar = k.GetAtomReserve(ctx)

	return qm, ar
}

func (k Keeper) IncreaseReserve(ctx sdk.Context, amount1 sdk.Int, amount2 sdk.Int) error {
	err := k.IncreaseAtomReserve(ctx, amount1)
	if err != nil {
		return err
	}
	err = k.IncreaseStablecoinSupply(ctx, amount2)
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) ReduceReserve(ctx sdk.Context, amount1 sdk.Int, amount2 sdk.Int) error {
	err := k.ReduceAtomReserve(ctx, amount1)
	if err != nil {
		return err
	}
	err = k.ReduceStablecoinSupply(ctx, amount2)
	if err != nil {
		return err
	}
	return nil
}
