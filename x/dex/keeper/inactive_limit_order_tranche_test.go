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

func createNInactiveLimitOrderTranche(
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
		keeper.SetInactiveLimitOrderTranche(ctx, items[i])
	}

	return items
}

func TestInactiveLimitOrderTrancheGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNInactiveLimitOrderTranche(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetInactiveLimitOrderTranche(ctx, item.Key)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestInactiveLimitOrderTrancheRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNInactiveLimitOrderTranche(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveInactiveLimitOrderTranche(ctx, item.Key)
		_, found := keeper.GetInactiveLimitOrderTranche(ctx, item.Key)
		require.False(t, found)
	}
}

func TestInactiveLimitOrderTrancheGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNInactiveLimitOrderTranche(keeper, ctx, 10)
	pointerItems := make([]*types.LimitOrderTranche, len(items))
	for i := range items {
		pointerItems[i] = items[i]
	}
	require.ElementsMatch(t,
		pointerItems,
		keeper.GetAllInactiveLimitOrderTranche(ctx),
	)
}
