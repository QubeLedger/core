package gmb_test

import (
	"testing"

	"github.com/QuadrateOrg/core/x/stable/gmb"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestCalculateBackingRatio(t *testing.T) {
	tests := []struct {
		afp         int64
		ar          int64
		qm          int64
		expectedRes int64
		expectedErr bool
	}{
		{
			95000,
			2000,
			18300,
			103,
			false,
		},
		{
			-95000,
			2000,
			18300,
			0,
			true,
		},
		{
			95000,
			-2000,
			18300,
			0,
			true,
		},
		{
			95000,
			2000,
			-18300,
			0,
			true,
		},
	}

	for _, tc := range tests {

		backing_ratio, err := gmb.CalculateBackingRatio(sdk.NewInt(tc.afp), sdk.NewInt(tc.ar), sdk.NewInt(tc.qm))

		if tc.expectedErr {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
			require.Equal(t, tc.expectedRes, backing_ratio.Int64())
		}
	}
}

func TestCalculateMintingFee(t *testing.T) {
	tests := []struct {
		backing_ratio sdk.Int
		expectedRes   int64
		expectedErr   bool
	}{
		{
			sdk.NewInt(100),
			3,
			false,
		},
		{
			sdk.NewInt(85),
			1,
			false,
		},
		{
			sdk.NewInt(141),
			0,
			true,
		},
		{
			sdk.NewInt(80),
			0,
			false,
		},
		{
			sdk.Int{},
			0,
			true,
		},
	}

	for _, tc := range tests {
		fee, err := gmb.CalculateMintingFee(tc.backing_ratio)
		if tc.expectedErr {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
			require.Equal(t, tc.expectedRes, fee.Int64())
		}
	}
}

func TestCalculateBurningFee(t *testing.T) {
	tests := []struct {
		backing_ratio sdk.Int
		expectedRes   int64
		expectedErr   bool
	}{
		{
			sdk.NewInt(100),
			2,
			false,
		},
		{
			sdk.NewInt(93),
			3,
			false,
		},
		{
			sdk.NewInt(84),
			0,
			true,
		},
		{
			sdk.NewInt(140),
			0,
			false,
		},
		{
			sdk.NewInt(141),
			0,
			false,
		},
		{
			sdk.NewInt(85),
			10,
			false,
		},
		{
			sdk.Int{},
			0,
			true,
		},
	}

	for _, tc := range tests {
		fee, err := gmb.CalculateBurningFee(tc.backing_ratio)

		if tc.expectedErr {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
			require.Equal(t, tc.expectedRes, fee.Int64())
		}
	}
}
