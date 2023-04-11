package keeper

import (
	"context"

	"github.com/QuadrateOrg/core/x/printer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Burn(goCtx context.Context, msg *types.MsgBurn) (*types.MsgBurnResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	/* Parse msg and check for correct */
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	amount, err := sdk.ParseCoinsNormalized(msg.Amount)
	if err != nil {
		panic(err)
	}

	if amount.GetDenomByIndex(0) != "usq" {
		panic(sdkerrors.Wrapf(types.ErrInvalidVersion, "got %s, expected usq", amount.GetDenomByIndex(0)))
	}

	/* Check delegation */
	validator := k.stakingKeeper.GetAllValidators(ctx)[0]
	validatorAddress := validator.GetOperator()

	price := sdk.NewInt(k.oracleKeeper.GetPrice(ctx) / 100000)

	usqAmount := (amount.AmountOfNoDenomValidation("usq")).Mul(sdk.NewInt(10))
	qubeAmountAfterPriceCalculation := usqAmount.Quo(price)

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, amount)
	if err != nil {
		return nil, err
	}

	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, amount)
	if err != nil {
		return nil, err
	}

	del, found := k.stakingKeeper.GetDelegation(ctx, creator, validatorAddress)
	if found == false {
		qubeTemp := sdk.NewCoin("qube", qubeAmountAfterPriceCalculation)
		err = k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(qubeTemp))
		if err != nil {
			return nil, err
		}
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creator, sdk.NewCoins(qubeTemp))
		if err != nil {
			return nil, err
		}
		return &types.MsgBurnResponse{}, nil
	}

	if (qubeAmountAfterPriceCalculation.Sub(sdk.Int(del.Shares))).IsPositive() {
		qubeTemp := sdk.NewCoin("qube", (qubeAmountAfterPriceCalculation).Sub(sdk.Int(del.Shares)))
		err = k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(qubeTemp))
		if err != nil {
			return nil, err
		}
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creator, sdk.NewCoins(qubeTemp))
		if err != nil {
			return nil, err
		}
	}
	_, err = k.stakingKeeper.Unbond(ctx, creator, validatorAddress, sdk.NewDecFromInt(qubeAmountAfterPriceCalculation))
	if err != nil {
		return nil, err
	}
	return &types.MsgBurnResponse{}, nil
}
