package keeper

import (
	"context"

	"github.com/QuadrateOrg/core/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.MsgServer = Keeper{}

func (k Keeper) DexDeposit(
	goCtx context.Context,
	msg *types.MsgDexDeposit,
) (*types.MsgDexDepositResponse, error) {
	callerAddr := sdk.MustAccAddressFromBech32(msg.Creator)
	receiverAddr := sdk.MustAccAddressFromBech32(msg.Receiver)

	pairID, err := types.NewPairIDFromUnsorted(msg.TokenA, msg.TokenB)
	if err != nil {
		return nil, err
	}

	// sort amounts
	amounts0, amounts1 := SortAmounts(msg.TokenA, pairID.Token0, msg.AmountsA, msg.AmountsB)

	tickIndexes := NormalizeAllTickIndexes(msg.TokenA, pairID.Token0, msg.TickIndexesAToB)

	Amounts0Deposit, Amounts1Deposit, _, err := k.DepositCore(
		goCtx,
		pairID,
		callerAddr,
		receiverAddr,
		amounts0,
		amounts1,
		tickIndexes,
		msg.Fees,
		msg.Options,
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgDexDepositResponse{
		Reserve0Deposited: Amounts0Deposit,
		Reserve1Deposited: Amounts1Deposit,
	}, nil
}

func (k Keeper) DexWithdrawal(
	goCtx context.Context,
	msg *types.MsgDexWithdrawal,
) (*types.MsgDexWithdrawalResponse, error) {
	callerAddr := sdk.MustAccAddressFromBech32(msg.Creator)
	receiverAddr := sdk.MustAccAddressFromBech32(msg.Receiver)

	pairID, err := types.NewPairIDFromUnsorted(msg.TokenA, msg.TokenB)
	if err != nil {
		return nil, err
	}

	tickIndexes := NormalizeAllTickIndexes(msg.TokenA, pairID.Token0, msg.TickIndexesAToB)

	err = k.WithdrawCore(
		goCtx,
		pairID,
		callerAddr,
		receiverAddr,
		msg.SharesToRemove,
		tickIndexes,
		msg.Fees,
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgDexWithdrawalResponse{}, nil
}

func (k Keeper) PlaceLimitOrder(
	goCtx context.Context,
	msg *types.MsgPlaceLimitOrder,
) (*types.MsgPlaceLimitOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	callerAddr := sdk.MustAccAddressFromBech32(msg.Creator)
	receiverAddr := sdk.MustAccAddressFromBech32(msg.Receiver)

	err := msg.ValidateGoodTilExpiration(ctx.BlockTime())
	if err != nil {
		return &types.MsgPlaceLimitOrderResponse{}, err
	}
	trancheKey, coinIn, _, coinOutSwap, err := k.PlaceLimitOrderCore(
		goCtx,
		msg.TokenIn,
		msg.TokenOut,
		msg.AmountIn,
		msg.TickIndexInToOut,
		msg.OrderType,
		msg.ExpirationTime,
		msg.MaxAmountOut,
		callerAddr,
		receiverAddr,
	)
	if err != nil {
		return &types.MsgPlaceLimitOrderResponse{}, err
	}

	return &types.MsgPlaceLimitOrderResponse{
		TrancheKey:   trancheKey,
		CoinIn:       coinIn,
		TakerCoinOut: coinOutSwap,
	}, nil
}

func (k Keeper) WithdrawFilledLimitOrder(
	goCtx context.Context,
	msg *types.MsgWithdrawFilledLimitOrder,
) (*types.MsgWithdrawFilledLimitOrderResponse, error) {
	callerAddr := sdk.MustAccAddressFromBech32(msg.Creator)

	err := k.WithdrawFilledLimitOrderCore(
		goCtx,
		msg.TrancheKey,
		callerAddr,
	)
	if err != nil {
		return &types.MsgWithdrawFilledLimitOrderResponse{}, err
	}

	return &types.MsgWithdrawFilledLimitOrderResponse{}, nil
}

func (k Keeper) CancelLimitOrder(
	goCtx context.Context,
	msg *types.MsgCancelLimitOrder,
) (*types.MsgCancelLimitOrderResponse, error) {
	callerAddr := sdk.MustAccAddressFromBech32(msg.Creator)

	err := k.CancelLimitOrderCore(
		goCtx,
		msg.TrancheKey,
		callerAddr,
	)
	if err != nil {
		return &types.MsgCancelLimitOrderResponse{}, err
	}

	return &types.MsgCancelLimitOrderResponse{}, nil
}

func (k Keeper) MultiHopSwap(
	goCtx context.Context,
	msg *types.MsgMultiHopSwap,
) (*types.MsgMultiHopSwapResponse, error) {
	callerAddr := sdk.MustAccAddressFromBech32(msg.Creator)
	receiverAddr := sdk.MustAccAddressFromBech32(msg.Receiver)
	ctx := sdk.UnwrapSDKContext(goCtx)

	coinOut, err := k.MultiHopSwapCore(
		ctx,
		msg.AmountIn,
		msg.Routes,
		msg.ExitLimitPrice,
		msg.PickBestRoute,
		callerAddr,
		receiverAddr,
	)
	if err != nil {
		return &types.MsgMultiHopSwapResponse{}, err
	}

	return &types.MsgMultiHopSwapResponse{CoinOut: coinOut}, nil
}
