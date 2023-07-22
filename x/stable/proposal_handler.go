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
		case *types.ChangeBaseTokenDenom:
			return handleChangeBaseTokenDenom(ctx, k, c)
		case *types.ChangeSendTokenDenom:
			return handleChangeSendTokenDenom(ctx, k, c)
		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s proposal content type: %T", types.ModuleName, c)
		}
	}
}

func NewChangeSendTokenDenomProposalHandler(k *keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.ChangeSendTokenDenom:
			return handleChangeSendTokenDenom(ctx, k, c)
		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s proposal content type: %T", types.ModuleName, c)
		}
	}
}

func handleChangeBaseTokenDenom(ctx sdk.Context, k *keeper.Keeper, p *types.ChangeBaseTokenDenom) error {
	for _, metadata := range p.Metadata {
		err := k.ChangeBaseTokenDenom(ctx, metadata)
		if err != nil {
			return err
		}
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventChangeBaseTokenDenom,
			),
		)
	}
	return nil
}

func handleChangeSendTokenDenom(ctx sdk.Context, k *keeper.Keeper, p *types.ChangeSendTokenDenom) error {
	for _, metadata := range p.Metadata {
		err := k.ChangeSendTokenDenom(ctx, metadata)
		if err != nil {
			return err
		}

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventChangeSendTokenDenom,
			),
		)
	}
	return nil
}
