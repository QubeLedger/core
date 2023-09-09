package keeper

import (
	"github.com/QuadrateOrg/core/x/grow/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) ExecuteCreateLiqPosition(ctx sdk.Context, msg *types.MsgCreateLiquidationPosition, LendAsset types.LendAsset) (error, string) {
	amountInCoins, err := sdk.ParseCoinsNormalized(msg.AmountIn)
	if err != nil {
		return err, ""
	}

	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return err, ""
	}

	if err := CheckCoinDenom(amountInCoins, types.DefaultDenom); err != nil {
		return err, ""
	}

	premium, err := k.ParseAndCheckPremium(msg.Premium)
	if err != nil {
		return err, ""
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, amountInCoins)
	if err != nil {
		return err, ""
	}

	liquidatorPositionId := k.GenerateLiquidatorPositionId(creator.String(), types.DefaultDenom, amountInCoins.String(), msg.Premium)

	liqPosition := types.LiquidatorPosition{
		LiquidatorPositionId: liquidatorPositionId,
		BorrowAssetId:        LendAsset.AssetMetadata.Name,
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

	amountInt := amountCoins.AmountOf(types.DefaultDenom)

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creator, amountCoins)
	if err != nil {
		return err, sdk.Coin{}
	}

	k.RemoveLiquidatorPosition(ctx, liqPosition.Id)

	return err, sdk.NewCoin(types.DefaultDenom, amountInt)
}
