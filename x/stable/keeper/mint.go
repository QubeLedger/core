package keeper

import (
	gmd "github.com/QuadrateOrg/core/x/stable/gmb"
	"github.com/QuadrateOrg/core/x/stable/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) ExecuteMint(ctx sdk.Context, msg *types.MsgMintUsq, atomPrice sdk.Int) (error, sdk.Coin) {
	// GMD math logic

	qm := k.GetStablecoinSupply(ctx)
	ar := k.GetAtomReserve(ctx)

	if qm.IsNil() && ar.IsNil() {
		k.InitAtomReserve(ctx)
		k.InitStablecoinSupply(ctx)
	}

	qm = k.GetStablecoinSupply(ctx)
	ar = k.GetAtomReserve(ctx)

	var backing_ratio sdk.Int
	var err error
	if qm.IsZero() && ar.IsZero() {
		backing_ratio = sdk.NewInt(100)
	} else {
		backing_ratio, err = gmd.CalculateBackingRatio(atomPrice, ar, qm)
		if err != nil {
			return err, sdk.Coin{}
		}
		if backing_ratio.IsNil() {
			return types.ErrSdkIntError, sdk.Coin{}
		}
	}

	mintingFee, allow, err := gmd.CalculateMintingFee(backing_ratio)

	if err != nil {
		return err, sdk.Coin{}
	}

	if !allow {
		return types.ErrMintBlocked, sdk.Coin{}
	}

	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
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
	if amount.GetDenomByIndex(0) != BaseTokenDenom {
		return types.ErrSendBaseTokenDenom, sdk.Coin{}
	}

	ibcBaseTokenDenomAmount := amount.AmountOf(BaseTokenDenom)

	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, amount)
	if sdkError != nil {
		return sdkError, sdk.Coin{}
	}
	//((atom * price) - (((atom * price) * fee) / 1000)) / 10000
	amountUsqToMint := k.CalculateAmountUsqToMint(ibcBaseTokenDenomAmount, atomPrice, mintingFee)

	if amountUsqToMint.IsNil() {
		return types.ErrSdkIntError, sdk.Coin{}
	}

	k.IncreaseAtomReserve(ctx, amount.AmountOf(BaseTokenDenom))
	k.IncreaseStablecoinSupply(ctx, amountUsqToMint)

	uusd := sdk.NewCoin(SendTokenDenom, amountUsqToMint)
	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(uusd))
	if err != nil {
		return err, sdk.Coin{}
	}

	if !mintingFee.IsZero() {
		// (((atom * price) * fee) / 1000) / 10000
		feeForStabilityFund := k.CalculateMintingFeeForStabilityFund(ibcBaseTokenDenomAmount, atomPrice, mintingFee)
		atomFeeForStabilityFund := sdk.NewCoin(amount.GetDenomByIndex(0), feeForStabilityFund)
		err = k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(atomFeeForStabilityFund))
		if err != nil {
			return err, sdk.Coin{}
		}
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, StabilityFundAddress, sdk.NewCoins(atomFeeForStabilityFund))
		if err != nil {
			return err, sdk.Coin{}
		}
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creator, sdk.NewCoins(uusd))
	if err != nil {
		return err, sdk.Coin{}
	}

	return nil, uusd
}
