package keeper_test

import (
	"testing"

	"github.com/QuadrateOrg/core/x/dex/keeper"
	"github.com/QuadrateOrg/core/x/dex/types"
	keepertest "github.com/QuadrateOrg/core/x/dex/utils/keeper"
	"github.com/QuadrateOrg/core/x/dex/utils/nullify"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNPoolReserves(k *keeper.Keeper, ctx sdk.Context, n int) []*types.PoolReserves {
	items := make([]*types.PoolReserves, n)
	pools := createNPools(k, ctx, n)
	for i, pool := range pools {
		items[i] = pool.LowerTick0
	}
	return items
}

func TestGetPoolReserves(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNPoolReserves(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPoolReserves(ctx, item.Key)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestRemovePoolReserves(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNPoolReserves(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePoolReserves(ctx, item.Key)
		_, found := keeper.GetPoolReserves(ctx, item.Key)
		require.False(t, found)
	}
}
