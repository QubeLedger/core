package keeper

import (
	gmd "github.com/QuadrateOrg/core/x/stable/gmb"
	"github.com/QuadrateOrg/core/x/stable/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

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

func VerificationDenomCoins(coins sdk.Coins) error {
	// TODO
	// Verification of denom and number of coins
	if coins.Len() != 1 {
		return types.ErrMultipleCoinsLockupNotSupported
	}
	if coins.GetDenomByIndex(0) != BaseTokenDenom {
		return types.ErrSendBaseTokenDenom
	}
	return nil
}
