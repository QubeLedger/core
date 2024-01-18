package types

import (
	math_utils "github.com/QuadrateOrg/core/x/dex/utils/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Liquidity interface {
	Swap(maxAmountTakerIn sdk.Int, maxAmountMakerOut *sdk.Int) (inAmount, outAmount sdk.Int)
	Price() math_utils.PrecDec
}
