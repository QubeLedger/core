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

	backing_ratio := (afp.Mul(ar).Quo(qm).Quo(sdk.NewInt(100)))

	if backing_ratio.IsNegative() {
		return sdk.Int{}, types.ErrBackingRatioNegative
	}
	if backing_ratio.IsNil() {
		return sdk.Int{}, types.ErrBackingRatioNil
	}
	return backing_ratio, nil
}

func CalculateMintingFee(backing_ratio sdk.Int) (sdk.Int, bool, error) {
	if backing_ratio.IsNil() {
		return sdk.Int{}, false, types.ErrCalculateMintingFee
	}
	switch {
	case backing_ratio.GT(sdk.NewInt(int64(120))) || backing_ratio.Equal(sdk.NewInt(int64(120))):
		return sdk.Int{}, false, nil

	case backing_ratio.GT(sdk.NewInt(int64(110))) || backing_ratio.Equal(sdk.NewInt(int64(110))):
		return sdk.NewInt(int64(20)), true, nil

	case backing_ratio.GT(sdk.NewInt(int64(93))) || backing_ratio.Equal(sdk.NewInt(int64(93))):
		return sdk.NewInt(int64(3)), true, nil

	case backing_ratio.GT(sdk.NewInt(int64(85))) || backing_ratio.Equal(sdk.NewInt(int64(90))):
		return sdk.NewInt(int64(3)), true, nil

	case sdk.NewInt(int64(85)).GT(backing_ratio):
		return sdk.NewInt(int64(0)), true, nil

	default:
		return sdk.Int{}, false, types.ErrCalculateMintingFee
	}

}

func CalculateBurningFee(backing_ratio sdk.Int) (sdk.Int, bool, error) {
	if backing_ratio.IsNil() {
		return sdk.Int{}, false, types.ErrCalculateBurningFee
	}
	switch {
	case backing_ratio.GT(sdk.NewInt(int64(120))) || backing_ratio.Equal(sdk.NewInt(int64(120))):
		return sdk.NewInt(int64(0)), true, nil

	case backing_ratio.GT(sdk.NewInt(int64(110))) || backing_ratio.Equal(sdk.NewInt(int64(110))):
		return sdk.NewInt(int64(0)), true, nil

	case backing_ratio.GT(sdk.NewInt(int64(93))) || backing_ratio.Equal(sdk.NewInt(int64(93))):
		return sdk.NewInt(int64(3)), true, nil

	case backing_ratio.GT(sdk.NewInt(int64(85))) || backing_ratio.Equal(sdk.NewInt(int64(85))):
		return sdk.NewInt(int64(3)), true, nil

	case sdk.NewInt(int64(85)).GT(backing_ratio):
		return sdk.NewInt(int64(0)), true, nil

	default:
		return sdk.Int{}, false, types.ErrCalculateBurningFee
	}

}
