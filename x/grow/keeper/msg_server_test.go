package keeper_test

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/QuadrateOrg/core/app/apptesting"
	"github.com/QuadrateOrg/core/x/grow/types"
	stabletypes "github.com/QuadrateOrg/core/x/stable/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

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

func (suite *GrowKeeperTestSuite) TestExecuteDepositCollateral() {
	testCases := []struct {
		name            string
		qStablePair     stabletypes.Pair
		gTokenPair      types.GTokenPair
		LendAsset       types.LendAsset
		sendTokenDenom  string
		sendTokenAmount int64
	}{
		{
			"ok-deposit-collateral",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			s.GetNormalLendAsset(0),
			"uosmo",
			1000 * 1000000,
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
		suite.app.GrowKeeper.AppendLendAsset(s.ctx, tc.LendAsset)

		suite.OracleAggregateExchangeRateFromInput("0.5", tc.LendAsset.AssetMetadata.Name)
		suite.AddTestCoinsToCustomAccount(sdk.NewInt(tc.sendTokenAmount), tc.sendTokenDenom, s.Address)

		suite.Run(fmt.Sprintf("Case---%s", tc.name), func() {
			suite.AddTestCoins(tc.sendTokenAmount, tc.sendTokenDenom)
			msg := types.NewMsgDepositCollateral(
				suite.Address.String(),
				sdk.NewInt(tc.sendTokenAmount).String()+tc.sendTokenDenom,
			)
			ctx := sdk.WrapSDKContext(suite.ctx)
			res, err := suite.app.GrowKeeper.DepositCollateral(ctx, msg)
			suite.Require().NoError(err)

			position, found := s.app.GrowKeeper.GetPositionByPositionId(s.ctx, res.PositionId)
			suite.Require().Equal(found, true)
			suite.Require().Equal(position.DepositId, res.PositionId)
			suite.Require().Equal(position.Amount, sdk.NewInt(tc.sendTokenAmount).String()+tc.sendTokenDenom)
			suite.Require().Equal(position.OracleTicker, tc.LendAsset.AssetMetadata.Name)
		})
	}
}

func (suite *GrowKeeperTestSuite) TestExecuteWithdrawalCollateral() {
	testCases := []struct {
		name            string
		qStablePair     stabletypes.Pair
		gTokenPair      types.GTokenPair
		LendAsset       types.LendAsset
		sendTokenDenom  string
		sendTokenAmount int64
	}{
		{
			"ok-withdrawl-collateral",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			s.GetNormalLendAsset(0),
			"uosmo",
			1000 * 1000000,
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
		suite.app.GrowKeeper.AppendLendAsset(s.ctx, tc.LendAsset)

		suite.OracleAggregateExchangeRateFromInput("0.5", tc.LendAsset.AssetMetadata.Name)
		suite.AddTestCoinsToCustomAccount(sdk.NewInt(tc.sendTokenAmount), tc.sendTokenDenom, s.Address)

		suite.Run(fmt.Sprintf("Case---%s", tc.name), func() {
			suite.AddTestCoins(tc.sendTokenAmount, tc.sendTokenDenom)
			msg := types.NewMsgDepositCollateral(
				suite.Address.String(),
				sdk.NewInt(tc.sendTokenAmount).String()+tc.sendTokenDenom,
			)
			ctx := sdk.WrapSDKContext(suite.ctx)
			res, err := suite.app.GrowKeeper.DepositCollateral(ctx, msg)
			suite.Require().NoError(err)

			position, found := s.app.GrowKeeper.GetPositionByPositionId(s.ctx, res.PositionId)
			suite.Require().Equal(found, true)
			suite.Require().Equal(position.DepositId, res.PositionId)
			suite.Require().Equal(position.Amount, sdk.NewInt(tc.sendTokenAmount).String()+tc.sendTokenDenom)
			suite.Require().Equal(position.OracleTicker, tc.LendAsset.AssetMetadata.Name)

			oldAccountBalance := s.app.BankKeeper.GetBalance(s.ctx, s.Address, tc.sendTokenDenom)

			msg1 := types.NewMsgWithdrawalCollateral(
				suite.Address.String(),
				tc.sendTokenDenom,
			)
			ctx = sdk.WrapSDKContext(suite.ctx)
			res1, err1 := suite.app.GrowKeeper.WithdrawalCollateral(ctx, msg1)
			suite.Require().NoError(err1)
			suite.Require().Equal(res1.AmountOut, sdk.NewInt(tc.sendTokenAmount).String()+tc.sendTokenDenom)

			newAccountBalance := s.app.BankKeeper.GetBalance(s.ctx, s.Address, tc.sendTokenDenom)

			suite.Require().Equal(newAccountBalance.Amount.Sub(oldAccountBalance.Amount), sdk.NewInt(tc.sendTokenAmount))
		})
	}

}

func (suite *GrowKeeperTestSuite) TestExecuteCreateLend() {
	testCases := []struct {
		name              string
		qStablePair       stabletypes.Pair
		gTokenPair        types.GTokenPair
		LendAsset         types.LendAsset
		sendTokenDenom    string
		sendTokenAmount   int64
		expectTokenDenom  string
		expectTokenAmount int64
	}{
		{
			"ok-create-lend",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			s.GetNormalLendAsset(0),
			"uosmo",
			1000 * 1000000,
			"uusd",
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
		suite.app.GrowKeeper.AppendLendAsset(s.ctx, tc.LendAsset)

		suite.OracleAggregateExchangeRateFromInput("0.5", tc.LendAsset.AssetMetadata.Name)

		suite.AddTestCoinsToCustomAccount(sdk.NewInt(100000*1000000), tc.qStablePair.AmountOutMetadata.Base, s.app.GrowKeeper.GetGrowStakingReserveAddress(s.ctx))

		suite.Run(fmt.Sprintf("Case---%s", tc.name), func() {
			suite.AddTestCoins(tc.sendTokenAmount, tc.sendTokenDenom)
			msg := types.NewMsgDepositCollateral(
				suite.Address.String(),
				sdk.NewInt(tc.sendTokenAmount).String()+tc.sendTokenDenom,
			)
			ctx := sdk.WrapSDKContext(suite.ctx)
			res, err := suite.app.GrowKeeper.DepositCollateral(ctx, msg)
			suite.Require().NoError(err)

			oldPosition, found := s.app.GrowKeeper.GetPositionByPositionId(s.ctx, res.PositionId)
			suite.Require().Equal(found, true)
			suite.Require().Equal(oldPosition.DepositId, res.PositionId)
			suite.Require().Equal(oldPosition.Amount, sdk.NewInt(tc.sendTokenAmount).String()+tc.sendTokenDenom)
			suite.Require().Equal(oldPosition.OracleTicker, tc.LendAsset.AssetMetadata.Name)

			oldAccountBalance := s.app.BankKeeper.GetBalance(s.ctx, s.Address, tc.expectTokenDenom)

			msg2 := types.NewMsgCreateLend(
				s.Address.String(),
				tc.sendTokenDenom,
				sdk.NewInt(tc.expectTokenAmount).String(),
			)
			res1, err1 := suite.app.GrowKeeper.CreateLend(ctx, msg2)
			suite.Require().NoError(err1)
			suite.Require().Equal(res1.AmountOut, sdk.NewInt(tc.expectTokenAmount).String()+tc.expectTokenDenom)

			newAccountBalance := s.app.BankKeeper.GetBalance(s.ctx, s.Address, tc.expectTokenDenom)
			suite.Require().Equal(newAccountBalance.Amount.Sub(oldAccountBalance.Amount).Int64(), tc.expectTokenAmount)

			position, found := s.app.GrowKeeper.GetPositionByPositionId(s.ctx, res.PositionId)
			suite.Require().Equal(found, true)

			suite.Require().Equal(len(position.LoanIds), 1)
			suite.Require().Equal(position.BorrowedAmountInUSD, uint64(tc.expectTokenAmount))
			loan, found := s.app.GrowKeeper.GetLoadByLoadId(s.ctx, res1.LoanId)
			suite.Require().Equal(found, true)
			suite.Require().Equal(loan.AmountOut, res1.AmountOut)
			suite.Require().Equal(loan.Borrower, s.Address.String())
			suite.Require().Equal(loan.StartTime, uint64(s.ctx.BlockTime().Unix()))

		})
	}
}

func (suite *GrowKeeperTestSuite) TestExecuteDeleteLend() {
	testCases := []struct {
		name              string
		qStablePair       stabletypes.Pair
		gTokenPair        types.GTokenPair
		LendAsset         types.LendAsset
		sendTokenDenom    string
		sendTokenAmount   int64
		expectTokenDenom  string
		expectTokenAmount int64
	}{
		{
			"ok-delete-lend",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			s.GetNormalLendAsset(0),
			"uosmo",
			1000 * 1000000,
			"uusd",
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
		suite.app.GrowKeeper.AppendLendAsset(s.ctx, tc.LendAsset)

		suite.OracleAggregateExchangeRateFromInput("0.5", tc.LendAsset.AssetMetadata.Name)

		suite.AddTestCoinsToCustomAccount(sdk.NewInt(100000*1000000), tc.qStablePair.AmountOutMetadata.Base, s.app.GrowKeeper.GetGrowStakingReserveAddress(s.ctx))

		suite.Run(fmt.Sprintf("Case---%s", tc.name), func() {
			suite.AddTestCoins(tc.sendTokenAmount, tc.sendTokenDenom)
			msg := types.NewMsgDepositCollateral(
				suite.Address.String(),
				sdk.NewInt(tc.sendTokenAmount).String()+tc.sendTokenDenom,
			)
			ctx := sdk.WrapSDKContext(suite.ctx)
			res, err := suite.app.GrowKeeper.DepositCollateral(ctx, msg)
			suite.Require().NoError(err)

			oldPosition, found := s.app.GrowKeeper.GetPositionByPositionId(s.ctx, res.PositionId)
			suite.Require().Equal(found, true)
			suite.Require().Equal(oldPosition.DepositId, res.PositionId)
			suite.Require().Equal(oldPosition.Amount, sdk.NewInt(tc.sendTokenAmount).String()+tc.sendTokenDenom)
			suite.Require().Equal(oldPosition.OracleTicker, tc.LendAsset.AssetMetadata.Name)

			oldAccountBalance := s.app.BankKeeper.GetBalance(s.ctx, s.Address, tc.expectTokenDenom)

			msg2 := types.NewMsgCreateLend(
				s.Address.String(),
				tc.sendTokenDenom,
				sdk.NewInt(tc.expectTokenAmount).String(),
			)
			ctx = sdk.WrapSDKContext(suite.ctx)
			res1, err1 := suite.app.GrowKeeper.CreateLend(ctx, msg2)
			suite.Require().NoError(err1)
			suite.Require().Equal(res1.AmountOut, sdk.NewInt(tc.expectTokenAmount).String()+tc.expectTokenDenom)

			newAccountBalance := s.app.BankKeeper.GetBalance(s.ctx, s.Address, tc.expectTokenDenom)
			suite.Require().Equal(newAccountBalance.Amount.Sub(oldAccountBalance.Amount).Int64(), tc.expectTokenAmount)

			position, found := s.app.GrowKeeper.GetPositionByPositionId(s.ctx, res.PositionId)
			suite.Require().Equal(found, true)

			suite.Require().Equal(len(position.LoanIds), 1)
			loan, found := s.app.GrowKeeper.GetLoadByLoadId(s.ctx, res1.LoanId)
			suite.Require().Equal(found, true)
			suite.Require().Equal(loan.AmountOut, res1.AmountOut)
			suite.Require().Equal(loan.Borrower, s.Address.String())
			suite.Require().Equal(loan.StartTime, uint64(s.ctx.BlockTime().Unix()))

			s.ctx = s.ctx.WithBlockTime(time.Unix((s.ctx.BlockTime().Unix() + 31536000), 0))

			borrowTime := sdk.NewIntFromUint64(uint64(s.ctx.BlockTime().Unix()) - loan.StartTime)

			sendAmountInt := s.app.GrowKeeper.CalculateNeedAmountToGet(sdk.NewInt(tc.expectTokenAmount), borrowTime)
			sendAmount := sendAmountInt.String() + tc.expectTokenDenom

			suite.AddTestCoins(sendAmountInt.Int64()-tc.expectTokenAmount, tc.expectTokenDenom)
			msg3 := types.NewMsgDeleteLend(
				s.Address.String(),
				sendAmount,
				res1.LoanId,
				tc.sendTokenDenom,
			)
			ctx = sdk.WrapSDKContext(suite.ctx)
			_, err2 := suite.app.GrowKeeper.DeleteLend(ctx, msg3)
			suite.Require().NoError(err2)

			newPosition, found := s.app.GrowKeeper.GetPositionByPositionId(s.ctx, res.PositionId)
			suite.Require().Equal(found, true)

			suite.Require().Equal(len(newPosition.LoanIds), 0)
			suite.Require().Equal(newPosition.BorrowedAmountInUSD, position.BorrowedAmountInUSD-uint64(tc.expectTokenAmount))
		})
	}
}
