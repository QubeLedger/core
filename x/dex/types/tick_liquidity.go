package types

import (
	math_utils "github.com/QuadrateOrg/core/x/dex/utils/math"
)

// NOTE: These methods should be avoided if possible.
// Generally default to dealing with LimitOrderTranche or PoolReserves explicitly

func (t TickLiquidity) Price() math_utils.PrecDec {
	switch liquidity := t.Liquidity.(type) {
	case *TickLiquidity_LimitOrderTranche:
		return liquidity.LimitOrderTranche.PriceTakerToMaker

	case *TickLiquidity_PoolReserves:
		return liquidity.PoolReserves.PriceTakerToMaker
	default:
		panic("Tick does not contain valid liqudityType")
	}
}

func (t TickLiquidity) TickIndex() int64 {
	switch liquidity := t.Liquidity.(type) {
	case *TickLiquidity_LimitOrderTranche:
		return liquidity.LimitOrderTranche.Key.TickIndexTakerToMaker

	case *TickLiquidity_PoolReserves:
		return liquidity.PoolReserves.Key.TickIndexTakerToMaker
	default:
		panic("Tick does not contain valid liqudityType")
	}
}

func (t TickLiquidity) HasToken() bool {
	switch liquidity := t.Liquidity.(type) {
	case *TickLiquidity_LimitOrderTranche:
		return liquidity.LimitOrderTranche.HasTokenIn()

	case *TickLiquidity_PoolReserves:
		return liquidity.PoolReserves.HasToken()
	default:
		panic("Tick does not contain valid liqudityType")
	}
}
