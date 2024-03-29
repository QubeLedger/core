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

func (k Keeper) GMB_ExecuteMint(ctx sdk.Context, msg *types.MsgMint, pair types.Pair) (error, sdk.Coin) {

	params := k.GetParams(ctx)
	ReserveFundAddress, _ := sdk.AccAddressFromBech32(params.ReserveFundAddress)
	BurningFundAddress, _ := sdk.AccAddressFromBech32(params.BurningFundAddress)

	atomPrice, err := k.GetAtomPrice(ctx, pair)
	if err != nil {
		return err, sdk.Coin{}
	}

	if k.AddressEmptyCheck(ctx) {
		return types.ErrReserveFundAddressEmpty, sdk.Coin{}
	}

	qm, ar := pair.Qm, pair.Ar

	backing_ratio, err = k.CalculateBackingRatio(qm, ar, atomPrice)
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

	amountIntCoins, err := sdk.ParseCoinsNormalized(msg.AmountIn)
	if err != nil {
		return err, sdk.Coin{}
	}

	err = VerificationMintDenomCoins(amountIntCoins, pair)
	if err != nil {
		return err, sdk.Coin{}
	}

	amountInt := amountIntCoins.AmountOf(pair.AmountInMetadata.DenomUnits[0].Denom)
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, amountIntCoins)
	if err != nil {
		return err, sdk.Coin{}
	}

	amountOutToMint := k.CalculateAmountToMint(amountInt, atomPrice, mintingFee)
	if amountOutToMint.IsNil() {
		return types.ErrSdkIntError, sdk.Coin{}
	}

	pair = k.IncreaseReserve(ctx, amountInt, amountOutToMint, pair)
	if err != nil {
		return err, sdk.Coin{}
	}

	amountOut := sdk.NewCoin(pair.AmountOutMetadata.DenomUnits[0].Denom, amountOutToMint)
	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(amountOut))
	if err != nil {
		return err, sdk.Coin{}
	}

	fee := sdk.NewInt(0)
	if !mintingFee.IsZero() {
		fee = k.CalculateMintingFeeForBurningFund(amountInt, atomPrice, mintingFee)
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, BurningFundAddress, types.CreateCoins(pair.AmountInMetadata.DenomUnits[0].Denom, fee))
		if err != nil {
			return err, sdk.Coin{}
		}
	}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, ReserveFundAddress, types.CreateCoins(pair.AmountInMetadata.DenomUnits[0].Denom, (amountInt.Sub(fee))))
	if err != nil {
		return err, sdk.Coin{}
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, sdk.NewCoins(amountOut))
	if err != nil {
		return err, sdk.Coin{}
	}

	k.SetPair(ctx, pair)

	return nil, amountOut
}
