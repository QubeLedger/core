package keeper_test

import (
	"fmt"
	"time"

	"github.com/QuadrateOrg/core/app/apptesting"
	"github.com/QuadrateOrg/core/x/grow"
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
	suite.SetupOracleKeeper("uatom")
	suite.RegisterValidator()
	suite.app.GrowKeeper.SetGrowStakingReserveAddress(s.ctx, apptesting.CreateRandomAccounts(1)[0])
	suite.app.GrowKeeper.SetGrowYieldReserveAddress(s.ctx, apptesting.CreateRandomAccounts(1)[0])
	suite.app.GrowKeeper.SetUSQReserveAddress(s.ctx, apptesting.CreateRandomAccounts(1)[0])
	suite.app.GrowKeeper.SetBorrowRate(s.ctx, sdk.NewInt(15))
	suite.app.GrowKeeper.SetRealRate(s.ctx, sdk.NewInt(15))
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
			denom := tc.qStablePair.AmountOutMetadata.Base

			oldBalance := s.app.BankKeeper.GetBalance(s.ctx, s.Address, denom)
			msg := types.NewMsgWithdrawal(
				suite.Address.String(),
				sdk.NewInt(tc.sendTokenAmount).String()+tc.sendTokenDenom,
				tc.gTokenPair.GTokenMetadata.Base,
			)
			ctx := sdk.WrapSDKContext(suite.ctx)
			res, err := suite.app.GrowKeeper.Withdrawal(ctx, msg)
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

func (suite *GrowKeeperTestSuite) TestExecuteDepositCollateral() {
	testCases := []struct {
		name            string
		qStablePair     stabletypes.Pair
		gTokenPair      types.GTokenPair
		LendAsset       types.LendAsset
		sendTokenDenom  string
		sendTokenAmount int64
		err             bool
		errString       string
	}{
		{
			"ok-deposit-collateral",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			s.GetNormalLendAsset(0),
			"uosmo",
			1000 * 1000000,
			false,
			"",
		},
		{
			"false-lend asset not found",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			s.GetWrongLendAsset(0),
			"uluna",
			1000 * 1000000,
			true,
			"ErrLendAssetNotFound err",
		},
		{
			"false-oracle id not found",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			s.GetWrongLendAsset(0),
			"uosmo",
			1000 * 1000000,
			true,
			"ErrOracleAssetIdNotFound err",
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
			if !tc.err {
				suite.Require().NoError(err)

				position, found := s.app.GrowKeeper.GetPositionByPositionId(s.ctx, res.PositionId)
				suite.Require().Equal(found, true)
				suite.Require().Equal(position.DepositId, res.PositionId)
				suite.Require().Equal(position.Collateral, sdk.NewInt(tc.sendTokenAmount).String()+tc.sendTokenDenom)
				suite.Require().Equal(position.OracleTicker, tc.LendAsset.AssetMetadata.Name)
			} else {
				suite.Require().Error(err, tc.errString)
			}
		})
	}
}

