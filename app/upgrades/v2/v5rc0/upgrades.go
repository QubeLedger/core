package v5rc0

import (
	"fmt"

	growmodulekeeper "github.com/QuadrateOrg/core/x/grow/keeper"
	"github.com/QuadrateOrg/core/x/grow/types"
	growmoduletypes "github.com/QuadrateOrg/core/x/grow/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

/* #nosec */
func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	growkeeper growmodulekeeper.Keeper,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		migrations, err := mm.RunMigrations(ctx, configurator, fromVM)
		if err != nil {
			return nil, err
		}
		// update already created asset

		// asset_id already created asset
		asset_id := growkeeper.GenerateAssetIdHash("factory/qube1t2ydw7r4asmk74ymuvykjshdzun8dxye0az5wz/uusdc")

		asset, found := growkeeper.GetAssetByAssetId(ctx, asset_id)
		if !found {
			return nil, growmoduletypes.ErrAssetNotFound
		}

		asset.Type = "stable"
		growkeeper.SetAsset(ctx, asset)

		ctx.Logger().Info(fmt.Sprintf("qLabs: set asset in x/grow with AssetId: %v", asset.AssetId))

		new_asset := types.Asset{
			AssetId: growkeeper.GenerateAssetIdHash("factory/qube1t2ydw7r4asmk74ymuvykjshdzun8dxye0az5wz/wbtc"),
			AssetMetadata: banktypes.Metadata{
				Description: "Wrapped Bitcoin (WBTC) is an ERC20 token backed 1:1 with Bitcoin. Completely transparent. 100% verifiable. Community led. (TESTNET)",
				DenomUnits: []*banktypes.DenomUnit{
					{
						Denom:    "factory/qube1t2ydw7r4asmk74ymuvykjshdzun8dxye0az5wz/wbtc",
						Exponent: uint32(0),
						Aliases:  []string{"wbtc"},
					},
					{
						Denom:    "wBTC",
						Exponent: uint32(6),
					},
				},
				Base:    "factory/qube1t2ydw7r4asmk74ymuvykjshdzun8dxye0az5wz/wbtc",
				Display: "wBTC",
				Name:    "Wrapped Bitcoin",
				Symbol:  "WBTC",
			},
			OracleAssetId:           "BTC",
			ProvideValue:            uint64(10000),
			CollectivelyBorrowValue: uint64(0),
			Type:                    "volatile",
		}

		new_id := growkeeper.AppendAsset(ctx, new_asset)
		if new_id != uint64(1) {
			return nil, types.ErrAssetNotFound
		}

		return migrations, nil
	}
}
