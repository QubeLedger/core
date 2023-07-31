package keeper

import (
	gmd "github.com/QuadrateOrg/core/x/stable/gmb"
	"github.com/QuadrateOrg/core/x/stable/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	backing_ratio sdk.Int
	err           error
)

func (k Keeper) ExecuteMint(ctx sdk.Context, msg *types.MsgMint) (error, sdk.Coin) {

	atomPrice, err := k.GetAtomPrice(ctx)

	if err != nil {
		return err, sdk.Coin{}
	}

	qm, ar := k.GetReserve(ctx)

	backing_ratio, err = CalculateBackingRatio(qm, ar, atomPrice)
	if err != nil {
		return err, sdk.Coin{}
	}

	mintingFee, allow, err := gmd.CalculateMintingFee(backing_ratio)
	if err != nil {
		return err, sdk.Coin{}
	}
	if !allow {
		return types.ErrMintBlocked, sdk.Coin{}
	}

	sender, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	amountIntCoin, err := sdk.ParseCoinsNormalized(msg.Amount)
	if err != nil {
		return err, sdk.Coin{}
	}

	// TODO
	// Verification of denom and number of coins
	if amountIntCoin.Len() != 1 {
		return types.ErrSend1Token, sdk.Coin{}
	}
	if amountIntCoin.GetDenomByIndex(0) != BaseTokenDenom {
		return types.ErrSendBaseTokenDenom, sdk.Coin{}
	}

	amountInt := amountIntCoin.AmountOf(BaseTokenDenom)
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, amountIntCoin)
	if err != nil {
		return err, sdk.Coin{}
	}

	amountOutToMint := k.CalculateAmountToMint(amountInt, atomPrice, mintingFee)
	if amountOutToMint.IsNil() {
		return types.ErrSdkIntError, sdk.Coin{}
	}

	err = k.IncreaseReserve(ctx, amountInt, amountOutToMint)
	if err != nil {
		return err, sdk.Coin{}
	}

	amountOut := sdk.NewCoin(SendTokenDenom, amountOutToMint)
	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(amountOut))
	if err != nil {
		return err, sdk.Coin{}
	}

	if !mintingFee.IsZero() {
		feeForStabilityFund := k.CalculateMintingFeeForStabilityFund(amountInt, atomPrice, mintingFee)
		err = k.bankKeeper.MintCoins(ctx, types.ModuleName, types.CreateCoins(BaseTokenDenom, feeForStabilityFund))
		if err != nil {
			return err, sdk.Coin{}
		}
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

func CalculateBackingRatio(qm sdk.Int, ar sdk.Int, atomPrice sdk.Int) (sdk.Int, error) {
	if qm.IsZero() && ar.IsZero() {
		backing_ratio = sdk.NewInt(100)
	} else {
		backing_ratio, err = gmd.CalculateBackingRatio(atomPrice, ar, qm)
		if err != nil {
			return sdk.Int{}, err
		}
		if backing_ratio.IsNil() {
			return sdk.Int{}, types.ErrSdkIntError
		}
	}
	return backing_ratio, nil
}
