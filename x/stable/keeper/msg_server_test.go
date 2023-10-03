package keeper_test

import (
	"fmt"
	"math/rand"
	"time"

	apptesting "github.com/QuadrateOrg/core/app/apptesting"
	"github.com/QuadrateOrg/core/x/stable/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *StableKeeperTestSuite) TestMint() {
	testCases := []struct {
		name            string
		pair            types.Pair
		sendTokenDenom  string
		sendTokenAmount int64
		getTokenAmount  int64
		price           int64
		err             bool
		errString       string
	}{
		{
			"ok-mint",
			s.GetNormalPair(0),
			"uatom",
			1000,
			9471,
			int64(95000),
			false,
			"",
		},
		{
			"fail-pair not found",
			s.GetNormalPair(0),
			"uqube",
			1000,
			9272,
			int64(93000),
			true,
			"ErrPairNotFound err",
		},
		{
			"fail-mint blocked",
			s.GetNormalPair(0),
			"uatom",
			1000,
			9471,
			int64(150000000),
			true,
			"Backing Ration >= 120%",
		},
		{
			"fail-amountIn less minAmountIn",
			s.GetNormalPair(0),
			"uatom",
			19,
			0,
			int64(93000),
			true,
			"ErrAmountInGTEminAmountIn err",
		},
	}

	suite.Setup()
	suite.Commit()

	for _, tc := range testCases {
		suite.app.StableKeeper.SetTestingMode(true)
		suite.Run(fmt.Sprintf("Case---%s", tc.name), func() {
			suite.app.StableKeeper.AppendPair(suite.ctx, tc.pair)
			suite.app.StableKeeper.UpdateAtomPriceTesting(suite.ctx, sdk.NewInt(tc.price))
			suite.AddTestCoins(tc.sendTokenAmount, tc.sendTokenDenom)

			msg := types.NewMsgMint(
				suite.Address.String(),
				sdk.NewInt(tc.sendTokenAmount).String()+tc.sendTokenDenom,
				tc.pair.AmountOutMetadata.Base,
			)
			ctx := sdk.WrapSDKContext(suite.ctx)
			_, err := suite.app.StableKeeper.Mint(ctx, msg)
			if !tc.err {
				suite.Require().NoError(err, tc.name)
				getTokenAmountFromBank := suite.app.BankKeeper.GetBalance(suite.ctx, suite.Address, tc.pair.AmountOutMetadata.Base)
				suite.Require().Equal(getTokenAmountFromBank.Amount, sdk.NewInt(int64(tc.getTokenAmount)))

				burningFundBalance := suite.app.BankKeeper.GetBalance(suite.ctx, suite.app.StableKeeper.GetBurningFundAddress(suite.ctx), tc.pair.AmountInMetadata.Base)
				feeForBurningFund := suite.app.StableKeeper.CalculateMintingFeeForBurningFund(sdk.NewInt(tc.sendTokenAmount), sdk.NewInt(tc.price), sdk.NewInt(3))
				suite.Require().Equal(burningFundBalance.Amount, feeForBurningFund)

				reserveFundBalance := suite.app.BankKeeper.GetBalance(suite.ctx, suite.app.StableKeeper.GetReserveFundAddress(suite.ctx), tc.pair.AmountInMetadata.Base)
				suite.Require().Equal(reserveFundBalance.Amount, (sdk.NewInt(tc.sendTokenAmount).Sub(feeForBurningFund)))
			} else {
				suite.Require().Error(err, tc.errString)
			}
		})
	}
}

