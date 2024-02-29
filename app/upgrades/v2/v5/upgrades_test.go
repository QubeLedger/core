package v5_test

import (
	"testing"
	"time"

	"github.com/QuadrateOrg/core/app"
	quadrateapptest "github.com/QuadrateOrg/core/app/helpers"
	apptypes "github.com/QuadrateOrg/core/types"
	growtypes "github.com/QuadrateOrg/core/x/grow/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/abci/types"
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

func (s *UpgradeTestSuite) runV_0_2_5_Upgrade() {
	s.ctx = s.ctx.WithBlockHeight(UpgradeHeight - 1)
	plan := upgradetypes.Plan{Name: "v0.2.5", Height: UpgradeHeight}
	err := s.App.UpgradeKeeper.ScheduleUpgrade(s.ctx, plan)
	s.Require().NoError(err)
	_, exists := s.App.UpgradeKeeper.GetUpgradePlan(s.ctx)
	s.Require().True(exists)

	s.ctx = s.ctx.WithBlockHeight(UpgradeHeight)
}

func (s *UpgradeTestSuite) TestUpgrade() {
	s.Setup()

	s.App.GrowKeeper.SetParams(s.ctx, growtypes.DefaultParams())

	s.App.GrowKeeper.SetAsset(s.ctx, growtypes.Asset{
		Id:      0,
		AssetId: "c5b4376538178084416cab617a2cace5a17db8ec762fed02a0ad35ef1a156e29",
		AssetMetadata: banktypes.Metadata{
			Description: "Wrapped Bitcoin (WBTC) is an ERC20 token backed 1:1 with Bitcoin. Completely transparent. 100% verifiable. Community led. (TESTNET)",
			DenomUnits: []*banktypes.DenomUnit{
				{Denom: "factory/qube1t2ydw7r4asmk74ymuvykjshdzun8dxye0az5wz/wbtc", Exponent: uint32(0), Aliases: []string{"wbtc"}},
			},
			Base:    "factory/qube1t2ydw7r4asmk74ymuvykjshdzun8dxye0az5wz/wbtc",
			Display: "wBTC",
			Name:    "Wrapped Bitcoin",
			Symbol:  "WBTC",
		},
		OracleAssetId: "BTC",
	})

	s.ctx = s.ctx.WithBlockTime(s.ctx.BlockTime().Add(time.Hour * 24))
	s.App.BeginBlocker(s.ctx, types.RequestBeginBlock{})

	asset, _ := s.App.GrowKeeper.GetAssetByAssetId(s.ctx, "c5b4376538178084416cab617a2cace5a17db8ec762fed02a0ad35ef1a156e29")
	s.Require().Equal("", asset.Type)
	s.Require().Equal(uint64(0), asset.CollectivelyBorrowValue)
	s.Require().Equal(uint64(0), asset.ProvideValue)

	params := s.App.GrowKeeper.GetParams(s.ctx)
	s.Require().Equal(uint64(0), params.LastTimeUpdateReserve)
	s.Require().Equal(uint64(0), params.UStaticStable)
	s.Require().Equal(uint64(0), params.UStaticVolatile)
	s.Require().Equal(uint64(0), params.MaxRateStable)
	s.Require().Equal(uint64(0), params.MaxRateVolatile)
	s.Require().Equal(uint64(0), params.Slope_1)
	s.Require().Equal(uint64(0), params.Slope_2)

	s.runV_0_2_5_Upgrade()
	s.ctx = s.ctx.WithBlockTime(s.ctx.BlockTime().Add(time.Hour * 24 * 7))
	s.App.BeginBlocker(s.ctx, types.RequestBeginBlock{})

	asset, _ = s.App.GrowKeeper.GetAssetByAssetId(s.ctx, "c5b4376538178084416cab617a2cace5a17db8ec762fed02a0ad35ef1a156e29")
	s.Require().Equal("volatile", asset.Type)

	params = s.App.GrowKeeper.GetParams(s.ctx)
	s.Require().Equal(uint64(s.ctx.BlockTime().Unix()), params.LastTimeUpdateReserve)
	s.Require().Equal(uint64(80), params.UStaticStable)
	s.Require().Equal(uint64(60), params.UStaticVolatile)
	s.Require().Equal(uint64(100), params.MaxRateStable)
	s.Require().Equal(uint64(300), params.MaxRateVolatile)
	s.Require().Equal(uint64(1), params.Slope_1)
	s.Require().Equal(uint64(8), params.Slope_2)

}
