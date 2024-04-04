package keeper

import (
	"context"

	"github.com/QuadrateOrg/core/x/perpetual/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.MsgServer = Keeper{}

func (k Keeper) Open(goCtx context.Context, msg *types.MsgOpen) (*types.MsgOpenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.OpenPosition(ctx, msg)
	if err != nil {
		return nil, err
	}

	return &types.MsgOpenResponse{}, nil
}

func (k Keeper) Close(goCtx context.Context, msg *types.MsgClose) (*types.MsgCloseResponse, error) {
	return &types.MsgCloseResponse{}, nil
}
