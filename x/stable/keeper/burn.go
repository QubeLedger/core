package keeper

import (
	gmd "github.com/QuadrateOrg/core/x/stable/gmb"
	"github.com/QuadrateOrg/core/x/stable/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) ExecuteBurn(ctx sdk.Context, msg *types.MsgBurn, pair types.Pair) (error, sdk.Coin) {

	ReserveFundAddress := k.GetReserveFundAddress(ctx)
	BurningFundAddress := k.GetBurningFundAddress(ctx)

	atomPrice, err := k.GetAtomPrice(ctx, pair)

	if err != nil {
		return err, sdk.Coin{}
	}

	if k.AddressEmptyCheck(ctx) {
		return types.ErrBurningFundAddressEmpty, sdk.Coin{}
	}

	qm, ar := pair.Qm, pair.Ar

	backing_ratio, err = k.CalculateBackingRatio(qm, ar, atomPrice)
	if err != nil {
		return err, sdk.Coin{}
	}

	burningFee, err := gmd.CalculateBurningFee(backing_ratio)
	if err != nil {
		return err, sdk.Coin{}
	}

	sender, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return err, sdk.Coin{}
	}
	amountIntCoins, err := sdk.ParseCoinsNormalized(msg.AmountIn)
	if err != nil {
		return err, sdk.Coin{}
	}

	err = VerificationBurnDenomCoins(amountIntCoins, pair)
	if err != nil {
		return err, sdk.Coin{}
	}

	amountInt := amountIntCoins.AmountOf(pair.AmountOutMetadata.DenomUnits[0].Denom)
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, amountIntCoins)
	if err != nil {
		return err, sdk.Coin{}
	}

	amountOutToSend := k.CalculateAmountToSend(amountInt, atomPrice, burningFee)
	if amountOutToSend.IsNil() {
		return types.ErrSdkIntError, sdk.Coin{}
	}

	pair = k.ReduceReserve(ctx, amountOutToSend, amountInt, pair)
	if err != nil {
		return err, sdk.Coin{}
	}

	amountOut := sdk.NewCoin(pair.AmountInMetadata.DenomUnits[0].Denom, amountOutToSend)

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, ReserveFundAddress, types.ModuleName, sdk.NewCoins(amountOut))
	if err != nil {
		return err, sdk.Coin{}
	}

	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, amountIntCoins)
	if err != nil {
		return err, sdk.Coin{}
	}

	if !burningFee.IsZero() {
		fee := k.CalculateBurningFeeForBurningFund(amountInt, atomPrice, burningFee)
		err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, ReserveFundAddress, types.ModuleName, types.CreateCoins(pair.AmountInMetadata.DenomUnits[0].Denom, fee))
		if err != nil {
			return err, sdk.Coin{}
		}
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, BurningFundAddress, types.CreateCoins(pair.AmountInMetadata.DenomUnits[0].Denom, fee))
		if err != nil {
			return err, sdk.Coin{}
		}
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, sdk.NewCoins(amountOut))
	if err != nil {
		return err, sdk.Coin{}
	}

	k.SetPair(ctx, pair)

	return nil, amountOut
}
