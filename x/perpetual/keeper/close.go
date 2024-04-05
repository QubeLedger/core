package keeper

import (
	"github.com/QuadrateOrg/core/x/perpetual/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) ClosePosition(ctx sdk.Context, msg *types.MsgClose) error {

	position, found := k.GetPositionByPositionId(ctx, msg.Id)
	if !found {
		return types.ErrPositionNotFound
	}

	vault, found := k.GetVaultByVaultId(ctx, k.GenerateVaultIdHash(position.CollateralDenom, position.ReturnDenom))
	if !found {
		return types.ErrVaultNotFound
	}

	err := k.CloseOrDecreasePosition(ctx, msg, vault, position)
	if err != nil {
		return err
	}

	event := sdk.NewEvent(types.EventClose,
		sdk.NewAttribute("id", msg.Id),
		sdk.NewAttribute("Creator", msg.Creator),
		sdk.NewAttribute("Amount", msg.Amount.String()),
	)
	ctx.EventManager().EmitEvent(event)

	return nil
}
