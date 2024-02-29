package keeper_test

import (
	"fmt"
	"time"

	"github.com/QuadrateOrg/core/x/grow/types"
	stabletypes "github.com/QuadrateOrg/core/x/stable/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *GrowKeeperTestSuite) TestExecuteCreateLend() {
	testCases := []struct {
		name        string
		qStablePair stabletypes.Pair
		gTokenPair  types.GTokenPair
		Asset       []Asset_type
		lends       []Lend_type
		err         bool
		errString   string
	}{
		{
			"ok-create-lend",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			[]Asset_type{
				{
					Asset: s.GetNormalAsset(0),
					Price: "45000",
				},
			},
			[]Lend_type{
				{
					amount:      1 * 1000000,
					denom:       "uwbtc",
					OracleDenom: "WBTC",
				},
			},
			false,
			"",
		},
		{
			"ok-create-2-lend",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			[]Asset_type{
				{
					Asset: s.GetNormalAsset(0),
					Price: "45000",
				},
			},
			[]Lend_type{
				{
					amount:      1 * 1000000,
					denom:       "uwbtc",
					OracleDenom: "WBTC",
				},
				{
					amount:      5 * 100000,
					denom:       "uwbtc",
					OracleDenom: "WBTC",
				},
			},
			false,
			"",
		},
		{
			"false-lend asset not found",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			[]Asset_type{
				{
					Asset: s.GetWrongAsset(0),
					Price: "0",
				},
			},
			[]Lend_type{
				{
					amount:      1 * 1000000,
					denom:       "uwbtc",
					OracleDenom: "WBTC",
				},
			},
			true,
			"ErrAssetNotFound err",
		},
		{
			"false-oracle id not found",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			[]Asset_type{
				{
					Asset: s.GetWrongAsset(0),
					Price: "0",
				},
			},
			[]Lend_type{
				{
					amount:      1 * 1000000,
					denom:       "uosmo",
					OracleDenom: "OSMO",
				},
			},
			true,
			"ErrOracleAssetIdNotFound err",
		},
	}
	for _, tc := range testCases {
		suite.Setup()
		suite.Commit()
		suite.RegisterValidator()
		suite.app.GrowKeeper.ChangeDepositMethodStatus(s.ctx)
		suite.app.GrowKeeper.ChangeCollateralMethodStatus(s.ctx)
		suite.app.GrowKeeper.ChangeBorrowMethodStatus(s.ctx)
		s.ctx = s.ctx.WithBlockTime(time.Now())

		suite.app.StableKeeper.AppendPair(s.ctx, tc.qStablePair)
		suite.app.GrowKeeper.AppendPair(s.ctx, tc.gTokenPair)

		oracle_denom := ""
		for i, asset_type := range tc.Asset {
			suite.app.GrowKeeper.AppendAsset(s.ctx, asset_type.Asset)
			suite.SetupOracleKeeper(asset_type.Asset.OracleAssetId)
			oracle_denom += asset_type.Price + asset_type.Asset.OracleAssetId
			if i != len(tc.Asset)-1 {
				oracle_denom += ","
			}
		}
		suite.OracleAggregateExchangeRateFromInput(oracle_denom)

		for _, lend_type := range tc.lends {
			suite.AddTestCoinsToCustomAccount(sdk.NewInt(lend_type.amount), lend_type.denom, s.Address)
			suite.AddTestCoins(lend_type.amount, lend_type.denom)
		}

		suite.Run(fmt.Sprintf("Case---%s", tc.name), func() {
			for _, lend_type := range tc.lends {
				msg := types.NewMsgCreateLend(
					suite.Address.String(),
					sdk.NewInt(lend_type.amount).String()+lend_type.denom,
				)
				ctx := sdk.WrapSDKContext(suite.ctx)
				res, err := suite.app.GrowKeeper.CreateLend(ctx, msg)
				if !tc.err {
					suite.Require().NoError(err)

					position, found := s.app.GrowKeeper.GetPositionByPositionId(s.ctx, res.PositionId)
					suite.Require().Equal(true, found)
					suite.Require().Equal(res.PositionId, position.DepositId)

					if len(tc.lends) == 1 {
						lend, found := s.app.GrowKeeper.GetLendByLendId(s.ctx, s.app.GrowKeeper.CalculateLendId(s.Address.String(), lend_type.denom, position.DepositId))
						suite.Require().Equal(true, found)
						suite.Require().Equal(sdk.NewInt(lend_type.amount).String()+lend_type.denom, lend.AmountIn)
						suite.Require().Equal(sdk.NewInt(lend_type.amount), lend.AmountInAmount.RoundInt())
					} else {
						lend, found := s.app.GrowKeeper.GetLendByLendId(s.ctx, s.app.GrowKeeper.CalculateLendId(s.Address.String(), lend_type.denom, position.DepositId))
						suite.Require().Equal(true, found)
						suite.Require().GreaterOrEqual(lend.AmountInAmount.RoundInt().Int64(), lend_type.amount)
					}

				} else {
					suite.Require().Error(err, tc.errString)
				}
			}

			if !tc.err {
				position, found := s.app.GrowKeeper.GetPositionByPositionId(s.ctx, s.app.GrowKeeper.CalculateDepositId(s.Address.String()))
				suite.Require().Equal(true, found)
				LendAmountInUsd := 0
				ProvideValue := 0

				for _, lend_type := range tc.lends {
					price, _ := s.app.GrowKeeper.GetPriceByDenom(s.ctx, lend_type.OracleDenom)
					LendAmountInUsd += (int(lend_type.amount) * int(price.Int64())) / 10000
					ProvideValue += int(lend_type.amount)
				}
				suite.Require().Equal(LendAmountInUsd, int(position.LendAmountInUSD))

				asset, _ := s.app.GrowKeeper.GetAssetByAssetId(s.ctx, s.GetNormalAsset(0).AssetId)
				suite.Require().Equal(ProvideValue, int(asset.ProvideValue))
			}

		})
	}
}

