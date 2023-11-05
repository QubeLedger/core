package stable

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/QuadrateOrg/core/x/stable/keeper"
	"github.com/QuadrateOrg/core/x/stable/types"
)

func NewStableProposalHandler(k *keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.RegisterPairProposal:
			return handleRegisterPairProposal(ctx, k, c)
		case *types.RegisterChangeBurningFundAddressProposal:
			return handleRegisterChangeBurningFundAddressProposal(ctx, k, c)
		case *types.RegisterChangeReserveFundAddressProposal:
			return handleRegisterChangeReserveFundAddressProposal(ctx, k, c)
		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s proposal content type: %T", types.ModuleName, c)
		}
	}
}

func handleRegisterPairProposal(ctx sdk.Context, k *keeper.Keeper, p *types.RegisterPairProposal) error {
	pair := types.Pair{
		AmountInMetadata:  p.AmountInMetadata,
		AmountOutMetadata: p.AmountOutMetadata,
		MinAmountIn:       p.MinAmountIn,
	}
	err := k.RegisterPair(ctx, pair)
	if err != nil {
		return err
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventRegisterCreateNewPairProposal,
		),
	)
	return nil
}

func handleRegisterChangeBurningFundAddressProposal(ctx sdk.Context, k *keeper.Keeper, p *types.RegisterChangeBurningFundAddressProposal) error {
	address, err := sdk.AccAddressFromBech32(p.Address)
	if err != nil {
		return err
	}
	err = k.ChangeBurningFundAddress(ctx, address)
	if err != nil {
		return err
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventRegisterChangeBurningFundAddressProposal,
		),
	)
	return nil
}

func handleRegisterChangeReserveFundAddressProposal(ctx sdk.Context, k *keeper.Keeper, p *types.RegisterChangeReserveFundAddressProposal) error {
	address, err := sdk.AccAddressFromBech32(p.Address)
	if err != nil {
		return err
	}
	err = k.ChangeReserveFundAddress(ctx, address)
	if err != nil {
		return err
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventRegisterChangeReserveFundAddressProposal,
		),
	)
	return nil
}
