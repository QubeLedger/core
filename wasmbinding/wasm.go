package wasmbinding

import (
	"github.com/CosmWasm/wasmd/x/wasm"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"

	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	tokenfactorykeeper "github.com/QubeLedger/core/x/tokenfactory/keeper"

	oraclekeeper "github.com/QubeLedger/core/x/oracle/keeper"
)

func RegisterCustomPlugins(
	bank *bankkeeper.BaseKeeper,
	tokenFactory *tokenfactorykeeper.Keeper,
	oracle *oraclekeeper.Keeper,
) []wasmkeeper.Option {
	wasmQueryPlugin := NewQueryPlugin(tokenFactory, oracle)

	queryPluginOpt := wasmkeeper.WithQueryPlugins(&wasmkeeper.QueryPlugins{
		Custom: CustomQuerier(wasmQueryPlugin),
	})
	messengerDecoratorOpt := wasmkeeper.WithMessageHandlerDecorator(
		CustomMessageDecorator(bank, tokenFactory),
	)

	return []wasm.Option{
		queryPluginOpt,
		messengerDecoratorOpt,
	}
}
