package keeper_test

import (
	"fmt"
	"time"

	dextypes "github.com/QuadrateOrg/core/x/dex/types"
	"github.com/QuadrateOrg/core/x/stable/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *StableKeeperTestSuite) TestDeltaMint() {
	testCases := []struct {
		name            string
		pair            types.Pair
		sendTokenDenom  string
		sendTokenAmount int64
		pricePerp       string
		priceStake      string
		priceSpot       string
		err             bool
		errString       string
	}{
		{
			"ok-delta: mint",
			s.GetNormalDeltaPair(0),
			"uatom",
			1000,
			"6.5",
			"0.8",
			"8.2",
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
			suite.app.StableKeeper.AppendPair(suite.ctx, tc.pair)
			suite.SetupOracleKeeper(tc.pair.OracleAssetId)
			suite.SetupOracleKeeper(tc.pair.PerpetualOracleAssetId)
			suite.OracleAggregateExchangeRateFromInput(tc.pricePerp + tc.pair.PerpetualOracleAssetId + "," + tc.priceSpot + tc.pair.OracleAssetId + "," + tc.priceStake + tc.pair.StakePriceOracleId)

			suite.app.PerpetualKeeper.AppendVault(s.ctx, *s.GetNormalTestVault(
				tc.pair.TokenStakeMetadata,
				tc.pair.TokenYMetadata,
				tc.pair.PerpetualOracleAssetId,
				100000,
			))

			suite.AddTestCoinsToCustomAccount(sdk.NewInt(100000*1000000), tc.pair.AmountInMetadata.Base, s.DexDepositAddress)
			suite.AddTestCoinsToCustomAccount((sdk.NewInt(100000 * 1000000).ToDec().Mul(sdk.MustNewDecFromStr(tc.priceStake))).RoundInt(), tc.pair.TokenStakeMetadata.Base, s.DexDepositAddress)

			_, err := suite.app.DexKeeper.DexDeposit(
				sdk.WrapSDKContext(s.ctx),
				dextypes.NewMsgDexDeposit(
					s.DexDepositAddress.String(),
					s.DexDepositAddress.String(),
					tc.pair.AmountInMetadata.Base,
					tc.pair.TokenStakeMetadata.Base,
					[]sdk.Int{
						sdk.NewInt(100000 * 1000000),
					},
					[]sdk.Int{
						(sdk.NewInt(100000 * 1000000).ToDec().Mul(sdk.MustNewDecFromStr(tc.priceStake))).RoundInt(),
					},
					[]int64{-2232},
					[]uint64{20},
					[]*dextypes.DepositOptions{
						{DisableAutoswap: false},
					},
				),
			)
			suite.Require().NoError(err)
			suite.AddTestCoins(tc.sendTokenAmount, tc.sendTokenDenom)

			old_balance := s.app.BankKeeper.GetBalance(s.ctx, suite.Address, tc.pair.AmountOutMetadata.Base)

			msg := types.NewMsgMint(
				suite.Address.String(),
				sdk.NewInt(tc.sendTokenAmount).String()+tc.sendTokenDenom,
				tc.pair.AmountOutMetadata.Base,
			)
			ctx := sdk.WrapSDKContext(suite.ctx)
			_, err = suite.app.StableKeeper.Mint(ctx, msg)
			if !tc.err {
				suite.Require().NoError(err, tc.name)
				new_balance := s.app.BankKeeper.GetBalance(s.ctx, suite.Address, tc.pair.AmountOutMetadata.Base)
				suite.Require().Equal(sdk.NewDec(tc.sendTokenAmount).Mul(sdk.MustNewDecFromStr(tc.priceSpot)), new_balance.Amount.Sub(old_balance.Amount).ToDec())
			} else {
				suite.Require().Error(err, tc.errString)
			}
		})
	}
}
