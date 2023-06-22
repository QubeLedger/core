package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	quadrateapptest "github.com/QuadrateOrg/core/app/helpers"
	apptypes "github.com/QuadrateOrg/core/types"

	"github.com/QuadrateOrg/core/x/tokenfactory/types"
)

func TestGenesis(t *testing.T) {
	apptypes.SetConfig()

	genesisState := types.GenesisState{
		FactoryDenoms: []types.GenesisDenom{
			{
				Denom: "factory/qube18ffd5mke4f8lav3eq2xrd4n94sn0xj2kr0c7uk/bitcoin",
				AuthorityMetadata: types.DenomAuthorityMetadata{
					Admin: "qube18ffd5mke4f8lav3eq2xrd4n94sn0xj2kr0c7uk",
				},
			},
			{
				Denom: "factory/qube18ffd5mke4f8lav3eq2xrd4n94sn0xj2kr0c7uk/litecoin",
				AuthorityMetadata: types.DenomAuthorityMetadata{
					Admin: "qube18ffd5mke4f8lav3eq2xrd4n94sn0xj2kr0c7uk",
				},
			},
		},
	}
	app := quadrateapptest.Setup(t, "quadrate_5120-1", false, 1)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	app.TokenFactoryKeeper.InitGenesis(ctx, genesisState)
	exportedGenesis := app.TokenFactoryKeeper.ExportGenesis(ctx)
	require.NotNil(t, exportedGenesis)
	require.Equal(t, genesisState, *exportedGenesis)
}
