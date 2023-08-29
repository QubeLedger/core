package keeper_test

import (
	"fmt"

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
			"ok-mint",
			s.GetNormalQStablePair(0),
			s.GetNormalGSTokenPair(0),
			"uusd",
			1000 * 1000000,
		},
	}

	suite.Setup()
	suite.Commit()
	suite.SetupOracleKeeper()
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
			"ok-mint",
			s.GetNormalQStablePair(0),
			s.GetNormalGSTokenPair(0),
			"ugusd",
			1000000,
		},
	}

	suite.Setup()
	suite.Commit()
	suite.SetupOracleKeeper()
	suite.RegisterValidator()
	suite.app.GrowKeeper.SetUSQStakingReserveAddress(s.ctx, apptesting.CreateRandomAccounts(1)[0])
	for _, tc := range testCases {

		suite.app.StableKeeper.AppendPair(s.ctx, tc.qStablePair)
		suite.app.GrowKeeper.AppendPair(s.ctx, tc.gTokenPair)

		suite.OracleAggregateExchangeRateFromNet()

		suite.AddTestCoins(10000, tc.qStablePair.AmountInMetadata.Base)
		suite.AddTestCoinsToCustomAccount(sdk.NewInt(1000*1000000), tc.qStablePair.AmountOutMetadata.Base, s.app.GrowKeeper.GetUSQStakingReserveAddress(s.ctx))

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
