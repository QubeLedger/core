package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	apptypes "github.com/QubeLedger/core/types"
	"github.com/QubeLedger/core/x/tokenfactory/types"
)

func TestGenesisState_Validate(t *testing.T) {
	apptypes.GetDefaultConfig()

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
				FactoryDenoms: []types.GenesisDenom{
					{
						Denom: "factory/qube18ffd5mke4f8lav3eq2xrd4n94sn0xj2kr0c7uk/bitcoin",
						AuthorityMetadata: types.DenomAuthorityMetadata{
							Admin: "qube18ffd5mke4f8lav3eq2xrd4n94sn0xj2kr0c7uk",
						},
					},
				},
			},
			valid: true,
		},
		{
			desc: "different admin from creator",
			genState: &types.GenesisState{
				FactoryDenoms: []types.GenesisDenom{
					{
						Denom: "factory/qube18ffd5mke4f8lav3eq2xrd4n94sn0xj2kr0c7uk/bitcoin",
						AuthorityMetadata: types.DenomAuthorityMetadata{
							Admin: "qube1l3pcj29m4dfuyfspa4dnct5myqtyhezlcalf54",
						},
					},
				},
			},
			valid: true,
		},
		{
			desc: "empty admin",
			genState: &types.GenesisState{
				FactoryDenoms: []types.GenesisDenom{
					{
						Denom: "factory/qube18ffd5mke4f8lav3eq2xrd4n94sn0xj2kr0c7uk/bitcoin",
						AuthorityMetadata: types.DenomAuthorityMetadata{
							Admin: "",
						},
					},
				},
			},
			valid: true,
		},
		{
			desc: "no admin",
			genState: &types.GenesisState{
				FactoryDenoms: []types.GenesisDenom{
					{
						Denom: "factory/qube18ffd5mke4f8lav3eq2xrd4n94sn0xj2kr0c7uk/bitcoin",
					},
				},
			},
			valid: true,
		},
		{
			desc: "invalid admin",
			genState: &types.GenesisState{
				FactoryDenoms: []types.GenesisDenom{
					{
						Denom: "factory/qube18ffd5mke4f8lav3eq2xrd4n94sn0xj2kr0c7uk/bitcoin",
						AuthorityMetadata: types.DenomAuthorityMetadata{
							Admin: "moose",
						},
					},
				},
			},
			valid: false,
		},
		{
			desc: "multiple denoms",
			genState: &types.GenesisState{
				FactoryDenoms: []types.GenesisDenom{
					{
						Denom: "factory/qube18ffd5mke4f8lav3eq2xrd4n94sn0xj2kr0c7uk/bitcoin",
						AuthorityMetadata: types.DenomAuthorityMetadata{
							Admin: "",
						},
					},
					{
						Denom: "factory/qube18ffd5mke4f8lav3eq2xrd4n94sn0xj2kr0c7uk/litecoin",
						AuthorityMetadata: types.DenomAuthorityMetadata{
							Admin: "",
						},
					},
				},
			},
			valid: true,
		},
		{
			desc: "duplicate denoms",
			genState: &types.GenesisState{
				FactoryDenoms: []types.GenesisDenom{
					{
						Denom: "factory/qube18ffd5mke4f8lav3eq2xrd4n94sn0xj2kr0c7uk/bitcoin",
						AuthorityMetadata: types.DenomAuthorityMetadata{
							Admin: "",
						},
					},
					{
						Denom: "factory/qube18ffd5mke4f8lav3eq2xrd4n94sn0xj2kr0c7uk/bitcoin",
						AuthorityMetadata: types.DenomAuthorityMetadata{
							Admin: "",
						},
					},
				},
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
