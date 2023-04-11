package keeper

import (
	"context"

	"github.com/QuadrateOrg/core/x/printer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

func (k msgServer) Mint(goCtx context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	amount, err := sdk.ParseCoinsNormalized(msg.Amount)
	if err != nil {
		panic(err)
	}

	if amount.GetDenomByIndex(0) != "qube" {
		panic(sdkerrors.Wrapf(types.ErrInvalidVersion, "got %s, expected qube", amount.GetDenomByIndex(0)))
	}

	price := sdk.NewInt(k.oracleKeeper.GetPrice(ctx) / 100000)

	qubeAmount := amount.AmountOfNoDenomValidation("qube")
	usqAmount := (qubeAmount.Mul(price)).Quo(sdk.NewInt(10))

	usq := sdk.NewCoin("usq", usqAmount)
	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(usq))
	if err != nil {
		return nil, err
	}

	//TODO: delegate to random validator

	validator := k.stakingKeeper.GetAllValidators(ctx)[0]
	_, err = k.stakingKeeper.Delegate(ctx, creator, qubeAmount, stakingtypes.Bonded, validator, false)
	if err != nil {
		panic(err)
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creator, sdk.NewCoins(usq))
	if err != nil {
		return nil, err
	}

	return &types.MsgMintResponse{}, nil
}
