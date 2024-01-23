package keeper

import (
	"context"
	"fmt"

	"github.com/QuadrateOrg/core/x/dex/types"
	dexutils "github.com/QuadrateOrg/core/x/dex/utils"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func FilteredPaginateAccountBalances(
	ctx sdk.Context,
	bankKeeper types.BankKeeper,
	address sdk.AccAddress,
	pageRequest *query.PageRequest,
	onResult func(coin sdk.Coin, accumulate bool) bool,
) (*query.PageResponse, error) {
	// if the PageRequest is nil, use default PageRequest
	if pageRequest == nil {
		pageRequest = &query.PageRequest{}
	}

	offset := pageRequest.Offset
	key := pageRequest.Key
	limit := pageRequest.Limit
	countTotal := pageRequest.CountTotal

	if pageRequest.Reverse {
		return nil, fmt.Errorf("invalid request, reverse pagination is not enabled")
	}
	if offset > 0 && key != nil {
		return nil, fmt.Errorf("invalid request, either offset or key is expected, got both")
	}

	if limit == 0 {
		limit = query.DefaultLimit

		// count total results when the limit is zero/not supplied
		countTotal = true
	}

	if len(key) != 0 {
		// paginate with key
		var (
			numHits uint64
			nextKey []byte
		)
		startAccum := false

		bankKeeper.IterateAccountBalances(ctx, address, func(coin sdk.Coin) bool {
			if coin.Denom == string(key) {
				startAccum = true
			}
			if numHits == limit {
				nextKey = []byte(coin.Denom)
				return true
			}
			if startAccum {
				hit := onResult(coin, true)
				if hit {
					numHits++
				}
			}

			return false
		})

		return &query.PageResponse{
			NextKey: nextKey,
		}, nil
	} else {
		// default pagination (with offset)

		end := offset + limit

		var (
			numHits uint64
			nextKey []byte
		)

		bankKeeper.IterateAccountBalances(ctx, address, func(coin sdk.Coin) bool {
			accumulate := numHits >= offset && numHits < end
			hit := onResult(coin, accumulate)

			if hit {
				numHits++
			}

			if numHits == end+1 {
				if nextKey == nil {
					nextKey = []byte(coin.Denom)
				}

				if !countTotal {
					return true
				}
			}

			return false
		})

		res := &query.PageResponse{NextKey: nextKey}
		if countTotal {
			res.Total = numHits
		}

		return res, nil
	}
}

func (k Keeper) UserDepositsAll(
	goCtx context.Context,
	req *types.QueryAllUserDepositsRequest,
) (*types.QueryAllUserDepositsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	var depositArr []*types.DepositRecord

	pageRes, err := FilteredPaginateAccountBalances(
		ctx,
		k.bankKeeper,
		addr,
		req.Pagination,
		func(poolCoinMaybe sdk.Coin, accumulate bool) bool {
			err := types.ValidatePoolDenom(poolCoinMaybe.Denom)
			if err != nil {
				return false
			}

			poolMetadata, err := k.GetPoolMetadataByDenom(ctx, poolCoinMaybe.Denom)
			if err != nil {
				panic("Can't get info for PoolDenom")
			}

			fee := dexutils.MustSafeUint64ToInt64(poolMetadata.Fee)

			if accumulate {
				depositRecord := &types.DepositRecord{
					PairID:          poolMetadata.PairID,
					SharesOwned:     poolCoinMaybe.Amount,
					CenterTickIndex: poolMetadata.Tick,
					LowerTickIndex:  poolMetadata.Tick - fee,
					UpperTickIndex:  poolMetadata.Tick + fee,
					Fee:             poolMetadata.Fee,
				}

				depositArr = append(depositArr, depositRecord)
			}

			return true
		})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllUserDepositsResponse{
		Deposits:   depositArr,
		Pagination: pageRes,
	}, nil
}