func (suite *StableKeeperTestSuite) TestBurn() {
	testCases := []struct {
		name            string
		pair            types.Pair
		sendTokenDenom  string
		sendTokenAmount int64
		getTokenAmount  int64
		price           int64
		err             bool
		errString       string
	}{
		{
			"ok - burn",
			s.GetNormalPair(0),
			"uusd",
			1000000,
			105052,
			int64(95000),
			false,
			"",
		},
		{
			"fail - wrong denom",
			s.GetNormalPair(0),
			"uqube",
			1000,
			104,
			int64(95000),
			true,
			"ErrSendBaseTokenDenom err",
		},
		{
			"fail - burn blocked",
			s.GetNormalPair(0),
			"uusd",
			1000,
			104,
			int64(3300),
			true,
			"Backing Ration < 85%",
		},
		{
			"fail-amountOut less minAnountOut",
			s.GetNormalPair(0),
			"uusd",
			19,
			0,
			int64(93000),
			true,
			"ErrAmountOutGTEminAmountOut err",
		},
	}
	suite.Setup()
	suite.Commit()

	for _, tc := range testCases {
		suite.app.StableKeeper.SetTestingMode(true)
		suite.app.StableKeeper.AppendPair(suite.ctx, tc.pair)

		suite.AddTestCoins(100000000, tc.pair.AmountInMetadata.Base)
		suite.MintStable(100000000, tc.pair)
		BurningFundBalanceBeforeBurn := suite.app.BankKeeper.GetBalance(suite.ctx, suite.app.StableKeeper.GetBurningFundAddress(suite.ctx), tc.pair.AmountInMetadata.Base)
		ReserveFundBalanceBeforeBurn := suite.app.BankKeeper.GetBalance(suite.ctx, suite.app.StableKeeper.GetReserveFundAddress(suite.ctx), tc.pair.AmountInMetadata.Base)
		suite.Run(fmt.Sprintf("Case---%s", tc.name), func() {
			suite.app.StableKeeper.UpdateAtomPriceTesting(suite.ctx, sdk.NewInt(tc.price))

			msg := types.NewMsgBurn(
				suite.Address.String(),
				sdk.NewInt(tc.sendTokenAmount).String()+tc.sendTokenDenom,
				tc.pair.AmountInMetadata.Base,
			)
			ctx := sdk.WrapSDKContext(suite.ctx)
			msgBurnResponse, err := suite.app.StableKeeper.Burn(ctx, msg)

			if !tc.err {
				suite.Require().NoError(err, tc.name)
				getTokenAmountFromBank := suite.app.BankKeeper.GetBalance(suite.ctx, suite.Address, tc.pair.AmountInMetadata.Base)
				suite.Require().Equal(getTokenAmountFromBank.Amount, sdk.NewInt(int64(tc.getTokenAmount)))

				burningFundBalance := suite.app.BankKeeper.GetBalance(suite.ctx, suite.app.StableKeeper.GetBurningFundAddress(suite.ctx), tc.pair.AmountInMetadata.Base)
				feeForBurningFund := suite.app.StableKeeper.CalculateBurningFeeForBurningFund(sdk.NewInt(tc.sendTokenAmount), sdk.NewInt(tc.price), sdk.NewInt(2))
				suite.Require().Equal(burningFundBalance.Amount.Sub(BurningFundBalanceBeforeBurn.Amount), feeForBurningFund)

				reserveFundBalance := suite.app.BankKeeper.GetBalance(suite.ctx, suite.app.StableKeeper.GetReserveFundAddress(suite.ctx), tc.pair.AmountInMetadata.Base)
				msgBurnResponseAmountOutCoins, _ := sdk.ParseCoinsNormalized(msgBurnResponse.AmountOut)
				suite.Require().Equal(ReserveFundBalanceBeforeBurn.Amount.Sub(reserveFundBalance.Amount), (msgBurnResponseAmountOutCoins.AmountOf(tc.pair.AmountInMetadata.Base)).Add(feeForBurningFund))
			} else {
				suite.Require().Error(err, tc.errString)
			}
		})
	}
}

