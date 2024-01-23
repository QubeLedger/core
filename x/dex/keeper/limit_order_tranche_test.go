package keeper_test

import (
	"strconv"
	"testing"

	"github.com/QuadrateOrg/core/x/dex/keeper"
	"github.com/QuadrateOrg/core/x/dex/types"
	keepertest "github.com/QuadrateOrg/core/x/dex/utils/keeper"
	"github.com/QuadrateOrg/core/x/dex/utils/nullify"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNLimitOrderTranches(
	keeper *keeper.Keeper,
	ctx sdk.Context,
	n int,
) []*types.LimitOrderTranche {
	items := make([]*types.LimitOrderTranche, n)
	for i := range items {
		items[i] = types.MustNewLimitOrderTranche(
			"TokenA",
			"TokenB",
			strconv.Itoa(i),
			int64(i),
			sdk.ZeroInt(),
			sdk.ZeroInt(),
			sdk.ZeroInt(),
			sdk.ZeroInt(),
		)
		keeper.SetLimitOrderTranche(ctx, items[i])
	}

	return items
}

func TestGetLimitOrderTranche(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNLimitOrderTranches(keeper, ctx, 10)
	for _, item := range items {
		rst := keeper.GetLimitOrderTranche(ctx, item.Key)
		require.NotNil(t, rst)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestRemoveLimitOrderTranche(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNLimitOrderTranches(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveLimitOrderTranche(ctx, item.Key)
		rst := keeper.GetLimitOrderTranche(ctx, item.Key)
		require.Nil(t, rst)
	}
}
