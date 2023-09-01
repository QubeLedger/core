package keeper_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/QuadrateOrg/core/app"
	apptesting "github.com/QuadrateOrg/core/app/apptesting"
	quadrateapptest "github.com/QuadrateOrg/core/app/helpers"
	"github.com/QuadrateOrg/core/x/grow/types"
	"github.com/QuadrateOrg/core/x/oracle"
	oraclekeeper "github.com/QuadrateOrg/core/x/oracle/keeper"
	oracletypes "github.com/QuadrateOrg/core/x/oracle/types"
	stabletypes "github.com/QuadrateOrg/core/x/stable/types"
	"github.com/buger/jsonparser"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/crypto"
)

type GrowKeeperTestSuite struct {
	suite.Suite
	ctx        sdk.Context
	app        *app.QuadrateApp
	genesis    types.GenesisState
	Address    sdk.AccAddress
	ValPubKeys []cryptotypes.PubKey
}

var s *GrowKeeperTestSuite

func (s *GrowKeeperTestSuite) Setup() {
	s.app = quadrateapptest.Setup(s.T(), "qube-1", false, 1)
	s.Address = apptesting.CreateRandomAccounts(1)[0]
	s.ValPubKeys = simapp.CreateTestPubKeys(1)
}

func TestStableKeeperTestSuite(t *testing.T) {
	s = new(GrowKeeperTestSuite)
	suite.Run(t, s)
	// Run Ginkgo integration tests
	RegisterFailHandler(Fail)
	RunSpecs(t, "Keeper Suite")
}

func (suite *GrowKeeperTestSuite) Commit() {
	header := suite.ctx.BlockHeader()
	suite.ctx = suite.app.BaseApp.NewContext(false, header)
}

func (suite *GrowKeeperTestSuite) MintStable(amount int64, pair stabletypes.Pair) error {
	suite.app.StableKeeper.UpdateAtomPriceTesting(suite.ctx, sdk.NewInt(95000))
	msg := stabletypes.NewMsgMint(
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

func (s *GrowKeeperTestSuite) GetNormalQStablePair(id uint64) stabletypes.Pair {
	pair := stabletypes.Pair{
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
		MinAmountOut: "20uusd",
	}

	return pair
}

func (s *GrowKeeperTestSuite) GetNormalGTokenPair(id uint64) types.GTokenPair {
	pair := types.GTokenPair{
		Id:            id,
		DenomID:       fmt.Sprintf("%x", crypto.Sha256(append([]byte("ugusd")))),
		QStablePairId: fmt.Sprintf("%x", crypto.Sha256(append([]byte("uatom"+"uusd")))),
		GTokenMetadata: banktypes.Metadata{
			Description: "",
			DenomUnits: []*banktypes.DenomUnit{
				{Denom: "ugusd", Exponent: uint32(0), Aliases: []string{"microgusd"}},
			},
			Base:    "ugusd",
			Display: "gusd",
			Name:    "gUSQ",
			Symbol:  "gUSQ",
		},
		MinAmountIn:                 "20uusd",
		MinAmountOut:                "20ugusd",
		GTokenLastPrice:             sdk.NewInt(1 * 1000000),
		GTokenLatestPriceUpdateTime: uint64(time.Now().Unix() - (31536000)),
	}

	return pair
}

func (s *GrowKeeperTestSuite) GetNormalBorrowAsset(id uint64) types.BorrowAsset {
	ba := types.BorrowAsset{
		Id:            id,
		BorrowAssetId: fmt.Sprintf("%x", crypto.Sha256(append([]byte("uosmo"+"uusd")))),
		AmountInAssetMetadata: banktypes.Metadata{
			Description: "",
			DenomUnits: []*banktypes.DenomUnit{
				{Denom: "uosmo", Exponent: uint32(0), Aliases: []string{"microosmo"}},
			},
			Base:    "uosmo",
			Display: "osmo",
			Name:    "OSMO",
			Symbol:  "OSMO",
		},
		AmountOutAssetMetadata: banktypes.Metadata{
			Description: "",
			DenomUnits: []*banktypes.DenomUnit{
				{Denom: "uusd", Exponent: uint32(0), Aliases: []string{"microusd"}},
			},
			Base:    "uusd",
			Display: "usd",
			Name:    "USQ",
			Symbol:  "USQ",
		},
		OracleAssetId: "OSMO",
	}
	return ba
}

func (s *GrowKeeperTestSuite) AddTestCoins(amount int64, denom string) {
	s.app.BankKeeper.MintCoins(s.ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin(denom, sdk.NewInt(amount))))
	s.app.BankKeeper.SendCoinsFromModuleToAccount(s.ctx, types.ModuleName, s.Address, sdk.NewCoins(sdk.NewCoin(denom, sdk.NewInt(amount))))
}

func (suite *GrowKeeperTestSuite) IncreaseModuleBalance(amount int64, denom string) {
	suite.app.BankKeeper.MintCoins(suite.ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin(denom, sdk.NewInt(amount))))
}

