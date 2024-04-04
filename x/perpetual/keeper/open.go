package keeper

import (
	"github.com/QuadrateOrg/core/x/perpetual/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) OpenPosition(ctx sdk.Context, msg *types.MsgOpen) error {
	vault, found := k.GetVaultByVaultId(ctx, msg.TradingAsset)
	if !found {
		return types.ErrVaultNotFound
	}

	sender, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return err
	}

	if position := k.CheckSamePosition(ctx, msg); position != nil {
		// if position already create - update
	} else {

	}

	collateral_coins, err := sdk.ParseCoinsNormalized(msg.Collateral)
	if err != nil {
		return err
	}

	err = k.CreateNewPosition(ctx, msg, vault)
	if err != nil {
		return err
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, collateral_coins)
	if err != nil {
		return err
	}

	event := sdk.NewEvent(types.EventOpen,
		sdk.NewAttribute("id", k.GenerateTraderPositionId(msg.Creator, vault.AmountXMetadata.Base, msg.TradingAsset, msg.TradeType)),
		sdk.NewAttribute("Creator", msg.Creator),
		sdk.NewAttribute("collateral", msg.Collateral),
		sdk.NewAttribute("leverage", msg.Leverage.String()),
	)
	ctx.EventManager().EmitEvent(event)
	return nil
}
