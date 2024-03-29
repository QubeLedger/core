package keeper

import (
	"github.com/QuadrateOrg/core/x/dex/types"
	math_utils "github.com/QuadrateOrg/core/x/dex/utils/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

///////////////////////////////////////////////////////////////////////////////
//                          STATE CALCULATIONS                               //
///////////////////////////////////////////////////////////////////////////////

func (k Keeper) GetCurrPrice(ctx sdk.Context, tradePairID *types.TradePairID) (math_utils.PrecDec, bool) {
	liq := k.GetCurrLiq(ctx, tradePairID)
	if liq != nil {
		return liq.Price(), true
	}
	return math_utils.ZeroPrecDec(), false
}

// Returns a takerToMaker tick index
func (k Keeper) GetCurrTickIndexTakerToMaker(
	ctx sdk.Context,
	tradePairID *types.TradePairID,
) (int64, bool) {
	liq := k.GetCurrLiq(ctx, tradePairID)
	if liq != nil {
		return liq.TickIndex(), true
	}
	return 0, false
}

// Returns a takerToMaker tick index
func (k Keeper) GetCurrTickIndexTakerToMakerNormalized(
	ctx sdk.Context,
	tradePairID *types.TradePairID,
) (int64, bool) {
	tickIndexTakerToMaker, found := k.GetCurrTickIndexTakerToMaker(ctx, tradePairID)
	if found {
		tickIndexTakerToMakerNormalized := tradePairID.TickIndexNormalized(tickIndexTakerToMaker)
		return tickIndexTakerToMakerNormalized, true
	}

	return 0, false
}

func (k Keeper) GetCurrLiq(ctx sdk.Context, tradePairID *types.TradePairID) *types.TickLiquidity {
	ti := k.NewTickIterator(ctx, tradePairID)
	defer ti.Close()
	for ; ti.Valid(); ti.Next() {
		tick := ti.Value()
		if tick.HasToken() {
			return &tick
		}
	}

	return nil
}

func (k Keeper) GetValidFees(ctx sdk.Context) []uint64 {
	return k.GetParams(ctx).FeeTiers
}

func contains(s []uint64, e uint64) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (k Keeper) ValidateFee(ctx sdk.Context, fee uint64) error {
	validFees := k.GetValidFees(ctx)
	if !contains(validFees, fee) { // #nosec G104
		return sdkerrors.Wrapf(types.ErrInvalidFee, "%d", validFees)
	}

	return nil
}

///////////////////////////////////////////////////////////////////////////////
//                            TOKENIZER UTILS                                //
///////////////////////////////////////////////////////////////////////////////

func (k Keeper) MintShares(ctx sdk.Context, addr sdk.AccAddress, shareCoin sdk.Coin) error {
	// mint share tokens
	sharesCoins := sdk.Coins{shareCoin}
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sharesCoins); err != nil {
		return err
	}
	// transfer them to addr
	err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, sharesCoins)

	return err
}

func (k Keeper) BurnShares(
	ctx sdk.Context,
	addr sdk.AccAddress,
	amount sdk.Int,
	sharesID string,
) error {
	sharesCoins := sdk.Coins{sdk.NewCoin(sharesID, amount)}
	// transfer tokens to module
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, addr, types.ModuleName, sharesCoins); err != nil {
		return err
	}
	// burn tokens
	err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, sharesCoins)

	return err
}