func (suite *GrowKeeperTestSuite) TestExecuteWithdrawalLend() {
	testCases := []struct {
		name        string
		qStablePair stabletypes.Pair
		gTokenPair  types.GTokenPair
		Asset       []Asset_type
		lends       []Lend_type
		withdrawals []Withdrawal_type
		err         bool
		errString   string
	}{
		{
			"ok-withdrawl-part-lend",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			[]Asset_type{
				{
					Asset: s.GetNormalAsset(0),
					Price: "45000",
				},
			},
			[]Lend_type{
				{
					amount:      1 * 1000000,
					denom:       "uwbtc",
					OracleDenom: "WBTC",
				},
			},
			[]Withdrawal_type{
				{
					amount:      5 * 100000,
					denom:       "uwbtc",
					OracleDenom: "WBTC",
				},
			},
			false,
			"",
		},
		{
			"ok-withdrawl-all-lend",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			[]Asset_type{
				{
					Asset: s.GetNormalAsset(0),
					Price: "45000",
				},
			},
			[]Lend_type{
				{
					amount:      1 * 1000000,
					denom:       "uwbtc",
					OracleDenom: "WBTC",
				},
			},
			[]Withdrawal_type{
				{
					amount:      1 * 1000000,
					denom:       "uwbtc",
					OracleDenom: "WBTC",
				},
			},
			false,
			"",
		},
		{
			"false-lend asstet not found",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			[]Asset_type{
				{
					Asset: s.GetNormalAsset(0),
					Price: "45000",
				},
			},
			[]Lend_type{
				{
					amount:      1 * 1000000,
					denom:       "uwbtc",
					OracleDenom: "WBTC",
				},
			},
			[]Withdrawal_type{
				{
					amount:      5 * 100000,
					denom:       "uluna",
					OracleDenom: "LUNA",
				},
			},
			true,
			"ErrAssetNotFound err",
		},
	}

	for _, tc := range testCases {
		suite.Setup()
		suite.Commit()
		suite.RegisterValidator()
		suite.app.GrowKeeper.ChangeDepositMethodStatus(s.ctx)
		suite.app.GrowKeeper.ChangeCollateralMethodStatus(s.ctx)
		suite.app.GrowKeeper.ChangeBorrowMethodStatus(s.ctx)
		s.ctx = s.ctx.WithBlockTime(time.Now())

		suite.app.StableKeeper.AppendPair(s.ctx, tc.qStablePair)
		suite.app.GrowKeeper.AppendPair(s.ctx, tc.gTokenPair)

		oracle_denom := ""
		for i, asset_type := range tc.Asset {
			suite.app.GrowKeeper.AppendAsset(s.ctx, asset_type.Asset)
			suite.SetupOracleKeeper(asset_type.Asset.OracleAssetId)
			oracle_denom += asset_type.Price + asset_type.Asset.OracleAssetId
			if i != len(tc.Asset)-1 {
				oracle_denom += ","
			}
		}
		suite.OracleAggregateExchangeRateFromInput(oracle_denom)

		for _, lend_type := range tc.lends {
			suite.AddTestCoinsToCustomAccount(sdk.NewInt(lend_type.amount), lend_type.denom, s.Address)
			suite.AddTestCoins(lend_type.amount, lend_type.denom)
		}

		suite.Run(fmt.Sprintf("Case---%s", tc.name), func() {
			for _, lend_type := range tc.lends {
				msg := types.NewMsgCreateLend(
					suite.Address.String(),
					sdk.NewInt(lend_type.amount).String()+lend_type.denom,
				)
				ctx := sdk.WrapSDKContext(suite.ctx)
				res, err := suite.app.GrowKeeper.CreateLend(ctx, msg)
				suite.Require().NoError(err)

				position, found := s.app.GrowKeeper.GetPositionByPositionId(s.ctx, res.PositionId)
				suite.Require().Equal(true, found)
				suite.Require().Equal(res.PositionId, position.DepositId)

				if len(tc.lends) == 1 {
					lend, found := s.app.GrowKeeper.GetLendByLendId(s.ctx, s.app.GrowKeeper.CalculateLendId(s.Address.String(), lend_type.denom, position.DepositId))
					suite.Require().Equal(true, found)
					suite.Require().Equal(sdk.NewInt(lend_type.amount).String()+lend_type.denom, lend.AmountIn)
					suite.Require().Equal(sdk.NewInt(lend_type.amount), lend.AmountInAmount.RoundInt())
				} else {
					lend, found := s.app.GrowKeeper.GetLendByLendId(s.ctx, s.app.GrowKeeper.CalculateLendId(s.Address.String(), lend_type.denom, position.DepositId))
					suite.Require().Equal(true, found)
					suite.Require().GreaterOrEqual(lend.AmountInAmount.RoundInt().Int64(), lend_type.amount)
				}

			}

			LendAmountInUsd := 0
			ProvideValue := 0
			if !tc.err {
				position, found := s.app.GrowKeeper.GetPositionByPositionId(s.ctx, s.app.GrowKeeper.CalculateDepositId(s.Address.String()))
				suite.Require().Equal(true, found)

				for _, lend_type := range tc.lends {
					price, _ := s.app.GrowKeeper.GetPriceByDenom(s.ctx, lend_type.OracleDenom)
					LendAmountInUsd += (int(lend_type.amount) * int(price.Int64())) / 10000
					ProvideValue += int(lend_type.amount)
				}

				suite.Require().Equal(LendAmountInUsd, int(position.LendAmountInUSD))

			}

			for _, w_type := range tc.withdrawals {
				lend_id := s.app.GrowKeeper.CalculateLendId(s.Address.String(), w_type.denom, s.app.GrowKeeper.CalculateDepositId(s.Address.String()))
				old_lend, _ := s.app.GrowKeeper.GetLendByLendId(s.ctx, lend_id)

				msg := types.NewMsgWithdrawalLend(
					suite.Address.String(),
					sdk.NewInt(w_type.amount).String()+w_type.denom,
					w_type.denom,
				)
				ctx := sdk.WrapSDKContext(suite.ctx)
				res, err := suite.app.GrowKeeper.WithdrawalLend(ctx, msg)
				if !tc.err {
					suite.Require().NoError(err)
					suite.Require().Equal(sdk.NewInt(w_type.amount).String()+w_type.denom, res.AmountOut)

					if w_type.amount == old_lend.AmountInAmount.RoundInt().Int64() {
						_, found := s.app.GrowKeeper.GetLendByLendId(s.ctx, lend_id)
						suite.Require().Equal(false, found)
					} else {
						lend, found := s.app.GrowKeeper.GetLendByLendId(s.ctx, lend_id)
						suite.Require().Equal(true, found)
						suite.Require().Equal(old_lend.AmountInAmount.RoundInt().Sub(sdk.NewInt(w_type.amount)), lend.AmountInAmount.RoundInt())
					}

				} else {
					suite.Require().Error(err, tc.errString)
				}
			}

			if !tc.err {
				position, found := s.app.GrowKeeper.GetPositionByPositionId(s.ctx, s.app.GrowKeeper.CalculateDepositId(s.Address.String()))
				suite.Require().Equal(true, found)

				for _, w_type := range tc.withdrawals {
					price, _ := s.app.GrowKeeper.GetPriceByDenom(s.ctx, w_type.OracleDenom)
					LendAmountInUsd -= (int(w_type.amount) * int(price.Int64())) / 10000
					ProvideValue -= int(w_type.amount)
				}

				suite.Require().Equal(LendAmountInUsd, int(position.LendAmountInUSD))

				asset, _ := s.app.GrowKeeper.GetAssetByAssetId(s.ctx, s.GetNormalAsset(0).AssetId)
				suite.Require().Equal(ProvideValue, int(asset.ProvideValue))
			}
		})
	}
}

