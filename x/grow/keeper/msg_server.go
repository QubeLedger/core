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
	if err != nil {
		return nil, err
	}
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
	if err != nil {
		return nil, err
	}
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

func (k Keeper) DepositCollateral(goCtx context.Context, msg *types.MsgDepositCollateral) (*types.MsgDepositCollateralResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	LendAssetId, err := k.GetLendAssetIdByCoins(msg.AmountIn)
	if err != nil {
		return nil, err
	}
	LendAsset, found := k.GetLendAssetByLendAssetId(ctx, LendAssetId)
	if !found {
		return nil, types.ErrLendAssetNotFound
	}

	err = k.CheckOracleAssetId(ctx, LendAsset)
	if err != nil {
		return nil, err
	}

	err, depositId := k.ExecuteDepositCollateral(ctx, msg, LendAsset)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Depositor),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.AttributeKeyActionDepositColletaral),
		),
	})
	return &types.MsgDepositCollateralResponse{
		Depositor:  msg.Depositor,
		PositionId: depositId,
	}, nil
}

func (k Keeper) WithdrawalCollateral(goCtx context.Context, msg *types.MsgWithdrawalCollateral) (*types.MsgWithdrawalCollateralResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	LendAssetId, err := k.GetLendAssetIdByDenom(msg.Denom)
	if err != nil {
		return nil, err
	}
	LendAsset, found := k.GetLendAssetByLendAssetId(ctx, LendAssetId)
	if !found {
		return nil, types.ErrLendAssetNotFound
	}

	err = k.CheckOracleAssetId(ctx, LendAsset)
	if err != nil {
		return nil, err
	}

	err, amountOut := k.ExecuteWithdrawalCollateral(ctx, msg, LendAsset)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Depositor),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.AttributeKeyActionWithdrawalColletaral),
		),
	})
	return &types.MsgWithdrawalCollateralResponse{
		Depositor: msg.Depositor,
		AmountOut: amountOut.String(),
	}, nil
}

func (k Keeper) CreateLend(goCtx context.Context, msg *types.MsgCreateLend) (*types.MsgCreateLendResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	LendAssetId, err := k.GetLendAssetIdByDenom(msg.DenomIn)
	if err != nil {
		return nil, err
	}
	LendAsset, found := k.GetLendAssetByLendAssetId(ctx, LendAssetId)
	if !found {
		return nil, types.ErrLendAssetNotFound
	}

	err = k.CheckOracleAssetId(ctx, LendAsset)
	if err != nil {
		return nil, err
	}

	err, amountOut, loanId := k.ExecuteLend(ctx, msg, LendAsset)
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
		DenomIn:   msg.DenomIn,
		AmountOut: amountOut.String(),
		LoanId:    loanId,
	}, nil
}

func (k Keeper) DeleteLend(goCtx context.Context, msg *types.MsgDeleteLend) (*types.MsgDeleteLendResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	LendAssetId, err := k.GetLendAssetIdByDenom(msg.DenomOut)
	if err != nil {
		return nil, err
	}
	LendAsset, found := k.GetLendAssetByLendAssetId(ctx, LendAssetId)
	if !found {
		return nil, types.ErrLendAssetNotFound
	}

	err = k.CheckOracleAssetId(ctx, LendAsset)
	if err != nil {
		return nil, err
	}

	err, loanId := k.ExecuteDeleteLend(ctx, msg, LendAsset)
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
		Borrower: msg.Borrower,
		LoanId:   loanId,
	}, nil
}

func (k Keeper) CreateLiquidationPosition(goCtx context.Context, msg *types.MsgCreateLiquidationPosition) (*types.MsgCreateLiquidationPositionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	LendAsset, err := k.GetLendAssetByOracleAssetId(ctx, msg.Asset)
	if err != nil {
		return nil, err
	}

	err = k.CheckOracleAssetId(ctx, LendAsset)
	if err != nil {
		return nil, err
	}

	err, id := k.ExecuteCreateLiqPosition(ctx, msg, LendAsset)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.AttributeKeyActionCreateLiqPosition),
		),
	})
	return &types.MsgCreateLiquidationPositionResponse{
		Creator:              msg.Creator,
		LiquidatorPositionId: id,
	}, nil
}

func (k Keeper) CloseLiquidationPosition(goCtx context.Context, msg *types.MsgCloseLiquidationPosition) (*types.MsgCloseLiquidationPositionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	err, amountOut := k.ExecuteCloseLiqPosition(ctx, msg)
	if err != nil {
		return nil, err
	}
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.AttributeKeyActionCloseLiqPosition),
		),
	})
	return &types.MsgCloseLiquidationPositionResponse{
		Creator:   msg.Creator,
		AmountOut: amountOut.String(),
	}, nil
}
