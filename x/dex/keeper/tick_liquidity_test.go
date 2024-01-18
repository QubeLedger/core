package keeper_test

import (
	"strconv"
	"testing"

	"github.com/QuadrateOrg/core/x/dex/keeper"
	"github.com/QuadrateOrg/core/x/dex/types"
	keepertest "github.com/QuadrateOrg/core/x/dex/utils/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func CreateNTickLiquidity(keeper *keeper.Keeper, ctx sdk.Context, n int) []*types.TickLiquidity {
	items := make([]*types.TickLiquidity, n)
	for i := range items {
		tick := types.TickLiquidity{
			Liquidity: &types.TickLiquidity_LimitOrderTranche{
				LimitOrderTranche: types.MustNewLimitOrderTranche(
					"TokenA",
					"TokenB",
					strconv.Itoa(i),
					int64(i),
					sdk.NewInt(10),
					sdk.NewInt(10),
					sdk.NewInt(10),
					sdk.NewInt(10),
				),
			},
		}
		keeper.SetLimitOrderTranche(ctx, tick.GetLimitOrderTranche())
		items[i] = &tick
	}

	return items
}

func TestTickLiquidityGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := CreateNTickLiquidity(keeper, ctx, 10)
	require.ElementsMatch(t,
		items,
		keeper.GetAllTickLiquidity(ctx),
	)
}
