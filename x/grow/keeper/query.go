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

func (k Keeper) LendAssetByLendAssetId(goCtx context.Context, req *types.QueryLendAssetByLendAssetIdRequest) (*types.QueryLendAssetByLendAssetIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	lendAsset, found := k.GetLendAssetByLendAssetId(ctx, req.Id)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryLendAssetByLendAssetIdResponse{
		LendAssetId:   lendAsset.LendAssetId,
		AssetMetadata: lendAsset.AssetMetadata,
		OracleAssetId: lendAsset.OracleAssetId,
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
		Creator:             pos.Creator,
		DepositId:           pos.DepositId,
		Collateral:          pos.Collateral,
		OracleTicker:        pos.OracleTicker,
		BorrowedAmountInUSD: pos.BorrowedAmountInUSD,
		LoanIds:             pos.LoanIds,
	}, nil
}

func (k Keeper) PositionByCreator(goCtx context.Context, req *types.QueryPositionByCreatorRequest) (*types.QueryPositionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	allPos := k.GetAllPosition(ctx)
	var pos types.Position
	for _, ps := range allPos {
		if ps.Creator == req.Creator {
			pos = ps
		}
	}

	return &types.QueryPositionResponse{
		Creator:             pos.Creator,
		DepositId:           pos.DepositId,
		Collateral:          pos.Collateral,
		OracleTicker:        pos.OracleTicker,
		BorrowedAmountInUSD: pos.BorrowedAmountInUSD,
		LoanIds:             pos.LoanIds,
	}, nil
}

func (k Keeper) AllPosition(goCtx context.Context, req *types.QueryAllPositionRequest) (*types.QueryAllPositionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	allPos := k.GetAllPosition(ctx)

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
	var pos types.LiquidatorPosition
	for _, ps := range allPos {
		if ps.Liquidator == req.Creator {
			pos = ps
		}
	}

	return &types.QueryLiquidatorPositionByCreatorResponse{
		LiquidatorPositionId: pos.LiquidatorPositionId,
		BorrowAssetId:        pos.BorrowAssetId,
		Liquidator:           pos.Liquidator,
		Amount:               pos.Amount,
		Premium:              pos.Premium,
	}, nil
}

func (k Keeper) LiquidatorPositionById(goCtx context.Context, req *types.QueryLiquidatorPositionByIdRequest) (*types.QueryLiquidatorPositionByCreatorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	pos, found := k.GetLiquidatorPositionByLiquidatorPositionId(ctx, req.Id)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryLiquidatorPositionByCreatorResponse{
		LiquidatorPositionId: pos.LiquidatorPositionId,
		BorrowAssetId:        pos.BorrowAssetId,
		Liquidator:           pos.Liquidator,
		Amount:               pos.Amount,
		Premium:              pos.Premium,
	}, nil
}

func (k Keeper) AllFundAddress(goCtx context.Context, req *types.QueryAllFundAddressRequest) (*types.QueryAllFundAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	USQReserveAddress := k.GetUSQReserveAddress(ctx)
	GrowYieldReserveAddress := k.GetGrowYieldReserveAddress(ctx)
	GrowStakingReserveAddress := k.GetGrowStakingReserveAddress(ctx)

	return &types.QueryAllFundAddressResponse{
		USQReserveAddress:         USQReserveAddress.String(),
		GrowYieldReserveAddress:   GrowYieldReserveAddress.String(),
		GrowStakingReserveAddress: GrowStakingReserveAddress.String(),
	}, nil
}

func (k Keeper) LoanById(goCtx context.Context, req *types.QueryLoanByIdRequest) (*types.QueryLoanByIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	loan, found := k.GetLoadByLoadId(ctx, req.Id)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryLoanByIdResponse{
		LoanId:       loan.LoanId,
		Borrower:     loan.Borrower,
		AmountOut:    loan.AmountOut,
		StartTime:    loan.StartTime,
		OracleTicker: loan.OracleTicker,
	}, nil
}

func (k Keeper) YieldPercentage(goCtx context.Context, req *types.QueryYieldPercentageRequest) (*types.QueryYieldPercentageResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	RealRate := k.GetRealRate(ctx)
	BorrowRate := k.GetBorrowRate(ctx)

	gTokenPair, found := k.GetPairByDenomID(ctx, req.Id)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

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
