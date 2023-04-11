package oracle

import (
	"github.com/QuadrateOrg/core/x/oracle/keeper"
	"github.com/QuadrateOrg/core/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the acData
	for _, elem := range genState.AcDataList {
		k.SetAcData(ctx, elem)
	}

	// Set acData count
	k.SetAcDataCount(ctx, genState.AcDataCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.AcDataList = k.GetAllAcData(ctx)
	genesis.AcDataCount = k.GetAcDataCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