func (suite *GrowKeeperTestSuite) TestExecuteCreateLendWithDifferentToken() {
	testCases := []struct {
		name        string
		qStablePair stabletypes.Pair
		gTokenPair  types.GTokenPair
		Asset       []Asset_type
		lends       []Lend_type
		err         bool
		errString   string
	}{
		{
			"ok-create-2-lend-with-different-token",
			s.GetNormalQStablePair(0),
			s.GetNormalGTokenPair(0),
			[]Asset_type{
				{
					Asset: s.GetNormalAsset(0),
					Price: "45000",
				},
				{
					Asset: s.GetSecondStableNormalAsset(0),
					Price: "1",
				},
			},
			[]Lend_type{
				{
					amount:      1 * 1000000,
					denom:       "uwbtc",
					OracleDenom: "WBTC",
				},
				{
					amount:      5 * 100000,
					denom:       "uwbtc",
					OracleDenom: "WBTC",
				},
				{
					amount:      10000 * 1000000,
					denom:       "uusdc",
					OracleDenom: "USDC",
				},
			},
			false,
			"",
		},
	}
	for _, tc := range testCases {
		suite.Setup()
		suite.Commit()
		suite.RegisterValidator()
		suite.app.GrowKeeper.ChangeDepositMethodStatus(s.ctx)
		suite.app.GrowKeeper.ChangeCollateralMethodStatus(s.ctx)
		suite.app.GrowKeeper.ChangeBorrowMethodStatus(s.ctx)
		s.ctx = s.ctx.WithBlockTime(time.Now())

		suite.app.StableKeeper.AppendPair(s.ctx, tc.qStablePair)
		suite.app.GrowKeeper.AppendPair(s.ctx, tc.gTokenPair)

		oracle_denom := ""
		for i, asset_type := range tc.Asset {
			suite.app.GrowKeeper.AppendAsset(s.ctx, asset_type.Asset)
			suite.SetupOracleKeeper(asset_type.Asset.OracleAssetId)
			oracle_denom += asset_type.Price + asset_type.Asset.OracleAssetId
			if i != len(tc.Asset)-1 {
				oracle_denom += ","
			}
		}
		suite.OracleAggregateExchangeRateFromInput(oracle_denom)

		for _, lend_type := range tc.lends {
			suite.AddTestCoinsToCustomAccount(sdk.NewInt(lend_type.amount), lend_type.denom, s.Address)
			suite.AddTestCoins(lend_type.amount, lend_type.denom)
		}

		suite.Run(fmt.Sprintf("Case---%s", tc.name), func() {
			for _, lend_type := range tc.lends {
				msg := types.NewMsgCreateLend(
					suite.Address.String(),
					sdk.NewInt(lend_type.amount).String()+lend_type.denom,
				)
				ctx := sdk.WrapSDKContext(suite.ctx)
				res, err := suite.app.GrowKeeper.CreateLend(ctx, msg)
				if !tc.err {
					suite.Require().NoError(err)

					position, found := s.app.GrowKeeper.GetPositionByPositionId(s.ctx, res.PositionId)
					suite.Require().Equal(true, found)
					suite.Require().Equal(res.PositionId, position.DepositId)

					if len(tc.lends) == 1 {
						lend, found := s.app.GrowKeeper.GetLendByLendId(s.ctx, s.app.GrowKeeper.CalculateLendId(s.Address.String(), lend_type.denom, position.DepositId))
						suite.Require().Equal(true, found)
						suite.Require().Equal(sdk.NewInt(lend_type.amount).String()+lend_type.denom, lend.AmountIn)
						suite.Require().Equal(sdk.NewInt(lend_type.amount), lend.AmountInAmount.RoundInt())
					} else {
						lend, found := s.app.GrowKeeper.GetLendByLendId(s.ctx, s.app.GrowKeeper.CalculateLendId(s.Address.String(), lend_type.denom, position.DepositId))
						suite.Require().Equal(true, found)
						suite.Require().GreaterOrEqual(lend.AmountInAmount.RoundInt().Int64(), lend_type.amount)
					}

				} else {
					suite.Require().Error(err, tc.errString)
				}
			}

			if !tc.err {
				position, found := s.app.GrowKeeper.GetPositionByPositionId(s.ctx, s.app.GrowKeeper.CalculateDepositId(s.Address.String()))
				suite.Require().Equal(true, found)
				LendAmountInUsd := 0

				for _, lend_type := range tc.lends {
					price, _ := s.app.GrowKeeper.GetPriceByDenom(s.ctx, lend_type.OracleDenom)
					LendAmountInUsd += (int(lend_type.amount) * int(price.Int64())) / 10000
				}

				suite.Require().Equal(LendAmountInUsd, int(position.LendAmountInUSD))
			}

		})
	}
}
