package keeper

import (
	"time"

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

	actualGTokenPrice := gTokenPair.GTokenLastPrice
	latestTime := gTokenPair.GTokenLatestPriceUpdateTime
	growRate, err := k.CalculateGrowRate(ctx, gTokenPair)
	if err != nil {
		return err
	}

	now := sdk.NewInt(time.Now().Unix())
	timeSinceLastUpdate := now.Sub(sdk.NewIntFromUint64(latestTime)).Quo(sdk.NewInt(86400))

	newGTokenPrice := k.CalculateGTokenAPY(actualGTokenPrice, growRate, timeSinceLastUpdate)

	gTokenPair.GTokenLastPrice = newGTokenPrice
	gTokenPair.GTokenLatestPriceUpdateTime = now.Uint64()
	k.SetPair(ctx, gTokenPair)
	return nil
}
