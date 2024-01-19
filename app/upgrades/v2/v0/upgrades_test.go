package v0_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/QuadrateOrg/core/app/apptesting"
	growtypes "github.com/QuadrateOrg/core/x/grow/types"
	stabletypes "github.com/QuadrateOrg/core/x/stable/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto"
)

const (
	v6UpgradeHeight = 15
)

type UpgradeTestSuite struct {
	apptesting.KeeperTestHelper
}

func (s *UpgradeTestSuite) SetupTest() {
	s.Setup()
}

func TestUpgradeTestSuite(t *testing.T) {
	suite.Run(t, new(UpgradeTestSuite))
}

func (s *UpgradeTestSuite) runV6Upgrade() {
	s.Ctx = s.Ctx.WithBlockHeight(v6UpgradeHeight - 1)
	plan := upgradetypes.Plan{Name: "v0.2.0", Height: v6UpgradeHeight}
	err := s.App.UpgradeKeeper.ScheduleUpgrade(s.Ctx, plan)
	s.Require().NoError(err)
	_, exists := s.App.UpgradeKeeper.GetUpgradePlan(s.Ctx)
	s.Require().True(exists)

	s.Ctx = s.Ctx.WithBlockHeight(v6UpgradeHeight)
}

func (s *UpgradeTestSuite) TestUpgrade() {
	s.Setup()
	s.App.GrowKeeper.SetPair(s.Ctx, growtypes.GTokenPair{
		Id:            0,
		DenomID:       fmt.Sprintf("%x", crypto.Sha256(append([]byte("ugusd")))),
		QStablePairId: fmt.Sprintf("%x", crypto.Sha256(append([]byte("uatom"+"uusd")))),
		GTokenMetadata: banktypes.Metadata{
			Description: "",
			DenomUnits: []*banktypes.DenomUnit{
				{Denom: "ugusd", Exponent: uint32(0), Aliases: []string{"microgusd"}},
			},
			Base:    "ugusd",
			Display: "gUSQ",
			Name:    "gUSQ",
			Symbol:  "gUSQ",
		},
		MinAmountIn:                 "20uusd",
		MinAmountOut:                "20ugusd",
		GTokenLastPrice:             sdk.NewInt(0),
		GTokenLatestPriceUpdateTime: uint64(time.Now().Unix() - (31536000)),
	})

	s.Ctx = s.Ctx.WithBlockTime(s.Ctx.BlockTime().Add(time.Hour * 24))
	s.App.BeginBlocker(s.Ctx, types.RequestBeginBlock{})

	reserve_addr := s.App.StableKeeper.GetReserveFundAddress(s.Ctx)
	s.Require().Equal("", reserve_addr.String())

	pair, _ := s.App.GrowKeeper.GetPairByDenomID(s.Ctx, s.App.GrowKeeper.GenerateDenomIdHash("ugusd"))
	s.Require().Equal(int64(0), pair.GTokenLastPrice.Int64())

	s.runV6Upgrade()
	s.Ctx = s.Ctx.WithBlockTime(s.Ctx.BlockTime().Add(time.Hour * 24 * 7))
	s.App.BeginBlocker(s.Ctx, types.RequestBeginBlock{})

	stable_params := s.App.StableKeeper.GetParams(s.Ctx)
	s.Require().Equal(stabletypes.DefaultParams(), stable_params)
	fmt.Printf("\n%s\n", stable_params)

	grow_params := s.App.GrowKeeper.GetParams(s.Ctx)
	s.Require().Equal(growtypes.DefaultParams(), grow_params)
	fmt.Printf("%s\n", grow_params)

	pair, _ = s.App.GrowKeeper.GetPairByDenomID(s.Ctx, s.App.GrowKeeper.GenerateDenomIdHash("ugusd"))
	s.Require().Equal(int64(1*1000000), pair.GTokenLastPrice.Int64())

	reserve_addr = s.App.StableKeeper.GetReserveFundAddress(s.Ctx)
	default_reserve_addr, _ := sdk.AccAddressFromBech32(stabletypes.DefaultParams().ReserveFundAddress)
	s.Require().Equal(default_reserve_addr, reserve_addr)

	burn_addr := s.App.StableKeeper.GetBurningFundAddress(s.Ctx)
	default_burn_addr, _ := sdk.AccAddressFromBech32(stabletypes.DefaultParams().BurningFundAddress)
	s.Require().Equal(burn_addr, default_burn_addr)

	usq_reserve_addr := s.App.GrowKeeper.GetUSQReserveAddress(s.Ctx)
	default_usq_reserve_addr, _ := sdk.AccAddressFromBech32(growtypes.DefaultParams().USQReserveAddress)
	s.Require().Equal(default_usq_reserve_addr, usq_reserve_addr)

	grow_staking_reserve_addr := s.App.GrowKeeper.GetGrowStakingReserveAddress(s.Ctx)
	default_grow_staking_reserve_addr, _ := sdk.AccAddressFromBech32(growtypes.DefaultParams().GrowStakingReserveAddress)
	s.Require().Equal(default_grow_staking_reserve_addr, grow_staking_reserve_addr)

	grow_yield_reserve_addr := s.App.GrowKeeper.GetGrowYieldReserveAddress(s.Ctx)
	default_grow_yield_reserve_addr, _ := sdk.AccAddressFromBech32(growtypes.DefaultParams().GrowYieldReserveAddress)
	s.Require().Equal(default_grow_yield_reserve_addr, grow_yield_reserve_addr)
}
