package keeper

import (
	"context"

	"github.com/QuadrateOrg/core/x/perpetual/types"
)

var _ types.MsgServer = Keeper{}

func (k Keeper) PerpetualDeposit(goCtx context.Context, msg *types.MsgPerpetualDeposit) (*types.MsgPerpetualDepositResponse, error) {
	return &types.MsgPerpetualDepositResponse{}, nil
}

func (k Keeper) PerpetualWithdraw(goCtx context.Context, msg *types.MsgPerpetualWithdraw) (*types.MsgPerpetualWithdrawResponse, error) {
	return &types.MsgPerpetualWithdrawResponse{}, nil
}

func (k Keeper) CreatePosition(goCtx context.Context, msg *types.MsgCreatePosition) (*types.MsgCreatePositionResponse, error) {
	return &types.MsgCreatePositionResponse{}, nil
}

func (k Keeper) ClosePosition(goCtx context.Context, msg *types.MsgClosePosition) (*types.MsgClosePositionResponse, error) {
	return &types.MsgClosePositionResponse{}, nil
}
