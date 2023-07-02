package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/QuadrateOrg/core/x/interquery/keeper"
)

func (s *InterQueryKeeperTestSuite) TestEndBlocker() {
	qvr := stakingtypes.QueryValidatorsResponse{
		Validators: s.GetSimApp(s.chainB).StakingKeeper.GetBondedValidatorsByPower(s.chainB.GetContext()),
	}

	bondedQuery := stakingtypes.QueryValidatorsRequest{Status: stakingtypes.BondStatusBonded}
	bz, err := bondedQuery.Marshal()
	s.NoError(err)

	id := keeper.GenerateQueryHash(s.path.EndpointB.ConnectionID, s.chainB.ChainID, "cosmos.staking.v1beta1.Query/Validators", bz, "")

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

	// call end blocker
	s.GetSimApp(s.chainA).InterQueryKeeper.EndBlocker(s.chainA.GetContext())

	err = s.GetSimApp(s.chainA).InterQueryKeeper.SetDatapointForID(
		s.chainA.GetContext(),
		id,
		s.GetSimApp(s.chainB).AppCodec().MustMarshalJSON(&qvr),
		sdk.NewInt(s.chainB.CurrentHeader.Height),
	)
	s.NoError(err)

	dataPoint, err := s.GetSimApp(s.chainA).InterQueryKeeper.GetDatapointForID(s.chainA.GetContext(), id)
	s.NoError(err)
	s.NotNil(dataPoint)

	// set the query
	s.GetSimApp(s.chainA).InterQueryKeeper.DeleteQuery(s.chainA.GetContext(), id)

	// call end blocker
	s.GetSimApp(s.chainA).InterQueryKeeper.EndBlocker(s.chainA.GetContext())
}
