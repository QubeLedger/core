package types

import (
	math_utils "github.com/QuadrateOrg/core/x/dex/utils/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type PoolLiquidity struct {
	TradePairID *TradePairID
	Pool        *Pool
}

func (pl *PoolLiquidity) Swap(
	maxAmountTakerDenomIn sdk.Int,
	maxAmountMakerDenomOut *sdk.Int,
) (inAmount, outAmount sdk.Int) {
	return pl.Pool.Swap(
		pl.TradePairID,
		maxAmountTakerDenomIn,
		maxAmountMakerDenomOut,
	)
}

func (pl *PoolLiquidity) Price() math_utils.PrecDec {
	return pl.Pool.Price(pl.TradePairID)
}
