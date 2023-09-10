package grow

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/QuadrateOrg/core/x/grow/keeper"
	"github.com/QuadrateOrg/core/x/grow/types"
)

func NewStableProposalHandler(k *keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.RegisterLendAssetProposal:
			return handleRegisterLendAssetProposal(ctx, k, c)
		case *types.RegisterGTokenPairProposal:
			return handleRegisterGTokenPairProposal(ctx, k, c)
		case *types.RegisterChangeUSQReserveAddressProposal:
			return handleRegisterChangeUSQReserveAddressProposal(ctx, k, c)
		case *types.RegisterChangeGrowYieldReserveAddressProposal:
			return handleRegisterChangeGrowYieldReserveAddressProposal(ctx, k, c)
		case *types.RegisterChangeGrowStakingReserveAddressProposal:
			return handleRegisterChangeGrowStakingReserveAddressProposal(ctx, k, c)
		case *types.RegisterChangeRealRateProposal:
			return handleRegisterChangeRealRateProposal(ctx, k, c)
		case *types.RegisterChangeBorrowRateProposal:
			return handleRegisterChangeBorrowRateProposal(ctx, k, c)
		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s proposal content type: %T", types.ModuleName, c)
		}
	}
}

func handleRegisterLendAssetProposal(ctx sdk.Context, k *keeper.Keeper, p *types.RegisterLendAssetProposal) error {
	la := types.LendAsset{
		LendAssetId:   k.GenerateLendAssetIdHash(p.AssetMetadata.Base),
		AssetMetadata: p.AssetMetadata,
		OracleAssetId: p.OracleAssetId,
	}
	err := k.RegisterLendAsset(ctx, la)
	if err != nil {
		return err
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventRegisterLendAssetProposal,
		),
	)
	return nil
}

func handleRegisterGTokenPairProposal(ctx sdk.Context, k *keeper.Keeper, p *types.RegisterGTokenPairProposal) error {
	pair := types.GTokenPair{
		GTokenMetadata: p.GTokenMetadata,
		QStablePairId:  p.QStablePairId,
		MinAmountIn:    p.MinAmountIn,
		MinAmountOut:   p.MinAmountOut,
	}
	err := k.RegisterPair(ctx, pair)
	if err != nil {
		return err
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventRegisterGTokenPairProposal,
		),
	)
	return nil
}

func handleRegisterChangeGrowYieldReserveAddressProposal(ctx sdk.Context, k *keeper.Keeper, p *types.RegisterChangeGrowYieldReserveAddressProposal) error {
	address, err := sdk.AccAddressFromBech32(p.Address)
	if err != nil {
		return err
	}
	err = k.ChangeGrowYieldReserveAddress(ctx, address)
	if err != nil {
		return err
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventRegisterChangeGrowYieldReserveAddressProposal,
		),
	)
	return nil
}

func handleRegisterChangeUSQReserveAddressProposal(ctx sdk.Context, k *keeper.Keeper, p *types.RegisterChangeUSQReserveAddressProposal) error {
	address, err := sdk.AccAddressFromBech32(p.Address)
	if err != nil {
		return err
	}
	err = k.ChangeUSQReserveAddress(ctx, address)
	if err != nil {
		return err
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventRegisterChangeUSQReserveAddressProposal,
		),
	)
	return nil
}

func handleRegisterChangeGrowStakingReserveAddressProposal(ctx sdk.Context, k *keeper.Keeper, p *types.RegisterChangeGrowStakingReserveAddressProposal) error {
	address, err := sdk.AccAddressFromBech32(p.Address)
	if err != nil {
		return err
	}
	err = k.ChangeGrowStakingReserveAddress(ctx, address)
	if err != nil {
		return err
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventRegisterChangeGrowStakingReserveAddressProposal,
		),
	)
	return nil
}

func handleRegisterChangeRealRateProposal(ctx sdk.Context, k *keeper.Keeper, p *types.RegisterChangeRealRateProposal) error {
	err := k.SetRealRate(ctx, sdk.NewIntFromUint64(p.Rate))
	if err != nil {
		return err
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventRegisterChangeRealRateProposal,
		),
	)
	return nil
}

func handleRegisterChangeBorrowRateProposal(ctx sdk.Context, k *keeper.Keeper, p *types.RegisterChangeBorrowRateProposal) error {
	err := k.SetBorrowRate(ctx, sdk.NewIntFromUint64(p.Rate))
	if err != nil {
		return err
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventRegisterChangeBorrowRateProposal,
		),
	)
	return nil
}