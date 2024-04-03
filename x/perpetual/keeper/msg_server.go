package keeper

import (
	"context"

	"github.com/QuadrateOrg/core/x/perpetual/types"
)

var _ types.MsgServer = Keeper{}

func (k Keeper) Open(goCtx context.Context, msg *types.MsgOpen) (*types.MsgOpenResponse, error) {
	return &types.MsgOpenResponse{}, nil
}

func (k Keeper) Close(goCtx context.Context, msg *types.MsgClose) (*types.MsgCloseResponse, error) {
	return &types.MsgCloseResponse{}, nil
}
