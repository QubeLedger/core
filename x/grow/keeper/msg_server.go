package keeper

import (
	"context"

	"github.com/QuadrateOrg/core/x/grow/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.MsgServer = Keeper{}

func (k Keeper) Deposit(goCtx context.Context, msg *types.MsgDeposit) (*types.MsgDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	pairId, err := k.GetPairIdDeposit(msg.AmountIn, msg.DenomOut)
	pair, found := k.GetPairByPairID(ctx, pairId)
	if !found {
		return nil, types.ErrPairNotFound
	}

	err = k.CheckDepositAmount(msg.AmountIn, pair)
	if err != nil {
		return nil, err
	}

	err, amountOut := k.ExecuteDeposit(ctx, msg, pair)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.AttributeKeyActionDeposit),
		),
	})

	return &types.MsgDepositResponse{
		Creator:   msg.Creator,
		AmountIn:  msg.AmountIn,
		AmountOut: amountOut.String(),
	}, nil
}

func (k Keeper) Withdrawal(goCtx context.Context, msg *types.MsgWithdrawal) (*types.MsgWithdrawalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	pairId, err := k.GetPairIdWithdrawal(msg.AmountIn, msg.DenomOut)
	pair, found := k.GetPairByPairID(ctx, pairId)
	if !found {
		return nil, types.ErrPairNotFound
	}

	err = k.CheckDepositAmount(msg.AmountIn, pair)
	if err != nil {
		return nil, err
	}

	err, amountOut := k.ExecuteWithdrawal(ctx, msg, pair)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.AttributeKeyActionWithdrawal),
		),
	})

	return &types.MsgWithdrawalResponse{
		Creator:   msg.Creator,
		AmountIn:  msg.AmountIn,
		AmountOut: amountOut.String(),
	}, nil
}

func (k Keeper) CreateLend(goCtx context.Context, msg *types.MsgCreateLend) (*types.MsgCreateLendResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.AttributeKeyActionCreateLend),
		),
	})

	return &types.MsgCreateLendResponse{}, nil
}

func (k Keeper) DeleteLend(goCtx context.Context, msg *types.MsgDeleteLend) (*types.MsgDeleteLendResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.AttributeKeyActionDeleteLend),
		),
	})

	return &types.MsgDeleteLendResponse{}, nil
}
