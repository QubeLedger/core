package gmb

import (
	"github.com/QuadrateOrg/core/x/stable/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func CalculateBackingRatio(afp sdk.Int, ar sdk.Int, qm sdk.Int) (sdk.Int, error) {
	if afp.IsNegative() {
		return sdk.Int{}, types.ErrAfpNegative
	}
	if ar.IsNegative() {
		return sdk.Int{}, types.ErrArNegative
	}
	if qm.IsNegative() {
		return sdk.Int{}, types.ErrQmNegative
	}

	backing_ratio := (afp.Mul(ar).Quo(qm).QuoRaw(100))

	if backing_ratio.IsNegative() {
		return sdk.Int{}, types.ErrBackingRatioNegative
	}
	if backing_ratio.IsNil() {
		return sdk.Int{}, types.ErrBackingRatioNil
	}
	return backing_ratio, nil
}

func CalculateMintingFee(backing_ratio sdk.Int) (sdk.Int, error) {
	if backing_ratio.IsNil() {
		return sdk.Int{}, types.ErrCalculateMintingFee
	}
	switch {
	case backing_ratio.GT(sdk.NewInt(140)) || backing_ratio.Equal(sdk.NewInt(140)):
		return sdk.Int{}, types.ErrMintBlocked

	case backing_ratio.GT(sdk.NewInt(120)) || backing_ratio.Equal(sdk.NewInt(120)):
		return sdk.NewInt(10), nil

	case backing_ratio.GT(sdk.NewInt(100)) || backing_ratio.Equal(sdk.NewInt(100)):
		return sdk.NewInt(3), nil

	case backing_ratio.GT(sdk.NewInt(93)) || backing_ratio.Equal(sdk.NewInt(93)):
		return sdk.NewInt(2), nil

	case backing_ratio.GT(sdk.NewInt(85)) || backing_ratio.Equal(sdk.NewInt(85)):
		return sdk.NewInt(1), nil

	case backing_ratio.LT(sdk.NewInt(85)):
		return sdk.NewInt(0), nil

	default:
		return sdk.Int{}, types.ErrCalculateMintingFee
	}

}

func CalculateBurningFee(backing_ratio sdk.Int) (sdk.Int, error) {
	if backing_ratio.IsNil() {
		return sdk.Int{}, types.ErrCalculateBurningFee
	}
	switch {
	case backing_ratio.GT(sdk.NewInt(140)) || backing_ratio.Equal(sdk.NewInt(140)):
		return sdk.NewInt(0), nil

	case backing_ratio.GT(sdk.NewInt(120)) || backing_ratio.Equal(sdk.NewInt(120)):
		return sdk.NewInt(1), nil

	case backing_ratio.GT(sdk.NewInt(100)) || backing_ratio.Equal(sdk.NewInt(100)):
		return sdk.NewInt(2), nil

	case backing_ratio.GT(sdk.NewInt(93)) || backing_ratio.Equal(sdk.NewInt(93)):
		return sdk.NewInt(3), nil

	case backing_ratio.GT(sdk.NewInt(85)) || backing_ratio.Equal(sdk.NewInt(85)):
		return sdk.NewInt(10), nil

	case backing_ratio.LT(sdk.NewInt(85)):
		return sdk.Int{}, types.ErrBurnBlocked

	default:
		return sdk.Int{}, types.ErrCalculateBurningFee
	}

}
