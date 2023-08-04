package cmd_test

import (
	"testing"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/stretchr/testify/require"

	app "github.com/QuadrateOrg/core/app"
	"github.com/QuadrateOrg/core/cmd/qubed/cmd"
)

func TestRootCmdConfig(t *testing.T) {
	rootCmd, _ := cmd.NewRootCmd()
	rootCmd.SetArgs([]string{
		"config",
		"keyring-backend",
		"test",
	})

	require.NoError(t, svrcmd.Execute(rootCmd, app.DefaultNodeHome))
}
