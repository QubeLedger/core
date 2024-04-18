package types

import (
	context "context"
	"time"

	dextypes "github.com/QuadrateOrg/core/x/dex/types"
	math_utils "github.com/QuadrateOrg/core/x/dex/utils/math"
	perptypes "github.com/QuadrateOrg/core/x/perpetual/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	GetModuleAccount(ctx sdk.Context, moduleName string) types.ModuleAccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromModuleToModule(ctx sdk.Context, senderModule, recipientModule string, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	HasSupply(ctx sdk.Context, denom string) bool
	IsSendEnabledCoin(ctx sdk.Context, coin sdk.Coin) bool
	BlockedAddr(addr sdk.AccAddress) bool
	GetDenomMetaData(ctx sdk.Context, denom string) (banktypes.Metadata, bool)
	SetDenomMetaData(ctx sdk.Context, denomMetaData banktypes.Metadata)
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
}

type OracleKeeper interface {
	GetExchangeRate(ctx sdk.Context, denom string) (sdk.Dec, error)
}

type DexKeeper interface {
	PlaceLimitOrderCore(
		goCtx context.Context,
		tokenIn string,
		tokenOut string,
		amountIn sdk.Int,
		tickIndexInToOut int64,
		orderType dextypes.LimitOrderType,
		goodTil *time.Time,
		maxAmountOut *sdk.Int,
		callerAddr sdk.AccAddress,
		receiverAddr sdk.AccAddress,
	) (trancheKey string, totalInCoin sdk.Coin, swapInCoin sdk.Coin, swapOutCoin sdk.Coin, err error)

	MultiHopSwapCore(
		ctx sdk.Context,
		amountIn sdk.Int,
		routes []*dextypes.MultiHopRoute,
		exitLimitPrice math_utils.PrecDec,
		pickBestRoute bool,
		callerAddr sdk.AccAddress,
		receiverAddr sdk.AccAddress,
	) (coinOut sdk.Coin, err error)
}

type PerpetualKeeper interface {
	GenerateVaultIdHash(denom1 string, denom2 string) string
	OpenPosition(ctx sdk.Context, msg *perptypes.MsgOpen) error
	ClosePosition(ctx sdk.Context, msg *perptypes.MsgClose) error
}
