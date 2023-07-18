package keeper

import (
	"context"

	gmd "github.com/QuadrateOrg/core/x/stable/gmb"
	"github.com/QuadrateOrg/core/x/stable/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.MsgServer = Keeper{}

func (k Keeper) MintUsq(goCtx context.Context, msg *types.MsgMintUsq) (*types.MsgMintUsqResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Oracle
	atomPrice := k.GetAtomPrice(ctx)

	// GMD math logic

	qm := k.GetStablecoinSupply(ctx)
	ar := k.GetAtomReserve(ctx)

	if qm.IsNil() && ar.IsNil() {
		k.InitAtomReserve(ctx)
		k.InitStablecoinSupply(ctx)
		//k.IncreaseAtomReserve(ctx, sdk.NewInt(1))
		//k.IncreaseStablecoinSupply(ctx, sdk.NewInt(1))
	}

	qm = k.GetStablecoinSupply(ctx)
	ar = k.GetAtomReserve(ctx)

	var backing_ratio sdk.Int
	if qm.IsZero() && ar.IsZero() {
		backing_ratio = sdk.NewInt(100)
	} else {
		backing_ratio = gmd.CalculateBackingRatio(atomPrice, ar, qm)
	}

	mintingFee, allow, err := gmd.CalculateMintingFee(backing_ratio)

	if !allow {
		return nil, types.ErrMintBlocked
	}

	if err != nil {
		return nil, err
	}

	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	amount, err := sdk.ParseCoinsNormalized(msg.Amount)
	if err != nil {
		return nil, err
	}

	// TODO
	if amount.Len() != 1 {
		return nil, types.ErrSend1Token
	}

	// TODO
	if amount.GetDenomByIndex(0) != BaseTokenDenom {
		return nil, types.ErrSendBaseTokenDenom
	}

	ibcBaseTokenDenomAmount := amount.AmountOf(BaseTokenDenom)

	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, amount)
	if sdkError != nil {
		return nil, sdkError
	}
	//((atom * price) - (((atom * price) * fee) / 1000)) / 10000
	amountUsqToMint := k.CalculateAmountUsqToMint(ibcBaseTokenDenomAmount, atomPrice, mintingFee)

	if amountUsqToMint.IsNil() {
		return nil, types.ErrSdkIntError
	}

	k.IncreaseAtomReserve(ctx, amount.AmountOf(BaseTokenDenom))
	k.IncreaseStablecoinSupply(ctx, amountUsqToMint)

	uusd := sdk.NewCoin("uusd", amountUsqToMint)
	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(uusd))
	if err != nil {
		return nil, err
	}

	if !mintingFee.IsZero() {
		// (((atom * price) * fee) / 1000) / 10000
		feeForStabilityFund := k.CalculateMintingFeeForStabilityFund(ibcBaseTokenDenomAmount, atomPrice, mintingFee)
		atomFeeForStabilityFund := sdk.NewCoin(amount.GetDenomByIndex(0), feeForStabilityFund)
		err = k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(atomFeeForStabilityFund))
		if err != nil {
			return nil, err
		}
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, StabilityFundAddress, sdk.NewCoins(atomFeeForStabilityFund))
		if err != nil {
			return nil, err
		}
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creator, sdk.NewCoins(uusd))
	if err != nil {
		return nil, err
	}

	return &types.MsgMintUsqResponse{}, nil
}

func (k Keeper) BurnUsq(goCtx context.Context, msg *types.MsgBurnUsq) (*types.MsgBurnUsqResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Oracle

	atomPrice := k.GetAtomPrice(ctx)

	// GMD math logic

	qm := k.GetStablecoinSupply(ctx)
	ar := k.GetAtomReserve(ctx)

	backing_ratio := gmd.CalculateBackingRatio(atomPrice, ar, qm)
	burningFee, allow, err := gmd.CalculateBurningFee(backing_ratio)

	if !allow {
		return nil, types.ErrBurnBlocked
	}

	if err != nil {
		return nil, err
	}

	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	amount, err := sdk.ParseCoinsNormalized(msg.Amount)
	if err != nil {
		panic(err)
	}

	// TODO
	if amount.Len() != 1 {
		return nil, types.ErrSend1Token
	}

	// TODO
	if amount.GetDenomByIndex(0) != "uusd" {
		return nil, types.ErrSendBaseTokenDenom
	}

	uusdTokenAmount := amount.AmountOf("uusd")

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, amount)
	if err != nil {
		return nil, err
	}
	// (( ((uusd * 1000000000) / price) ) - ((((uusd * 1000000000) / price) * fee) / 1000) )/ 100000
	amountAtomToSend := k.CalculateAmountAtomToSend(uusdTokenAmount, atomPrice, burningFee)

	if amountAtomToSend.IsNil() {
		return nil, types.ErrSdkIntError
	}

	k.ReduceAtomReserve(ctx, amountAtomToSend)
	k.ReduceStablecoinSupply(ctx, amount.AmountOf("uusd"))

	atom := sdk.NewCoin(BaseTokenDenom, amountAtomToSend)
	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, amount)
	if err != nil {
		return nil, err
	}

	if !burningFee.IsZero() {
		// ((((uusd * 1000000000) / price) * fee) / 1000) / 100000
		feeForStabilityFund := k.CalculateBurningFeeForStabilityFund(uusdTokenAmount, atomPrice, burningFee)
		atomFeeForStabilityFund := sdk.NewCoin(BaseTokenDenom, feeForStabilityFund)
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, StabilityFundAddress, sdk.NewCoins(atomFeeForStabilityFund))
		if err != nil {
			return nil, err
		}
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creator, sdk.NewCoins(atom))
	if err != nil {
		return nil, err
	}

	return &types.MsgBurnUsqResponse{}, nil
}
