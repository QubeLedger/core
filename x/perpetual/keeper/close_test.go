package keeper_test

import (
	"fmt"
	"time"

	"github.com/QuadrateOrg/core/x/perpetual/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type OpenLong struct {
	sendTokenDenom  string
	sendTokenAmount int64
	leverage        sdk.Dec
}

type CloseLong struct {
	id              string
	sendTokenAmount int64
}

func (suite *PerpetualKeeperTestSuite) TestCloseLongPosition() {
	suite.Setup()
	suite.Commit()

	testCases := []struct {
		name          string
		open_long     []OpenLong
		close_long    []CloseLong
		oracleAssetId string
		price         string
		err           bool
		errString     string
	}{
		{
			"ok-open",
			[]OpenLong{
				{
					sendTokenDenom:  TestDefaultXDenom,
					sendTokenAmount: 2000,
					leverage:        sdk.NewDec(2),
				},
			},
			[]CloseLong{
				{
					id: s.app.PerpetualKeeper.GenerateTraderPositionId(
						suite.Address.String(),
						TestDefaultXDenom,
						s.app.PerpetualKeeper.GenerateVaultIdHash(TestDefaultXDenom, TestDefaultYDenom),
						types.PerpetualTradeType_PERPETUAL_LONG_POSITION,
						sdk.NewDec(2),
					),
					sendTokenAmount: 385,
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
				suite.AddTestCoinsToCustomAccount(sdk.NewInt(open_position.sendTokenAmount), open_position.sendTokenDenom, suite.Address)

				msg := types.NewMsgOpen(
					suite.Address.String(),
					types.PerpetualTradeType_PERPETUAL_LONG_POSITION,
					sdk.NewDec(2),
					s.app.PerpetualKeeper.GenerateVaultIdHash(TestDefaultXDenom, TestDefaultYDenom),
					sdk.NewInt(open_position.sendTokenAmount).String()+open_position.sendTokenDenom,
				)
				ctx := sdk.WrapSDKContext(suite.ctx)
				_, err := suite.app.PerpetualKeeper.Open(ctx, msg)
				suite.Require().NoError(err)

				position, f := s.app.PerpetualKeeper.GetPositionByPositionId(
					s.ctx,
					s.app.PerpetualKeeper.GenerateTraderPositionId(suite.Address.String(),
						open_position.sendTokenDenom,
						s.app.PerpetualKeeper.GenerateVaultIdHash(TestDefaultXDenom, TestDefaultYDenom),
						types.PerpetualTradeType_PERPETUAL_LONG_POSITION,
						sdk.NewDec(2),
					),
				)
				suite.Require().Equal(true, f)
				suite.Require().Greater(position.ReturnAmount.RoundInt().Int64(), int64(0))

				vault, _ := s.app.PerpetualKeeper.GetVaultByVaultId(s.ctx, s.GetNormalTestVault().VaultId)
				fmt.Printf("price in vAMM: %f\n", float64(vault.X.MulRaw(10000).Quo(vault.Y).Int64())/10000)
			}

			for _, close_position := range tc.close_long {

				old_balance := (s.app.BankKeeper.GetBalance(s.ctx, s.Address, TestDefaultXDenom)).Amount

				msg := types.NewMsgClose(
					suite.Address.String(),
					close_position.id,
					sdk.NewInt(close_position.sendTokenAmount),
				)
				ctx := sdk.WrapSDKContext(suite.ctx)
				_, err := suite.app.PerpetualKeeper.Close(ctx, msg)
				suite.Require().NoError(err)

				_, f := s.app.PerpetualKeeper.GetPositionByPositionId(
					s.ctx,
					s.app.PerpetualKeeper.GenerateTraderPositionId(suite.Address.String(),
						TestDefaultXDenom,
						s.app.PerpetualKeeper.GenerateVaultIdHash(TestDefaultXDenom, TestDefaultYDenom),
						types.PerpetualTradeType_PERPETUAL_LONG_POSITION,
						sdk.NewDec(2),
					),
				)
				suite.Require().Equal(false, f)

				new_balance := (s.app.BankKeeper.GetBalance(s.ctx, s.Address, TestDefaultXDenom)).Amount

				vault, _ := s.app.PerpetualKeeper.GetVaultByVaultId(s.ctx, s.GetNormalTestVault().VaultId)
				fmt.Printf("price in vAMM: %f\n", float64(vault.X.MulRaw(10000).Quo(vault.Y).Int64())/10000)

				fmt.Printf("difference between: %v", new_balance.Sub(old_balance).Int64())
			}
		})
	}
}
