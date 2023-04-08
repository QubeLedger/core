package keeper

import (
	"context"

	"github.com/QuadrateOrg/core/x/printer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Burn(goCtx context.Context, msg *types.MsgBurn) (*types.MsgBurnResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

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

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, amount)
	if err != nil {
		return nil, err
	}

	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, amount)
	if err != nil {
		return nil, err
	}

	//TODO
	/*
	* Undelegate qube and send to creator
	* Get price from oracle
	 */

	return &types.MsgBurnResponse{}, nil
}
