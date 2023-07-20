package keeper_test

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/QuadrateOrg/core/x/stable/types"

	apptesting "github.com/QuadrateOrg/core/app/apptesting"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *StableKeeperTestSuite) TestMintUsq() {
	testCases := []struct {
		name           string
		baseTokenDenom string
		sendTokenDenom string
		atomAmount     int64
		uusdAmount     int64
		atomPrice      int64
		err            bool
		errString      string
	}{
		{
			"ok-mint",
			"uatom",
			"uatom",
			1000,
			9471,
			int64(95000),
			false,
			"",
		},
		{
			"fail-wrong denom",
			"uatom",
			"ukuji",
			1000,
			9272,
			int64(93000),
			true,
			"ErrSendBaseTokenDenom err",
		},
		{
			"fail-mint blocked",
			"uatom",
			"uatom",
			1000,
			9471,
			int64(1500000),
			true,
			"Backing Ration >= 120%",
		},
	}

	for _, tc := range testCases {
		suite.Setup()
		suite.Commit()
		suite.app.StableKeeper.SetTestingMode(true)
		suite.Run(fmt.Sprintf("Case---%s", tc.name), func() {
			suite.app.StableKeeper.SetBaseTokenDenom(suite.ctx, tc.baseTokenDenom)
			suite.app.StableKeeper.UpdateAtomPriceTesting(suite.ctx, sdk.NewInt(tc.atomPrice))

			suite.app.BankKeeper.MintCoins(suite.ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin("uatom", sdk.NewInt(tc.atomAmount))))
			suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, suite.Address, sdk.NewCoins(sdk.NewCoin("uatom", sdk.NewInt(tc.atomAmount))))

			msg := types.NewMsgMintUsq(
				suite.Address.String(),
				sdk.NewInt(tc.atomAmount).String()+tc.sendTokenDenom,
			)
			ctx := sdk.WrapSDKContext(suite.ctx)
			_, err := suite.app.StableKeeper.MintUsq(ctx, msg)

			if !tc.err {
				suite.Require().NoError(err, tc.name)
				uusdAmount := suite.app.BankKeeper.GetBalance(suite.ctx, suite.Address, "uusd")
				suite.Require().Equal(uusdAmount.Amount, sdk.NewInt(int64(tc.uusdAmount)))
			} else {
				suite.Require().Error(err, tc.errString)
			}
		})
	}
}

func (suite *StableKeeperTestSuite) TestBurnUsq() {
	testCases := []struct {
		name           string
		baseTokenDenom string
		sendTokenDenom string
		uusdAmount     int64
		atomAmount     int64
		atomPrice      int64
		err            bool
		errString      string
	}{
		{
			"ok - burn",
			"uatom",
			"uusd",
			1000,
			104,
			int64(95000),
			false,
			"",
		},
		{
			"fail - wrong denom",
			"uatom",
			"ukuji",
			1000,
			104,
			int64(95000),
			true,
			"ErrSendBaseTokenDenom err",
		},
		{
			"fail - burn blocked",
			"uatom",
			"uusd",
			1000,
			104,
			int64(3300),
			true,
			"Backing Ration < 85%",
		},
	}
	for _, tc := range testCases {
		suite.Setup()
		suite.Commit()

		suite.app.StableKeeper.SetTestingMode(true)

		suite.app.BankKeeper.MintCoins(suite.ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin("uatom", sdk.NewInt(10000))))
		suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, suite.Address, sdk.NewCoins(sdk.NewCoin("uatom", sdk.NewInt(10000))))

		atomAmount := suite.app.BankKeeper.GetBalance(suite.ctx, suite.Address, "uatom")
		suite.Require().Equal(atomAmount.Amount, sdk.NewInt(10000))
		suite.Commit()

		err := suite.MintUsq()
		suite.Require().NoError(err)

		suite.Run(fmt.Sprintf("Case---%s", tc.name), func() {
			suite.app.StableKeeper.SetBaseTokenDenom(suite.ctx, tc.baseTokenDenom)
			suite.app.StableKeeper.UpdateAtomPriceTesting(suite.ctx, sdk.NewInt(tc.atomPrice))

			msg := types.NewMsgBurnUsq(
				suite.Address.String(),
				sdk.NewInt(tc.uusdAmount).String()+tc.sendTokenDenom,
			)
			ctx := sdk.WrapSDKContext(suite.ctx)
			_, err = suite.app.StableKeeper.BurnUsq(ctx, msg)

			if !tc.err {
				suite.Require().NoError(err, tc.name)
				atomAmount := suite.app.BankKeeper.GetBalance(suite.ctx, suite.Address, "uatom")
				suite.Require().Equal(atomAmount.Amount, sdk.NewInt(int64(tc.atomAmount)))
			} else {
				suite.Require().Error(err, tc.errString)
			}
		})
	}
}

