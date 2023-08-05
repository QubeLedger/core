package keeper

import (
	"context"

	"github.com/QuadrateOrg/core/x/stable/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.MsgServer = Keeper{}

func (k Keeper) Mint(goCtx context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	pairId, err := k.GetPairIdMint(msg.AmountInt, msg.DenomOut)
	pair, found := k.GetPairByPairID(ctx, pairId)
	if !found {
		return nil, types.ErrPairNotFound
	}

	err, amountOut := k.ExecuteMint(ctx, msg, pair)
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

	return &types.MsgMintResponse{
		Creator:   msg.Creator,
		AmountInt: msg.AmountInt,
		AmountOut: amountOut.String(),
	}, nil
}

func (k Keeper) Burn(goCtx context.Context, msg *types.MsgBurn) (*types.MsgBurnResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	pairId, err := k.GetPairIdBurn(msg.AmountInt, msg.DenomOut)
	pair, found := k.GetPairByPairID(ctx, pairId)
	if !found {
		return nil, types.ErrPairNotFound
	}
	err, amountOut := k.ExecuteBurn(ctx, msg, pair)

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

	return &types.MsgBurnResponse{
		Creator:   msg.Creator,
		AmountInt: msg.AmountInt,
		AmountOut: amountOut.String(),
	}, nil
}
