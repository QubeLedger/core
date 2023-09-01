package keeper

import (
	"context"

	"github.com/QuadrateOrg/core/x/grow/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.MsgServer = Keeper{}

func (k Keeper) Deposit(goCtx context.Context, msg *types.MsgDeposit) (*types.MsgDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	denomID, err := k.GetDenomIdDeposit(msg.DenomOut)
	gTokenPair, found := k.GetPairByDenomID(ctx, denomID)
	if !found {
		return nil, types.ErrPairNotFound
	}

	err = k.CheckDepositAmount(ctx, msg.AmountIn, gTokenPair)
	if err != nil {
		return nil, err
	}

	err, amountOut := k.ExecuteDeposit(ctx, msg, gTokenPair)
	if err != nil {
		return nil, err
	}

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

	denomID, err := k.GetDenomIdWithdrawal(msg.AmountIn)
	gTokenPair, found := k.GetPairByDenomID(ctx, denomID)
	if !found {
		return nil, types.ErrPairNotFound
	}

	err = k.CheckWithdrawalAmount(msg.AmountIn, gTokenPair)
	if err != nil {
		return nil, err
	}

	err, amountOut := k.ExecuteWithdrawal(ctx, msg, gTokenPair)
	if err != nil {
		return nil, err
	}

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

	borrowAssetId, err := k.GetBorrowAssetIdCreateLend(msg.AmountIn, msg.DenomOut)
	borrowAsset, found := k.GetBorrowAssetByBorrowAssetId(ctx, borrowAssetId)
	if !found {
		return nil, types.ErrPairNotFound
	}

	err, amountOut, loanId := k.ExecuteLend(ctx, msg, borrowAsset)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Borrower),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.AttributeKeyActionCreateLend),
		),
	})

	return &types.MsgCreateLendResponse{
		Borrower:  msg.Borrower,
		AmountIn:  msg.AmountIn,
		AmountOut: amountOut.String(),
		LoanId:    loanId,
	}, nil
}

func (k Keeper) DeleteLend(goCtx context.Context, msg *types.MsgDeleteLend) (*types.MsgDeleteLendResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	borrowAssetId, err := k.GetBorrowAssetIdDeleteLend(msg.AmountIn, msg.DenomOut)
	borrowAsset, found := k.GetBorrowAssetByBorrowAssetId(ctx, borrowAssetId)
	if !found {
		return nil, types.ErrPairNotFound
	}

	err, amountOut, loanId := k.ExecuteDeleteLend(ctx, msg, borrowAsset)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Borrower),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.AttributeKeyActionDeleteLend),
		),
	})

	return &types.MsgDeleteLendResponse{
		Borrower:  msg.Borrower,
		AmountIn:  msg.AmountIn,
		AmountOut: amountOut.String(),
		LoanId:    loanId,
	}, nil
}