func (suite *StableKeeperTestSuite) TestMintGetPriceFromOracle() {
	testCases := []struct {
		name            string
		pair            types.Pair
		sendTokenDenom  string
		sendTokenAmount int64
		err             bool
		errString       string
	}{
		{
			"ok-mint№1",
			s.GetNormalPair(0),
			"uatom",
			1000,
			false,
			"",
		},
		{
			"ok-mint№2",
			s.GetNormalPair(0),
			"uatom",
			300,
			false,
			"",
		},
		{
			"ok-mint№3",
			s.GetNormalPair(0),
			"uatom",
			730,
			false,
			"",
		},
	}
	suite.Setup()
	suite.Commit()
	suite.app.StableKeeper.SetTestingMode(false)
	suite.SetupOracleKeeper()
	suite.RegisterValidator()
	for _, tc := range testCases {
		suite.app.StableKeeper.AppendPair(s.ctx, tc.pair)
		suite.OracleAggregateExchangeRateFromNet()
		price, _ := suite.app.StableKeeper.GetAtomPrice(suite.ctx, tc.pair)
		suite.Run(fmt.Sprintf("Case---%s---price---%f", tc.name, float64(float64(price.Int64())/10000)), func() {
			suite.AddTestCoins(tc.sendTokenAmount, tc.sendTokenDenom)
			msg := types.NewMsgMint(
				suite.Address.String(),
				sdk.NewInt(tc.sendTokenAmount).String()+tc.sendTokenDenom,
				tc.pair.AmountOutMetadata.Base,
			)
			ctx := sdk.WrapSDKContext(suite.ctx)
			_, err := suite.app.StableKeeper.Mint(ctx, msg)
			suite.Require().NoError(err)

			getTokenAmountFromBank := suite.app.BankKeeper.GetBalance(suite.ctx, suite.Address, tc.pair.AmountOutMetadata.Base)
			suite.Require().Greater(getTokenAmountFromBank.Amount.Int64(), int64(0))
		})
	}
}

func (suite *StableKeeperTestSuite) TestBurnGetPriceFromOracle() {
	testCases := []struct {
		name            string
		pair            types.Pair
		sendTokenDenom  string
		sendTokenAmount int64
		err             bool
		errString       string
	}{
		{
			"ok-burn№1",
			s.GetNormalPair(0),
			"uusd",
			1000,
			false,
			"",
		},
		{
			"ok-burn№2",
			s.GetNormalPair(0),
			"uusd",
			300,
			false,
			"",
		},
		{
			"ok-burn№3",
			s.GetNormalPair(0),
			"uusd",
			730,
			false,
			"",
		},
	}
	suite.Setup()
	suite.Commit()
	suite.app.StableKeeper.SetTestingMode(false)
	suite.SetupOracleKeeper()
	suite.RegisterValidator()
	for _, tc := range testCases {
		suite.OracleAggregateExchangeRateFromNet()
		suite.app.StableKeeper.AppendPair(suite.ctx, tc.pair)
		suite.AddTestCoins(10000, tc.pair.AmountInMetadata.Base)
		err := suite.MintStable(10000, s.GetNormalPair(0))
		suite.Require().NoError(err)
		price, _ := suite.app.StableKeeper.GetAtomPrice(suite.ctx, tc.pair)
		suite.Run(fmt.Sprintf("Case---%s---price---%f", tc.name, float64(float64(price.Int64())/10000)), func() {
			msg := types.NewMsgBurn(
				suite.Address.String(),
				sdk.NewInt(tc.sendTokenAmount).String()+tc.sendTokenDenom,
				tc.pair.AmountInMetadata.Base,
			)
			ctx := sdk.WrapSDKContext(suite.ctx)
			_, err := suite.app.StableKeeper.Burn(ctx, msg)
			suite.Require().NoError(err)
			uatomSuply := suite.app.BankKeeper.GetBalance(suite.ctx, suite.Address, tc.pair.AmountInMetadata.Base)
			suite.Require().Greater(uatomSuply.Amount.Int64(), int64(0))
		})
	}
}

