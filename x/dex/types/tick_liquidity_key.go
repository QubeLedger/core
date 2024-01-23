package types

import math_utils "github.com/QuadrateOrg/core/x/dex/utils/math"

type TickLiquidityKey interface {
	KeyMarshal() []byte
	PriceTakerToMaker() (priceTakerToMaker math_utils.PrecDec, err error)
}
