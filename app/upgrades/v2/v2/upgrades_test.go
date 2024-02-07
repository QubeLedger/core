package v2_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/QuadrateOrg/core/app"
	quadrateapptest "github.com/QuadrateOrg/core/app/helpers"
	apptypes "github.com/QuadrateOrg/core/types"
	stabletypes "github.com/QuadrateOrg/core/x/stable/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto"
)

const (
	UpgradeHeight = 15
)

type UpgradeTestSuite struct {
	suite.Suite
	App *app.QuadrateApp
	ctx sdk.Context
}

func (s *UpgradeTestSuite) SetupTest() {
	s.Setup()
}

var s *UpgradeTestSuite

func (s *UpgradeTestSuite) Setup() {
	apptypes.SetConfig()
	s.App = quadrateapptest.Setup(s.T(), "qube-1", false, 1)
	s.ctx = s.App.BaseApp.NewContext(false, s.ctx.BlockHeader())
	s.ctx = s.ctx.WithBlockHeight(0)
}

func TestUpgradeTestSuite(t *testing.T) {
	suite.Run(t, new(UpgradeTestSuite))
}

func (s *UpgradeTestSuite) runV_0_2_2_Upgrade() {
	s.ctx = s.ctx.WithBlockHeight(UpgradeHeight - 1)
	plan := upgradetypes.Plan{Name: "v0.2.2", Height: UpgradeHeight}
	err := s.App.UpgradeKeeper.ScheduleUpgrade(s.ctx, plan)
	s.Require().NoError(err)
	_, exists := s.App.UpgradeKeeper.GetUpgradePlan(s.ctx)
	s.Require().True(exists)

	s.ctx = s.ctx.WithBlockHeight(UpgradeHeight)
}

func (s *UpgradeTestSuite) TestUpgrade() {
	s.Setup()
	s.App.StableKeeper.SetPair(s.ctx, stabletypes.Pair{
		Id:     0,
		PairId: fmt.Sprintf("%x", crypto.Sha256(append([]byte("uatom"+"uusd")))),
		AmountInMetadata: banktypes.Metadata{
			Description: "",
			DenomUnits: []*banktypes.DenomUnit{
				{Denom: "uatom", Exponent: uint32(0), Aliases: []string{"microatom"}},
			},
			Base:    "uatom",
			Display: "ATOM",
			Name:    "ATOM",
			Symbol:  "ATOM",
		},
		AmountOutMetadata: banktypes.Metadata{
			Description: "",
			DenomUnits: []*banktypes.DenomUnit{
				{Denom: "uusd", Exponent: uint32(0), Aliases: []string{"microusd"}},
			},
			Base:    "uusd",
			Display: "usd",
			Name:    "USQ",
			Symbol:  "USQ",
		},
		Qm:           sdk.NewInt(0),
		Ar:           sdk.NewInt(0),
		MinAmountIn:  "",
		MinAmountOut: "20uusd",
	})

	s.ctx = s.ctx.WithBlockTime(s.ctx.BlockTime().Add(time.Hour * 24))
	s.App.BeginBlocker(s.ctx, types.RequestBeginBlock{})

	pair, _ := s.App.StableKeeper.GetPairByPairID(s.ctx, s.App.StableKeeper.GeneratePairIdHash("uatom", "uusd"))
	s.Require().Equal("", pair.Model)
	s.Require().Equal("", pair.MinAmountIn)

	s.runV_0_2_2_Upgrade()
	s.ctx = s.ctx.WithBlockTime(s.ctx.BlockTime().Add(time.Hour * 24 * 7))
	s.App.BeginBlocker(s.ctx, types.RequestBeginBlock{})

	pair, _ = s.App.StableKeeper.GetPairByPairID(s.ctx, s.App.StableKeeper.GeneratePairIdHash("uatom", "uusd"))
	s.Require().Equal("gmb", pair.Model)
	s.Require().Equal("20ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2", pair.MinAmountIn)
}
