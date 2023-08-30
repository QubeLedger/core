package keeper

import (
	"github.com/QuadrateOrg/core/x/grow/types"
	gmb "github.com/QuadrateOrg/core/x/stable/gmb"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	realRate              sdk.Int
	lastTimeUpdateReserve sdk.Int
)

func (k Keeper) SetLastTimeUpdateReserve(ctx sdk.Context, val sdk.Int) error {
	if val.IsNil() || val.IsZero() || val.IsNegative() {
		return types.ErrIntNegativeOrZero
	}
	lastTimeUpdateReserve = val
	return nil
}

func (k Keeper) SetRealRate(ctx sdk.Context, val sdk.Int) error {
	if val.IsNil() || val.IsZero() || val.IsNegative() {
		return types.ErrIntNegativeOrZero
	}
	realRate = val
	return nil
}

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

func (k Keeper) CalculateGrowYield(ctx sdk.Context, gTokenPair types.GTokenPair) (sdk.Int, error) {
	growRate, err := k.CalculateGrowRate(ctx, gTokenPair)
	if err != nil {
		return sdk.Int{}, err
	}
	if gTokenPair.St.IsNil() || gTokenPair.St.IsZero() {
		return sdk.Int{}, types.ErrIntNegativeOrZero
	}

	return (gTokenPair.St.Mul(growRate)).Quo(sdk.NewInt(1000)), nil
}

func (k Keeper) CalculateRealYield(ctx sdk.Context, gTokenPair types.GTokenPair) (sdk.Int, error) {
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

	qm := qStablePair.Qm
	res := ((qm.Mul(br)).Mul(realRate)).QuoRaw(10000)

	return res, nil
}

func (k Keeper) CheckYieldRate(ctx sdk.Context, gTokenPair types.GTokenPair) (string, sdk.Int, error) {
	growYield, err := k.CalculateGrowYield(ctx, gTokenPair)
	if err != nil {
		return "", sdk.Int{}, err
	}
	realYield, err := k.CalculateRealYield(ctx, gTokenPair)
	if err != nil {
		return "", sdk.Int{}, err
	}

	if realYield.GT(growYield) {
		return types.SendToReserveAction, realYield.Sub(growYield), nil
	}

	if growYield.GT(realYield) {
		return types.SendFromReserveAction, growYield.Sub(realYield), nil
	}

	return "", sdk.Int{}, nil
}

func (k Keeper) CalculateAddToReserveValue(ctx sdk.Context, val sdk.Int, gTokenPair types.GTokenPair) (sdk.Int, bool) {
	diff := sdk.NewInt(ctx.BlockTime().Unix()).Sub(lastTimeUpdateReserve)
	if diff.LT(sdk.NewInt(10)) {
		return sdk.Int{}, false
	}

	if (sdk.NewInt(31536000).Quo(diff)).IsNil() || (sdk.NewInt(31536000).Quo(diff)).IsZero() {
		return sdk.Int{}, false
	}
	return val.Quo(sdk.NewInt(31536000).Quo(diff)), true
}
