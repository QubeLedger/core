package keeper_test

import (
	"fmt"

	"github.com/QuadrateOrg/core/x/stable/types"

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
			"ok - mint",
			"uatom",
			"uatom",
			1000,
			9471,
			int64(95000),
			false,
			"",
		},
		{
			"fail - wrong denom",
			"uatom",
			"ukuji",
			1000,
			9272,
			int64(93000),
			true,
			"ErrSendBaseTokenDenom err",
		},
		{
			"fail - mint blocked",
			"uatom",
			"uatom",
			1000,
			9471,
			int64(1500000),
			true,
			"Backing Ration >= 225%",
		},
	}

	for _, tc := range testCases {
		suite.Setup()
		suite.Commit()
		suite.app.StableKeeper.SetTestingMode(true)
		suite.Run(fmt.Sprintf("Case    %s", tc.name), func() {
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
			"Backing Ration < 75%",
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

		suite.Run(fmt.Sprintf("Case    %s", tc.name), func() {
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

func (suite *StableKeeperTestSuite) MintUsq() error {
	suite.app.StableKeeper.SetBaseTokenDenom(suite.ctx, "uatom")
	suite.app.StableKeeper.UpdateAtomPriceTesting(suite.ctx, sdk.NewInt(95000))
	msg := types.NewMsgMintUsq(
		suite.Address.String(),
		sdk.NewInt(10000).String()+"uatom",
	)
	ctx := sdk.WrapSDKContext(suite.ctx)
	_, err := suite.app.StableKeeper.MintUsq(ctx, msg)
	if err != nil {
		return err
	}
	return nil
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
	suite.Setup()
	suite.Commit()
	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
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
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
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
