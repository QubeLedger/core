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

	if position := k.CheckSamePosition(ctx, msg, vault); position != nil {
		err := k.IncreasePosition(ctx, msg, vault, *position)
		if err != nil {
			return err
		}
	} else {
		err := k.CreateNewPosition(ctx, msg, vault)
		if err != nil {
			return err
		}
	}

	event := sdk.NewEvent(types.EventOpen,
		sdk.NewAttribute("id", k.GenerateTraderPositionId(msg.Creator, vault.AmountXMetadata.Base, msg.TradingAsset, msg.TradeType, msg.Leverage)),
		sdk.NewAttribute("Creator", msg.Creator),
		sdk.NewAttribute("collateral", msg.Collateral),
		sdk.NewAttribute("leverage", msg.Leverage.String()),
	)
	ctx.EventManager().EmitEvent(event)
	return nil
}
