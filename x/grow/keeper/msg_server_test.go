package keeper_test

import (
	"fmt"
	"time"

	"github.com/QuadrateOrg/core/app/apptesting"
	"github.com/QuadrateOrg/core/x/grow/types"
	stabletypes "github.com/QuadrateOrg/core/x/stable/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *GrowKeeperTestSuite) TestExecuteDeposit() {
	testCases := []struct {
		name            string
		qStablePair     stabletypes.Pair
		gTokenPair      types.GTokenPair
		sendTokenDenom  string
		sendTokenAmount int64
	}{
		{
			"ok-deposit",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			"uusd",
			1000 * 1000000,
		},
	}

	suite.Setup()
	suite.Commit()
	suite.SetupOracleKeeper("uatom")
	suite.RegisterValidator()
	for _, tc := range testCases {

		suite.app.StableKeeper.AppendPair(s.ctx, tc.qStablePair)
		suite.app.GrowKeeper.AppendPair(s.ctx, tc.gTokenPair)

		suite.OracleAggregateExchangeRateFromNet()

		suite.AddTestCoins(10000, tc.qStablePair.AmountInMetadata.Base)

		err := suite.MintStable(10000, s.GetNormalQStablePair(0))
		suite.Require().NoError(err)

		suite.Run(fmt.Sprintf("Case---%s", tc.name), func() {
			suite.AddTestCoins(tc.sendTokenAmount, tc.sendTokenDenom)
			msg := types.NewMsgDeposit(
				suite.Address.String(),
				sdk.NewInt(tc.sendTokenAmount).String()+tc.sendTokenDenom,
				tc.gTokenPair.GTokenMetadata.Base,
			)
			ctx := sdk.WrapSDKContext(suite.ctx)
			_, err := suite.app.GrowKeeper.Deposit(ctx, msg)
			suite.Require().NoError(err)
		})
	}
}

func (suite *GrowKeeperTestSuite) TestExecuteWithdrawal() {
	testCases := []struct {
		name            string
		qStablePair     stabletypes.Pair
		gTokenPair      types.GTokenPair
		sendTokenDenom  string
		sendTokenAmount int64
	}{
		{
			"ok-withdrawal",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			"ugusd",
			1000000,
		},
	}

	suite.Setup()
	suite.Commit()
	suite.SetupOracleKeeper("uatom")
	suite.RegisterValidator()
	suite.app.GrowKeeper.SetGrowStakingReserveAddress(s.ctx, apptesting.CreateRandomAccounts(1)[0])
	for _, tc := range testCases {

		suite.app.StableKeeper.AppendPair(s.ctx, tc.qStablePair)
		suite.app.GrowKeeper.AppendPair(s.ctx, tc.gTokenPair)

		suite.OracleAggregateExchangeRateFromNet()

		suite.AddTestCoins(10000, tc.qStablePair.AmountInMetadata.Base)
		suite.AddTestCoinsToCustomAccount(sdk.NewInt(1000*1000000), tc.qStablePair.AmountOutMetadata.Base, s.app.GrowKeeper.GetGrowStakingReserveAddress(s.ctx))

		err := suite.MintStable(10000, s.GetNormalQStablePair(0))
		suite.Require().NoError(err)

		suite.Run(fmt.Sprintf("Case---%s", tc.name), func() {
			suite.AddTestCoins(tc.sendTokenAmount, tc.sendTokenDenom)
			msg := types.NewMsgWithdrawal(
				suite.Address.String(),
				sdk.NewInt(tc.sendTokenAmount).String()+tc.sendTokenDenom,
				tc.gTokenPair.GTokenMetadata.Base,
			)
			ctx := sdk.WrapSDKContext(suite.ctx)
			res, err := suite.app.GrowKeeper.Withdrawal(ctx, msg)
			suite.Require().NoError(err)
			fmt.Printf("%s", res.AmountOut)
		})
	}
}

func (suite *GrowKeeperTestSuite) TestExecuteCreateLend() {
	testCases := []struct {
		name              string
		qStablePair       stabletypes.Pair
		gTokenPair        types.GTokenPair
		borrowAsset       types.BorrowAsset
		sendTokenDenom    string
		sendTokenAmount   int64
		expectTokenAmount int64
	}{
		{
			"ok-create-lend",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			s.GetNormalBorrowAsset(0),
			"uosmo",
			1000 * 1000000,
			500 * 1000000,
		},
	}

	suite.Setup()
	suite.Commit()
	suite.SetupOracleKeeper("OSMO")
	suite.RegisterValidator()
	s.ctx = s.ctx.WithBlockTime(time.Now())
	for _, tc := range testCases {

		suite.app.StableKeeper.AppendPair(s.ctx, tc.qStablePair)
		suite.app.GrowKeeper.AppendPair(s.ctx, tc.gTokenPair)
		suite.app.GrowKeeper.AppendBorrowAsset(s.ctx, tc.borrowAsset)

		suite.OracleAggregateExchangeRateFromInput("0.5", tc.borrowAsset.AmountInAssetMetadata.Name)

		suite.AddTestCoinsToCustomAccount(sdk.NewInt(1000*1000000), tc.qStablePair.AmountOutMetadata.Base, s.app.GrowKeeper.GetGrowStakingReserveAddress(s.ctx))

		suite.Run(fmt.Sprintf("Case---%s", tc.name), func() {
			suite.AddTestCoins(tc.sendTokenAmount, tc.sendTokenDenom)
			msg := types.NewMsgCreateLend(
				suite.Address.String(),
				sdk.NewInt(tc.sendTokenAmount).String()+tc.sendTokenDenom,
				tc.borrowAsset.AmountOutAssetMetadata.Base,
			)
			ctx := sdk.WrapSDKContext(suite.ctx)
			res, err := suite.app.GrowKeeper.CreateLend(ctx, msg)
			suite.Require().NoError(err)

			price, err := s.app.GrowKeeper.GetPriceByDenom(s.ctx, tc.borrowAsset.AmountInAssetMetadata.Name)
			suite.Require().NoError(err)

			expectAmountOut := s.app.GrowKeeper.CalculateCreateLendAmountOut(sdk.NewInt(tc.sendTokenAmount), price)
			balanceUser := s.app.BankKeeper.GetBalance(s.ctx, s.Address, tc.qStablePair.AmountOutMetadata.Base)

			s.Require().Equal(balanceUser.Amount, expectAmountOut)
			s.Require().Equal(balanceUser.Amount, sdk.NewInt(tc.expectTokenAmount))

			loan, found := s.app.GrowKeeper.GetLoadByLoadId(s.ctx, res.LoanId)
			s.Require().Equal(found, true)

			s.Require().Equal(loan.StartTime, uint64(s.ctx.BlockTime().Unix()))
		})
	}
}
