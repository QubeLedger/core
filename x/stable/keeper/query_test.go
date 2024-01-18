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
			s.GetNormalGMBPair(0).PairId,
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
			suite.app.StableKeeper.AppendPair(suite.ctx, s.GetNormalGMBPair(0))
			suite.MintStable(10000000000, s.GetNormalGMBPair(0))

			req := types.PairByPairIdRequest{
				PairId: tc.pairID,
			}

			pair, err := suite.app.StableKeeper.PairByPairId(ctx, &req)
			if !tc.err {
				suite.NoError(err)
				suite.Equal(&types.PairRequestResponse{
					PairId:            s.GetNormalGMBPair(0).PairId,
					AmountInMetadata:  s.GetNormalGMBPair(0).AmountInMetadata,
					AmountOutMetadata: s.GetNormalGMBPair(0).AmountOutMetadata,
					Qm:                s.GetNormalGMBPair(0).Qm,
					Ar:                s.GetNormalGMBPair(0).Ar,
					MinAmountIn:       s.GetNormalGMBPair(0).MinAmountIn,
					MinAmountOut:      s.GetNormalGMBPair(0).MinAmountOut,
					BackingRatio:      100,
					MintingFee:        3,
					BurningFee:        2,
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
			s.GetNormalGMBPair(0).Id,
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
			suite.app.StableKeeper.AppendPair(suite.ctx, s.GetNormalGMBPair(0))
			suite.MintStable(10000000000, s.GetNormalGMBPair(0))

			req := types.PairByIdRequest{
				Id: tc.id,
			}

			pair, err := suite.app.StableKeeper.PairById(ctx, &req)
			if !tc.err {
				suite.NoError(err)
				suite.Equal(&types.PairRequestResponse{
					PairId:            s.GetNormalGMBPair(0).PairId,
					AmountInMetadata:  s.GetNormalGMBPair(0).AmountInMetadata,
					AmountOutMetadata: s.GetNormalGMBPair(0).AmountOutMetadata,
					Qm:                s.GetNormalGMBPair(0).Qm,
					Ar:                s.GetNormalGMBPair(0).Ar,
					MinAmountIn:       s.GetNormalGMBPair(0).MinAmountIn,
					MinAmountOut:      s.GetNormalGMBPair(0).MinAmountOut,
					BackingRatio:      100,
					MintingFee:        3,
					BurningFee:        2,
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
			s.GetNormalGMBPair(0),
			1000,
			"mint",
		},
		{
			"ok-burn",
			s.GetNormalGMBPair(0),
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
			suite.app.StableKeeper.AppendPair(suite.ctx, suite.GetNormalGMBPair(0))
			suite.MintStable(10000000000, s.GetNormalGMBPair(0))

			req := types.GetAmountOutByAmountInRequest{
				PairId:   tc.pair.PairId,
				AmountIn: tc.amount,
				Action:   tc.action,
			}

			_, err := suite.app.StableKeeper.GetAmountOutByAmountIn(ctx, &req)
			suite.NoError(err)
		})
	}
}

func (suite *StableKeeperTestSuite) TestAllPairs() {
	testCases := []struct {
		name   string
		pair   types.Pair
		amount uint64
	}{
		{
			"ok",
			s.GetNormalGMBPair(0),
			1,
		},
		{
			"ok",
			s.GetNormalGMBPair(0),
			2,
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case---%s", tc.name), func() {
			suite.Setup()
			suite.Commit()
			ctx := sdk.WrapSDKContext(suite.ctx)
			for i := 0; i < int(tc.amount); i++ {
				suite.app.StableKeeper.AppendPair(suite.ctx, suite.GetNormalGMBPair(0))
			}

			req := types.AllPairsRequest{}

			res, err := suite.app.StableKeeper.AllPairs(ctx, &req)
			suite.NoError(err)

			s.Require().Equal(len(res.Pairs), int(tc.amount))
		})
	}
}
