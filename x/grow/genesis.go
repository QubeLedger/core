package grow

import (
	"github.com/QuadrateOrg/core/x/grow/keeper"
	"github.com/QuadrateOrg/core/x/grow/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the loan
	for _, elem := range genState.GTokenPairList {
		k.AppendPair(ctx, elem)
	}
	k.SetParams(ctx, genState.Params)
	k.SetRealRate(ctx, sdk.NewIntFromUint64(genState.RealRate))
	k.SetGrowStakingReserveAddress(ctx, sdk.AccAddress(genState.GrowStakingReserveAddress))
	k.SetUSQReserveAddress(ctx, sdk.AccAddress(genState.USQReserveAddress))
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.GTokenPairList = k.GetAllPair(ctx)
	genesis.RealRate = k.GetRealRate(ctx).Uint64()

	genesis.GrowStakingReserveAddress = string(k.GetGrowStakingReserveAddress(ctx))
	genesis.USQReserveAddress = string(k.GetUSQReserveAddress(ctx))

	return genesis
}
