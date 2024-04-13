package keeper

import (
	"github.com/QuadrateOrg/core/x/perpetual/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) CheckSamePosition(ctx sdk.Context, msg *types.MsgOpen, vault types.Vault) *types.TradePosition {
	positions := k.GetAllPositions(ctx)
	for _, position := range positions {
		if position.Creator == msg.Creator && position.TradingAsset == msg.TradingAsset && position.TradePositionId == k.GenerateTraderPositionId(msg.Creator, vault.AmountXMetadata.Base, msg.TradingAsset, msg.TradeType, msg.Leverage) {
			return &position
		}
	}

	return nil
}
