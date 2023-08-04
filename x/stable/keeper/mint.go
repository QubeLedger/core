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

	mintingFee, err := gmd.CalculateMintingFee(backing_ratio)
	if err != nil {
		return err, sdk.Coin{}
	}

	sender, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return err, sdk.Coin{}
	}

	amountIntCoins, err := sdk.ParseCoinsNormalized(msg.Amount)
	if err != nil {
		return err, sdk.Coin{}
	}

	err = VerificationBaseDenomCoins(amountIntCoins)
	if err != nil {
		return err, sdk.Coin{}
	}

	amountInt := amountIntCoins.AmountOf(BaseTokenDenom)
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, amountIntCoins)
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
