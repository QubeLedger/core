package keeper

import (
	"context"

	"github.com/QuadrateOrg/core/x/stable/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.MsgServer = Keeper{}

//nolint:all
func (k Keeper) Mint(goCtx context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	pairId, err := k.GetPairIdMint(msg.AmountIn, msg.DenomOut)
	pair, found := k.GetPairByPairID(ctx, pairId)
	if !found {
		return nil, types.ErrPairNotFound
	}

	err = k.CheckMinAmount(msg.AmountIn, pair)
	if err != nil {
		return nil, err
	}

	var (
		mint_err  error
		amountOut sdk.Coin
	)

	switch pair.Model {
	case "gmb":
		mint_err, amountOut = k.GMB_ExecuteMint(ctx, msg, pair)
	}

	if mint_err != nil {
		return nil, mint_err
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
		AmountIn:  msg.AmountIn,
		AmountOut: amountOut.String(),
	}, nil
}

//nolint:all
func (k Keeper) Burn(goCtx context.Context, msg *types.MsgBurn) (*types.MsgBurnResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	pairId, err := k.GetPairIdBurn(msg.AmountIn, msg.DenomOut)
	pair, found := k.GetPairByPairID(ctx, pairId)
	if !found {
		return nil, types.ErrPairNotFound
	}

	err = k.CheckBurnAmount(msg.AmountIn, pair)
	if err != nil {
		return nil, err
	}

	var (
		burn_err  error
		amountOut sdk.Coin
	)

	switch pair.Model {
	case "gmb":
		burn_err, amountOut = k.GMB_ExecuteBurn(ctx, msg, pair)
	}

	if burn_err != nil {
		return nil, burn_err
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
		AmountIn:  msg.AmountIn,
		AmountOut: amountOut.String(),
	}, nil
}
