package gmb

import (
	"github.com/QuadrateOrg/core/x/stable/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func CalculateBackingRatio(afp sdk.Int, ar sdk.Int, qm sdk.Int) sdk.Int {
	if afp.IsNegative() {
		panic(types.ErrAfpNegative)
	}
	if ar.IsNegative() {
		panic(types.ErrArNegative)
	}
	if qm.IsNegative() {
		panic(types.ErrQmNegative)
	}

	// ERROR
	// backing_ratio always 100 or 200
	backing_ratio := afp.Mul(ar).Quo(qm).Quo(sdk.NewInt(100)) //((afp.Mul(ar)).Quo(qm)).Mul(sdk.NewInt(int64(100)))
	if backing_ratio.IsNegative() {
		panic(types.ErrBackingRatioNegative)
	}
	if backing_ratio.IsNil() {
		panic(types.ErrBackingRatioNil)
	}
	return backing_ratio
}

func CalculateMintingFee(backing_ratio sdk.Int) (sdk.Int, bool, error) {
	switch {
	case backing_ratio.GT(sdk.NewInt(int64(225))) || backing_ratio.Equal(sdk.NewInt(int64(225))):
		return sdk.Int{}, false, nil

	case backing_ratio.GT(sdk.NewInt(int64(150))) || backing_ratio.Equal(sdk.NewInt(int64(150))):
		return sdk.NewInt(int64(20)), true, nil

	case backing_ratio.GT(sdk.NewInt(int64(100))) || backing_ratio.Equal(sdk.NewInt(int64(100))):
		return sdk.NewInt(int64(3)), true, nil

	case backing_ratio.GT(sdk.NewInt(int64(90))) || backing_ratio.Equal(sdk.NewInt(int64(90))):
		return sdk.NewInt(int64(3)), true, nil

	case backing_ratio.GT(sdk.NewInt(int64(75))) || backing_ratio.Equal(sdk.NewInt(int64(75))):
		return sdk.NewInt(int64(0)), true, nil

	case backing_ratio.GTE(sdk.NewInt(int64(75))):
		return sdk.NewInt(int64(0)), true, nil

	default:
		return sdk.Int{}, false, types.ErrCalculateMintingFee
	}

}

func CalculateBurningFee(backing_ratio sdk.Int) (sdk.Int, bool, error) {
	switch {
	case backing_ratio.GT(sdk.NewInt(int64(225))) || backing_ratio.Equal(sdk.NewInt(int64(225))):
		return sdk.NewInt(int64(0)), true, nil

	case backing_ratio.GT(sdk.NewInt(int64(150))) || backing_ratio.Equal(sdk.NewInt(int64(150))):
		return sdk.NewInt(int64(0)), true, nil

	case backing_ratio.GT(sdk.NewInt(int64(100))) || backing_ratio.Equal(sdk.NewInt(int64(100))):
		return sdk.NewInt(int64(3)), true, nil

	case backing_ratio.GT(sdk.NewInt(int64(90))) || backing_ratio.Equal(sdk.NewInt(int64(90))):
		return sdk.NewInt(int64(3)), true, nil

	case backing_ratio.GT(sdk.NewInt(int64(75))) || backing_ratio.Equal(sdk.NewInt(int64(75))):
		return sdk.NewInt(int64(20)), true, nil

	case backing_ratio.GTE(sdk.NewInt(int64(75))):
		return sdk.Int{}, false, nil

	default:
		return sdk.Int{}, false, types.ErrCalculateMintingFee
	}

}
