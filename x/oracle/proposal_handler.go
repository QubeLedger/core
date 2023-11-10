package oracle

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/QuadrateOrg/core/x/oracle/keeper"
	"github.com/QuadrateOrg/core/x/oracle/types"
)

func NewOracleProposalHandler(k *keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.RegisterAddNewDenomProposal:
			return handleRegisterAddNewDenomProposal(ctx, k, c)
		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s proposal content type: %T", types.ModuleName, c)
		}
	}
}

func handleRegisterAddNewDenomProposal(ctx sdk.Context, k *keeper.Keeper, p *types.RegisterAddNewDenomProposal) error {
	denom := types.Denom{
		Name: p.Denom,
	}
	params := k.GetParams(ctx)

	params.Whitelist = append(params.Whitelist, denom)

	k.SetParams(ctx, params)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventRegisterAddNewDenomProposal,
		),
	)
	return nil
}
