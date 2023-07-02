package interquery_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	ibctesting "github.com/cosmos/ibc-go/v3/testing"
	"github.com/stretchr/testify/suite"

	"github.com/QuadrateOrg/core/app"
	apphelpers "github.com/QuadrateOrg/core/app/app_helpers"
	"github.com/QuadrateOrg/core/x/interquery"
	"github.com/QuadrateOrg/core/x/interquery/keeper"
	"github.com/QuadrateOrg/core/x/interquery/types"
)

func init() {
	ibctesting.DefaultTestingAppInit = apphelpers.SetupTestingApp
}

func TestInterQueryTestSuite(t *testing.T) {
	suite.Run(t, new(InterQueryTestSuite))
}

type InterQueryTestSuite struct {
	suite.Suite

	coordinator *ibctesting.Coordinator

	chainA *ibctesting.TestChain
	chainB *ibctesting.TestChain
	path   *ibctesting.Path
}

func (s *InterQueryTestSuite) GetSimApp(chain *ibctesting.TestChain) *app.QuadrateApp {
	qa, ok := chain.App.(*app.QuadrateApp)
	if !ok {
		panic("not quadrate app")
	}

	return qa
}

func (s *InterQueryTestSuite) SetupTest() {
	s.coordinator = ibctesting.NewCoordinator(s.T(), 2)
	s.chainA = s.coordinator.GetChain(ibctesting.GetChainID(1))
	s.chainB = s.coordinator.GetChain(ibctesting.GetChainID(2))

	s.path = newSimAppPath(s.chainA, s.chainB)
	s.coordinator.SetupConnections(s.path)
}

func (s *InterQueryTestSuite) TestInitGenesis() {
	bondedQuery := stakingtypes.QueryValidatorsRequest{Status: stakingtypes.BondStatusBonded}
	bz, err := bondedQuery.Marshal()
	s.NoError(err)

	query := s.GetSimApp(s.chainA).InterQueryKeeper.NewQuery(
		"",
		s.path.EndpointB.ConnectionID,
		s.chainB.ChainID,
		"cosmos.staking.v1beta1.Query/Validators",
		bz,
		sdk.NewInt(200),
		"",
		0,
	)

	interquery.InitGenesis(
		s.chainA.GetContext(),
		s.GetSimApp(s.chainA).InterQueryKeeper,
		types.GenesisState{
			Queries: []types.Query{
				*query,
			},
		},
	)

	id := keeper.GenerateQueryHash(s.path.EndpointB.ConnectionID, s.chainB.ChainID, "cosmos.staking.v1beta1.Query/Validators", bz, "")
	queryResponse, found := s.GetSimApp(s.chainA).InterQueryKeeper.GetQuery(s.chainA.GetContext(), id)
	s.True(found)
	s.Equal(s.path.EndpointB.ConnectionID, queryResponse.ConnectionId)
	s.Equal(s.chainB.ChainID, queryResponse.ChainId)
	s.Equal("cosmos.staking.v1beta1.Query/Validators", queryResponse.QueryType)
	s.Equal(sdk.NewInt(200), queryResponse.Period)
	s.Equal(uint64(0), queryResponse.Ttl)
	s.Equal("", queryResponse.CallbackId)
}

func newSimAppPath(chainA, chainB *ibctesting.TestChain) *ibctesting.Path {
	path := ibctesting.NewPath(chainA, chainB)
	path.EndpointA.ChannelConfig.PortID = ibctesting.TransferPort
	path.EndpointB.ChannelConfig.PortID = ibctesting.TransferPort

	return path
}
