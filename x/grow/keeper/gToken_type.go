package keeper

import (
	"github.com/QuadrateOrg/core/x/grow/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetGTokenPrice(ctx sdk.Context, denomID string) (sdk.Int, error) {
	gTokenPair, found := k.GetPairByDenomID(ctx, denomID)
	if !found {
		return sdk.Int{}, types.ErrPairNotFound
	}
	return gTokenPair.GTokenLastPrice, nil
}

func (k Keeper) UpdateGTokenPrice(ctx sdk.Context, gTokenPair types.GTokenPair) error {
	latestTime := gTokenPair.GTokenLatestPriceUpdateTime

	growRate, err := k.CalculateGrowRate(ctx, gTokenPair)
	if err != nil {
		return err
	}
	if growRate.IsNil() || growRate.IsZero() {
		return types.ErrIntNegativeOrZero
	}

	now := sdk.NewInt(ctx.BlockTime().Unix())
	if now.IsNil() || now.IsZero() {
		return types.ErrIntNegativeOrZero
	}

	timeSinceLastUpdate := now.Sub(sdk.NewIntFromUint64(latestTime)).Quo(sdk.NewInt(86400))

	if timeSinceLastUpdate.IsZero() {
		return nil
	}

	newGTokenPrice := k.CalculateGTokenAPY(sdk.NewInt(1*1000000), growRate, timeSinceLastUpdate)
	if newGTokenPrice.IsNil() || newGTokenPrice.IsZero() {
		return types.ErrIntNegativeOrZero
	}

	gTokenPair.GTokenLastPrice = newGTokenPrice
	gTokenPair.GTokenLatestPriceUpdateTime = now.Uint64()
	k.SetPair(ctx, gTokenPair)
	return nil
}
