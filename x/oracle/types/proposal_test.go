package types_test

import (
	"testing"

	"github.com/QuadrateOrg/core/x/oracle/types"
	"github.com/stretchr/testify/require"
)

func TestRegisterRegisterAddNewDenomProposal(t *testing.T) {
	tests := []struct {
		title       string
		description string
		denom       string
		expectedErr bool
	}{
		{
			"test",
			"test",
			"OSMO",
			false,
		},
		{
			"test",
			"test",
			"",
			true,
		},
	}

	for _, tc := range tests {
		msg := types.NewRegisterAddNewDenomProposal(tc.title, tc.description, tc.denom)
		err := msg.ValidateBasic()
		if tc.expectedErr {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
	}
}
