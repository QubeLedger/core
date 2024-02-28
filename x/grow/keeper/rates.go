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

	borrow_interest_rate := 0.0
	supply_interest_rate := 0.0

	slope_1 := float64(params.Slope_1) / 100
	slope_2 := float64(params.Slope_2) / 100

	if utilization_rate < u_static {
		borrow_interest_rate = slope_1 + (utilization_rate * ((slope_2 - slope_1) / u_static))
		supply_interest_rate = (borrow_interest_rate) * utilization_rate
	} else {
		borrow_interest_rate = slope_1 + ((utilization_rate - u_static) * ((max_rate - slope_2) / (1 - u_static)))
		supply_interest_rate = (borrow_interest_rate) * utilization_rate
	}

	return borrow_interest_rate, supply_interest_rate, nil
}
