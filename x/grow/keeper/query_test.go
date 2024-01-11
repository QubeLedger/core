package keeper_test

import (
	"fmt"
	"time"

	apptesting "github.com/QuadrateOrg/core/app/apptesting"
	"github.com/QuadrateOrg/core/x/grow/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *GrowKeeperTestSuite) TestLendAssetByLendAssetId() {
	testCases := []struct {
		name      string
		lendAsset types.LendAsset
		id        string
		err       bool
		errString string
	}{
		{
			"ok-get lend asset",
			s.GetNormalLendAsset(0),
			s.GetNormalLendAsset(0).LendAssetId,
			false,
			"",
		},
		{
			"false-lend asset not found",
			s.GetNormalLendAsset(0),
			"test",
			true,
			"not found",
		},
	}
	suite.Setup()
	suite.Commit()
	suite.SetupOracleKeeper("ATOM")
	suite.RegisterValidator()
	suite.app.GrowKeeper.ChangeGrowStatus()
	for _, tc := range testCases {
		suite.app.GrowKeeper.AppendLendAsset(s.ctx, tc.lendAsset)
		suite.Run(fmt.Sprintf("Case---%s", tc.name), func() {
			ctx := sdk.WrapSDKContext(suite.ctx)

			req := types.QueryLendAssetByLendAssetIdRequest{
				Id: tc.id,
			}

			_, err := suite.app.GrowKeeper.LendAssetByLendAssetId(ctx, &req)
			if !tc.err {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err, tc.errString)
			}
		})
	}
}

func (suite *GrowKeeperTestSuite) TestPositionById() {
	suite.Setup()
	suite.Commit()
	testCases := []struct {
		name      string
		position  types.Position
		id        string
		err       bool
		errString string
	}{
		{
			"ok-get position",
			s.GetNormalPosition(),
			s.GetNormalPosition().DepositId,
			false,
			"",
		},
		{
			"false-position not found",
			s.GetNormalPosition(),
			"test",
			true,
			"not found",
		},
	}
	suite.SetupOracleKeeper("ATOM")
	suite.RegisterValidator()
	suite.app.GrowKeeper.ChangeGrowStatus()
	for _, tc := range testCases {
		suite.app.GrowKeeper.AppendPosition(s.ctx, tc.position)
		suite.Run(fmt.Sprintf("Case---%s", tc.name), func() {
			ctx := sdk.WrapSDKContext(suite.ctx)

			req := types.QueryPositionByIdRequest{
				Id: tc.id,
			}

			_, err := suite.app.GrowKeeper.PositionById(ctx, &req)
			if !tc.err {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err, tc.errString)
			}
		})
	}
}

func (suite *GrowKeeperTestSuite) TestPositionByCreator() {
	suite.Setup()
	suite.Commit()
	testCases := []struct {
		name      string
		position  types.Position
		address   string
		err       bool
		errString string
	}{
		{
			"ok-get position",
			s.GetNormalPosition(),
			s.GetNormalPosition().Creator,
			false,
			"",
		},
		{
			"false-position not found",
			s.GetNormalPosition(),
			"test",
			true,
			"not found",
		},
	}
	suite.SetupOracleKeeper("ATOM")
	suite.RegisterValidator()
	suite.app.GrowKeeper.ChangeGrowStatus()
	for _, tc := range testCases {
		suite.app.GrowKeeper.AppendPosition(s.ctx, tc.position)
		suite.Run(fmt.Sprintf("Case---%s", tc.name), func() {
			ctx := sdk.WrapSDKContext(suite.ctx)

			req := types.QueryPositionByCreatorRequest{
				Creator: tc.address,
			}

			_, err := suite.app.GrowKeeper.PositionByCreator(ctx, &req)
			if !tc.err {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err, tc.errString)
			}
		})
	}
}

func (suite *GrowKeeperTestSuite) TestLiquidatorPositionByCreator() {
	suite.Setup()
	suite.Commit()
	testCases := []struct {
		name      string
		position  types.LiquidatorPosition
		address   string
		err       bool
		errString string
	}{
		{
			"ok-get position",
			s.GetNormalLiqPosition(),
			s.GetNormalLiqPosition().Liquidator,
			false,
			"",
		},
		{
			"false-position not found",
			s.GetNormalLiqPosition(),
			"test",
			true,
			"not found",
		},
	}
	suite.SetupOracleKeeper("ATOM")
	suite.RegisterValidator()
	suite.app.GrowKeeper.ChangeGrowStatus()
	for _, tc := range testCases {
		suite.app.GrowKeeper.AppendLiquidatorPosition(s.ctx, tc.position)
		suite.Run(fmt.Sprintf("Case---%s", tc.name), func() {
			ctx := sdk.WrapSDKContext(suite.ctx)

			req := types.QueryLiquidatorPositionByCreatorRequest{
				Creator: tc.address,
			}

			_, err := suite.app.GrowKeeper.LiquidatorPositionByCreator(ctx, &req)
			if !tc.err {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err, tc.errString)
			}
		})
	}
}