func (suite *StableKeeperTestSuite) TestExtremeMarketSituations() {
	user1 := apptesting.CreateRandomAccounts(1)[0]
	user2 := apptesting.CreateRandomAccounts(1)[0]

	testCases := []struct {
		name                string
		pair                types.Pair
		sendTokenDenom      string
		sendTokenAmount     int64
		expectedTokenAmount int64
		price               int64
		err                 bool
		errString           string
		action              string
		address             sdk.AccAddress
	}{
		{
			"mint-user1",
			s.GetNormalPair(0),
			"uatom",
			1000,
			100,
			95000, // 9.5 * 10000
			false,
			"",
			"mint",
			user1,
		},
		{
			"mint-user2",
			s.GetNormalPair(0),
			"uatom",
			1000,
			100,
			95670,
			false,
			"",
			"mint",
			user2,
		},
		{
			"mint№2-user1",
			s.GetNormalPair(0),
			"uatom",
			1000,
			50,
			98530,
			false,
			"",
			"mint",
			user1,
		},
		{
			"mint№2-user2",
			s.GetNormalPair(0),
			"uatom",
			700,
			50,
			99410,
			false,
			"",
			"mint",
			user2,
		},
		{
			"burn-user1",
			s.GetNormalPair(0),
			"uusd",
			9000,
			50,
			92133,
			true,
			"",
			"burn",
			user1,
		},
		{
			"burn-user1",
			s.GetNormalPair(0),
			"uusd",
			9000,
			50,
			90312,
			true,
			"",
			"burn",
			user1,
		},
	}

	suite.Setup()
	suite.Commit()
	suite.app.StableKeeper.SetTestingMode(true)
	for _, tc := range testCases {
		if tc.action == "mint" {
			suite.AddTestCoinsToCustomAccount(sdk.NewInt(tc.sendTokenAmount), tc.sendTokenDenom, tc.address)
		}
	}
	for _, tc := range testCases {
		suite.app.StableKeeper.AppendPair(suite.ctx, tc.pair)
		suite.Run(fmt.Sprintf("Case---%s", tc.name), func() {
			suite.app.StableKeeper.UpdateAtomPriceTesting(suite.ctx, sdk.NewInt(tc.price))
			switch tc.action {
			case "mint":
				msg := types.NewMsgMint(
					tc.address.String(),
					sdk.NewInt(tc.sendTokenAmount).String()+tc.sendTokenDenom,
					tc.pair.AmountOutMetadata.Base,
				)
				ctx := sdk.WrapSDKContext(suite.ctx)
				_, err := suite.app.StableKeeper.Mint(ctx, msg)
				suite.Require().NoError(err)
				balance := suite.app.BankKeeper.GetBalance(suite.ctx, tc.address, tc.pair.AmountOutMetadata.Base)
				suite.Require().Greater(balance.Amount.Int64(), int64(0))
			case "burn":
				msg := types.NewMsgBurn(
					tc.address.String(),
					sdk.NewInt(tc.sendTokenAmount).String()+tc.sendTokenDenom,
					tc.pair.AmountInMetadata.Base,
				)
				ctx := sdk.WrapSDKContext(suite.ctx)
				_, err := suite.app.StableKeeper.Burn(ctx, msg)
				suite.Require().NoError(err)
				balance := suite.app.BankKeeper.GetBalance(suite.ctx, tc.address, tc.pair.AmountInMetadata.Base)
				suite.Require().Greater(balance.Amount.Int64(), int64(0))
			default:
				suite.Error(nil)
			}
		})
	}
}

