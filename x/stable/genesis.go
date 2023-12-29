package stable

import (
	"github.com/QuadrateOrg/core/x/stable/keeper"
	"github.com/QuadrateOrg/core/x/stable/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.SetPort(ctx, genState.PortId)
	if !k.IsBound(ctx, genState.PortId) {
		err := k.BindPort(ctx, genState.PortId)
		if err != nil {
			panic("could not claim port capability: " + err.Error())
		}
	}
	k.SetParams(ctx, genState.Params)

	for _, pair := range genState.Pairs {
		k.AppendPair(ctx, pair)
	}
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.PortId = k.GetPort(ctx)

	genesis.Pairs = k.GetAllPair(ctx)

	return genesis
}
