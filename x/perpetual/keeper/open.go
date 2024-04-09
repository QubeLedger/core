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

	if position := k.CheckSamePosition(ctx, msg); position != nil {
		// if position already create - update
	} else {
		err := k.CreateNewPosition(ctx, msg, vault)
		if err != nil {
			return err
		}
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
