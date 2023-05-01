package wasmbinding

import (
	"github.com/CosmWasm/wasmd/x/wasm"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"

	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	interchainqueriesmodulekeeper "github.com/QuadrateOrg/core/x/interchainqueries/keeper"
	interchaintransactionsmodulekeeper "github.com/QuadrateOrg/core/x/interchaintxs/keeper"
	tokenfactorykeeper "github.com/QuadrateOrg/core/x/tokenfactory/keeper"
	transfer "github.com/QuadrateOrg/core/x/transfer/keeper"
	evmkeeper "github.com/evmos/ethermint/x/evm/keeper"
)

func RegisterCustomPlugins(
	bank *bankkeeper.BaseKeeper,
	tokenFactory *tokenfactorykeeper.Keeper,
	ictxKeeper *interchaintransactionsmodulekeeper.Keeper,
	icqKeeper *interchainqueriesmodulekeeper.Keeper,
	transfer transfer.KeeperTransferWrapper,
	evm evmkeeper.Keeper,
) []wasmkeeper.Option {
	wasmQueryPlugin := NewQueryPlugin(tokenFactory, ictxKeeper, icqKeeper, evm)

	queryPluginOpt := wasmkeeper.WithQueryPlugins(&wasmkeeper.QueryPlugins{
		Custom: CustomQuerier(wasmQueryPlugin),
	})
	messengerDecoratorOpt := wasmkeeper.WithMessageHandlerDecorator(
		CustomMessageDecorator(bank, tokenFactory, ictxKeeper, icqKeeper, transfer, evm),
	)

	return []wasm.Option{
		queryPluginOpt,
		messengerDecoratorOpt,
	}
}