func (suite *GrowKeeperTestSuite) TestExecuteWithdrawalCollateral() {
	testCases := []struct {
		name           string
		qStablePair    stabletypes.Pair
		gTokenPair     types.GTokenPair
		LendAsset      types.LendAsset
		sendTokenDenom string
		err            bool
		errString      string
	}{
		{
			"ok-withdrawl-collateral",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			s.GetNormalLendAsset(0),
			"uosmo",
			false,
			"",
		},
		{
			"false-lend asstet not found",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			s.GetNormalLendAsset(0),
			"uluna",
			true,
			"ErrLendAssetNotFound err",
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

		config := s.GetNormalConfig()

		suite.AddTestCoins(config.collateralAmount, config.collateralDenom)
		msg := types.NewMsgDepositCollateral(
			suite.Address.String(),
			sdk.NewInt(config.collateralAmount).String()+config.collateralDenom,
		)
		ctx := sdk.WrapSDKContext(suite.ctx)
		_, err := suite.app.GrowKeeper.DepositCollateral(ctx, msg)
		suite.Require().NoError(err)

		suite.Run(fmt.Sprintf("Case---%s", tc.name), func() {
			oldAccountBalance := s.app.BankKeeper.GetBalance(s.ctx, s.Address, config.collateralDenom)
			msg1 := types.NewMsgWithdrawalCollateral(
				suite.Address.String(),
				tc.sendTokenDenom,
			)
			ctx = sdk.WrapSDKContext(suite.ctx)
			res1, err1 := suite.app.GrowKeeper.WithdrawalCollateral(ctx, msg1)
			if !tc.err {
				suite.Require().NoError(err1)
				suite.Require().Equal(res1.AmountOut, sdk.NewInt(config.sendTokenAmount).String()+tc.sendTokenDenom)

				newAccountBalance := s.app.BankKeeper.GetBalance(s.ctx, s.Address, tc.sendTokenDenom)

				suite.Require().Equal(newAccountBalance.Amount.Sub(oldAccountBalance.Amount), sdk.NewInt(config.sendTokenAmount))
			} else {
				suite.Require().Error(err1, tc.errString)
			}
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
		err               bool
		errString         string
	}{
		{
			"ok-create-lend",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			s.GetNormalLendAsset(0),
			"uosmo",
			1000 * 1000000,
			"uusd",
			250 * 1000000,
			false,
			"",
		},
		{
			"false-not found position",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			s.GetNormalLendAsset(0),
			"uluna",
			1000 * 1000000,
			"uusd",
			250 * 1000000,
			true,
			"ErrPositionNotFound err",
		},
		{
			"false-risk rate",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			s.GetNormalLendAsset(0),
			"uosmo",
			1000 * 1000000,
			"uusd",
			1200 * 1000000,
			true,
			"ErrRiskRateIsGreaterThenShouldBe err",
		},
	}

	suite.Setup()
	suite.Commit()
	suite.SetupOracleKeeper("OSMO")
	suite.RegisterValidator()
	s.ctx = s.ctx.WithBlockTime(time.Now())

	suite.app.StableKeeper.AppendPair(s.ctx, s.GetNormalQStablePair(0))
	suite.app.GrowKeeper.AppendPair(s.ctx, s.GetNormalGTokenPair(0))
	suite.app.GrowKeeper.AppendLendAsset(s.ctx, s.GetNormalLendAsset(0))

	suite.OracleAggregateExchangeRateFromInput("0.5", s.GetNormalLendAsset(0).AssetMetadata.Name)

	suite.AddTestCoinsToCustomAccount(sdk.NewInt(100000*1000000), s.GetNormalQStablePair(0).AmountOutMetadata.Base, s.app.GrowKeeper.GetGrowStakingReserveAddress(s.ctx))

	config := s.GetNormalConfig()

	suite.AddTestCoins(config.collateralAmount, config.collateralDenom)
	msg := types.NewMsgDepositCollateral(
		suite.Address.String(),
		sdk.NewInt(config.collateralAmount).String()+config.collateralDenom,
	)
	ctx := sdk.WrapSDKContext(suite.ctx)
	res, err := suite.app.GrowKeeper.DepositCollateral(ctx, msg)
	suite.Require().NoError(err)

	_, found := s.app.GrowKeeper.GetPositionByPositionId(s.ctx, res.PositionId)
	suite.Require().Equal(found, true)

	oldAccountBalance := s.app.BankKeeper.GetBalance(s.ctx, s.Address, config.sendTokenDenom)

	for _, tc := range testCases {

		suite.Run(fmt.Sprintf("Case---%s", tc.name), func() {

			msg2 := types.NewMsgCreateLend(
				s.Address.String(),
				tc.sendTokenDenom,
				sdk.NewInt(tc.expectTokenAmount).String(),
			)
			res1, err1 := suite.app.GrowKeeper.CreateLend(ctx, msg2)

			if !tc.err {
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
			} else {
				suite.Require().Error(err1, tc.errString)
			}

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
		err               bool
		errString         string
	}{
		{
			"ok-delete-lend",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			s.GetNormalLendAsset(0),
			"uosmo",
			1000 * 1000000,
			"uusd",
			250 * 1000000,
			false,
			"",
		},
		{
			"false-not enough amountIn",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			s.GetNormalLendAsset(0),
			"uosmo",
			1000 * 1000000,
			"uusd",
			250,
			true,
			"ErrNotEnoughAmountIn err",
		},
		{
			"false-not found position",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			s.GetNormalLendAsset(0),
			"uluna",
			1000 * 1000000,
			"uusd",
			250,
			true,
			"ErrPositionNotFound err",
		},
		{
			"false-need send usq",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			s.GetNormalLendAsset(0),
			"uosmo",
			1000 * 1000000,
			"ueur",
			250,
			true,
			"ErrNeedSendUSQ err",
		},
	}

	suite.Setup()
	suite.Commit()
	suite.SetupOracleKeeper("OSMO")
	suite.RegisterValidator()
	s.ctx = s.ctx.WithBlockTime(time.Now())

	suite.app.StableKeeper.AppendPair(s.ctx, s.GetNormalQStablePair(0))
	suite.app.GrowKeeper.AppendPair(s.ctx, s.GetNormalGTokenPair(0))
	suite.app.GrowKeeper.AppendLendAsset(s.ctx, s.GetNormalLendAsset(0))

	suite.OracleAggregateExchangeRateFromInput("0.5", s.GetNormalLendAsset(0).AssetMetadata.Name)

	suite.AddTestCoinsToCustomAccount(sdk.NewInt(100000*1000000), s.GetNormalQStablePair(0).AmountOutMetadata.Base, s.app.GrowKeeper.GetGrowStakingReserveAddress(s.ctx))

	config := s.GetNormalConfig()

	suite.AddTestCoins(config.collateralAmount, config.collateralDenom)
	msg := types.NewMsgDepositCollateral(
		suite.Address.String(),
		sdk.NewInt(config.collateralAmount).String()+config.collateralDenom,
	)
	ctx := sdk.WrapSDKContext(suite.ctx)
	res, err := suite.app.GrowKeeper.DepositCollateral(ctx, msg)
	suite.Require().NoError(err)

	msg2 := types.NewMsgCreateLend(
		s.Address.String(),
		config.collateralDenom,
		sdk.NewInt(config.lendTokenAmount).String(),
	)
	ctx = sdk.WrapSDKContext(suite.ctx)
	res1, err1 := suite.app.GrowKeeper.CreateLend(ctx, msg2)
	suite.Require().NoError(err1)

	position, foundPos := s.app.GrowKeeper.GetPositionByPositionId(s.ctx, res.PositionId)
	loan, foundLoan := s.app.GrowKeeper.GetLoadByLoadId(s.ctx, res1.LoanId)
	suite.Require().Equal(foundPos, true)
	suite.Require().Equal(foundLoan, true)

	for _, tc := range testCases {

		suite.Run(fmt.Sprintf("Case---%s", tc.name), func() {
			s.ctx = s.ctx.WithBlockTime(time.Unix((s.ctx.BlockTime().Unix() + 31536000), 0))
			borrowTime := sdk.NewIntFromUint64(uint64(s.ctx.BlockTime().Unix()) - loan.StartTime)
			sendAmountInt := s.app.GrowKeeper.CalculateNeedAmountToGet(sdk.NewInt(tc.expectTokenAmount), borrowTime)
			sendAmount := sendAmountInt.String() + tc.expectTokenDenom

			config := s.GetNormalConfig()

			oldUsqReserveBalance := s.app.BankKeeper.GetBalance(s.ctx, s.app.GrowKeeper.GetUSQReserveAddress(s.ctx), config.sendTokenDenom)
			oldBurningFundBalance := s.app.BankKeeper.GetBalance(s.ctx, s.app.StableKeeper.GetBurningFundAddress(s.ctx), config.sendTokenDenom)

			suite.AddTestCoins(sendAmountInt.Int64()-tc.expectTokenAmount, tc.expectTokenDenom)

			msg3 := types.NewMsgDeleteLend(
				s.Address.String(),
				sendAmount,
				res1.LoanId,
				tc.sendTokenDenom,
			)
			ctx = sdk.WrapSDKContext(suite.ctx)
			_, err2 := suite.app.GrowKeeper.DeleteLend(ctx, msg3)

			if !tc.err {
				suite.Require().NoError(err2)
				newPosition, found := s.app.GrowKeeper.GetPositionByPositionId(s.ctx, res.PositionId)
				suite.Require().Equal(found, true)
				suite.Require().Equal(len(newPosition.LoanIds), 0)

				loanAmountInt, _, _ := s.app.GrowKeeper.GetAmountIntFromCoins(loan.AmountOut)
				collateralAmount, collateralDenom, _ := s.app.GrowKeeper.GetAmountIntFromCoins(position.Collateral)
				price, _ := s.app.GrowKeeper.GetPriceByDenom(s.ctx, position.OracleTicker)
				collateralReduceValue := s.app.GrowKeeper.CalculateAmountForRemoveFromCollateral(sendAmountInt.Sub(loanAmountInt), price)

				suite.Require().Equal(newPosition.BorrowedAmountInUSD, position.BorrowedAmountInUSD-uint64(tc.expectTokenAmount))
				suite.Require().Equal(newPosition.Collateral, s.app.GrowKeeper.FastCoins(collateralDenom, collateralAmount.Sub(collateralReduceValue)).String())

				newUsqReserveBalance := s.app.BankKeeper.GetBalance(s.ctx, s.app.GrowKeeper.GetUSQReserveAddress(s.ctx), config.sendTokenDenom)
				newBurningFundBalance := s.app.BankKeeper.GetBalance(s.ctx, s.app.StableKeeper.GetBurningFundAddress(s.ctx), config.sendTokenDenom)

				suite.Require().Equal(newUsqReserveBalance.Amount.Sub(oldUsqReserveBalance.Amount), sendAmountInt.Sub(loanAmountInt).QuoRaw(2))
				suite.Require().Equal(newBurningFundBalance.Amount.Sub(oldBurningFundBalance.Amount), sendAmountInt.Sub(loanAmountInt).QuoRaw(2))
			} else {
				suite.Require().Error(err2, tc.errString)
			}
		})
	}
}

func (suite *GrowKeeperTestSuite) TestExecuteCreateLiqPosition() {
	testCases := []struct {
		name             string
		qStablePair      stabletypes.Pair
		gTokenPair       types.GTokenPair
		LendAsset        types.LendAsset
		sendTokenDenom   string
		sendTokenAmount  int64
		expectTokenDenom string
		premium          string
	}{
		{
			"ok-create-liq-position",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			s.GetNormalLendAsset(0),
			"uusd",
			500 * 1000000,
			"uosmo",
			"5",
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
			msg := types.NewMsgCreateLiquidationPosition(
				suite.Address.String(),
				sdk.NewInt(tc.sendTokenAmount).String()+tc.sendTokenDenom,
				tc.LendAsset.AssetMetadata.Name,
				tc.premium,
			)
			ctx := sdk.WrapSDKContext(suite.ctx)
			res, err := suite.app.GrowKeeper.CreateLiquidationPosition(ctx, msg)
			suite.Require().NoError(err)

			position, found := s.app.GrowKeeper.GetLiquidatorPositionByLiquidatorPositionId(s.ctx, res.LiquidatorPositionId)
			suite.Require().Equal(found, true)

			suite.Require().Equal(position.Liquidator, s.Address.String())
			suite.Require().Equal(position.Amount, sdk.NewInt(tc.sendTokenAmount).String()+tc.sendTokenDenom)
			premiumInt, _ := s.app.GrowKeeper.ParseAndCheckPremium(tc.premium)
			suite.Require().Equal(position.Premium, premiumInt.Uint64())
		})
	}
}

func (suite *GrowKeeperTestSuite) TestExecuteCloseLiqPosition() {
	testCases := []struct {
		name             string
		qStablePair      stabletypes.Pair
		gTokenPair       types.GTokenPair
		LendAsset        types.LendAsset
		sendTokenDenom   string
		sendTokenAmount  int64
		expectTokenDenom string
		premium          string
	}{
		{
			"ok-close-liq-position",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			s.GetNormalLendAsset(0),
			"uusd",
			500 * 1000000,
			"uosmo",
			"5",
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
			msg := types.NewMsgCreateLiquidationPosition(
				suite.Address.String(),
				sdk.NewInt(tc.sendTokenAmount).String()+tc.sendTokenDenom,
				tc.LendAsset.AssetMetadata.Name,
				tc.premium,
			)
			ctx := sdk.WrapSDKContext(suite.ctx)
			res, err := suite.app.GrowKeeper.CreateLiquidationPosition(ctx, msg)
			suite.Require().NoError(err)

			position, found := s.app.GrowKeeper.GetLiquidatorPositionByLiquidatorPositionId(s.ctx, res.LiquidatorPositionId)
			suite.Require().Equal(found, true)

			suite.Require().Equal(position.Liquidator, s.Address.String())
			suite.Require().Equal(position.Amount, sdk.NewInt(tc.sendTokenAmount).String()+tc.sendTokenDenom)
			premiumInt, _ := s.app.GrowKeeper.ParseAndCheckPremium(tc.premium)
			suite.Require().Equal(position.Premium, premiumInt.Uint64())

			oldBalance := s.app.BankKeeper.GetBalance(s.ctx, s.Address, tc.sendTokenDenom)

			msg1 := types.NewMsgCloseLiquidationPosition(
				suite.Address.String(),
				res.LiquidatorPositionId,
			)
			ctx = sdk.WrapSDKContext(suite.ctx)
			res1, err1 := suite.app.GrowKeeper.CloseLiquidationPosition(ctx, msg1)
			suite.Require().NoError(err1)

			_, found = s.app.GrowKeeper.GetLiquidatorPositionByLiquidatorPositionId(s.ctx, res.LiquidatorPositionId)
			suite.Require().Equal(found, false)

			newBalance := s.app.BankKeeper.GetBalance(s.ctx, s.Address, tc.sendTokenDenom)

			amountOut, _ := sdk.ParseCoinsNormalized(res1.AmountOut)

			suite.Require().Equal(newBalance.Amount.Sub(oldBalance.Amount), amountOut.AmountOf(tc.sendTokenDenom))
		})
	}
}

func (suite *GrowKeeperTestSuite) TestLiquidatePositionFull() {
	testCases := []struct {
		name              string
		qStablePair       stabletypes.Pair
		gTokenPair        types.GTokenPair
		LendAsset         types.LendAsset
		sendTokenDenom    string
		sendTokenAmount   int64
		expectTokenDenom  string
		expectTokenAmount int64
		premium           string
	}{
		{
			"ok-execute-liquidation",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			s.GetNormalLendAsset(0),
			"uosmo",
			1000 * 1000000,
			"uusd",
			250 * 1000000,
			"5",
		},
	}

	suite.Setup()
	suite.Commit()
	suite.SetupOracleKeeper("OSMO")
	suite.RegisterValidator()
	s.ctx = s.ctx.WithBlockTime(time.Now())
	for _, tc := range testCases {

		suite.app.StableKeeper.AppendPair(s.ctx, tc.qStablePair)
		//suite.app.GrowKeeper.AppendPair(s.ctx, tc.gTokenPair)
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
			suite.Require().Equal(oldPosition.Collateral, sdk.NewInt(tc.sendTokenAmount).String()+tc.sendTokenDenom)
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

			suite.AddTestCoinsToCustomAccount(sdk.NewInt(tc.sendTokenAmount), tc.expectTokenDenom, suite.LiquidatorAddress)
			msg3 := types.NewMsgCreateLiquidationPosition(
				suite.LiquidatorAddress.String(),
				sdk.NewInt(tc.sendTokenAmount).String()+tc.expectTokenDenom,
				tc.LendAsset.AssetMetadata.Name,
				tc.premium,
			)
			ctx = sdk.WrapSDKContext(suite.ctx)
			res3, err3 := suite.app.GrowKeeper.CreateLiquidationPosition(ctx, msg3)
			suite.Require().NoError(err3)

			suite.OracleAggregateExchangeRateFromInput("0.4", tc.LendAsset.AssetMetadata.Name)

			position, _ = s.app.GrowKeeper.GetPositionByPositionId(s.ctx, res.PositionId)
			fmt.Printf("position Collateral msg_server: %s\n", position.Collateral)
			fmt.Printf("position BorrowedAmountInUSD msg_server: %d\n", position.BorrowedAmountInUSD)

			liqPos, _ := s.app.GrowKeeper.GetLiquidatorPositionByLiquidatorPositionId(s.ctx, res3.LiquidatorPositionId)
			fmt.Printf("liqPos Amount msg_server: %s\n", liqPos.Amount)

			liqBalance := s.app.BankKeeper.GetBalance(s.ctx, s.LiquidatorAddress, tc.sendTokenDenom)
			fmt.Printf("\nliqBalance msg_server: %s\n\n", liqBalance.String())

			err4 := grow.EndBlocker(s.ctx, s.app.GrowKeeper)
			suite.Require().NoError(err4)

			position, _ = s.app.GrowKeeper.GetPositionByPositionId(s.ctx, res.PositionId)
			fmt.Printf("position Collateral msg_server: %s\n", position.Collateral)
			fmt.Printf("position BorrowedAmountInUSD msg_server: %d\n", position.BorrowedAmountInUSD)

			liqPos, _ = s.app.GrowKeeper.GetLiquidatorPositionByLiquidatorPositionId(s.ctx, res3.LiquidatorPositionId)
			fmt.Printf("liqPos Amount msg_server: %s\n", liqPos.Amount)

			liqBalance1 := s.app.BankKeeper.GetBalance(s.ctx, s.LiquidatorAddress, tc.sendTokenDenom)
			fmt.Printf("\nliqBalance msg_server: %s\n\n", liqBalance1.String())
		})
	}
}

func (suite *GrowKeeperTestSuite) TestManyLiquidator() {
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
			"ok-execute-liquidation-many-liquidators",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			s.GetNormalLendAsset(0),
			"uosmo",
			100 * 1000000,
			"uusd",
			60 * 1000000,
		},
	}

	liquidatorCases := []struct {
		address sdk.AccAddress
		amount  int64
		denom   string
		asset   string
		premium string
	}{
		{
			apptesting.CreateRandomAccounts(1)[0],
			0.97 * 1000000,
			"uusd",
			"OSMO",
			"3",
		},
		{
			apptesting.CreateRandomAccounts(1)[0],
			4 * 1000000,
			"uusd",
			"OSMO",
			"4",
		},
		{
			apptesting.CreateRandomAccounts(1)[0],
			2 * 1000000,
			"uusd",
			"OSMO",
			"4",
		},
		{
			apptesting.CreateRandomAccounts(1)[0],
			9 * 1000000,
			"uusd",
			"OSMO",
			"9",
		},
	}

	suite.Setup()
	suite.Commit()
	suite.SetupOracleKeeper("OSMO")
	suite.RegisterValidator()
	s.ctx = s.ctx.WithBlockTime(time.Now())
	for _, tc := range testCases {

		suite.app.StableKeeper.AppendPair(s.ctx, tc.qStablePair)
		suite.app.GrowKeeper.AppendLendAsset(s.ctx, tc.LendAsset)

		suite.OracleAggregateExchangeRateFromInput("1.1", tc.LendAsset.AssetMetadata.Name)

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
			suite.Require().NotEmpty(res)

			msg2 := types.NewMsgCreateLend(
				s.Address.String(),
				tc.sendTokenDenom,
				sdk.NewInt(tc.expectTokenAmount).String(),
			)
			res1, err1 := suite.app.GrowKeeper.CreateLend(ctx, msg2)
			suite.Require().NoError(err1)
			suite.Require().NotEmpty(res1)

			for _, lc := range liquidatorCases {
				suite.AddTestCoinsToCustomAccount(sdk.NewInt(lc.amount), lc.denom, lc.address)
				msg := types.NewMsgCreateLiquidationPosition(
					lc.address.String(),
					sdk.NewInt(lc.amount).String()+lc.denom,
					lc.asset,
					lc.premium,
				)
				ctx = sdk.WrapSDKContext(suite.ctx)
				_, err := suite.app.GrowKeeper.CreateLiquidationPosition(ctx, msg)
				suite.Require().NoError(err)
			}

			suite.OracleAggregateExchangeRateFromInput("1", tc.LendAsset.AssetMetadata.Name)

			err2 := grow.EndBlocker(s.ctx, s.app.GrowKeeper)
			suite.Require().NoError(err2)

			for i, lc := range liquidatorCases {
				balance := s.app.BankKeeper.GetBalance(s.ctx, lc.address, tc.sendTokenDenom)
				fmt.Printf("Balance liquidator %d: %f OSMO\n", i, float64(balance.Amount.Int64())/1000000)
			}
		})
	}
}
