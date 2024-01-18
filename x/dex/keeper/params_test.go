package keeper_test

import (
	"testing"

	"github.com/QuadrateOrg/core/x/dex/types"
	testkeeper "github.com/QuadrateOrg/core/x/dex/utils/keeper"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.DexKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}

func TestValidateParams(t *testing.T) {
	goodFees := []uint64{1, 2, 3, 4, 5, 200}
	require.NoError(t, types.Params{FeeTiers: goodFees}.Validate())

	badFees := []uint64{1, 2, 3, 3}
	require.Error(t, types.Params{FeeTiers: badFees}.Validate())
}