func (suite *GrowKeeperTestSuite) TestLiquidatorPositionById() {
	suite.Setup()
	suite.Commit()
	testCases := []struct {
		name      string
		position  types.LiquidatorPosition
		id        string
		err       bool
		errString string
	}{
		{
			"ok-get position",
			s.GetNormalLiqPosition(),
			s.GetNormalLiqPosition().LiquidatorPositionId,
			false,
			"",
		},
		{
			"false-position not found",
			s.GetNormalLiqPosition(),
			"test",
			true,
			"not found",
		},
	}
	suite.SetupOracleKeeper("ATOM")
	suite.RegisterValidator()
	suite.app.GrowKeeper.ChangeGrowStatus()
	for _, tc := range testCases {
		suite.app.GrowKeeper.AppendLiquidatorPosition(s.ctx, tc.position)
		suite.Run(fmt.Sprintf("Case---%s", tc.name), func() {
			ctx := sdk.WrapSDKContext(suite.ctx)

			req := types.QueryLiquidatorPositionByIdRequest{
				Id: tc.id,
			}

			_, err := suite.app.GrowKeeper.LiquidatorPositionById(ctx, &req)
			if !tc.err {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err, tc.errString)
			}
		})
	}
}

func (suite *GrowKeeperTestSuite) TestAllFundAddress() {
	suite.Setup()
	suite.Commit()
	suite.SetupOracleKeeper("ATOM")
	suite.RegisterValidator()

	address := []sdk.AccAddress{
		apptesting.CreateRandomAccounts(1)[0],
		apptesting.CreateRandomAccounts(1)[0],
		apptesting.CreateRandomAccounts(1)[0],
	}

	suite.app.GrowKeeper.ChangeGrowStatus()
	suite.app.GrowKeeper.SetUSQReserveAddress(s.ctx, address[0])
	suite.app.GrowKeeper.SetGrowYieldReserveAddress(s.ctx, address[1])
	suite.app.GrowKeeper.SetGrowStakingReserveAddress(s.ctx, address[2])

	suite.Run(fmt.Sprintf("Case---%s", "found all address"), func() {
		ctx := sdk.WrapSDKContext(suite.ctx)

		req := types.QueryAllFundAddressRequest{}

		res, err := suite.app.GrowKeeper.AllFundAddress(ctx, &req)
		suite.Require().NoError(err)

		suite.Require().Equal(res.USQReserveAddress, address[0].String())
		suite.Require().Equal(res.GrowYieldReserveAddress, address[1].String())
		suite.Require().Equal(res.GrowStakingReserveAddress, address[2].String())

	})
}

func (suite *GrowKeeperTestSuite) TestLoanById() {
	suite.Setup()
	suite.Commit()
	testCases := []struct {
		name      string
		loan      types.Loan
		id        string
		err       bool
		errString string
	}{
		{
			"ok-get loan",
			s.GetNormalLoan(),
			s.GetNormalLoan().LoanId,
			false,
			"",
		},
		{
			"false-loan not found",
			s.GetNormalLoan(),
			"test",
			true,
			"not found",
		},
	}
	suite.SetupOracleKeeper("ATOM")
	suite.RegisterValidator()
	suite.app.GrowKeeper.ChangeGrowStatus()
	for _, tc := range testCases {
		suite.app.GrowKeeper.AppendLoan(s.ctx, tc.loan)
		suite.Run(fmt.Sprintf("Case---%s", tc.name), func() {
			ctx := sdk.WrapSDKContext(suite.ctx)

			req := types.QueryLoanByIdRequest{
				Id: tc.id,
			}

			_, err := suite.app.GrowKeeper.LoanById(ctx, &req)
			if !tc.err {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err, tc.errString)
			}
		})
	}
}

