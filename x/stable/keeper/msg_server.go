package keeper

import (
	"context"

	"github.com/QuadrateOrg/core/x/stable/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.MsgServer = Keeper{}

func (k Keeper) MintUsq(goCtx context.Context, msg *types.MsgMintUsq) (*types.MsgMintUsqResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.UpdateAtomPrice(ctx)
	atomPrice := k.GetAtomPrice(ctx)

	err, amountOut := k.ExecuteMint(ctx, msg, atomPrice)

	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.AttributeKeyActionMint),
		),
	})

	return &types.MsgMintUsqResponse{
		Creator:   msg.Creator,
		AmountInt: msg.Amount,
		AmountOut: amountOut.String(),
	}, nil
}

func (k Keeper) BurnUsq(goCtx context.Context, msg *types.MsgBurnUsq) (*types.MsgBurnUsqResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.UpdateAtomPrice(ctx)
	atomPrice := k.GetAtomPrice(ctx)

	err, amountOut := k.ExecuteBurn(ctx, msg, atomPrice)

	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.AttributeKeyActionBurn),
		),
	})

	return &types.MsgBurnUsqResponse{
		Creator:   msg.Creator,
		AmountInt: msg.Amount,
		AmountOut: amountOut.String(),
	}, nil
}
