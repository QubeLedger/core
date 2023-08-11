package keeper_test

import (
	"fmt"
	"testing"

	"github.com/QuadrateOrg/core/app"
	apptesting "github.com/QuadrateOrg/core/app/apptesting"
	quadrateapptest "github.com/QuadrateOrg/core/app/helpers"
	"github.com/QuadrateOrg/core/x/stable/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/crypto"
)

type StableKeeperTestSuite struct {
	suite.Suite
	ctx     sdk.Context
	app     *app.QuadrateApp
	genesis types.GenesisState
	Address sdk.AccAddress
}

var s *StableKeeperTestSuite

func (s *StableKeeperTestSuite) Setup() {
	s.app = quadrateapptest.Setup(s.T(), "qube-1", false, 1)
	s.Address = apptesting.CreateRandomAccounts(1)[0]
}

func TestStableKeeperTestSuite(t *testing.T) {
	s = new(StableKeeperTestSuite)
	suite.Run(t, s)
	// Run Ginkgo integration tests
	RegisterFailHandler(Fail)
	RunSpecs(t, "Keeper Suite")
}

func (suite *StableKeeperTestSuite) Commit() {
	header := suite.ctx.BlockHeader()
	// update ctx
	suite.ctx = suite.app.BaseApp.NewContext(false, header)
}

func (suite *StableKeeperTestSuite) MintStable(amount int64, pair types.Pair) error {
	suite.app.StableKeeper.UpdateAtomPriceTesting(suite.ctx, sdk.NewInt(95000))
	msg := types.NewMsgMint(
		suite.Address.String(),
		sdk.NewInt(amount).String()+pair.AmountInMetadata.Base,
		pair.AmountOutMetadata.Base,
	)
	ctx := sdk.WrapSDKContext(suite.ctx)
	_, err := suite.app.StableKeeper.Mint(ctx, msg)
	if err != nil {
		return err
	}
	return nil
}

func (s *StableKeeperTestSuite) GetNormalPair(id uint64) types.Pair {
	pair := types.Pair{
		Id:     id,
		PairId: fmt.Sprintf("%x", crypto.Sha256(append([]byte("uatom"+"uusd")))),
		AmountInMetadata: banktypes.Metadata{
			Description: "",
			DenomUnits: []*banktypes.DenomUnit{
				{Denom: "uatom", Exponent: uint32(0), Aliases: []string{"microatom"}},
			},
			Base:    "uatom",
			Display: "atom",
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
		MinAmountIn:  "20uatom",
		MinAmountOut: "20uuusd",
	}

	return pair
}

func (s *StableKeeperTestSuite) AddTestCoins(amount int64, denom string) {
	s.app.BankKeeper.MintCoins(s.ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin(denom, sdk.NewInt(amount))))
	s.app.BankKeeper.SendCoinsFromModuleToAccount(s.ctx, types.ModuleName, s.Address, sdk.NewCoins(sdk.NewCoin(denom, sdk.NewInt(amount))))
}

func (suite *StableKeeperTestSuite) IncreaseModuleBalance(amount int64, denom string) {
	suite.app.BankKeeper.MintCoins(suite.ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin(denom, sdk.NewInt(amount))))
}

func (s *StableKeeperTestSuite) AddTestCoinsToAccount(amount int64, denom string, acc sdk.AccAddress) {
	s.app.BankKeeper.MintCoins(s.ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin(denom, sdk.NewInt(amount))))
	s.app.BankKeeper.SendCoinsFromModuleToAccount(s.ctx, types.ModuleName, acc, sdk.NewCoins(sdk.NewCoin(denom, sdk.NewInt(amount))))
}
