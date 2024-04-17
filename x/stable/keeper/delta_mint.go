package keeper

import (
	"time"

	dextypes "github.com/QuadrateOrg/core/x/dex/types"
	perptypes "github.com/QuadrateOrg/core/x/perpetual/types"
	"github.com/QuadrateOrg/core/x/stable/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) DeltaMint(ctx sdk.Context, msg *types.MsgMint, pair types.Pair) (error, sdk.Coin) {

	sender, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return err, sdk.Coin{}
	}

	amountIntCoins, err := sdk.ParseCoinsNormalized(msg.AmountIn)
	if err != nil {
		return err, sdk.Coin{}
	}

	err = VerificationMintDenomCoins(amountIntCoins, pair)
	if err != nil {
		return err, sdk.Coin{}
	}

	amountIn := amountIntCoins.AmountOf(pair.AmountInMetadata.Base)
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, amountIntCoins)
	if err != nil {
		return err, sdk.Coin{}
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, k.GetSystemModuleAccount(ctx).GetAddress(), amountIntCoins)
	if err != nil {
		return err, sdk.Coin{}
	}

	maxAmountOut := sdk.NewInt(0)
	_, _, _, coinOutSwap, err := k.dexKeeper.PlaceLimitOrderCore(
		ctx.Context(),
		pair.AmountInMetadata.Base,
		pair.TokenStakeMetadata.Base,
		amountIn,
		0, // TODO
		dextypes.LimitOrderType_FILL_OR_KILL,
		&time.Time{},
		&maxAmountOut,
		k.GetSystemModuleAccount(ctx).GetAddress(),
		k.GetSystemModuleAccount(ctx).GetAddress(),
	)

	if err != nil {
		return err, sdk.Coin{}
	}

	amountOutAfterSwap := coinOutSwap.Amount

	err = k.perpetualKeeper.OpenPosition(ctx, perptypes.NewMsgOpen(
		k.GetSystemModuleAccount(ctx).GetAddress().String(),
		perptypes.PerpetualTradeType_PERPETUAL_SHORT_POSITION,
		sdk.NewDec(1),
		k.perpetualKeeper.GenerateVaultIdHash(pair.TokenStakeMetadata.Base, pair.TokenYMetadata.Base),
		sdk.NewCoins(sdk.NewCoin(
			coinOutSwap.Denom,
			amountOutAfterSwap.QuoRaw(2),
		)).String(),
	))
	if err != nil {
		return err, sdk.Coin{}
	}

	price, err := k.oracleKeeper.GetExchangeRate(ctx, pair.OracleAssetId)
	if err != nil {
		return err, sdk.Coin{}
	}

	amountOutToMint := (amountIn.ToDec().Mul(price)).RoundInt()

	amountOut := sdk.NewCoin(pair.AmountOutMetadata.DenomUnits[0].Denom, amountOutToMint)
	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(amountOut))
	if err != nil {
		return err, sdk.Coin{}
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, sdk.NewCoins(amountOut))
	if err != nil {
		return err, sdk.Coin{}
	}

	return nil, amountOut
}