func (s *GrowKeeperTestSuite) AddTestCoinsToCustomAccount(amount sdk.Int, denom string, acc sdk.AccAddress) {
	s.app.BankKeeper.MintCoins(s.ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin(denom, amount)))
	s.app.BankKeeper.SendCoinsFromModuleToAccount(s.ctx, types.ModuleName, acc, sdk.NewCoins(sdk.NewCoin(denom, amount)))
}

// NewTestMsgCreateValidator test msg creator
func NewTestMsgCreateValidator(address sdk.ValAddress, pubKey cryptotypes.PubKey, amt sdk.Int) *stakingtypes.MsgCreateValidator {
	commission := stakingtypes.NewCommissionRates(sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec())
	msg, _ := stakingtypes.NewMsgCreateValidator(
		address, pubKey, sdk.NewCoin("stake", amt),
		stakingtypes.Description{}, commission, sdk.OneInt(),
	)

	return msg
}

func (s *GrowKeeperTestSuite) RegisterValidator() error {
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

func (s *GrowKeeperTestSuite) PrevoteVotePrice(exchangeRatesStr string) error {
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

func GetTokensActualPrice() (string, error) {

	var atomPriceString string

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}

	res, err := client.Get("https://api.coinbase.com/v2/exchange-rates?currency=ATOM")
	if err != nil {
		return "", err
	}
	body, _ := ioutil.ReadAll(res.Body)

	if value, err := jsonparser.GetString(body, "data", "rates", "USD"); err == nil {
		atomPriceString = fmt.Sprintf("%v", value)
	} else {
		return "", err
	}

	return atomPriceString, nil
}

func (s *GrowKeeperTestSuite) OracleAggregateExchangeRateFromNet() {
	params := s.app.OracleKeeper.GetParams(s.ctx)
	price, err := GetTokensActualPrice()
	s.Require().NoError(err)
	err = s.PrevoteVotePrice(price + params.Whitelist[0].Name)
	s.Require().NoError(err)
}

func (s *GrowKeeperTestSuite) OracleAggregateExchangeRateFromInput(price string, denom string) {
	err := s.PrevoteVotePrice(price + denom)
	s.Require().NoError(err)
}

func (s *GrowKeeperTestSuite) SetupOracleKeeper(denom string) {
	params := s.app.OracleKeeper.GetParams(s.ctx)
	params.Whitelist = oracletypes.DenomList{
		{
			Name: denom,
		},
	}
	params.VotePeriod = 1
	params.SlashWindow = 100
	params.RewardDistributionWindow = 100
	s.app.OracleKeeper.SetParams(s.ctx, params)
}

func (s *GrowKeeperTestSuite) TestTime() {
	s.Setup()
	s.Commit()
	s.ctx = s.ctx.WithBlockTime(time.Now())
	fmt.Printf("%d\n", s.ctx.BlockTime().Unix())
}
