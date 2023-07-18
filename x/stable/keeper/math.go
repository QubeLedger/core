package keeper

import (
	"github.com/QuadrateOrg/core/x/stable/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) CalculateAmountUsqToMint(ibcBaseTokenDenomAmount sdk.Int, atomPrice sdk.Int, mintingFee sdk.Int) sdk.Int {
	return (ibcBaseTokenDenomAmount.Mul(atomPrice).Sub((ibcBaseTokenDenomAmount.Mul(atomPrice).Mul(mintingFee)).Quo(types.MintUsqMultiplier))).Quo(types.Multiplier)
}

func (k Keeper) CalculateMintingFeeForStabilityFund(ibcBaseTokenDenomAmount sdk.Int, atomPrice sdk.Int, mintingFee sdk.Int) sdk.Int {
	return (((ibcBaseTokenDenomAmount.Mul(atomPrice)).Mul(mintingFee)).Sub(types.MintUsqMultiplier)).Sub(types.Multiplier)
}

func (k Keeper) CalculateAmountAtomToSend(uusdTokenAmount sdk.Int, atomPrice sdk.Int, burningFee sdk.Int) sdk.Int {
	//  (( ((uusd * 1000000000) / price) ) - ( ( ((uusd * 1000000000) / price) ) * fee) / 1000 ) / 100000
	fee := (((uusdTokenAmount.Mul(types.BurnUsqMultiplier)).Quo(atomPrice)).Mul(burningFee)).Quo(sdk.NewInt(1000))
	return (((uusdTokenAmount.Mul(types.BurnUsqMultiplier)).Quo(atomPrice)).Sub(fee)).Quo(sdk.NewInt(100000))
}

func (k Keeper) CalculateBurningFeeForStabilityFund(uusdTokenAmount sdk.Int, atomPrice sdk.Int, burningFee sdk.Int) sdk.Int {
	// ((((uusd * 1000000000) / price) * fee) / 1000) / 100000
	return ((((uusdTokenAmount.Mul(types.BurnUsqMultiplier)).Quo(atomPrice)).Mul(burningFee)).Quo(types.MintUsqMultiplier)).Quo(types.Multiplier)
}
