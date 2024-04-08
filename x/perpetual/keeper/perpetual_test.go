package keeper_test

import (
	"fmt"
	"time"

	apptesting "github.com/QuadrateOrg/core/app/apptesting"
	"github.com/QuadrateOrg/core/x/perpetual/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type OpenLongWithUser struct {
	sendTokenDenom  string
	sendTokenAmount int64
	type_trade      types.PerpetualTradeType
	leverage        sdk.Dec
	user            sdk.AccAddress
}

type CloseLongWithUser struct {
	id              string
	sendTokenAmount int64
	user            sdk.AccAddress
}

func (suite *PerpetualKeeperTestSuite) TestOpenAndCloseLongDifferenceUsers() {
	suite.Setup()
	suite.Commit()

	user_1 := apptesting.CreateRandomAccounts(1)[0]
	user_2 := apptesting.CreateRandomAccounts(1)[0]

	testCases := []struct {
		name          string
		open_long     []OpenLongWithUser
		close_long    []CloseLongWithUser
		oracleAssetId string
		price         string
		err           bool
		errString     string
	}{
		{
			"ok-open",
			[]OpenLongWithUser{
				{
					sendTokenDenom:  TestDefaultXDenom,
					sendTokenAmount: 2000,
					leverage:        sdk.NewDec(2),
					user:            user_1,
				},
				{
					sendTokenDenom:  TestDefaultXDenom,
					sendTokenAmount: 4000,
					leverage:        sdk.NewDec(1),
					user:            user_2,
				},
			},
			[]CloseLongWithUser{
				{
					id: s.app.PerpetualKeeper.GenerateTraderPositionId(
						user_1.String(),
						TestDefaultXDenom,
						s.app.PerpetualKeeper.GenerateVaultIdHash(TestDefaultXDenom, TestDefaultYDenom),
						types.PerpetualTradeType_PERPETUAL_LONG_POSITION,
					),
					user: user_1,
				},
				{
					id: s.app.PerpetualKeeper.GenerateTraderPositionId(
						user_2.String(),
						TestDefaultXDenom,
						s.app.PerpetualKeeper.GenerateVaultIdHash(TestDefaultXDenom, TestDefaultYDenom),
						types.PerpetualTradeType_PERPETUAL_LONG_POSITION,
					),
					user: user_2,
				},
			},
			TestDefaultOracleAssetId,
			"10",
			false,
			"",
		},
	}

	suite.RegisterValidator()
	s.ctx = s.ctx.WithBlockTime(time.Now())

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case---%s", tc.name), func() {

			suite.SetupOracleKeeper(tc.oracleAssetId)
			suite.OracleAggregateExchangeRateFromInput(tc.price + tc.oracleAssetId)
			s.app.PerpetualKeeper.AppendVault(s.ctx, *s.GetNormalTestVault())

			for _, open_position := range tc.open_long {
				suite.AddTestCoinsToCustomAccount(sdk.NewInt(open_position.sendTokenAmount), open_position.sendTokenDenom, open_position.user)

				msg := types.NewMsgOpen(
					open_position.user.String(),
					types.PerpetualTradeType_PERPETUAL_LONG_POSITION,
					open_position.leverage,
					s.app.PerpetualKeeper.GenerateVaultIdHash(TestDefaultXDenom, TestDefaultYDenom),
					sdk.NewInt(open_position.sendTokenAmount).String()+open_position.sendTokenDenom,
				)
				ctx := sdk.WrapSDKContext(suite.ctx)
				_, err := suite.app.PerpetualKeeper.Open(ctx, msg)
				suite.Require().NoError(err)

				position, f := s.app.PerpetualKeeper.GetPositionByPositionId(
					s.ctx,
					s.app.PerpetualKeeper.GenerateTraderPositionId(
						open_position.user.String(),
						open_position.sendTokenDenom,
						s.app.PerpetualKeeper.GenerateVaultIdHash(TestDefaultXDenom, TestDefaultYDenom),
						types.PerpetualTradeType_PERPETUAL_LONG_POSITION,
					),
				)
				suite.Require().Equal(true, f)
				suite.Require().Greater(position.ReturnAmount.Int64(), int64(0))
			}

			for _, close_position := range tc.close_long {

				position, f := s.app.PerpetualKeeper.GetPositionByPositionId(
					s.ctx,
					close_position.id,
				)
				suite.Require().Equal(true, f)

				msg := types.NewMsgClose(
					close_position.user.String(),
					close_position.id,
					position.ReturnAmount,
				)
				ctx := sdk.WrapSDKContext(suite.ctx)
				_, err := suite.app.PerpetualKeeper.Close(ctx, msg)
				suite.Require().NoError(err)

				_, f = s.app.PerpetualKeeper.GetPositionByPositionId(
					s.ctx,
					s.app.PerpetualKeeper.GenerateTraderPositionId(
						close_position.user.String(),
						TestDefaultXDenom,
						s.app.PerpetualKeeper.GenerateVaultIdHash(TestDefaultXDenom, TestDefaultYDenom),
						types.PerpetualTradeType_PERPETUAL_LONG_POSITION,
					),
				)
				suite.Require().Equal(false, f)
			}

			for _, open_position := range tc.open_long {
				new_balance := (s.app.BankKeeper.GetBalance(s.ctx, open_position.user, TestDefaultXDenom)).Amount
				fmt.Printf("difference between: %v\n", new_balance.SubRaw(open_position.sendTokenAmount).Int64())
			}
		})
	}
}

