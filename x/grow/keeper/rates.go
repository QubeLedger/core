package keeper

import (
	"github.com/QuadrateOrg/core/x/grow/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetRatesByUtilizationRate(ctx sdk.Context, utilization_rate float64, asset types.Asset) (float64, float64, error) {
	params := k.GetParams(ctx)

	if utilization_rate == 0.0 {
		return 0.0, 0.0, types.ErrSdkIntError
	}

	var u_static float64
	var max_rate float64

	switch {
	case asset.Type == "volatile":
		u_static = float64(params.UStaticVolatile) / 100
		max_rate = float64(params.MaxRateVolatile) / 100
	case asset.Type == "stable":
		u_static = float64(params.UStaticStable) / 100
		max_rate = float64(params.MaxRateStable) / 100
	}

	slope := float64(params.Slope) / 100

	borrow_interest_rate := slope + ((utilization_rate-u_static)/(1-u_static))*max_rate
	supply_interest_rate := (borrow_interest_rate) * utilization_rate

	return borrow_interest_rate, supply_interest_rate, nil
}
