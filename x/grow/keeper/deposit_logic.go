package keeper

import (
	"github.com/QuadrateOrg/core/x/grow/types"
	stabletypes "github.com/QuadrateOrg/core/x/stable/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) ExecuteDeposit(ctx sdk.Context, msg *types.MsgDeposit, gTokenPair types.GTokenPair) (error, sdk.Coin) {

	amountInCoins, err := sdk.ParseCoinsNormalized(msg.AmountIn)
	if err != nil {
		return err, sdk.Coin{}
	}

	sender, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return err, sdk.Coin{}
	}

	qStablePair, found := k.stableKeeper.GetPairByPairID(ctx, gTokenPair.QStablePairId)
	if !found {
		return stabletypes.ErrPairNotFound, sdk.Coin{}
	}

	amountInInt := amountInCoins.AmountOf(qStablePair.AmountOutMetadata.Base)

	gTokenPrice, err := k.GetGTokenPrice(ctx, gTokenPair.DenomID)
	if err != nil {
		return err, sdk.Coin{}
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, amountInCoins)
	if err != nil {
		return err, sdk.Coin{}
	}

	gTokenPair, err = k.IncreaseGrowStakingReserve(ctx, amountInCoins, gTokenPair, qStablePair)
	if err != nil {
		return err, sdk.Coin{}
	}

	amountOutInt := k.CalculateGTokenAmountOut(amountInInt, gTokenPrice)
	if amountOutInt.IsNil() || amountOutInt.IsZero() {
		return types.ErrIntNegativeOrZero, sdk.Coin{}
	}

	amountOut := sdk.NewCoin(gTokenPair.GTokenMetadata.Base, amountOutInt)

	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(amountOut))
	if err != nil {
		return err, sdk.Coin{}
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, sdk.NewCoins(amountOut))
	if err != nil {
		return err, sdk.Coin{}
	}

	k.SetPair(ctx, gTokenPair)

	return nil, amountOut
}

func (k Keeper) ExecuteWithdrawal(ctx sdk.Context, msg *types.MsgWithdrawal, gTokenPair types.GTokenPair) (error, sdk.Coin) {

	amountInCoins, err := sdk.ParseCoinsNormalized(msg.AmountIn)
	if err != nil {
		return err, sdk.Coin{}
	}

	sender, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return err, sdk.Coin{}
	}

	amountInInt := amountInCoins.AmountOf(gTokenPair.GTokenMetadata.Base)

	qStablePair, found := k.stableKeeper.GetPairByPairID(ctx, gTokenPair.QStablePairId)
	if !found {
		return stabletypes.ErrPairNotFound, sdk.Coin{}
	}

	gTokenPrice, err := k.GetGTokenPrice(ctx, gTokenPair.DenomID)
	if err != nil {
		return err, sdk.Coin{}
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, amountInCoins)
	if err != nil {
		return err, sdk.Coin{}
	}

	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, amountInCoins)
	if err != nil {
		return err, sdk.Coin{}
	}

	amountOutInt := k.CalculateReturnQubeStableAmountOut(amountInInt, gTokenPrice)
	if amountOutInt.IsNil() || amountOutInt.IsZero() {
		return types.ErrIntNegativeOrZero, sdk.Coin{}
	}

	amountOut := sdk.NewCoin(qStablePair.AmountOutMetadata.Base, amountOutInt)

	gTokenPair, err = k.ReduceGrowStakingReserve(ctx, sdk.NewCoins(amountOut), gTokenPair)
	if err != nil {
		return err, sdk.Coin{}
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, sdk.NewCoins(amountOut))
	if err != nil {
		return err, sdk.Coin{}
	}

	k.SetPair(ctx, gTokenPair)

	return nil, amountOut

}
