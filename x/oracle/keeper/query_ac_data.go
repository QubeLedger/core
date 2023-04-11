package keeper

import (
	"context"

	"github.com/QuadrateOrg/core/x/oracle/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AcDataAll(goCtx context.Context, req *types.QueryAllAcDataRequest) (*types.QueryAllAcDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var acDatas []types.AcData
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	acDataStore := prefix.NewStore(store, types.KeyPrefix(types.AcDataKey))

	pageRes, err := query.Paginate(acDataStore, req.Pagination, func(key []byte, value []byte) error {
		var acData types.AcData
		if err := k.cdc.Unmarshal(value, &acData); err != nil {
			return err
		}

		acDatas = append(acDatas, acData)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAcDataResponse{AcData: acDatas, Pagination: pageRes}, nil
}

func (k Keeper) AcData(goCtx context.Context, req *types.QueryGetAcDataRequest) (*types.QueryGetAcDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	acData, found := k.GetAcData(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetAcDataResponse{AcData: acData}, nil
}
