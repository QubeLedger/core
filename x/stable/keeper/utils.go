package keeper

import (
	gmd "github.com/QuadrateOrg/core/x/stable/gmb"
	"github.com/QuadrateOrg/core/x/stable/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) CalculateBackingRatio(qm sdk.Int, ar sdk.Int, atomPrice sdk.Int) (sdk.Int, error) {
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

func VerificationMintDenomCoins(coins sdk.Coins, pair types.Pair) error {
	// TODO
	// Verification of denom and number of coins
	if coins.Len() != 1 {
		return types.ErrMultipleCoinsLockupNotSupported
	}
	if coins.GetDenomByIndex(0) != pair.AmountInMetadata.DenomUnits[0].Denom {
		return types.ErrSendBaseTokenDenom
	}
	return nil
}

func VerificationBurnDenomCoins(coins sdk.Coins, pair types.Pair) error {
	// TODO
	// Verification of denom and number of coins
	if coins.Len() != 1 {
		return types.ErrMultipleCoinsLockupNotSupported
	}
	if coins.GetDenomByIndex(0) != pair.AmountOutMetadata.DenomUnits[0].Denom {
		return types.ErrSendBaseTokenDenom
	}
	return nil
}

func (k Keeper) CheckMinAmount(msgAmountIn string, pair types.Pair) error {
	msgAmountInCoins, err := sdk.ParseCoinsNormalized(msgAmountIn)
	if err != nil {
		return err
	}

	pairMinAmountInCoins, err := sdk.ParseCoinsNormalized(pair.MinAmountIn)
	if err != nil {
		return err
	}
	if !msgAmountInCoins.AmountOf(pair.AmountInMetadata.Base).GT(pairMinAmountInCoins.AmountOf(pair.AmountInMetadata.Base)) {
		return types.ErrAmountInGTEminAmountIn
	}

	return nil
}

func (k Keeper) CheckBurnAmount(msgAmountIn string, pair types.Pair) error {
	msgAmountOutCoins, err := sdk.ParseCoinsNormalized(msgAmountIn)
	if err != nil {
		return err
	}

	pairMinAmountoutCoins, err := sdk.ParseCoinsNormalized(pair.MinAmountOut)
	if err != nil {
		return err
	}
	if !msgAmountOutCoins.AmountOf(pair.AmountOutMetadata.Base).GT(pairMinAmountoutCoins.AmountOf(pair.AmountOutMetadata.Base)) {
		return types.ErrAmountOutGTEminAmountOut
	}

	return nil
}
