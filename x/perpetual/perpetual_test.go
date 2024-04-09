package perpetual_test

import (
	"testing"

	"github.com/QuadrateOrg/core/app"
	apptesting "github.com/QuadrateOrg/core/app/apptesting"
	quadrateapptest "github.com/QuadrateOrg/core/app/helpers"
	"github.com/QuadrateOrg/core/app/simapp"
	apptypes "github.com/QuadrateOrg/core/types"
	"github.com/QuadrateOrg/core/x/oracle"
	oraclekeeper "github.com/QuadrateOrg/core/x/oracle/keeper"
	oracletypes "github.com/QuadrateOrg/core/x/oracle/types"
	"github.com/QuadrateOrg/core/x/perpetual/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/suite"
)

type PerpetualTestSuite struct {
	suite.Suite
	ctx        sdk.Context
	app        *app.QuadrateApp
	genesis    types.GenesisState
	Address    sdk.AccAddress
	ValPubKeys []cryptotypes.PubKey
}

var s *PerpetualTestSuite

const (
	TestDefaultXDenom        = "uatom"
	TestDefaultYDenom        = "uusq"
	TestDefaultOracleAssetId = "USQ/ATOM"
	TestInitXLiq             = 100000
)

func (s *PerpetualTestSuite) Setup() {
	apptypes.SetConfig()
	s.app = quadrateapptest.Setup(s.T(), "qube-1", false, 1)
	s.Address = apptesting.CreateRandomAccounts(1)[0]
	s.ValPubKeys = simapp.CreateTestPubKeys(1)
	s.Commit()
}

func TestPerpetualKeeperTestSuite(t *testing.T) {
	s = new(PerpetualTestSuite)
	suite.Run(t, s)
	// Run Ginkgo integration tests
	RegisterFailHandler(Fail)
	RunSpecs(t, "Keeper Suite")
}

func (suite *PerpetualTestSuite) Commit() {
	header := suite.ctx.BlockHeader()
	suite.ctx = suite.app.BaseApp.NewContext(false, header)
}

func (suite *PerpetualTestSuite) GetNormalTestVault() *types.Vault {
	price, _ := s.app.OracleKeeper.GetExchangeRate(s.ctx, TestDefaultOracleAssetId)
	return &types.Vault{
		VaultId: s.app.PerpetualKeeper.GenerateVaultIdHash(TestDefaultXDenom, TestDefaultYDenom),
		AmountXMetadata: banktypes.Metadata{
			Description: "",
			DenomUnits: []*banktypes.DenomUnit{
				{Denom: "uatom", Exponent: uint32(0), Aliases: []string{"microatom"}},
			},
			Base:    "uatom",
			Display: "ATOM",
			Name:    "ATOM",
			Symbol:  "ATOM",
		},
		AmountYMetadata: banktypes.Metadata{
			Description: "",
			DenomUnits: []*banktypes.DenomUnit{
				{Denom: "uusq", Exponent: uint32(0), Aliases: []string{"microuusq"}},
			},
			Base:    "uusq",
			Display: "uusq",
			Name:    "USQ",
			Symbol:  "USQ",
		},
		X:               sdk.NewInt(TestInitXLiq),
		Y:               (sdk.NewDec(TestInitXLiq).Mul(price)).RoundInt(),
		K:               sdk.NewInt(TestInitXLiq).Mul((sdk.NewDec(TestInitXLiq).Mul(price)).RoundInt()),
		OracleAssetId:   TestDefaultOracleAssetId,
		LongPositionId:  []string{},
		ShortPositionId: []string{},
	}
}

func (s *PerpetualTestSuite) SetupOracleKeeper(denom string) {
	params := s.app.OracleKeeper.GetParams(s.ctx)
	params.Whitelist = append(params.Whitelist, oracletypes.Denom{Name: denom})
	params.VotePeriod = 1
	params.SlashWindow = 100
	params.RewardDistributionWindow = 100
	s.app.OracleKeeper.SetParams(s.ctx, params)
}

func (s *PerpetualTestSuite) PrevoteVotePrice(exchangeRatesStr string) error {
	salt := "1"
	hash := oracletypes.GetAggregateVoteHash(salt, exchangeRatesStr, sdk.ValAddress(s.ValPubKeys[0].Address()))

	msgServer := oraclekeeper.NewMsgServerImpl(s.app.OracleKeeper)

	s.ctx = s.ctx.WithBlockHeight(0)
	aggregateExchangeRatePrevoteMsg := oracletypes.NewMsgAggregateExchangeRatePrevote(hash, sdk.AccAddress(s.ValPubKeys[0].Address()), sdk.ValAddress(s.ValPubKeys[0].Address()))
	_, err := msgServer.AggregateExchangeRatePrevote(sdk.WrapSDKContext(s.ctx), aggregateExchangeRatePrevoteMsg)
	if err != nil {
		return err
	}
	s.ctx = s.ctx.WithBlockHeight(1)
	aggregateExchangeRateVoteMsg := oracletypes.NewMsgAggregateExchangeRateVote(salt, exchangeRatesStr, sdk.AccAddress(s.ValPubKeys[0].Address()), sdk.ValAddress(s.ValPubKeys[0].Address()))
	_, err = msgServer.AggregateExchangeRateVote(sdk.WrapSDKContext(s.ctx), aggregateExchangeRateVoteMsg)
	if err != nil {
		return err
	}

	s.ctx = s.ctx.WithBlockHeight(2)
	oracle.EndBlocker(s.ctx, s.app.OracleKeeper)

	return nil
}

func (s *PerpetualTestSuite) AddTestCoinsToCustomAccount(amount sdk.Int, denom string, acc sdk.AccAddress) {
	s.app.BankKeeper.MintCoins(s.ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin(denom, amount)))
	s.app.BankKeeper.SendCoinsFromModuleToAccount(s.ctx, types.ModuleName, acc, sdk.NewCoins(sdk.NewCoin(denom, amount)))
}

// NewTestMsgCreateValidator test msg creator
func NewTestMsgCreateValidator(address sdk.ValAddress, pubKey cryptotypes.PubKey, amt sdk.Int) *stakingtypes.MsgCreateValidator {
	commission := stakingtypes.NewCommissionRates(sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec())
	msg, _ := stakingtypes.NewMsgCreateValidator(
		address, pubKey, sdk.NewCoin("stake", amt),
		stakingtypes.Description{}, commission,
	)

	return msg
}

func (s *PerpetualTestSuite) RegisterValidator() error {
	for _, vp := range s.ValPubKeys {
		s.AddTestCoinsToCustomAccount(sdk.TokensFromConsensusPower(10, sdk.DefaultPowerReduction), "stake", sdk.AccAddress(vp.Address()))
	}
	sh := stakingkeeper.NewMsgServerImpl(s.app.StakingKeeper)
	stakingAmt := sdk.TokensFromConsensusPower(10, sdk.DefaultPowerReduction)
	for _, vp := range s.ValPubKeys {
		_, err := sh.CreateValidator(sdk.WrapSDKContext(s.ctx), NewTestMsgCreateValidator(sdk.ValAddress(vp.Address()), vp, stakingAmt))
		if err != nil {
			return err
		}
	}
	staking.EndBlocker(s.ctx, s.app.StakingKeeper)
	return nil
}

func (s *PerpetualTestSuite) OracleAggregateExchangeRateFromInput(denom string) {
	err := s.PrevoteVotePrice(denom)
	s.Require().NoError(err)
}
