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
		err             bool
		errString       string
	}{
		{
			"ok-deposit",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			"uusd",
			1000 * 1000000,
			false,
			"",
		},
		{
			"false-pair not found",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			"ueur",
			1000 * 1000000,
			true,
			"ErrPairNotFound err",
		},
		{
			"fail-amountIn less minAmountIn",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			"uusd",
			10,
			true,
			"ErrAmountInGTEminAmountIn err",
		},
	}

	suite.Setup()
	suite.Commit()
	suite.SetupOracleKeeper("ATOM")
	suite.RegisterValidator()
	suite.app.GrowKeeper.ChangeDepositMethodStatus(s.ctx)
	suite.app.GrowKeeper.ChangeCollateralMethodStatus(s.ctx)
	suite.app.GrowKeeper.ChangeBorrowMethodStatus(s.ctx)
	for _, tc := range testCases {

		suite.app.StableKeeper.AppendPair(s.ctx, tc.qStablePair)
		suite.app.GrowKeeper.AppendPair(s.ctx, tc.gTokenPair)

		suite.OracleAggregateExchangeRateFromNet()

		suite.AddTestCoins(10000, tc.qStablePair.AmountInMetadata.Base)

		err := suite.MintStable(10000, s.GetNormalQStablePair(0))
		suite.Require().NoError(err)

		suite.Run(fmt.Sprintf("Case---%s", tc.name), func() {
			suite.AddTestCoins(tc.sendTokenAmount, tc.sendTokenDenom)
			msg := types.NewMsgGrowDeposit(
				suite.Address.String(),
				sdk.NewInt(tc.sendTokenAmount).String()+tc.sendTokenDenom,
				tc.gTokenPair.GTokenMetadata.Base,
			)
			ctx := sdk.WrapSDKContext(suite.ctx)
			_, err := suite.app.GrowKeeper.GrowDeposit(ctx, msg)
			if !tc.err {
				suite.Require().NoError(err)
				getTokenAmountFromBank := suite.app.BankKeeper.GetBalance(suite.ctx, suite.Address, tc.gTokenPair.GTokenMetadata.Base)
				suite.Require().Equal(getTokenAmountFromBank.Amount, sdk.NewInt(tc.sendTokenAmount))
			} else {
				suite.Require().Error(err, tc.errString)
			}
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
		err             bool
		errString       string
	}{
		{
			"ok-withdrawal",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			"ugusd",
			1000000,
			false,
			"",
		},
		{
			"false-pair not found",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			"ugeur",
			1000000,
			true,
			"ErrPairNotFound err",
		},
		{
			"fail-amountIn less minAmountOut",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			"ugusd",
			10,
			true,
			"ErrAmountOutGTEminAmountOut err",
		},
	}

	suite.Setup()
	suite.Commit()
	suite.SetupOracleKeeper("ATOM")
	suite.RegisterValidator()
	suite.app.GrowKeeper.ChangeDepositMethodStatus(s.ctx)
	suite.app.GrowKeeper.ChangeCollateralMethodStatus(s.ctx)
	suite.app.GrowKeeper.ChangeBorrowMethodStatus(s.ctx)
	suite.app.GrowKeeper.SetGrowStakingReserveAddress(s.ctx, apptesting.CreateRandomAccounts(1)[0])
	suite.app.GrowKeeper.SetGrowYieldReserveAddress(s.ctx, apptesting.CreateRandomAccounts(1)[0])
	suite.app.GrowKeeper.SetUSQReserveAddress(s.ctx, apptesting.CreateRandomAccounts(1)[0])
	for _, tc := range testCases {

		suite.app.StableKeeper.AppendPair(s.ctx, tc.qStablePair)
		suite.app.GrowKeeper.AppendPair(s.ctx, tc.gTokenPair)

		suite.app.GrowKeeper.SetBorrowRate(s.ctx, sdk.NewInt(15), tc.gTokenPair.DenomID)
		suite.app.GrowKeeper.SetRealRate(s.ctx, sdk.NewInt(15), tc.gTokenPair.DenomID)

		suite.OracleAggregateExchangeRateFromNet()

		suite.AddTestCoins(10000, tc.qStablePair.AmountInMetadata.Base)
		suite.AddTestCoinsToCustomAccount(sdk.NewInt(1000*1000000), tc.qStablePair.AmountOutMetadata.Base, s.app.GrowKeeper.GetGrowStakingReserveAddress(s.ctx))

		err := suite.MintStable(10000, s.GetNormalQStablePair(0))
		suite.Require().NoError(err)

		suite.Run(fmt.Sprintf("Case---%s", tc.name), func() {
			suite.AddTestCoins(tc.sendTokenAmount, tc.sendTokenDenom)
			denom := tc.qStablePair.AmountOutMetadata.Base

			oldBalance := s.app.BankKeeper.GetBalance(s.ctx, s.Address, denom)
			msg := types.NewMsgGrowWithdrawal(
				suite.Address.String(),
				sdk.NewInt(tc.sendTokenAmount).String()+tc.sendTokenDenom,
			)
			ctx := sdk.WrapSDKContext(suite.ctx)
			res, err := suite.app.GrowKeeper.GrowWithdrawal(ctx, msg)
			if !tc.err {
				suite.Require().NoError(err)
				newBalance := s.app.BankKeeper.GetBalance(s.ctx, s.Address, denom)
				amountOut, _ := sdk.ParseCoinsNormalized(res.AmountOut)
				suite.Require().Equal(newBalance.Amount.Sub(oldBalance.Amount), amountOut.AmountOf(denom))
			} else {
				suite.Require().Error(err, tc.errString)
			}
		})
	}
}
