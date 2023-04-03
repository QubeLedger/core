package keeper

import (
	"context"

	"github.com/0xknstntn/quadrate/x/converter/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Cw20Erc20(goCtx context.Context, msg *types.MsgCw20Erc20) (*types.MsgCw20Erc20Response, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCw20Erc20Response{}, nil
}
