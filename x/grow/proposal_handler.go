package grow

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/QuadrateOrg/core/x/grow/keeper"
	"github.com/QuadrateOrg/core/x/grow/types"
)

func NewGrowProposalHandler(k *keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.RegisterLendAssetProposal:
			return handleRegisterAssetProposal(ctx, k, c)
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
		case *types.RegisterChangeDepositMethodStatusProposal:
			return handelRegisterChangeDepositMethodStatusProposal(ctx, k, c)
		case *types.RegisterChangeCollateralMethodStatusProposal:
			return handelRegisterChangeCollateralMethodStatusProposal(ctx, k, c)
		case *types.RegisterChangeBorrowMethodStatusProposal:
			return handleRegisterChangeBorrowMethodStatusProposal(ctx, k, c)
		case *types.RegisterRemoveLendAssetProposal:
			return handleRegisterRemoveAssetProposal(ctx, k, c)
		case *types.RegisterRemoveGTokenPairProposal:
			return handleRegisterRemoveGTokenPairProposal(ctx, k, c)
		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s proposal content type: %T", types.ModuleName, c)
		}
	}
}

func handleRegisterAssetProposal(ctx sdk.Context, k *keeper.Keeper, p *types.RegisterLendAssetProposal) error {
	la := types.Asset{
		AssetId:                 k.GenerateAssetIdHash(p.AssetMetadata.Base),
		AssetMetadata:           p.AssetMetadata,
		OracleAssetId:           p.OracleAssetId,
		ProvideValue:            0,
		CollectivelyBorrowValue: 0,
		Type:                    p.Type,
	}
	err := k.RegisterAsset(ctx, la)
	if err != nil {
		return err
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventRegisterAssetProposal,
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
	err := k.SetRealRate(ctx, sdk.NewIntFromUint64(p.Rate), p.Id)
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
	err := k.SetBorrowRate(ctx, sdk.NewIntFromUint64(p.Rate), p.Id)
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

func handleRegisterRemoveAssetProposal(ctx sdk.Context, k *keeper.Keeper, p *types.RegisterRemoveLendAssetProposal) error {
	asset, found := k.GetAssetByAssetId(ctx, p.LendAssetId)
	if !found {
		return types.ErrAssetNotFound
	}
	k.RemoveAsset(ctx, asset.Id)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventRegisterRemoveAssetProposal,
		),
	)
	return nil
}

func handleRegisterRemoveGTokenPairProposal(ctx sdk.Context, k *keeper.Keeper, p *types.RegisterRemoveGTokenPairProposal) error {
	pair, found := k.GetPairByDenomID(ctx, p.GTokenPairID)
	if !found {
		return types.ErrPairNotFound
	}
	k.RemovePair(ctx, pair.Id)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventRegisterRemoveGTokenPairProposal,
		),
	)
	return nil
}

func handelRegisterChangeDepositMethodStatusProposal(ctx sdk.Context, k *keeper.Keeper, p *types.RegisterChangeDepositMethodStatusProposal) error {
	k.ChangeDepositMethodStatus(ctx)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventRegisterChangeDepositMethodStatusProposal,
		),
	)
	return nil
}

func handelRegisterChangeCollateralMethodStatusProposal(ctx sdk.Context, k *keeper.Keeper, p *types.RegisterChangeCollateralMethodStatusProposal) error {
	k.ChangeCollateralMethodStatus(ctx)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventRegisterChangeCollateralMethodStatusProposal,
		),
	)
	return nil
}

func handleRegisterChangeBorrowMethodStatusProposal(ctx sdk.Context, k *keeper.Keeper, p *types.RegisterChangeBorrowMethodStatusProposal) error {
	k.ChangeBorrowMethodStatus(ctx)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventRegisterChangeBorrowMethodStatusProposal,
		),
	)
	return nil
}