func (suite *GrowKeeperTestSuite) TestYieldPercentage() {
	suite.Setup()
	suite.Commit()
	suite.SetupOracleKeeper("ATOM")
	suite.RegisterValidator()
	s.ctx = s.ctx.WithBlockTime(time.Now())

	address := []sdk.AccAddress{
		apptesting.CreateRandomAccounts(1)[0],
		apptesting.CreateRandomAccounts(1)[0],
		apptesting.CreateRandomAccounts(1)[0],
	}

	suite.app.GrowKeeper.ChangeGrowStatus()
	suite.app.GrowKeeper.SetUSQReserveAddress(s.ctx, address[0])
	suite.app.GrowKeeper.SetGrowYieldReserveAddress(s.ctx, address[1])
	suite.app.GrowKeeper.SetGrowStakingReserveAddress(s.ctx, address[2])

	suite.OracleAggregateExchangeRateFromNet()

	suite.app.StableKeeper.AppendPair(s.ctx, s.GetNormalQStablePair(0))
	suite.app.GrowKeeper.AppendPair(s.ctx, s.GetNormalGTokenPair(0))

	err := suite.app.GrowKeeper.SetBorrowRate(s.ctx, sdk.NewInt(15), s.GetNormalGTokenPair(0).DenomID)
	suite.Require().NoError(err)
	err = suite.app.GrowKeeper.SetRealRate(s.ctx, sdk.NewInt(15), s.GetNormalGTokenPair(0).DenomID)
	suite.Require().NoError(err)

	suite.AddTestCoins(1000*1000000, s.GetNormalQStablePair(0).AmountInMetadata.Base)
	suite.MintStable(1000*1000000, s.GetNormalQStablePair(0))

	msg := types.NewMsgDeposit(
		s.Address.String(),
		sdk.NewInt(1000*1000000).String()+"uusd",
		s.GetNormalGTokenPair(0).GTokenMetadata.Base,
	)
	ctx := sdk.WrapSDKContext(suite.ctx)
	_, err = suite.app.GrowKeeper.Deposit(ctx, msg)
	suite.Require().NoError(err)

	suite.Run(fmt.Sprintf("Case---%s", "found all address"), func() {
		ctx := sdk.WrapSDKContext(suite.ctx)

		req := types.QueryYieldPercentageRequest{
			Id: s.GetNormalGTokenPair(0).DenomID,
		}

		res, err := suite.app.GrowKeeper.YieldPercentage(ctx, &req)
		suite.Require().NoError(err)
		fmt.Printf("Real Rate: %d\nBorrowRate: %d\nGrowYield: %f\nRealYield: %f\n", res.RealRate, res.BorrowRate, float64(res.GrowYield)/1000000, float64(res.RealYield)/1000000)
		fmt.Printf("Action: %s\nDiff between RealYield and GrowYield: %f\n", res.ActualAction, float64(res.Difference)/1000000)

		gTokenPair, _ := s.app.GrowKeeper.GetPairByDenomID(s.ctx, s.GetNormalGTokenPair(0).DenomID)
		s.app.GrowKeeper.SetLastTimeUpdateReserve(s.ctx, sdk.NewInt(s.ctx.BlockTime().Unix()))
		s.ctx = s.ctx.WithBlockTime(time.Unix((s.ctx.BlockTime().Unix() + 10), 0))

		realValue, blocked := s.app.GrowKeeper.CalculateAddToReserveValue(s.ctx, sdk.NewInt(res.Difference), gTokenPair)
		s.Require().Equal(blocked, false)

		fmt.Printf("Real send to/from reserve: %f\n", float64(realValue.Int64())/1000000)
	})
}

func (suite *GrowKeeperTestSuite) TestPairByDenomId() {
	suite.Setup()
	suite.Commit()
	testCases := []struct {
		name      string
		pair      types.GTokenPair
		denomId   string
		err       bool
		errString string
	}{
		{
			"ok-get position",
			s.GetNormalGTokenPair(1),
			s.GetNormalGTokenPair(1).DenomID,
			false,
			"",
		},
		{
			"false-position not found",
			s.GetNormalGTokenPair(1),
			"test",
			true,
			"not found",
		},
	}
	suite.SetupOracleKeeper("ATOM")
	suite.RegisterValidator()
	suite.app.GrowKeeper.ChangeGrowStatus()
	for _, tc := range testCases {
		suite.app.GrowKeeper.AppendPair(s.ctx, tc.pair)
		suite.Run(fmt.Sprintf("Case---%s", tc.name), func() {
			ctx := sdk.WrapSDKContext(suite.ctx)

			req := types.PairByDenomIdRequest{
				DenomId: tc.denomId,
			}

			_, err := suite.app.GrowKeeper.PairByDenomId(ctx, &req)
			if !tc.err {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err, tc.errString)
			}
		})
	}
}

func (suite *GrowKeeperTestSuite) TestAllPairs() {
	testCases := []struct {
		name   string
		pair   types.GTokenPair
		amount uint64
	}{
		{
			"ok",
			s.GetNormalGTokenPair(0),
			1,
		},
		{
			"ok",
			s.GetNormalGTokenPair(0),
			2,
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case---%s", tc.name), func() {
			suite.Setup()
			suite.Commit()
			ctx := sdk.WrapSDKContext(suite.ctx)
			for i := 0; i < int(tc.amount); i++ {
				suite.app.GrowKeeper.AppendPair(suite.ctx, suite.GetNormalGTokenPair(0))
			}

			req := types.AllPairsRequest{}

			res, err := suite.app.GrowKeeper.AllPairs(ctx, &req)
			suite.NoError(err)

			s.Require().Equal(len(res.Pairs), int(tc.amount))
		})
	}
}
