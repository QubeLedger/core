package keeper_test

import (
	"fmt"
	"time"

	"github.com/QuadrateOrg/core/x/perpetual/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *PerpetualKeeperTestSuite) TestOpenLongPosition() {
	testCases := []struct {
		name            string
		sendTokenDenom  string
		sendTokenAmount int64
		leverage        sdk.Dec
		oracleAssetId   string
		price           string
		err             bool
		errString       string
	}{
		{
			"ok-open",
			TestDefaultXDenom,
			2000,
			sdk.NewDec(2),
			TestDefaultOracleAssetId,
			"8",
			false,
			"",
		},
	}

	suite.Setup()
	suite.Commit()
	suite.RegisterValidator()
	s.ctx = s.ctx.WithBlockTime(time.Now())

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case---%s", tc.name), func() {
			suite.SetupOracleKeeper(tc.oracleAssetId)
			suite.OracleAggregateExchangeRateFromInput(tc.price + tc.oracleAssetId)

			suite.AddTestCoinsToCustomAccount(sdk.NewInt(tc.sendTokenAmount), tc.sendTokenDenom, suite.Address)

			s.app.PerpetualKeeper.AppendVault(s.ctx, *s.GetNormalTestVault())
			vault, _ := s.app.PerpetualKeeper.GetVaultByVaultId(s.ctx, s.GetNormalTestVault().VaultId)
			fmt.Printf("price in vAMM: %f\n", float64(vault.X.MulRaw(10000).Quo(vault.Y).Int64())/10000)
			msg := types.NewMsgOpen(
				suite.Address.String(),
				types.PerpetualTradeType_PERPETUAL_LONG_POSITION,
				tc.leverage,
				s.app.PerpetualKeeper.GenerateVaultIdHash(TestDefaultXDenom, TestDefaultYDenom),
				sdk.NewInt(tc.sendTokenAmount).String()+tc.sendTokenDenom,
			)
			ctx := sdk.WrapSDKContext(suite.ctx)
			_, err := suite.app.PerpetualKeeper.Open(ctx, msg)
			suite.Require().NoError(err)

			position, f := s.app.PerpetualKeeper.GetPositionByPositionId(
				s.ctx,
				s.app.PerpetualKeeper.GenerateTraderPositionId(suite.Address.String(),
					tc.sendTokenDenom,
					s.app.PerpetualKeeper.GenerateVaultIdHash(TestDefaultXDenom, TestDefaultYDenom),
					types.PerpetualTradeType_PERPETUAL_LONG_POSITION,
					tc.leverage,
				),
			)
			suite.Require().Equal(true, f)
			suite.Require().Greater(position.ReturnAmount.RoundInt().Int64(), int64(0))

			vault, _ = s.app.PerpetualKeeper.GetVaultByVaultId(s.ctx, s.GetNormalTestVault().VaultId)
			fmt.Printf("price in vAMM after LONG: %f\n", float64(vault.X.MulRaw(10000).Quo(vault.Y).Int64())/10000)
		})
	}
}

func (suite *PerpetualKeeperTestSuite) TestOpenShortPosition() {
	testCases := []struct {
		name            string
		sendTokenDenom  string
		sendTokenAmount int64
		oracleAssetId   string
		price           string
		err             bool
		errString       string
	}{
		{
			"ok-open",
			TestDefaultXDenom,
			1000,
			TestDefaultOracleAssetId,
			"8",
			false,
			"",
		},
	}

	suite.Setup()
	suite.Commit()
	suite.RegisterValidator()
	s.ctx = s.ctx.WithBlockTime(time.Now())

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case---%s", tc.name), func() {
			suite.SetupOracleKeeper(tc.oracleAssetId)
			suite.OracleAggregateExchangeRateFromInput(tc.price + tc.oracleAssetId)

			suite.AddTestCoinsToCustomAccount(sdk.NewInt(tc.sendTokenAmount), tc.sendTokenDenom, suite.Address)

			s.app.PerpetualKeeper.AppendVault(s.ctx, *s.GetNormalTestVault())
			vault, _ := s.app.PerpetualKeeper.GetVaultByVaultId(s.ctx, s.GetNormalTestVault().VaultId)
			fmt.Printf("price in vAMM: %f\n", float64(vault.X.MulRaw(10000).Quo(vault.Y).Int64())/10000)
			msg := types.NewMsgOpen(
				suite.Address.String(),
				types.PerpetualTradeType_PERPETUAL_SHORT_POSITION,
				sdk.NewDec(2),
				s.app.PerpetualKeeper.GenerateVaultIdHash(TestDefaultXDenom, TestDefaultYDenom),
				sdk.NewInt(tc.sendTokenAmount).String()+tc.sendTokenDenom,
			)
			ctx := sdk.WrapSDKContext(suite.ctx)
			_, err := suite.app.PerpetualKeeper.Open(ctx, msg)
			suite.Require().NoError(err)

			position, f := s.app.PerpetualKeeper.GetPositionByPositionId(
				s.ctx,
				s.app.PerpetualKeeper.GenerateTraderPositionId(suite.Address.String(),
					tc.sendTokenDenom,
					s.app.PerpetualKeeper.GenerateVaultIdHash(TestDefaultXDenom, TestDefaultYDenom),
					types.PerpetualTradeType_PERPETUAL_SHORT_POSITION,
					sdk.NewDec(2),
				),
			)
			suite.Require().Equal(true, f)
			suite.Require().Greater(position.ReturnAmount.RoundInt().Int64(), int64(0))

			vault, _ = s.app.PerpetualKeeper.GetVaultByVaultId(s.ctx, s.GetNormalTestVault().VaultId)
			fmt.Printf("price in vAMM after SHORT: %f\n", float64(vault.X.MulRaw(10000).Quo(vault.Y).Int64())/10000)
		})
	}
}
