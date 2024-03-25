package keeper

import (
	"context"

	"github.com/QuadrateOrg/core/x/grow/types"
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

func (k Keeper) AssetByAssetId(goCtx context.Context, req *types.QueryAssetByAssetIdRequest) (*types.QueryAssetByAssetIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	Asset, found := k.GetAssetByAssetId(ctx, req.Id)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	sir := 0.0
	bir := 0.0

	utilization_rate := (float64(Asset.CollectivelyBorrowValue) / float64(Asset.ProvideValue))
	if utilization_rate > 0 {
		bir_temp, sir_temp, err := k.GetRatesByUtilizationRate(ctx, utilization_rate, Asset)
		if err != nil {
			return nil, types.ErrCalculateBIROrSIR
		}
		sir = sir_temp
		bir = bir_temp
	}

	return &types.QueryAssetByAssetIdResponse{
		Asset:              Asset,
		SupplyInterestRate: sir,
		BorrowInterestRate: bir,
	}, nil
}

func (k Keeper) GetAllAssets(goCtx context.Context, req *types.QueryGetAllAssetsRequest) (*types.QueryGetAllAssetsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	assets := k.GetAllAsset(ctx)

	return &types.QueryGetAllAssetsResponse{
		Assets: assets,
	}, nil
}

func (k Keeper) PositionById(goCtx context.Context, req *types.QueryPositionByIdRequest) (*types.QueryPositionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	pos, found := k.GetPositionByPositionId(ctx, req.Id)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryPositionResponse{
		Position: pos,
	}, nil
}

func (k Keeper) PositionByCreator(goCtx context.Context, req *types.QueryPositionByCreatorRequest) (*types.QueryPositionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	allPos := k.GetAllPosition(ctx)
	var pos types.Position
	var found = false
	for _, ps := range allPos {
		if ps.Creator == req.Creator {
			pos = ps
			found = true
		}
	}

	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryPositionResponse{
		Position: pos,
	}, nil
}

func (k Keeper) AllPosition(goCtx context.Context, req *types.QueryAllPositionRequest) (*types.QueryAllPositionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	allPos := k.GetAllPosition(ctx)

	if len(allPos) == 0 {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryAllPositionResponse{
		Positions: allPos,
	}, nil
}

func (k Keeper) AllLiquidatorPosition(goCtx context.Context, req *types.QueryAllLiquidatorPositionRequest) (*types.QueryAllLiquidatorPositionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	allPos := k.GetAllLiquidatorPosition(ctx)

	return &types.QueryAllLiquidatorPositionResponse{
		LiquidatorsPosition: allPos,
	}, nil
}

func (k Keeper) LiquidatorPositionByCreator(goCtx context.Context, req *types.QueryLiquidatorPositionByCreatorRequest) (*types.QueryLiquidatorPositionByCreatorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	allPos := k.GetAllLiquidatorPosition(ctx)
	var pos []types.LiquidatorPosition
	var found = false
	for _, ps := range allPos {
		if ps.Liquidator == req.Creator {
			pos = append(pos, ps)
			found = true
		}
	}

	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryLiquidatorPositionByCreatorResponse{
		Position: pos,
	}, nil
}

func (k Keeper) LiquidatorPositionById(goCtx context.Context, req *types.QueryLiquidatorPositionByIdRequest) (*types.QueryLiquidatorPositionByIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	pos, found := k.GetLiquidatorPositionByLiquidatorPositionId(ctx, req.Id)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryLiquidatorPositionByIdResponse{
		LiquidatorsPosition: pos,
	}, nil
}

func (k Keeper) LoanById(goCtx context.Context, req *types.QueryLoanByIdRequest) (*types.QueryLoanByIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	loan, found := k.GetLoadByLoanId(ctx, req.Id)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryLoanByIdResponse{
		Loan: loan,
	}, nil
}

func (k Keeper) YieldPercentage(goCtx context.Context, req *types.QueryYieldPercentageRequest) (*types.QueryYieldPercentageResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	gTokenPair, found := k.GetPairByDenomID(ctx, req.Id)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	RealRate := k.GetRealRate(ctx, gTokenPair.DenomID)
	BorrowRate := k.GetBorrowRate(ctx, gTokenPair.DenomID)

	growYield, err := k.CalculateGrowYield(ctx, gTokenPair)
	if err != nil {
		return nil, err
	}
	realYield, err := k.CalculateRealYield(ctx, gTokenPair)
	if err != nil {
		return nil, err
	}
	action, value, err := k.CheckYieldRate(ctx, gTokenPair)
	if err != nil {
		return nil, err
	}
	return &types.QueryYieldPercentageResponse{
		RealRate:     RealRate.Int64(),
		BorrowRate:   BorrowRate.Int64(),
		RealYield:    realYield.Int64(),
		GrowYield:    growYield.Int64(),
		ActualAction: action,
		Difference:   value.Int64(),
	}, nil
}

func (k Keeper) PairByDenomId(goCtx context.Context, req *types.PairByDenomIdRequest) (*types.PairByDenomIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	pair, found := k.GetPairByDenomID(ctx, req.DenomId)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.PairByDenomIdResponse{
		Pair: &pair,
	}, nil
}

func (k Keeper) AllPairs(goCtx context.Context, req *types.AllPairsRequest) (*types.AllPairsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	allPair := k.GetAllPair(ctx)

	if len(allPair) == 0 {
		return nil, status.Error(codes.NotFound, "not found")
	}

	pairs := []*types.GTokenPair{}

	for _, pair := range allPair {
		pairs = append(pairs, &pair) // #nosec
	}

	return &types.AllPairsResponse{
		Pairs: pairs,
	}, nil
}

func (k Keeper) LendById(goCtx context.Context, req *types.QueryLendByIdRequest) (*types.QueryLendByIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	lend, found := k.GetLendByLendId(ctx, req.Id)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryLendByIdResponse{
		Lend: lend,
	}, nil
}
