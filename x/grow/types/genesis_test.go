package types_test

import (
	"testing"

	"github.com/QuadrateOrg/core/x/grow/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
				GTokenPairList:            []types.GTokenPair{},
				RealRate:                  10,
				BorrowRate:                10,
				GrowStakingReserveAddress: "qube17ca7p2gvf6qcg0n6ucnkjpe3estscfdhaj9ep9",
				USQReserveAddress:         "qube17ca7p2gvf6qcg0n6ucnkjpe3estscfdhaj9ep9",
			},
			valid: true,
		},
		{
			desc: "invalid genesis state - zero Real Rate",
			genState: &types.GenesisState{
				GTokenPairList:            []types.GTokenPair{},
				RealRate:                  0,
				BorrowRate:                10,
				GrowStakingReserveAddress: "qube17ca7p2gvf6qcg0n6ucnkjpe3estscfdhaj9ep9",
				USQReserveAddress:         "qube17ca7p2gvf6qcg0n6ucnkjpe3estscfdhaj9ep9",
			},
			valid: false,
		},
		{
			desc: "invalid genesis state - zero Borrow Rate",
			genState: &types.GenesisState{
				GTokenPairList:            []types.GTokenPair{},
				RealRate:                  10,
				BorrowRate:                0,
				GrowStakingReserveAddress: "qube17ca7p2gvf6qcg0n6ucnkjpe3estscfdhaj9ep9",
				USQReserveAddress:         "qube17ca7p2gvf6qcg0n6ucnkjpe3estscfdhaj9ep9",
			},
			valid: false,
		},
		{
			desc: "invalid genesis state - wrong address",
			genState: &types.GenesisState{
				GTokenPairList:            []types.GTokenPair{},
				RealRate:                  10,
				BorrowRate:                10,
				GrowStakingReserveAddress: "",
				USQReserveAddress:         "qube17ca7p2gvf6qcg0n6ucnkjpe3estscfdhaj9ep9",
			},
			valid: false,
		},
		{
			desc: "invalid genesis state - wrong address",
			genState: &types.GenesisState{
				GTokenPairList:            []types.GTokenPair{},
				RealRate:                  10,
				BorrowRate:                10,
				GrowStakingReserveAddress: "qube17ca7p2gvf6qcg0n6ucnkjpe3estscfdhaj9ep9",
				USQReserveAddress:         "",
			},
			valid: false,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
