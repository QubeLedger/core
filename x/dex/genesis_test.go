package dex_test

/* TODO */
/*
import (
	"testing"

	"github.com/QuadrateOrg/core/x/dex"
	"github.com/QuadrateOrg/core/x/dex/types"
	keepertest "github.com/QuadrateOrg/core/x/dex/utils/keeper"
	"github.com/QuadrateOrg/core/x/dex/utils/nullify"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		LimitOrderTrancheUserList: []*types.LimitOrderTrancheUser{
			{
				TradePairID: &types.TradePairID{
					TakerDenom: "TokenA",
					MakerDenom: "TokenB",
				},
				TickIndexTakerToMaker: 1,
				TrancheKey:            "0",
				Address:               "fakeAddr",
				SharesOwned:           sdk.NewInt(10),
				SharesWithdrawn:       sdk.NewInt(0),
				SharesCancelled:       sdk.NewInt(0),
			},
			{
				TradePairID: &types.TradePairID{
					TakerDenom: "TokenB",
					MakerDenom: "TokenA",
				},
				TickIndexTakerToMaker: 20,
				TrancheKey:            "0",
				Address:               "fakeAddr",
				SharesOwned:           sdk.NewInt(10),
				SharesWithdrawn:       sdk.NewInt(0),
				SharesCancelled:       sdk.NewInt(0),
			},
		},
		TickLiquidityList: []*types.TickLiquidity{
			{
				Liquidity: &types.TickLiquidity_LimitOrderTranche{
					LimitOrderTranche: types.MustNewLimitOrderTranche(
						"TokenB",
						"TokenA",
						"0",
						0,
						sdk.ZeroInt(),
						sdk.ZeroInt(),
						sdk.ZeroInt(),
						sdk.ZeroInt(),
					),
				},
			},
			{
				Liquidity: &types.TickLiquidity_LimitOrderTranche{
					LimitOrderTranche: types.MustNewLimitOrderTranche(
						"TokenB",
						"TokenA",
						"0",
						0,
						sdk.ZeroInt(),
						sdk.ZeroInt(),
						sdk.ZeroInt(),
						sdk.ZeroInt(),
					),
				},
			},
		},
		InactiveLimitOrderTrancheList: []*types.LimitOrderTranche{
			{
				Key: &types.LimitOrderTrancheKey{
					TradePairID: &types.TradePairID{
						TakerDenom: "TokenA",
						MakerDenom: "TokenB",
					},
					TickIndexTakerToMaker: 0,
					TrancheKey:            "0",
				},
			},
			{
				Key: &types.LimitOrderTrancheKey{
					TradePairID: &types.TradePairID{
						TakerDenom: "TokenA",
						MakerDenom: "TokenB",
					},
					TickIndexTakerToMaker: 1,
					TrancheKey:            "1",
				},
			},
		},
		PoolMetadataList: []types.PoolMetadata{
			{
				ID: 0,
			},
			{
				ID: 1,
			},
		},
		PoolCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DexKeeper(t)
	dex.InitGenesis(ctx, *k, genesisState)
	got := dex.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.LimitOrderTrancheUserList, got.LimitOrderTrancheUserList)
	require.ElementsMatch(t, genesisState.TickLiquidityList, got.TickLiquidityList)
	require.ElementsMatch(
		t,
		genesisState.InactiveLimitOrderTrancheList,
		got.InactiveLimitOrderTrancheList,
	)
	require.ElementsMatch(t, genesisState.PoolMetadataList, got.PoolMetadataList)
	require.Equal(t, genesisState.PoolCount, got.PoolCount)
	// this line is used by starport scaffolding # genesis/test/assert
}

*/