func (suite *PerpetualKeeperTestSuite) TestOpenAndCloseLongDifferenceUsersWithNotMaxAmount() {
	suite.Setup()
	suite.Commit()

	user_1 := apptesting.CreateRandomAccounts(1)[0]
	user_2 := apptesting.CreateRandomAccounts(1)[0]

	testCases := []struct {
		name          string
		open_long     []OpenLongWithUser
		close_long    []CloseLongWithUser
		oracleAssetId string
		price         string
		err           bool
		errString     string
	}{
		{
			"ok-open",
			[]OpenLongWithUser{
				{
					sendTokenDenom:  TestDefaultXDenom,
					sendTokenAmount: 2000,
					leverage:        sdk.NewDec(2),
					user:            user_1,
				},
				{
					sendTokenDenom:  TestDefaultXDenom,
					sendTokenAmount: 4000,
					leverage:        sdk.NewDec(1),
					user:            user_2,
				},
			},
			[]CloseLongWithUser{
				{
					id: s.app.PerpetualKeeper.GenerateTraderPositionId(
						user_1.String(),
						TestDefaultXDenom,
						s.app.PerpetualKeeper.GenerateVaultIdHash(TestDefaultXDenom, TestDefaultYDenom),
						types.PerpetualTradeType_PERPETUAL_LONG_POSITION,
					),
					sendTokenAmount: 50,
					user:            user_1,
				},
				{
					id: s.app.PerpetualKeeper.GenerateTraderPositionId(
						user_1.String(),
						TestDefaultXDenom,
						s.app.PerpetualKeeper.GenerateVaultIdHash(TestDefaultXDenom, TestDefaultYDenom),
						types.PerpetualTradeType_PERPETUAL_LONG_POSITION,
					),
					sendTokenAmount: 335,
					user:            user_1,
				},
				{
					id: s.app.PerpetualKeeper.GenerateTraderPositionId(
						user_2.String(),
						TestDefaultXDenom,
						s.app.PerpetualKeeper.GenerateVaultIdHash(TestDefaultXDenom, TestDefaultYDenom),
						types.PerpetualTradeType_PERPETUAL_LONG_POSITION,
					),
					sendTokenAmount: 356,
					user:            user_2,
				},
			},
			TestDefaultOracleAssetId,
			"10",
			false,
			"",
		},
	}

	suite.RegisterValidator()
	s.ctx = s.ctx.WithBlockTime(time.Now())

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case---%s", tc.name), func() {

			suite.SetupOracleKeeper(tc.oracleAssetId)
			suite.OracleAggregateExchangeRateFromInput(tc.price + tc.oracleAssetId)
			s.app.PerpetualKeeper.AppendVault(s.ctx, *s.GetNormalTestVault())

			for _, open_position := range tc.open_long {
				suite.AddTestCoinsToCustomAccount(sdk.NewInt(open_position.sendTokenAmount), open_position.sendTokenDenom, open_position.user)

				msg := types.NewMsgOpen(
					open_position.user.String(),
					types.PerpetualTradeType_PERPETUAL_LONG_POSITION,
					open_position.leverage,
					s.app.PerpetualKeeper.GenerateVaultIdHash(TestDefaultXDenom, TestDefaultYDenom),
					sdk.NewInt(open_position.sendTokenAmount).String()+open_position.sendTokenDenom,
				)
				ctx := sdk.WrapSDKContext(suite.ctx)
				_, err := suite.app.PerpetualKeeper.Open(ctx, msg)
				suite.Require().NoError(err)

				position, f := s.app.PerpetualKeeper.GetPositionByPositionId(
					s.ctx,
					s.app.PerpetualKeeper.GenerateTraderPositionId(
						open_position.user.String(),
						open_position.sendTokenDenom,
						s.app.PerpetualKeeper.GenerateVaultIdHash(TestDefaultXDenom, TestDefaultYDenom),
						types.PerpetualTradeType_PERPETUAL_LONG_POSITION,
					),
				)
				suite.Require().Equal(true, f)
				suite.Require().Greater(position.ReturnAmount.Int64(), int64(0))
			}

			for _, close_position := range tc.close_long {
				msg := types.NewMsgClose(
					close_position.user.String(),
					close_position.id,
					sdk.NewInt(close_position.sendTokenAmount),
				)
				ctx := sdk.WrapSDKContext(suite.ctx)
				_, err := suite.app.PerpetualKeeper.Close(ctx, msg)
				suite.Require().NoError(err)
			}

			for _, open_position := range tc.open_long {
				new_balance := (s.app.BankKeeper.GetBalance(s.ctx, open_position.user, TestDefaultXDenom)).Amount
				fmt.Printf("difference between: %v\n", new_balance.SubRaw(open_position.sendTokenAmount).Int64())
			}
		})
	}
}

