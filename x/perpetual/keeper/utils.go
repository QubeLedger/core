package keeper

import (
	"github.com/QuadrateOrg/core/x/perpetual/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) CheckSamePosition(ctx sdk.Context, msg *types.MsgOpen) *types.TradePosition {
	positions := k.GetAllPositions(ctx)
	for _, position := range positions {
		if position.Creator == msg.Creator {
			return &position
		}
	}

	return nil
}
