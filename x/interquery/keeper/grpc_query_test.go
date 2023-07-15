package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	icqtypes "github.com/QuadrateOrg/core/x/interquery/types"
)

func (s *InterQueryKeeperTestSuite) TestQueries() {
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

	// set the query
	s.GetSimApp(s.chainA).InterQueryKeeper.SetQuery(s.chainA.GetContext(), *query)

	icqsrvSrv := icqtypes.QuerySrvrServer(s.GetSimApp(s.chainA).InterQueryKeeper)

	res, err := icqsrvSrv.Queries(sdk.WrapSDKContext(s.chainA.GetContext()), &icqtypes.QueryRequestsRequest{ChainId: s.chainB.ChainID})
	s.NoError(err)
	s.Len(res.Queries, 1)
	s.Equal(s.path.EndpointB.ConnectionID, res.Queries[0].ConnectionId)
	s.Equal(s.chainB.ChainID, res.Queries[0].ChainId)
	s.Equal("cosmos.staking.v1beta1.Query/Validators", res.Queries[0].QueryType)
	s.Equal(int64(200), res.Queries[0].Period)
	s.Equal("", res.Queries[0].CallbackId)
}
