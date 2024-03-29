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

func createNPoolMetadata(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.PoolMetadata {
	items := make([]types.PoolMetadata, n)
	for i := range items {
		items[i].ID = uint64(i)
		keeper.SetPoolMetadata(ctx, items[i])
	}

	return items
}

func TestPoolMetadataGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNPoolMetadata(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetPoolMetadata(ctx, item.ID)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestPoolMetadataRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNPoolMetadata(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePoolMetadata(ctx, item.ID)
		_, found := keeper.GetPoolMetadata(ctx, item.ID)
		require.False(t, found)
	}
}

func TestPoolMetadataGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNPoolMetadata(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPoolMetadata(ctx)),
	)
}
