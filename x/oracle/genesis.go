package oracle

import (
	"github.com/QubeLedger/core/x/oracle/keeper"
	"github.com/QubeLedger/core/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the price
	for _, elem := range genState.PriceList {
		k.SetPrice(ctx, elem)
	}

	// Set price count
	k.SetPriceCount(ctx, genState.PriceCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.PriceList = k.GetAllPrice(ctx)
	genesis.PriceCount = k.GetPriceCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
