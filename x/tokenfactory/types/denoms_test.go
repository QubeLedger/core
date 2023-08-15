package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	apptypes "github.com/QubeLedger/core/types"
	"github.com/QubeLedger/core/x/tokenfactory/types"
)

func TestDecomposeDenoms(t *testing.T) {
	apptypes.GetDefaultConfig()
	for _, tc := range []struct {
		desc  string
		denom string
		valid bool
	}{
		{
			desc:  "empty is invalid",
			denom: "",
			valid: false,
		},
		{
			desc:  "normal",
			denom: "factory/qube18ffd5mke4f8lav3eq2xrd4n94sn0xj2kr0c7uk/bitcoin",
			valid: true,
		},
		{
			desc:  "multiple slashes in subdenom",
			denom: "factory/qube18ffd5mke4f8lav3eq2xrd4n94sn0xj2kr0c7uk/bitcoin/1",
			valid: true,
		},
		{
			desc:  "no subdenom",
			denom: "factory/qube18ffd5mke4f8lav3eq2xrd4n94sn0xj2kr0c7uk/",
			valid: true,
		},
		{
			desc:  "incorrect prefix",
			denom: "ibc/qube18ffd5mke4f8lav3eq2xrd4n94sn0xj2kr0c7uk/bitcoin",
			valid: false,
		},
		{
			desc:  "subdenom of only slashes",
			denom: "factory/qube18ffd5mke4f8lav3eq2xrd4n94sn0xj2kr0c7uk/////",
			valid: true,
		},
		{
			desc:  "too long name",
			denom: "factory/qube18ffd5mke4f8lav3eq2xrd4n94sn0xj2kr0c7uk/adsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsf",
			valid: false,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			_, _, err := types.DeconstructDenom(tc.denom)
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
