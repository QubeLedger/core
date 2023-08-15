package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	qubeapptest "github.com/QubeLedger/core/app/helpers"
	apptypes "github.com/QubeLedger/core/types"

	"github.com/QubeLedger/core/x/tokenfactory/types"
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
	app := qubeapptest.Setup(t, "Qube_5120-1", false, 1)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	app.TokenFactoryKeeper.InitGenesis(ctx, genesisState)
	exportedGenesis := app.TokenFactoryKeeper.ExportGenesis(ctx)
	require.NotNil(t, exportedGenesis)
	require.Equal(t, genesisState, *exportedGenesis)
}