func (suite *StableKeeperTestSuite) TestMintUsqGetPriceFromOracle() {
	testCases := []struct {
		name           string
		baseTokenDenom string
		sendTokenDenom string
		atomAmount     int64
		err            bool
		errString      string
	}{
		{
			"ok - mint",
			"uatom",
			"uatom",
			1000,
			false,
			"",
		},
		{
			"ok - mint",
			"uatom",
			"uatom",
			300,
			false,
			"",
		},
		{
			"ok - mint",
			"uatom",
			"uatom",
			730,
			false,
			"",
		},
	}
	for _, tc := range testCases {
		suite.Setup()
		suite.Commit()
		suite.app.StableKeeper.SetTestingMode(false)
		suite.Run(fmt.Sprintf("Case---%s", tc.name), func() {
			suite.app.BankKeeper.MintCoins(suite.ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin("uatom", sdk.NewInt(10000))))
			suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, suite.Address, sdk.NewCoins(sdk.NewCoin("uatom", sdk.NewInt(10000))))
			suite.app.StableKeeper.SetBaseTokenDenom(suite.ctx, tc.baseTokenDenom)
			msg := types.NewMsgMintUsq(
				suite.Address.String(),
				sdk.NewInt(tc.atomAmount).String()+tc.sendTokenDenom,
			)
			ctx := sdk.WrapSDKContext(suite.ctx)
			_, err := suite.app.StableKeeper.MintUsq(ctx, msg)
			suite.Require().NoError(err)
			uusdSuply := suite.app.BankKeeper.GetBalance(suite.ctx, suite.Address, "uusd")
			suite.Require().Greater(uusdSuply.Amount.Int64(), int64(0))
		})
	}
}

func (suite *StableKeeperTestSuite) TestBurnUsqGetPriceFromOracle() {
	testCases := []struct {
		name           string
		baseTokenDenom string
		sendTokenDenom string
		uusdAmount     int64
		err            bool
		errString      string
	}{
		{
			"ok - burn",
			"uatom",
			"uusd",
			1000,
			false,
			"",
		},
		{
			"ok - burn",
			"uatom",
			"uusd",
			300,
			false,
			"",
		},
		{
			"ok - burn",
			"uatom",
			"uusd",
			730,
			false,
			"",
		},
	}
	suite.Setup()
	suite.Commit()

	suite.app.BankKeeper.MintCoins(suite.ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin("uatom", sdk.NewInt(10000))))
	suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, suite.Address, sdk.NewCoins(sdk.NewCoin("uatom", sdk.NewInt(10000))))

	atomAmount := suite.app.BankKeeper.GetBalance(suite.ctx, suite.Address, "uatom")
	suite.Require().Equal(atomAmount.Amount, sdk.NewInt(10000))
	suite.Commit()

	err := suite.MintUsq()
	suite.Require().NoError(err)

	suite.app.StableKeeper.SetBaseTokenDenom(suite.ctx, suite.app.StableKeeper.GetBaseTokenDenom(suite.ctx))
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case---%s", tc.name), func() {
			msg := types.NewMsgBurnUsq(
				suite.Address.String(),
				sdk.NewInt(tc.uusdAmount).String()+tc.sendTokenDenom,
			)
			ctx := sdk.WrapSDKContext(suite.ctx)
			_, err := suite.app.StableKeeper.BurnUsq(ctx, msg)
			suite.Require().NoError(err)
			uatomSuply := suite.app.BankKeeper.GetBalance(suite.ctx, suite.Address, "uatom")
			fmt.Printf("uatomSuply: %s \n", uatomSuply)
			suite.Require().Greater(uatomSuply.Amount.Int64(), int64(0))
		})
	}
}

func (suite *StableKeeperTestSuite) TestExtremeMarketSituations() {
	user1 := apptesting.CreateRandomAccounts(1)[0]
	user2 := apptesting.CreateRandomAccounts(1)[0]
	testCases := []struct {
		name                string
		baseTokenDenom      string
		sendTokenDenom      string
		expectedTokenDenom  string
		sendTokenAmount     int64
		expectedTokenAmount int64
		atomPrice           int64
		err                 bool
		errString           string
		action              string
		address             sdk.AccAddress
	}{
		{
			"mint-user1",
			"uatom",
			"uatom",
			"uusd",
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
			"uatom",
			"uatom",
			"uusd",
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
			"uatom",
			"uatom",
			"uusd",
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
			"uatom",
			"uatom",
			"uusd",
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
			"uatom",
			"uusd",
			"uatom",
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
			"uatom",
			"uusd",
			"uatom",
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
			suite.IncreaseBalance(tc.address, tc.baseTokenDenom, sdk.NewInt(tc.sendTokenAmount))
			suite.app.StableKeeper.SetBaseTokenDenom(suite.ctx, tc.baseTokenDenom)
		}
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case---%s", tc.name), func() {
			suite.app.StableKeeper.UpdateAtomPriceTesting(suite.ctx, sdk.NewInt(tc.atomPrice))
			switch tc.action {
			case "mint":
				msg := types.NewMsgMintUsq(
					tc.address.String(),
					sdk.NewInt(tc.sendTokenAmount).String()+tc.sendTokenDenom,
				)
				ctx := sdk.WrapSDKContext(suite.ctx)
				_, err := suite.app.StableKeeper.MintUsq(ctx, msg)
				suite.Require().NoError(err)
				balance := suite.app.BankKeeper.GetBalance(suite.ctx, tc.address, tc.expectedTokenDenom)
				suite.Require().Greater(balance.Amount.Int64(), int64(0))
			case "burn":
				msg := types.NewMsgBurnUsq(
					tc.address.String(),
					sdk.NewInt(tc.sendTokenAmount).String()+tc.sendTokenDenom,
				)
				ctx := sdk.WrapSDKContext(suite.ctx)
				_, err := suite.app.StableKeeper.BurnUsq(ctx, msg)
				suite.Require().NoError(err)
				balance := suite.app.BankKeeper.GetBalance(suite.ctx, tc.address, tc.expectedTokenDenom)
				suite.Require().Greater(balance.Amount.Int64(), int64(0))
			default:
				suite.Error(nil)
			}
		})
	}
}

