package keeper

import (
	"github.com/QuadrateOrg/core/x/grow/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) ExecuteCreateLiqPosition(ctx sdk.Context, msg *types.MsgCreateLiquidationPosition) (error, string) {
	return nil, ""
}

func (k Keeper) ExecuteCloseLiqPosition(ctx sdk.Context, msg *types.MsgCloseLiquidationPosition) (error, string) {
	return nil, ""
}
