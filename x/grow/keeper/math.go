package keeper

import (
	"math"

	"github.com/QuadrateOrg/core/x/grow/types"
	gmb "github.com/QuadrateOrg/core/x/stable/gmb"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func CalculatGrowRatePercent(backing_ratio sdk.Int) (sdk.Int, error) {
	if backing_ratio.IsNil() {
		return sdk.Int{}, types.ErrCalculatGrowRate
	}
	switch {
	case backing_ratio.GT(sdk.NewInt(int64(110))):
		return sdk.NewInt(75), nil

	case backing_ratio.GT(sdk.NewInt(93)):
		return sdk.NewInt(150), nil

	case backing_ratio.GT(sdk.NewInt(85)):
		return sdk.NewInt(200), nil

	case sdk.NewInt(int64(85)).GT(backing_ratio) || backing_ratio.Equal(sdk.NewInt(85)):
		return sdk.NewInt(250), nil

	default:
		return sdk.Int{}, types.ErrCalculatGrowRate
	}

}

func (k Keeper) CalculateGrowRate(ctx sdk.Context, gTokenPair types.GTokenPair) (sdk.Int, error) {
	qStablePair, found := k.stableKeeper.GetPairByPairID(ctx, gTokenPair.QStablePairId)
	if !found {
		return sdk.Int{}, types.ErrPairNotFound
	}
	atomPrice, err := k.oracleKeeper.GetExchangeRate(ctx, qStablePair.AmountInMetadata.Base)
	if err != nil {
		return sdk.Int{}, err
	}
	br, err := gmb.CalculateBackingRatio(atomPrice.MulInt64(10000).RoundInt(), qStablePair.Ar, qStablePair.Qm)
	if err != nil {
		return sdk.Int{}, err
	}

	growRate, err := CalculatGrowRatePercent(br)
	if err != nil {
		return sdk.Int{}, err
	}

	return growRate, nil
}

func (k Keeper) CalculateGTokenAmountOut(amount sdk.Int, price sdk.Int) sdk.Int {
	return ((amount.Mul(types.Multiplier)).Quo(price))
}

func (k Keeper) CalculateReturnQubeStableAmountOut(amount sdk.Int, price sdk.Int) sdk.Int {
	return ((amount.Mul(price)).Quo(types.Multiplier))
}

func (k Keeper) CalculateGTokenAPY(lastAmount sdk.Int, growRate sdk.Int, day sdk.Int) sdk.Int {
	lastAmountInt := lastAmount.Int64()
	growRateInt := growRate.Int64()
	dayInt := day.Int64()

	res := float64(lastAmountInt) * (math.Pow((1 + (float64(growRateInt)/1000)/365), (float64(dayInt) - 1)))
	return sdk.NewInt(int64(res))
}
