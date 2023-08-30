package keeper

import (
	"github.com/QuadrateOrg/core/x/grow/types"
	stabletypes "github.com/QuadrateOrg/core/x/stable/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) IncreaseGrowStakingReserve(ctx sdk.Context, amountIn sdk.Coins, gTokenPair types.GTokenPair, qStablePair stabletypes.Pair) (types.GTokenPair, error) {
	err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, k.GetGrowStakingReserveAddress(ctx), amountIn)
	if err != nil {
		return gTokenPair, err
	}

	gTokenPair.St = gTokenPair.St.Add(amountIn.AmountOf(qStablePair.AmountOutMetadata.Base))

	return gTokenPair, nil
}

func (k Keeper) ReduceGrowStakingReserve(ctx sdk.Context, amountIn sdk.Coins, gTokenPair types.GTokenPair) (types.GTokenPair, error) {
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, k.GetGrowStakingReserveAddress(ctx), types.ModuleName, amountIn)
	if err != nil {
		return gTokenPair, err
	}

	gTokenPair.St = gTokenPair.St.Sub(amountIn.AmountOf(gTokenPair.GTokenMetadata.Base))

	return gTokenPair, nil
}
