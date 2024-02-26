package keeper

import (
	"context"

	"github.com/QuadrateOrg/core/x/grow/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.MsgServer = Keeper{}

// Msg for deposit
func (k Keeper) GrowDeposit(goCtx context.Context, msg *types.MsgGrowDeposit) (*types.MsgGrowDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.CheckDepositMethodStatus(ctx)
	if err != nil {
		return nil, err
	}

	denomID := k.GenerateDenomIdHash(msg.DenomOut)

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

	return &types.MsgGrowDepositResponse{
		Creator:   msg.Creator,
		AmountIn:  msg.AmountIn,
		AmountOut: amountOut.String(),
	}, nil
}

// Msg for withdrawal
func (k Keeper) GrowWithdrawal(goCtx context.Context, msg *types.MsgGrowWithdrawal) (*types.MsgGrowWithdrawalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.CheckDepositMethodStatus(ctx)
	if err != nil {
		return nil, err
	}

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

	return &types.MsgGrowWithdrawalResponse{
		Creator:   msg.Creator,
		AmountIn:  msg.AmountIn,
		AmountOut: amountOut.String(),
	}, nil
}

// Msg of deposit collateral for borrowing money from x/grow
func (k Keeper) CreateLend(goCtx context.Context, msg *types.MsgCreateLend) (*types.MsgCreateLendResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.CheckCollateralMethodStatus(ctx)
	if err != nil {
		return nil, err
	}

	AssetId, err := k.GetAssetIdByCoins(msg.AmountIn)
	if err != nil {
		return nil, err
	}
	Asset, found := k.GetAssetByAssetId(ctx, AssetId)
	if !found {
		return nil, types.ErrAssetNotFound
	}

	err = k.CheckOracleAssetId(ctx, Asset)
	if err != nil {
		return nil, err
	}

	err, depositId := k.ExecuteCreateLend(ctx, msg, Asset)
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
	return &types.MsgCreateLendResponse{
		Depositor:  msg.Depositor,
		PositionId: depositId,
	}, nil
}

// Msg of withdrawal collateral from x/grow
func (k Keeper) WithdrawalLend(goCtx context.Context, msg *types.MsgWithdrawalLend) (*types.MsgWithdrawalLendResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.CheckCollateralMethodStatus(ctx)
	if err != nil {
		return nil, err
	}

	AssetId := k.GenerateAssetIdHash(msg.DenomOut)
	if err != nil {
		return nil, err
	}
	Asset, found := k.GetAssetByAssetId(ctx, AssetId)
	if !found {
		return nil, types.ErrAssetNotFound
	}

	err = k.CheckOracleAssetId(ctx, Asset)
	if err != nil {
		return nil, err
	}

	err, amountOut := k.ExecuteWithdrawalLend(ctx, msg, Asset)
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
	return &types.MsgWithdrawalLendResponse{
		Depositor: msg.Depositor,
		AmountOut: amountOut.String(),
	}, nil
}

// Msg for lend asset
func (k Keeper) CreateBorrow(goCtx context.Context, msg *types.MsgCreateBorrow) (*types.MsgCreateBorrowResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.CheckBorrowMethodStatus(ctx)
	if err != nil {
		return nil, err
	}

	AssetId := k.GenerateAssetIdHash(msg.DenomIn)
	if err != nil {
		return nil, err
	}

	Asset, found := k.GetAssetByAssetId(ctx, AssetId)
	if !found {
		return nil, types.ErrAssetNotFound
	}

	err = k.CheckOracleAssetId(ctx, Asset)
	if err != nil {
		return nil, err
	}

	err, amountOut, loanId := k.ExecuteCreateBorrow(ctx, msg, Asset)
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

	return &types.MsgCreateBorrowResponse{
		Borrower:  msg.Borrower,
		DenomIn:   msg.DenomIn,
		AmountOut: amountOut.String(),
		LoanId:    loanId,
	}, nil
}

// Msg for delete lend
func (k Keeper) DeleteBorrow(goCtx context.Context, msg *types.MsgDeleteBorrow) (*types.MsgDeleteBorrowResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.CheckBorrowMethodStatus(ctx)
	if err != nil {
		return nil, err
	}

	AssetId := k.GenerateAssetIdHash(msg.DenomOut)
	if err != nil {
		return nil, err
	}

	Asset, found := k.GetAssetByAssetId(ctx, AssetId)
	if !found {
		return nil, types.ErrAssetNotFound
	}

	err = k.CheckOracleAssetId(ctx, Asset)
	if err != nil {
		return nil, err
	}

	err, loanId := k.ExecuteDeleteBorrow(ctx, msg, Asset)
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

	return &types.MsgDeleteBorrowResponse{
		Borrower: msg.Borrower,
		LoanId:   loanId,
	}, nil
}

// Msg for open liquidation postion
func (k Keeper) OpenLiquidationPosition(goCtx context.Context, msg *types.MsgOpenLiquidationPosition) (*types.MsgOpenLiquidationPositionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.CheckBorrowMethodStatus(ctx)
	if err != nil {
		return nil, err
	}

	Asset, err := k.GetAssetByOracleAssetId(ctx, msg.Asset)
	if err != nil {
		return nil, err
	}

	err = k.CheckOracleAssetId(ctx, Asset)
	if err != nil {
		return nil, err
	}

	err, id := k.ExecuteCreateLiqPosition(ctx, msg, Asset)
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
	return &types.MsgOpenLiquidationPositionResponse{
		Creator:              msg.Creator,
		LiquidatorPositionId: id,
	}, nil
}

// Msg for close liquidation postion
func (k Keeper) CloseLiquidationPosition(goCtx context.Context, msg *types.MsgCloseLiquidationPosition) (*types.MsgCloseLiquidationPositionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.CheckBorrowMethodStatus(ctx)
	if err != nil {
		return nil, err
	}

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