func (suite *StableKeeperTestSuite) TestMarketDropOf40() {
	rand.Seed(time.Now().UnixNano())
	actionUser := apptesting.CreateRandomAccounts(1)[0]
	type testCase struct {
		name                string
		pair                types.Pair
		sendTokenDenom      string
		sendTokenAmount     int64
		expectedTokenAmount int64
		price               int64
		err                 bool
		errString           string
		action              string
		address             sdk.AccAddress
	}
	testCases := []testCase{}

	for i := 0; i < 10; i++ {
		newTestCase := testCase{
			name:                fmt.Sprintf("mint-user%d", i),
			pair:                suite.GetNormalPair(0),
			sendTokenDenom:      "uatom",
			sendTokenAmount:     1000,
			expectedTokenAmount: 100,
			price:               int64(rand.Intn(96000-90000) + 90000),
			err:                 false,
			errString:           "",
			action:              "mint",
			address:             apptesting.CreateRandomAccounts(1)[0],
		}
		testCases = append(testCases, newTestCase)
	}

	burnTestCaseFail := testCase{
		name:                fmt.Sprintf("burn"),
		pair:                suite.GetNormalPair(0),
		sendTokenDenom:      "uusd",
		sendTokenAmount:     8000,
		expectedTokenAmount: 50,
		price:               int64(57000),
		err:                 true,
		errString:           "Backing Ration < 85%",
		action:              "burn",
		address:             actionUser,
	}
	mintForTest := testCase{
		name:                fmt.Sprintf("mint"),
		pair:                suite.GetNormalPair(0),
		sendTokenDenom:      "uatom",
		sendTokenAmount:     1000,
		expectedTokenAmount: 100,
		price:               int64(rand.Intn(96000-90000) + 90000),
		err:                 false,
		errString:           "",
		action:              "mint",
		address:             actionUser,
	}
	burnTestCase1Succses := testCase{
		name:                fmt.Sprintf("burn"),
		pair:                suite.GetNormalPair(0),
		sendTokenDenom:      "uusd",
		sendTokenAmount:     8000,
		expectedTokenAmount: 50,
		price:               int64(84000),
		err:                 false,
		errString:           "Backing Ration < 85%",
		action:              "burn",
		address:             actionUser,
	}
	testCases = append(testCases, burnTestCaseFail)
	testCases = append(testCases, mintForTest)
	testCases = append(testCases, burnTestCase1Succses)

	suite.Setup()
	suite.Commit()
	suite.app.StableKeeper.SetTestingMode(true)
	for _, tc := range testCases {
		if tc.action == "mint" {
			suite.AddTestCoinsToCustomAccount(sdk.NewInt(tc.sendTokenAmount), tc.sendTokenDenom, tc.address)
		}
	}

	for _, tc := range testCases {
		suite.app.StableKeeper.AppendPair(suite.ctx, tc.pair)
		suite.Run(fmt.Sprintf("Case---%s--Price---%f", tc.name, float64(float64(tc.price)/10000)), func() {
			suite.app.StableKeeper.UpdateAtomPriceTesting(suite.ctx, sdk.NewInt(tc.price))
			switch tc.action {
			case "mint":
				msg := types.NewMsgMint(
					tc.address.String(),
					sdk.NewInt(tc.sendTokenAmount).String()+tc.sendTokenDenom,
					tc.pair.AmountOutMetadata.Base,
				)
				ctx := sdk.WrapSDKContext(suite.ctx)
				_, err := suite.app.StableKeeper.Mint(ctx, msg)
				if tc.err {
					suite.Require().Error(err, tc.errString)
				} else {
					suite.Require().NoError(err)
					balance := suite.app.BankKeeper.GetBalance(suite.ctx, tc.address, tc.pair.AmountOutMetadata.Base)
					suite.Require().Greater(balance.Amount.Int64(), int64(0))
				}
			case "burn":
				msg := types.NewMsgBurn(
					tc.address.String(),
					sdk.NewInt(tc.sendTokenAmount).String()+tc.sendTokenDenom,
					tc.pair.AmountInMetadata.Base,
				)
				ctx := sdk.WrapSDKContext(suite.ctx)
				_, err := suite.app.StableKeeper.Burn(ctx, msg)
				if tc.err {
					suite.Require().Error(err, tc.errString)
				} else {
					suite.Require().NoError(err)
					balance := suite.app.BankKeeper.GetBalance(suite.ctx, tc.address, tc.pair.AmountInMetadata.Base)
					suite.Require().Greater(balance.Amount.Int64(), int64(0))
				}
			default:
				suite.Error(nil)
			}
		})
	}
}
