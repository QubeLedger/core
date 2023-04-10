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

	//TODO: get price from oracle
	//QUBE=$2
	price := int64(2 * 1000000 / 1000000)

	usqAmount := (amount.AmountOfNoDenomValidation("usq")).Int64()

	qubeAmountAfterPriceCalculation := usqAmount / price

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
		qubeTemp := sdk.NewCoin("qube", sdk.NewInt(qubeAmountAfterPriceCalculation))
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

	if (sdk.NewInt(qubeAmountAfterPriceCalculation).Sub(sdk.Int(del.Shares))).IsPositive() {
		qubeTemp := sdk.NewCoin("qube", sdk.NewInt(qubeAmountAfterPriceCalculation).Sub(sdk.Int(del.Shares)))
		err = k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(qubeTemp))
		if err != nil {
			return nil, err
		}
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creator, sdk.NewCoins(qubeTemp))
		if err != nil {
			return nil, err
		}
	}

	_, err = k.stakingKeeper.Unbond(ctx, creator, validatorAddress, sdk.NewDec(qubeAmountAfterPriceCalculation))
	if err != nil {
		return nil, err
	}
	return &types.MsgBurnResponse{}, nil
}
