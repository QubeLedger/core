package grow_test

import (
	"fmt"
	"time"

	"github.com/QuadrateOrg/core/app/apptesting"
	grow "github.com/QuadrateOrg/core/x/grow"
	"github.com/QuadrateOrg/core/x/grow/types"
	"github.com/QuadrateOrg/core/x/stable/gmb"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (s *GrowAbciTestSuite) TestGrowPriceChangeWhenBlockEnd() {

	amt := 1000 * (types.Multiplier).Int64()

	{
		s.Setup()
		s.Commit()
		s.SetupOracleKeeper()
		s.RegisterValidator()
		s.app.StableKeeper.AppendPair(s.ctx, s.GetNormalQStablePair(0))
		s.app.GrowKeeper.AppendPair(s.ctx, s.GetNormalGTokenPair(0))
		s.app.GrowKeeper.SetGrowStakingReserveAddress(s.ctx, apptesting.CreateRandomAccounts(1)[0])
		s.app.GrowKeeper.SetUSQReserveAddress(s.ctx, apptesting.CreateRandomAccounts(1)[0])
		s.app.GrowKeeper.SetRealRate(s.ctx, sdk.NewInt(15))
		s.app.GrowKeeper.SetLastTimeUpdateReserve(s.ctx, sdk.NewInt(s.ctx.BlockTime().Unix()))
	}

	s.OracleAggregateExchangeRateFromNet()

	{
		s.AddTestCoins(amt, s.GetNormalQStablePair(0).AmountInMetadata.Base)
		err := s.MintStable(amt, s.GetNormalQStablePair(0))
		s.Require().NoError(err)

		msg := types.NewMsgDeposit(
			s.Address.String(),
			sdk.NewInt(amt).String()+s.GetNormalQStablePair(0).AmountOutMetadata.Base,
			s.GetNormalGTokenPair(0).GTokenMetadata.Base,
		)
		ctx := sdk.WrapSDKContext(s.ctx)
		res, err := s.app.GrowKeeper.Deposit(ctx, msg)
		s.Require().NoError(err)
		s.Require().Equal(res.AmountOut, sdk.NewCoin(s.GetNormalGTokenPair(0).GTokenMetadata.Base, sdk.NewInt(amt)).String())
	}

	s.ctx = s.ctx.WithBlockHeight(2)
	s.ctx = s.ctx.WithBlockTime(time.Unix((s.ctx.BlockTime().Unix() + 86350), 0))
	err := grow.EndBlocker(s.ctx, s.app.GrowKeeper)
	s.Require().NoError(err)

	price, err := s.app.GrowKeeper.GetGTokenPrice(s.ctx, s.GetNormalGTokenPair(0).DenomID)
	s.Require().NoError(err)
	s.Require().Equal(price.Int64(), int64(1000000))

	s.ctx = s.ctx.WithBlockHeight(3)
	s.ctx = s.ctx.WithBlockTime(time.Unix((s.ctx.BlockTime().Unix() + 86500), 0))
	err = grow.EndBlocker(s.ctx, s.app.GrowKeeper)
	s.Require().NoError(err)

	price, err = s.app.GrowKeeper.GetGTokenPrice(s.ctx, s.GetNormalGTokenPair(0).DenomID)
	s.Require().NoError(err)
	s.Require().Greater(price.Int64(), int64(1000000))

	s.ctx = s.ctx.WithBlockHeight(4)
	s.ctx = s.ctx.WithBlockTime(time.Unix((s.ctx.BlockTime().Unix() + 31536000), 0))
	err = grow.EndBlocker(s.ctx, s.app.GrowKeeper)
	s.Require().NoError(err)

	price, err = s.app.GrowKeeper.GetGTokenPrice(s.ctx, s.GetNormalGTokenPair(0).DenomID)
	s.Require().NoError(err)
	s.Require().Equal(price.Int64(), int64(1161321))
}

func (s *GrowAbciTestSuite) TestGrowReserveMath() {

	amt := 1000 * (types.Multiplier).Int64()

	{
		s.Setup()
		s.Commit()
		s.SetupOracleKeeper()
		s.RegisterValidator()
		s.app.StableKeeper.AppendPair(s.ctx, s.GetNormalQStablePair(0))
		s.app.GrowKeeper.AppendPair(s.ctx, s.GetNormalGTokenPair(0))
		s.app.GrowKeeper.SetGrowStakingReserveAddress(s.ctx, apptesting.CreateRandomAccounts(1)[0])
		s.app.GrowKeeper.SetUSQReserveAddress(s.ctx, apptesting.CreateRandomAccounts(1)[0])
		s.app.GrowKeeper.SetRealRate(s.ctx, sdk.NewInt(15))
		s.app.GrowKeeper.SetLastTimeUpdateReserve(s.ctx, sdk.NewInt(s.ctx.BlockTime().Unix()))
	}

	s.OracleAggregateExchangeRateFromNet()

	s.AddTestCoins(amt, s.GetNormalQStablePair(0).AmountInMetadata.Base)
	err := s.MintStable(amt, s.GetNormalQStablePair(0))
	s.Require().NoError(err)

	msg := types.NewMsgDeposit(
		s.Address.String(),
		sdk.NewInt(amt).String()+s.GetNormalQStablePair(0).AmountOutMetadata.Base,
		s.GetNormalGTokenPair(0).GTokenMetadata.Base,
	)
	ctx := sdk.WrapSDKContext(s.ctx)
	res, err := s.app.GrowKeeper.Deposit(ctx, msg)
	s.Require().NoError(err)
	s.Require().Equal(res.AmountOut, sdk.NewCoin(s.GetNormalGTokenPair(0).GTokenMetadata.Base, sdk.NewInt(amt)).String())

	balanceUSQStakingReserveAddress := s.app.BankKeeper.GetBalance(s.ctx, s.app.GrowKeeper.GetGrowStakingReserveAddress(s.ctx), s.GetNormalQStablePair(0).AmountOutMetadata.Base)
	updatedPair, _ := s.app.GrowKeeper.GetPairByDenomID(s.ctx, s.GetNormalGTokenPair(0).DenomID)
	s.Require().Equal(balanceUSQStakingReserveAddress.Amount, updatedPair.St)

	updatedqStablePair, _ := s.app.StableKeeper.GetPairByPairID(s.ctx, s.GetNormalQStablePair(0).PairId)

	atomPrice, _ := s.app.OracleKeeper.GetExchangeRate(s.ctx, updatedqStablePair.AmountInMetadata.Base)
	br, _ := gmb.CalculateBackingRatio(atomPrice.MulInt64(10000).RoundInt(), updatedqStablePair.Ar, updatedqStablePair.Qm)

	qm := updatedqStablePair.Qm

	gy, err := s.app.GrowKeeper.CalculateGrowYield(s.ctx, updatedPair)
	s.Require().NoError(err)

	ry, err := s.app.GrowKeeper.CalculateRealYield(s.ctx, updatedPair)
	s.Require().NoError(err)

	fmt.Printf("qm: %d\nst: %d\nBackingRatio: %d\nGrowYield: %d\nRealYield: %d\n", qm.Int64()/1000000, updatedPair.St.Int64()/1000000, br.Int64(), gy.Int64()/1000000, ry.Int64()/1000000)

	action, value, err := s.app.GrowKeeper.CheckYieldRate(s.ctx, updatedPair)
	s.Require().NoError(err)
	fmt.Printf("Action: %s\nDiff between RealYield and GrowYield: %d\n", action, value.Int64()/1000000)

	_, found := s.app.GrowKeeper.CalculateAddToReserveValue(s.ctx, value, updatedPair)
	s.Require().Equal(found, false)

	s.ctx = s.ctx.WithBlockHeight(2)
	s.ctx = s.ctx.WithBlockTime(time.Unix((s.ctx.BlockTime().Unix() + 10), 0))

	realValue, found := s.app.GrowKeeper.CalculateAddToReserveValue(s.ctx, value, updatedPair)
	s.Require().Equal(found, true)

	fmt.Printf("Real send to/from reserve: %f\n", float64(realValue.Int64())/1000000)
}

func (s *GrowAbciTestSuite) TestGrowIncreaseUSQReserve() {

	amt := 1000 * (types.Multiplier).Int64()

	{
		s.Setup()
		s.Commit()
		s.SetupOracleKeeper()
		s.RegisterValidator()
		s.app.StableKeeper.AppendPair(s.ctx, s.GetNormalQStablePair(0))
		s.app.GrowKeeper.AppendPair(s.ctx, s.GetNormalGTokenPair(0))
		s.app.GrowKeeper.SetGrowStakingReserveAddress(s.ctx, apptesting.CreateRandomAccounts(1)[0])
		s.app.GrowKeeper.SetUSQReserveAddress(s.ctx, apptesting.CreateRandomAccounts(1)[0])
		s.app.GrowKeeper.SetRealRate(s.ctx, sdk.NewInt(15))
		s.app.GrowKeeper.SetLastTimeUpdateReserve(s.ctx, sdk.NewInt(s.ctx.BlockTime().Unix()))
	}

	s.OracleAggregateExchangeRateFromNet()

	s.AddTestCoins(amt, s.GetNormalQStablePair(0).AmountInMetadata.Base)
	err := s.MintStable(amt, s.GetNormalQStablePair(0))
	s.Require().NoError(err)

	msg := types.NewMsgDeposit(
		s.Address.String(),
		sdk.NewInt(amt).String()+s.GetNormalQStablePair(0).AmountOutMetadata.Base,
		s.GetNormalGTokenPair(0).GTokenMetadata.Base,
	)
	ctx := sdk.WrapSDKContext(s.ctx)
	res, err := s.app.GrowKeeper.Deposit(ctx, msg)
	s.Require().NoError(err)
	s.Require().Equal(res.AmountOut, sdk.NewCoin(s.GetNormalGTokenPair(0).GTokenMetadata.Base, sdk.NewInt(amt)).String())

	balanceGrowStakingReserveOld := s.app.BankKeeper.GetBalance(s.ctx, s.app.GrowKeeper.GetGrowStakingReserveAddress(s.ctx), s.GetNormalQStablePair(0).AmountOutMetadata.Base)
	balanceUSQReserveOld := s.app.BankKeeper.GetBalance(s.ctx, s.app.GrowKeeper.GetUSQReserveAddress(s.ctx), s.GetNormalQStablePair(0).AmountOutMetadata.Base)
	updatedPair, _ := s.app.GrowKeeper.GetPairByDenomID(s.ctx, s.GetNormalGTokenPair(0).DenomID)
	s.Require().Equal(balanceGrowStakingReserveOld.Amount, updatedPair.St)

	s.NewBlock_IncreaseBlockTime10Sec()

	err = grow.EndBlocker(s.ctx, s.app.GrowKeeper)
	s.Require().NoError(err)

	action, value, err := s.app.GrowKeeper.CheckYieldRate(s.ctx, updatedPair)
	s.Require().NoError(err)

	realValue, found := s.app.GrowKeeper.CalculateAddToReserveValue(s.ctx, value, updatedPair)
	s.Require().Equal(found, true)

	if action == types.SendToReserveAction {
		balanceGrowStakingReserve := s.app.BankKeeper.GetBalance(s.ctx, s.app.GrowKeeper.GetGrowStakingReserveAddress(s.ctx), s.GetNormalQStablePair(0).AmountOutMetadata.Base)
		s.Require().Equal((balanceGrowStakingReserveOld.Amount).Sub(balanceGrowStakingReserve.Amount), realValue)

		balanceUSQReserve := s.app.BankKeeper.GetBalance(s.ctx, s.app.GrowKeeper.GetUSQReserveAddress(s.ctx), s.GetNormalQStablePair(0).AmountOutMetadata.Base)
		s.Require().Equal((balanceUSQReserve.Amount).Sub(balanceUSQReserveOld.Amount), realValue)
	}

	if action == types.SendFromReserveAction {
		balanceGrowStakingReserve := s.app.BankKeeper.GetBalance(s.ctx, s.app.GrowKeeper.GetGrowStakingReserveAddress(s.ctx), s.GetNormalQStablePair(0).AmountOutMetadata.Base)
		s.Require().Equal((balanceGrowStakingReserve.Amount).Sub(balanceGrowStakingReserveOld.Amount), realValue)

		balanceUSQReserve := s.app.BankKeeper.GetBalance(s.ctx, s.app.GrowKeeper.GetUSQReserveAddress(s.ctx), s.GetNormalQStablePair(0).AmountOutMetadata.Base)
		s.Require().Equal((balanceUSQReserveOld.Amount).Sub(balanceUSQReserve.Amount), realValue)
	}

}
