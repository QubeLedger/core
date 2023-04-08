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

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, amount)
	if err != nil {
		return nil, err
	}

	//TODO: get price from oracle
	price := sdk.NewInt(2 * 1000000)

	qubeAmount := amount.AmountOfNoDenomValidation("qube")
	usqAmount := qubeAmount.Mul(price.Sub(sdk.NewInt(1000000)))

	usq := sdk.NewCoin("usq", usqAmount)
	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(usq))
	if err != nil {
		return nil, err
	}

	validatorAddress, err := sdk.ValAddressFromBech32(types.ModuleName)
	validator, err1 := k.stakingKeeper.GetValidator(ctx, validatorAddress)
	if err1 != true {
		return nil, sdkerrors.Wrapf(types.ErrValNotFound, "validator not found")
	}

	//TODO: delegate to random validator
	_, err = k.stakingKeeper.Delegate(ctx, sdk.AccAddress(types.ModuleName), qubeAmount, stakingtypes.Unbonded, validator, true)
	if err != nil {
		panic(err)
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creator, sdk.NewCoins(usq))
	if err != nil {
		return nil, err
	}

	return &types.MsgMintResponse{}, nil
}
