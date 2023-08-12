package keeper_test

import (
	"fmt"

	"github.com/QuadrateOrg/core/x/stable/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *StableKeeperTestSuite) TestPairByPairId() {
	testCases := []struct {
		name   string
		pairID string
		amount uint64
		err    bool
	}{
		{
			"ok-found",
			s.GetNormalPair(0).PairId,
			1000,
			false,
		},
		{
			"fail-not find",
			"test",
			1000,
			true,
		},
	}
	suite.Setup()
	suite.Commit()

	for _, tc := range testCases {
		suite.app.StableKeeper.SetTestingMode(true)
		suite.Run(fmt.Sprintf("Case---%s", tc.name), func() {
			ctx := sdk.WrapSDKContext(suite.ctx)
			suite.app.StableKeeper.AppendPair(suite.ctx, s.GetNormalPair(0))
			suite.MintStable(10000000000, s.GetNormalPair(0))

			req := types.PairByPairIdRequest{
				PairId: tc.pairID,
			}

			pair, err := suite.app.StableKeeper.PairByPairId(ctx, &req)
			if !tc.err {
				suite.NoError(err)
				suite.Equal(&types.PairRequestResponse{
					PairId:            s.GetNormalPair(0).PairId,
					AmountInMetadata:  s.GetNormalPair(0).AmountInMetadata,
					AmountOutMetadata: s.GetNormalPair(0).AmountOutMetadata,
					Qm:                s.GetNormalPair(0).Qm,
					Ar:                s.GetNormalPair(0).Ar,
					MinAmountIn:       s.GetNormalPair(0).MinAmountIn,
					MinAmountOut:      s.GetNormalPair(0).MinAmountOut,
				}, pair)
			} else {
				suite.Error(err)
			}
		})
	}
}

func (suite *StableKeeperTestSuite) TestPairById() {
	testCases := []struct {
		name   string
		id     uint64
		amount uint64
		err    bool
	}{
		{
			"ok-found",
			s.GetNormalPair(0).Id,
			1000,
			false,
		},
		{
			"fail-not find",
			10,
			1000,
			true,
		},
	}
	suite.Setup()
	suite.Commit()

	for _, tc := range testCases {
		suite.app.StableKeeper.SetTestingMode(true)
		suite.Run(fmt.Sprintf("Case---%s", tc.name), func() {
			ctx := sdk.WrapSDKContext(suite.ctx)
			suite.app.StableKeeper.AppendPair(suite.ctx, s.GetNormalPair(0))
			suite.MintStable(10000000000, s.GetNormalPair(0))

			req := types.PairByIdRequest{
				Id: tc.id,
			}

			pair, err := suite.app.StableKeeper.PairById(ctx, &req)
			if !tc.err {
				suite.NoError(err)
				suite.Equal(&types.PairRequestResponse{
					PairId:            s.GetNormalPair(0).PairId,
					AmountInMetadata:  s.GetNormalPair(0).AmountInMetadata,
					AmountOutMetadata: s.GetNormalPair(0).AmountOutMetadata,
					Qm:                s.GetNormalPair(0).Qm,
					Ar:                s.GetNormalPair(0).Ar,
					MinAmountIn:       s.GetNormalPair(0).MinAmountIn,
					MinAmountOut:      s.GetNormalPair(0).MinAmountOut,
				}, pair)
			} else {
				suite.Error(err)
			}
		})
	}
}

func (suite *StableKeeperTestSuite) TestGetAmountOutByAmountIn() {
	testCases := []struct {
		name   string
		pair   types.Pair
		amount uint64
		action string
	}{
		{
			"ok-mint",
			s.GetNormalPair(0),
			1000,
			"mint",
		},
		{
			"ok-burn",
			s.GetNormalPair(0),
			1000,
			"burn",
		},
	}
	suite.Setup()
	suite.Commit()

	for _, tc := range testCases {
		suite.app.StableKeeper.SetTestingMode(true)
		suite.Run(fmt.Sprintf("Case---%s", tc.name), func() {
			ctx := sdk.WrapSDKContext(suite.ctx)
			suite.app.StableKeeper.AppendPair(suite.ctx, suite.GetNormalPair(0))
			suite.MintStable(10000000000, s.GetNormalPair(0))

			req := types.GetAmountOutByAmountIn{
				PairId:   tc.pair.PairId,
				AmountIn: tc.amount,
				Action:   tc.action,
			}

			_, err := suite.app.StableKeeper.GetAmountOutByAmountIn(ctx, &req)
			suite.NoError(err)
		})
	}
}
