package keeper

import (
	gmd "github.com/QuadrateOrg/core/x/stable/gmb"
	"github.com/QuadrateOrg/core/x/stable/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) ExecuteBurn(ctx sdk.Context, msg *types.MsgBurn, atomPrice sdk.Int) (error, sdk.Coin) {
	qm := k.GetStablecoinSupply(ctx)
	ar := k.GetAtomReserve(ctx)

	backing_ratio, err := gmd.CalculateBackingRatio(atomPrice, ar, qm)
	if err != nil {
		return err, sdk.Coin{}
	}
	if backing_ratio.IsNil() {
		return types.ErrSdkIntError, sdk.Coin{}
	}
	burningFee, allow, err := gmd.CalculateBurningFee(backing_ratio)

	if err != nil {
		return err, sdk.Coin{}
	}

	if !allow {
		return types.ErrBurnBlocked, sdk.Coin{}
	}

	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return err, sdk.Coin{}
	}
	amount, err := sdk.ParseCoinsNormalized(msg.Amount)
	if err != nil {
		return err, sdk.Coin{}
	}

	// TODO
	if amount.Len() != 1 {
		return types.ErrSend1Token, sdk.Coin{}
	}

	// TODO
	if amount.GetDenomByIndex(0) != SendTokenDenom {
		return types.ErrSendBaseTokenDenom, sdk.Coin{}
	}

	uusdTokenAmount := amount.AmountOf(SendTokenDenom)

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, amount)
	if err != nil {
		return err, sdk.Coin{}
	}
	// (( ((uusd * 1000000000) / price) ) - ((((uusd * 1000000000) / price) * fee) / 1000) )/ 100000
	amountAtomToSend := k.CalculateAmountAtomToSend(uusdTokenAmount, atomPrice, burningFee)

	if amountAtomToSend.IsNil() {
		return types.ErrSdkIntError, sdk.Coin{}
	}

	k.ReduceAtomReserve(ctx, amountAtomToSend)
	k.ReduceStablecoinSupply(ctx, amount.AmountOf(SendTokenDenom))

	atom := sdk.NewCoin(BaseTokenDenom, amountAtomToSend)
	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, amount)
	if err != nil {
		return err, sdk.Coin{}
	}

	if !burningFee.IsZero() {
		// ((((uusd * 1000000000) / price) * fee) / 1000) / 100000
		feeForStabilityFund := k.CalculateBurningFeeForStabilityFund(uusdTokenAmount, atomPrice, burningFee)
		atomFeeForStabilityFund := sdk.NewCoin(BaseTokenDenom, feeForStabilityFund)
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, StabilityFundAddress, sdk.NewCoins(atomFeeForStabilityFund))
		if err != nil {
			return err, sdk.Coin{}
		}
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creator, sdk.NewCoins(atom))
	if err != nil {
		return err, sdk.Coin{}
	}
	return nil, atom
}
