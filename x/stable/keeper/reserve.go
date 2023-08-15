package keeper

import (
	"github.com/QubeLedger/core/x/stable/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) IncreaseReserve(ctx sdk.Context, amount1 sdk.Int, amount2 sdk.Int, pair types.Pair) types.Pair {
	pair.Ar, pair.Qm = pair.Ar.Add(amount1), pair.Qm.Add(amount2)
	return pair
}

func (k Keeper) ReduceReserve(ctx sdk.Context, amount1 sdk.Int, amount2 sdk.Int, pair types.Pair) types.Pair {
	pair.Ar, pair.Qm = pair.Ar.Sub(amount1), pair.Qm.Sub(amount2)
	return pair
}
