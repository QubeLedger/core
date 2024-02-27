package keeper

import (
	"github.com/QuadrateOrg/core/x/grow/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) ExecuteCreateLiqPosition(ctx sdk.Context, msg *types.MsgOpenLiquidationPosition, Asset types.Asset) (error, string) {
	amountInCoins, err := sdk.ParseCoinsNormalized(msg.AmountIn)
	if err != nil {
		return err, ""
	}

	if k.AddressEmptyCheck(ctx) {
		return types.ErrReserveAddressEmpty, ""
	}

	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return err, ""
	}

	DenomIn := amountInCoins.GetDenomByIndex(0)

	asset, found := k.GetAssetByAssetId(ctx, k.GenerateAssetIdHash(DenomIn))
	if !found {
		return types.ErrAssetNotFound, ""
	}

	premium, err := k.ParseAndCheckPremium(msg.Premium)
	if err != nil {
		return err, ""
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, amountInCoins)
	if err != nil {
		return err, ""
	}

	liquidatorPositionId := k.GenerateLiquidatorPositionId(creator.String(), DenomIn, Asset.OracleAssetId, amountInCoins.String(), msg.Premium)

	liqPosition := types.LiquidatorPosition{
		LiquidatorPositionId: liquidatorPositionId,
		ProvidedAssetId:      asset.OracleAssetId,
		WantAssetId:          Asset.OracleAssetId,
		Liquidator:           creator.String(),
		Amount:               amountInCoins.String(),
		Premium:              premium.Uint64(),
	}

	k.AppendLiquidatorPosition(ctx, liqPosition)

	return nil, liquidatorPositionId
}

func (k Keeper) ExecuteCloseLiqPosition(ctx sdk.Context, msg *types.MsgCloseLiquidationPosition) (error, sdk.Coin) {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return err, sdk.Coin{}
	}

	if k.AddressEmptyCheck(ctx) {
		return types.ErrReserveAddressEmpty, sdk.Coin{}
	}

	liqPosition, found := k.GetLiquidatorPositionByLiquidatorPositionId(ctx, msg.LiquidatorPositionId)
	if !found {
		return types.ErrLiqPositionNotFound, sdk.Coin{}
	}

	err = k.CheckLiquidator(creator, liqPosition)
	if err != nil {
		return err, sdk.Coin{}
	}

	amountCoins, err := sdk.ParseCoinsNormalized(liqPosition.Amount)
	if err != nil {
		return err, sdk.Coin{}
	}

	amountInt := amountCoins.AmountOf(amountCoins.GetDenomByIndex(0))

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creator, amountCoins)
	if err != nil {
		return err, sdk.Coin{}
	}

	k.RemoveLiquidatorPosition(ctx, liqPosition.Id)

	return nil, sdk.NewCoin(amountCoins.GetDenomByIndex(0), amountInt)
}
