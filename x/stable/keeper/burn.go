package keeper

import (
	gmd "github.com/QuadrateOrg/core/x/stable/gmb"
	"github.com/QuadrateOrg/core/x/stable/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) ExecuteBurn(ctx sdk.Context, msg *types.MsgBurn) (error, sdk.Coin) {
	atomPrice, err := k.GetAtomPrice(ctx)

	if err != nil {
		return err, sdk.Coin{}
	}

	qm, ar := k.GetReserve(ctx)

	backing_ratio, err = CalculateBackingRatio(qm, ar, atomPrice)
	if err != nil {
		return err, sdk.Coin{}
	}

	burningFee, allow, err := gmd.CalculateBurningFee(backing_ratio)
	if err != nil {
		return err, sdk.Coin{}
	}
	if !allow {
		return types.ErrBurnBlocked, sdk.Coin{}
	}

	sender, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return err, sdk.Coin{}
	}
	amountIntCoins, err := sdk.ParseCoinsNormalized(msg.Amount)
	if err != nil {
		return err, sdk.Coin{}
	}

	// TODO
	// Verification of denom and number of coins
	if amountIntCoins.Len() != 1 {
		return types.ErrSend1Token, sdk.Coin{}
	}
	if amountIntCoins.GetDenomByIndex(0) != SendTokenDenom {
		return types.ErrSendBaseTokenDenom, sdk.Coin{}
	}

	amountInt := amountIntCoins.AmountOf(SendTokenDenom)
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, amountIntCoins)
	if err != nil {
		return err, sdk.Coin{}
	}

	amountOutToSend := k.CalculateAmountToSend(amountInt, atomPrice, burningFee)
	if amountOutToSend.IsNil() {
		return types.ErrSdkIntError, sdk.Coin{}
	}

	err = k.ReduceReserve(ctx, amountOutToSend, amountInt)
	if err != nil {
		return err, sdk.Coin{}
	}

	amountOut := sdk.NewCoin(BaseTokenDenom, amountOutToSend)
	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, amountIntCoins)
	if err != nil {
		return err, sdk.Coin{}
	}

	if !burningFee.IsZero() {
		feeForStabilityFund := k.CalculateBurningFeeForStabilityFund(amountOutToSend, atomPrice, burningFee)
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, StabilityFundAddress, types.CreateCoins(BaseTokenDenom, feeForStabilityFund))
		if err != nil {
			return err, sdk.Coin{}
		}
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, sdk.NewCoins(amountOut))
	if err != nil {
		return err, sdk.Coin{}
	}
	return nil, amountOut
}
