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
		expectedAllow bool
		expectedErr   bool
	}{
		{
			sdk.NewInt(100),
			3,
			true,
			false,
		},
		{
			sdk.NewInt(121),
			0,
			false,
			false,
		},
		{
			sdk.NewInt(84),
			0,
			true,
			false,
		},
		{
			sdk.Int{},
			0,
			false,
			true,
		},
	}

	for _, tc := range tests {
		fee, allow, err := gmb.CalculateMintingFee(tc.backing_ratio)

		if tc.expectedErr {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}

		if tc.expectedAllow {
			require.Equal(t, true, allow)
			require.Equal(t, tc.expectedRes, fee.Int64())
		}
	}
}

func TestCalculateBurningFee(t *testing.T) {
	tests := []struct {
		backing_ratio sdk.Int
		expectedRes   int64
		expectedAllow bool
		expectedErr   bool
	}{
		{
			sdk.NewInt(100),
			3,
			true,
			false,
		},
		{
			sdk.NewInt(121),
			0,
			true,
			false,
		},
		{
			sdk.NewInt(84),
			0,
			false,
			false,
		},
		{
			sdk.Int{},
			0,
			false,
			true,
		},
	}

	for _, tc := range tests {
		fee, allow, err := gmb.CalculateBurningFee(tc.backing_ratio)

		if tc.expectedErr {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}

		if tc.expectedAllow {
			require.Equal(t, true, allow)
			require.Equal(t, tc.expectedRes, fee.Int64())
		}
	}
}