func (suite *PerpetualKeeperTestSuite) TestOpenAndCloseShortDifferenceUsersWithMaxAmount() {
	suite.Setup()
	suite.Commit()
	user_1 := apptesting.CreateRandomAccounts(1)[0]
	user_2 := apptesting.CreateRandomAccounts(1)[0]

	testCases := []struct {
		name          string
		open_long     []OpenLongWithUser
		close_long    []CloseLongWithUser
		oracleAssetId string
		price         string
		err           bool
		errString     string
	}{
		{
			"ok-open big long",
			[]OpenLongWithUser{
				{
					sendTokenDenom:  TestDefaultXDenom,
					sendTokenAmount: 2000,
					leverage:        sdk.NewDec(2),
					type_trade:      types.PerpetualTradeType_PERPETUAL_SHORT_POSITION,
					user:            user_1,
				},
				{
					sendTokenDenom:  TestDefaultXDenom,
					sendTokenAmount: 4000,
					leverage:        sdk.NewDec(3),
					type_trade:      types.PerpetualTradeType_PERPETUAL_LONG_POSITION,
					user:            user_2,
				},
			},
			[]CloseLongWithUser{
				{
					id: s.app.PerpetualKeeper.GenerateTraderPositionId(
						user_1.String(),
						TestDefaultXDenom,
						s.app.PerpetualKeeper.GenerateVaultIdHash(TestDefaultXDenom, TestDefaultYDenom),
						types.PerpetualTradeType_PERPETUAL_SHORT_POSITION,
					),
					user: user_1,
				},
				{
					id: s.app.PerpetualKeeper.GenerateTraderPositionId(
						user_2.String(),
						TestDefaultXDenom,
						s.app.PerpetualKeeper.GenerateVaultIdHash(TestDefaultXDenom, TestDefaultYDenom),
						types.PerpetualTradeType_PERPETUAL_LONG_POSITION,
					),
					user: user_2,
				},
			},
			TestDefaultOracleAssetId,
			"10",
			false,
			"",
		},
		{
			"ok-open bit short",
			[]OpenLongWithUser{
				{
					sendTokenDenom:  TestDefaultXDenom,
					sendTokenAmount: 2000,
					leverage:        sdk.NewDec(2),
					type_trade:      types.PerpetualTradeType_PERPETUAL_LONG_POSITION,
					user:            user_1,
				},
				{
					sendTokenDenom:  TestDefaultXDenom,
					sendTokenAmount: 4000,
					leverage:        sdk.NewDec(3),
					type_trade:      types.PerpetualTradeType_PERPETUAL_SHORT_POSITION,
					user:            user_2,
				},
			},
			[]CloseLongWithUser{
				{
					id: s.app.PerpetualKeeper.GenerateTraderPositionId(
						user_1.String(),
						TestDefaultXDenom,
						s.app.PerpetualKeeper.GenerateVaultIdHash(TestDefaultXDenom, TestDefaultYDenom),
						types.PerpetualTradeType_PERPETUAL_LONG_POSITION,
					),
					user: user_1,
				},
				{
					id: s.app.PerpetualKeeper.GenerateTraderPositionId(
						user_2.String(),
						TestDefaultXDenom,
						s.app.PerpetualKeeper.GenerateVaultIdHash(TestDefaultXDenom, TestDefaultYDenom),
						types.PerpetualTradeType_PERPETUAL_SHORT_POSITION,
					),
					user: user_2,
				},
			},
			TestDefaultOracleAssetId,
			"10",
			false,
			"",
		},
		{
			"ok-open small short",
			[]OpenLongWithUser{
				{
					sendTokenDenom:  TestDefaultXDenom,
					sendTokenAmount: 2000,
					leverage:        sdk.NewDec(2),
					type_trade:      types.PerpetualTradeType_PERPETUAL_LONG_POSITION,
					user:            user_1,
				},
				{
					sendTokenDenom:  TestDefaultXDenom,
					sendTokenAmount: 1000,
					leverage:        sdk.NewDec(2),
					type_trade:      types.PerpetualTradeType_PERPETUAL_SHORT_POSITION,
					user:            user_2,
				},
			},
			[]CloseLongWithUser{
				{
					id: s.app.PerpetualKeeper.GenerateTraderPositionId(
						user_1.String(),
						TestDefaultXDenom,
						s.app.PerpetualKeeper.GenerateVaultIdHash(TestDefaultXDenom, TestDefaultYDenom),
						types.PerpetualTradeType_PERPETUAL_LONG_POSITION,
					),
					user: user_1,
				},
				{
					id: s.app.PerpetualKeeper.GenerateTraderPositionId(
						user_2.String(),
						TestDefaultXDenom,
						s.app.PerpetualKeeper.GenerateVaultIdHash(TestDefaultXDenom, TestDefaultYDenom),
						types.PerpetualTradeType_PERPETUAL_SHORT_POSITION,
					),
					user: user_2,
				},
			},
			TestDefaultOracleAssetId,
			"10",
			false,
			"",
		},
		{
			"ok-open small short",
			[]OpenLongWithUser{
				{
					sendTokenDenom:  TestDefaultXDenom,
					sendTokenAmount: 3500,
					leverage:        sdk.NewDec(2),
					type_trade:      types.PerpetualTradeType_PERPETUAL_LONG_POSITION,
					user:            user_1,
				},
				{
					sendTokenDenom:  TestDefaultXDenom,
					sendTokenAmount: 500,
					leverage:        sdk.NewDec(1),
					type_trade:      types.PerpetualTradeType_PERPETUAL_SHORT_POSITION,
					user:            user_2,
				},
			},
			[]CloseLongWithUser{
				{
					id: s.app.PerpetualKeeper.GenerateTraderPositionId(
						user_1.String(),
						TestDefaultXDenom,
						s.app.PerpetualKeeper.GenerateVaultIdHash(TestDefaultXDenom, TestDefaultYDenom),
						types.PerpetualTradeType_PERPETUAL_LONG_POSITION,
					),
					user: user_1,
				},
				{
					id: s.app.PerpetualKeeper.GenerateTraderPositionId(
						user_2.String(),
						TestDefaultXDenom,
						s.app.PerpetualKeeper.GenerateVaultIdHash(TestDefaultXDenom, TestDefaultYDenom),
						types.PerpetualTradeType_PERPETUAL_SHORT_POSITION,
					),
					user: user_2,
				},
			},
			TestDefaultOracleAssetId,
			"10",
			false,
			"",
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case---%s", tc.name), func() {

			suite.Setup()
			suite.Commit()

			suite.RegisterValidator()
			s.ctx = s.ctx.WithBlockTime(time.Now())

			suite.SetupOracleKeeper(tc.oracleAssetId)
			suite.OracleAggregateExchangeRateFromInput(tc.price + tc.oracleAssetId)
			s.app.PerpetualKeeper.AppendVault(s.ctx, *s.GetNormalTestVault())

			for _, open_position := range tc.open_long {
				suite.AddTestCoinsToCustomAccount(sdk.NewInt(open_position.sendTokenAmount), open_position.sendTokenDenom, open_position.user)

				msg := types.NewMsgOpen(
					open_position.user.String(),
					open_position.type_trade,
					open_position.leverage,
					s.app.PerpetualKeeper.GenerateVaultIdHash(TestDefaultXDenom, TestDefaultYDenom),
					sdk.NewInt(open_position.sendTokenAmount).String()+open_position.sendTokenDenom,
				)
				ctx := sdk.WrapSDKContext(suite.ctx)
				_, err := suite.app.PerpetualKeeper.Open(ctx, msg)
				suite.Require().NoError(err)

				vault, _ := s.app.PerpetualKeeper.GetVaultByVaultId(s.ctx, s.GetNormalTestVault().VaultId)
				fmt.Printf("price in vAMM: %f\n", float64(vault.X.MulRaw(10000).Quo(vault.Y).Int64())/10000)

				position, f := s.app.PerpetualKeeper.GetPositionByPositionId(
					s.ctx,
					s.app.PerpetualKeeper.GenerateTraderPositionId(
						open_position.user.String(),
						open_position.sendTokenDenom,
						s.app.PerpetualKeeper.GenerateVaultIdHash(TestDefaultXDenom, TestDefaultYDenom),
						open_position.type_trade,
					),
				)
				suite.Require().Equal(true, f)
				suite.Require().Greater(position.ReturnAmount.Int64(), int64(0))
			}

			for _, close_position := range tc.close_long {

				position, f := s.app.PerpetualKeeper.GetPositionByPositionId(
					s.ctx,
					close_position.id,
				)
				suite.Require().Equal(true, f)

				msg := types.NewMsgClose(
					close_position.user.String(),
					close_position.id,
					position.ReturnAmount,
				)
				ctx := sdk.WrapSDKContext(suite.ctx)
				_, err := suite.app.PerpetualKeeper.Close(ctx, msg)
				suite.Require().NoError(err)
			}

			for _, open_position := range tc.open_long {
				new_balance := (s.app.BankKeeper.GetBalance(s.ctx, open_position.user, TestDefaultXDenom)).Amount
				fmt.Printf("difference between: %v type_trade: %v\n", new_balance.SubRaw(open_position.sendTokenAmount).Int64(), open_position.type_trade)
			}

			fmt.Printf("\n")
		})
	}
}
