package perpetual_test

import (
	"fmt"
	"time"

	apptesting "github.com/QuadrateOrg/core/app/apptesting"
	"github.com/QuadrateOrg/core/x/perpetual"
	"github.com/QuadrateOrg/core/x/perpetual/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *PerpetualTestSuite) TestCalculateFundingPaymentByBlock() {
	suite.Setup()
	suite.Commit()
	suite.RegisterValidator()
	s.ctx = s.ctx.WithBlockTime(time.Now())

	user_1 := apptesting.CreateRandomAccounts(1)[0]
	user_2 := apptesting.CreateRandomAccounts(1)[0]
	user_3 := apptesting.CreateRandomAccounts(1)[0]
	user_4 := apptesting.CreateRandomAccounts(1)[0]

	suite.SetupOracleKeeper(s.GetNormalTestVault().GetOracleAssetId())
	suite.OracleAggregateExchangeRateFromInput("10" + s.GetNormalTestVault().GetOracleAssetId())

	s.app.PerpetualKeeper.AppendVault(s.ctx, *s.GetNormalTestVault())

	// Open long position user_1
	suite.AddTestCoinsToCustomAccount(sdk.NewInt(100), TestDefaultXDenom, user_1)
	msg := types.NewMsgOpen(
		user_1.String(),
		types.PerpetualTradeType_PERPETUAL_LONG_POSITION,
		sdk.NewDec(10),
		s.app.PerpetualKeeper.GenerateVaultIdHash(TestDefaultXDenom, TestDefaultYDenom),
		sdk.NewInt(10).String()+TestDefaultXDenom,
	)
	ctx := sdk.WrapSDKContext(suite.ctx)
	_, err := suite.app.PerpetualKeeper.Open(ctx, msg)
	suite.Require().NoError(err)

	// Open long position user_2
	suite.AddTestCoinsToCustomAccount(sdk.NewInt(50), TestDefaultXDenom, user_2)
	msg2 := types.NewMsgOpen(
		user_2.String(),
		types.PerpetualTradeType_PERPETUAL_LONG_POSITION,
		sdk.NewDec(2),
		s.app.PerpetualKeeper.GenerateVaultIdHash(TestDefaultXDenom, TestDefaultYDenom),
		sdk.NewInt(5).String()+TestDefaultXDenom,
	)
	ctx = sdk.WrapSDKContext(suite.ctx)
	_, err2 := suite.app.PerpetualKeeper.Open(ctx, msg2)
	suite.Require().NoError(err2)

	// Open short position user_3
	suite.AddTestCoinsToCustomAccount(sdk.NewInt(50), TestDefaultXDenom, user_3)
	msg3 := types.NewMsgOpen(
		user_3.String(),
		types.PerpetualTradeType_PERPETUAL_SHORT_POSITION,
		sdk.NewDec(3),
		s.app.PerpetualKeeper.GenerateVaultIdHash(TestDefaultXDenom, TestDefaultYDenom),
		sdk.NewInt(5).String()+TestDefaultXDenom,
	)
	ctx = sdk.WrapSDKContext(suite.ctx)
	_, err3 := suite.app.PerpetualKeeper.Open(ctx, msg3)
	suite.Require().NoError(err3)

	// Open short position user_4
	suite.AddTestCoinsToCustomAccount(sdk.NewInt(30), TestDefaultXDenom, user_4)
	msg4 := types.NewMsgOpen(
		user_4.String(),
		types.PerpetualTradeType_PERPETUAL_SHORT_POSITION,
		sdk.NewDec(2),
		s.app.PerpetualKeeper.GenerateVaultIdHash(TestDefaultXDenom, TestDefaultYDenom),
		sdk.NewInt(3).String()+TestDefaultXDenom,
	)
	ctx = sdk.WrapSDKContext(suite.ctx)
	_, err4 := suite.app.PerpetualKeeper.Open(ctx, msg4)
	suite.Require().NoError(err4)

	old_position1, f := s.app.PerpetualKeeper.GetPositionByPositionId(
		s.ctx,
		s.app.PerpetualKeeper.GenerateTraderPositionId(
			user_1.String(),
			TestDefaultXDenom,
			s.app.PerpetualKeeper.GenerateVaultIdHash(TestDefaultXDenom, TestDefaultYDenom),
			types.PerpetualTradeType_PERPETUAL_LONG_POSITION,
		),
	)
	suite.Require().Equal(true, f)

	old_position2, f := s.app.PerpetualKeeper.GetPositionByPositionId(
		s.ctx,
		s.app.PerpetualKeeper.GenerateTraderPositionId(
			user_2.String(),
			TestDefaultXDenom,
			s.app.PerpetualKeeper.GenerateVaultIdHash(TestDefaultXDenom, TestDefaultYDenom),
			types.PerpetualTradeType_PERPETUAL_LONG_POSITION,
		),
	)
	suite.Require().Equal(true, f)

	old_position3, f := s.app.PerpetualKeeper.GetPositionByPositionId(
		s.ctx,
		s.app.PerpetualKeeper.GenerateTraderPositionId(
			user_3.String(),
			TestDefaultXDenom,
			s.app.PerpetualKeeper.GenerateVaultIdHash(TestDefaultXDenom, TestDefaultYDenom),
			types.PerpetualTradeType_PERPETUAL_SHORT_POSITION,
		),
	)
	suite.Require().Equal(true, f)

	old_position4, f := s.app.PerpetualKeeper.GetPositionByPositionId(
		s.ctx,
		s.app.PerpetualKeeper.GenerateTraderPositionId(
			user_4.String(),
			TestDefaultXDenom,
			s.app.PerpetualKeeper.GenerateVaultIdHash(TestDefaultXDenom, TestDefaultYDenom),
			types.PerpetualTradeType_PERPETUAL_SHORT_POSITION,
		),
	)
	suite.Require().Equal(true, f)

	vault, _ := s.app.PerpetualKeeper.GetVaultByVaultId(s.ctx, s.GetNormalTestVault().VaultId)
	fmt.Printf("price in vAMM: %f\n", float64(vault.X.MulRaw(10000).Quo(vault.Y).Int64())/10000)

	price_for_calc := ((0.001 * 24) - float64(vault.X.MulRaw(10000).Quo(vault.Y).Int64())/10000) * -1

	suite.OracleAggregateExchangeRateFromInput(fmt.Sprintf("%f", price_for_calc) + s.GetNormalTestVault().GetOracleAssetId())

	perpetual.EndBlocker(s.ctx, s.app.PerpetualKeeper)

	new_position1, f := s.app.PerpetualKeeper.GetPositionByPositionId(
		s.ctx,
		s.app.PerpetualKeeper.GenerateTraderPositionId(
			user_1.String(),
			TestDefaultXDenom,
			s.app.PerpetualKeeper.GenerateVaultIdHash(TestDefaultXDenom, TestDefaultYDenom),
			types.PerpetualTradeType_PERPETUAL_LONG_POSITION,
		),
	)
	suite.Require().Equal(true, f)

	new_position2, f := s.app.PerpetualKeeper.GetPositionByPositionId(
		s.ctx,
		s.app.PerpetualKeeper.GenerateTraderPositionId(
			user_2.String(),
			TestDefaultXDenom,
			s.app.PerpetualKeeper.GenerateVaultIdHash(TestDefaultXDenom, TestDefaultYDenom),
			types.PerpetualTradeType_PERPETUAL_LONG_POSITION,
		),
	)
	suite.Require().Equal(true, f)

	new_position3, f := s.app.PerpetualKeeper.GetPositionByPositionId(
		s.ctx,
		s.app.PerpetualKeeper.GenerateTraderPositionId(
			user_3.String(),
			TestDefaultXDenom,
			s.app.PerpetualKeeper.GenerateVaultIdHash(TestDefaultXDenom, TestDefaultYDenom),
			types.PerpetualTradeType_PERPETUAL_SHORT_POSITION,
		),
	)
	suite.Require().Equal(true, f)

	new_position4, f := s.app.PerpetualKeeper.GetPositionByPositionId(
		s.ctx,
		s.app.PerpetualKeeper.GenerateTraderPositionId(
			user_4.String(),
			TestDefaultXDenom,
			s.app.PerpetualKeeper.GenerateVaultIdHash(TestDefaultXDenom, TestDefaultYDenom),
			types.PerpetualTradeType_PERPETUAL_SHORT_POSITION,
		),
	)
	suite.Require().Equal(true, f)

	fmt.Printf("difference between new and old: %v\n", (new_position1.ReturnAmount.Sub(old_position1.ReturnAmount)))
	fmt.Printf("difference between new and old: %v\n", (new_position2.ReturnAmount.Sub(old_position2.ReturnAmount)))
	fmt.Printf("difference between new and old: %v\n", (new_position3.ReturnAmount.Sub(old_position3.ReturnAmount)))
	fmt.Printf("difference between new and old: %v\n", (new_position4.ReturnAmount.Sub(old_position4.ReturnAmount)))
}
