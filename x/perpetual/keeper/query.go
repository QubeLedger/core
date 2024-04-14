package keeper

import (
	"context"

	"github.com/QuadrateOrg/core/x/perpetual/types"
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

	return &types.QueryParamsResponse{
		Params: &params,
	}, nil
}

func (k Keeper) VaultById(goCtx context.Context, req *types.QueryVaultByIdRequest) (*types.QueryVaultByIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	vault, f := k.GetVaultByVaultId(ctx, req.Id)
	if !f {
		return nil, types.ErrVaultNotFound
	}

	return &types.QueryVaultByIdResponse{
		Vault: &vault,
	}, nil
}

func (k Keeper) PositionById(goCtx context.Context, req *types.QueryPositionByIdRequest) (*types.QueryPositionByIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	pos, f := k.GetPositionByPositionId(ctx, req.Id)
	if !f {
		return nil, types.ErrPositionNotFound
	}

	return &types.QueryPositionByIdResponse{
		Position: &pos,
	}, nil
}

func (k Keeper) AllPositions(goCtx context.Context, req *types.QueryEmpty) (*types.QueryAllPositionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	poss := k.GetAllPositions(ctx)

	positions := []*types.TradePosition{}

	for _, ps := range poss {
		positions = append(positions, &ps) // #nosec
	}

	return &types.QueryAllPositionsResponse{
		Positions: positions,
	}, nil
}

func (k Keeper) AllVaults(goCtx context.Context, req *types.QueryEmpty) (*types.QueryAllVaultsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	vs := k.GetAllVault(ctx)

	vaults := []*types.Vault{}

	for _, vs := range vs {
		vaults = append(vaults, &vs) // #nosec
	}

	return &types.QueryAllVaultsResponse{
		Vaults: vaults,
	}, nil
}