func (suite *StableKeeperTestSuite) TestMarketDropOf40() {
	rand.Seed(time.Now().UnixNano())
	burnUser := apptesting.CreateRandomAccounts(1)[0]
	type testCase struct {
		name                string
		baseTokenDenom      string
		sendTokenDenom      string
		expectedTokenDenom  string
		sendTokenAmount     int64
		expectedTokenAmount int64
		atomPrice           int64
		err                 bool
		errString           string
		action              string
		address             sdk.AccAddress
	}
	testCases := []testCase{}

	for i := 0; i < 10; i++ {
		newTestCase := testCase{
			name:                fmt.Sprintf("mint-user%d", i),
			baseTokenDenom:      "uatom",
			sendTokenDenom:      "uatom",
			expectedTokenDenom:  "uusd",
			sendTokenAmount:     1000,
			expectedTokenAmount: 100,
			atomPrice:           int64(rand.Intn(96000-90000) + 90000),
			err:                 false,
			errString:           "",
			action:              "mint",
			address:             apptesting.CreateRandomAccounts(1)[0],
		}
		testCases = append(testCases, newTestCase)
	}

	burnTestCase := testCase{
		name:                fmt.Sprintf("burn"),
		baseTokenDenom:      "uatom",
		sendTokenDenom:      "uusd",
		expectedTokenDenom:  "uatom",
		sendTokenAmount:     8000,
		expectedTokenAmount: 50,
		atomPrice:           int64(57000),
		err:                 true,
		errString:           "Backing Ration < 85%",
		action:              "burn",
		address:             burnUser,
	}
	testCases = append(testCases, burnTestCase)

	suite.Setup()
	suite.Commit()
	suite.app.StableKeeper.SetTestingMode(true)
	for _, tc := range testCases {
		if tc.action == "mint" {
			suite.IncreaseBalance(tc.address, tc.baseTokenDenom, sdk.NewInt(tc.sendTokenAmount))
			suite.app.StableKeeper.SetBaseTokenDenom(suite.ctx, tc.baseTokenDenom)
		}
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case---%s--Price---%f", tc.name, float64(float64(tc.atomPrice)/10000)), func() {
			suite.app.StableKeeper.UpdateAtomPriceTesting(suite.ctx, sdk.NewInt(tc.atomPrice))
			switch tc.action {
			case "mint":
				msg := types.NewMsgMintUsq(
					tc.address.String(),
					sdk.NewInt(tc.sendTokenAmount).String()+tc.sendTokenDenom,
				)
				ctx := sdk.WrapSDKContext(suite.ctx)
				_, err := suite.app.StableKeeper.MintUsq(ctx, msg)
				if tc.err {
					suite.Require().Error(err, tc.errString)
				} else {
					suite.Require().NoError(err)
					balance := suite.app.BankKeeper.GetBalance(suite.ctx, tc.address, tc.expectedTokenDenom)
					suite.Require().Greater(balance.Amount.Int64(), int64(0))
				}
			case "burn":
				msg := types.NewMsgBurnUsq(
					tc.address.String(),
					sdk.NewInt(tc.sendTokenAmount).String()+tc.sendTokenDenom,
				)
				ctx := sdk.WrapSDKContext(suite.ctx)
				_, err := suite.app.StableKeeper.BurnUsq(ctx, msg)
				if tc.err {
					suite.Require().Error(err, tc.errString)
				} else {
					suite.Require().NoError(err)
					balance := suite.app.BankKeeper.GetBalance(suite.ctx, tc.address, tc.expectedTokenDenom)
					suite.Require().Greater(balance.Amount.Int64(), int64(0))
				}
			default:
				suite.Error(nil)
			}
		})
	}
}
