package keeper

import (
	"context"

	gmd "github.com/QuadrateOrg/core/x/stable/gmb"
	"github.com/QuadrateOrg/core/x/stable/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Params(goCtx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := k.GetParams(ctx)

	return &types.QueryParamsResponse{Params: &params}, nil
}

func (k Keeper) PairByPairId(goCtx context.Context, req *types.PairByPairIdRequest) (*types.PairRequestResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	pair, found := k.GetPairByPairID(ctx, req.PairId)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	atomPrice, err := k.GetAtomPrice(ctx)
	if err != nil {
		return nil, err
	}
	qm, ar := pair.Qm, pair.Ar

	backing_ratio, err = CalculateBackingRatio(qm, ar, atomPrice)
	if err != nil {
		return nil, err
	}

	mintingFee, err := gmd.CalculateMintingFee(backing_ratio)
	if err != nil {
		return nil, err
	}

	burning_fee, err := gmd.CalculateBurningFee(backing_ratio)
	if err != nil {
		return nil, err
	}

	return &types.PairRequestResponse{
		PairId:            pair.PairId,
		AmountInMetadata:  pair.AmountInMetadata,
		AmountOutMetadata: pair.AmountOutMetadata,
		Qm:                pair.Qm,
		Ar:                pair.Ar,
		MinAmountIn:       pair.MinAmountIn,
		MinAmountOut:      pair.MinAmountOut,
		BackingRatio:      backing_ratio.Uint64(),
		MintingFee:        mintingFee.Uint64(),
		BurningFee:        burning_fee.Uint64(),
	}, nil
}

func (k Keeper) PairById(goCtx context.Context, req *types.PairByIdRequest) (*types.PairRequestResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	pair, found := k.GetPairByID(ctx, req.Id)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	atomPrice, err := k.GetAtomPrice(ctx)
	if err != nil {
		return nil, err
	}
	qm, ar := pair.Qm, pair.Ar

	backing_ratio, err = CalculateBackingRatio(qm, ar, atomPrice)
	if err != nil {
		return nil, err
	}

	mintingFee, err := gmd.CalculateMintingFee(backing_ratio)
	if err != nil {
		return nil, err
	}

	burning_fee, err := gmd.CalculateBurningFee(backing_ratio)
	if err != nil {
		return nil, err
	}

	return &types.PairRequestResponse{
		PairId:            pair.PairId,
		AmountInMetadata:  pair.AmountInMetadata,
		AmountOutMetadata: pair.AmountOutMetadata,
		Qm:                pair.Qm,
		Ar:                pair.Ar,
		MinAmountIn:       pair.MinAmountIn,
		MinAmountOut:      pair.MinAmountOut,
		BackingRatio:      backing_ratio.Uint64(),
		MintingFee:        mintingFee.Uint64(),
		BurningFee:        burning_fee.Uint64(),
	}, nil
}

func (k Keeper) GetAmountOutByAmountIn(goCtx context.Context, req *types.GetAmountOutByAmountInRequest) (*types.GetAmountOutByAmountInResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	pair, found := k.GetPairByPairID(ctx, req.PairId)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	atomPrice, err := k.GetAtomPrice(ctx)
	if err != nil {
		return nil, err
	}
	qm, ar := pair.Qm, pair.Ar

	backing_ratio, err = CalculateBackingRatio(qm, ar, atomPrice)
	if err != nil {
		return nil, err
	}
	switch req.Action {
	case "mint":
		mintingFee, err := gmd.CalculateMintingFee(backing_ratio)
		if err != nil {
			return nil, err
		}
		amountOutToMint := k.CalculateAmountToMint(sdk.NewIntFromUint64(req.AmountIn), atomPrice, mintingFee)
		return &types.GetAmountOutByAmountInResponse{
			PairId:    req.PairId,
			AmountOut: amountOutToMint.Uint64(),
			Action:    req.Action,
		}, nil
	case "burn":
		burningFee, err := gmd.CalculateBurningFee(backing_ratio)
		if err != nil {
			return nil, err
		}
		amountOutToSend := k.CalculateAmountToSend(sdk.NewIntFromUint64(req.AmountIn), atomPrice, burningFee)
		return &types.GetAmountOutByAmountInResponse{
			PairId:    req.PairId,
			AmountOut: amountOutToSend.Uint64(),
			Action:    req.Action,
		}, nil
	default:
		return nil, status.Error(codes.NotFound, "action not found")
	}
}
