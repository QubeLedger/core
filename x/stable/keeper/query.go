package keeper

import (
	"context"

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
	return &types.PairRequestResponse{
		PairId:            pair.PairId,
		AmountInMetadata:  pair.AmountInMetadata,
		AmountOutMetadata: pair.AmountOutMetadata,
		Qm:                pair.Qm,
		Ar:                pair.Ar,
		MinAmountIn:       pair.MinAmountIn,
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
	return &types.PairRequestResponse{
		PairId:            pair.PairId,
		AmountInMetadata:  pair.AmountInMetadata,
		AmountOutMetadata: pair.AmountOutMetadata,
		Qm:                pair.Qm,
		Ar:                pair.Ar,
		MinAmountIn:       pair.MinAmountIn,
	}, nil
}
