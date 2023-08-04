package types_test

import (
	"testing"

	"github.com/QuadrateOrg/core/x/stable/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/require"
)

func TestChangeBaseTokenDenomProposal(t *testing.T) {
	tests := []struct {
		title       string
		description string
		metadata    banktypes.Metadata
		expectedErr bool
	}{
		{
			"test",
			"test",
			banktypes.Metadata{
				Name:        "Cosmos Hub Atom",
				Symbol:      "ATOM",
				Description: "The native staking token of the Cosmos Hub.",
				DenomUnits: []*banktypes.DenomUnit{
					{"uatom", uint32(0), []string{"microatom"}},
				},
				Base:    "uatom",
				Display: "atom",
			},
			false,
		},
		{
			"test",
			"test",
			banktypes.Metadata{
				Name: "",
			},
			true,
		},
		{
			"test",
			"test",
			banktypes.Metadata{
				Name:   "Cosmos Hub Atom",
				Symbol: "",
			},
			true,
		},
		{
			"test",
			"test",
			banktypes.Metadata{
				Name:   "Cosmos Hub Atom",
				Symbol: "ATOM",
				Base:   "",
			},
			true,
		},
		{
			"test",
			"test",
			banktypes.Metadata{
				Name:    "Cosmos Hub Atom",
				Symbol:  "ATOM",
				Base:    "uatom",
				Display: "",
			},
			true,
		},
	}

	for _, tc := range tests {
		msg := types.NewRegisterChangeBaseTokenDenomProposal(tc.title, tc.description, tc.metadata)
		err := msg.ValidateBasic()
		if tc.expectedErr {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
	}
}

func TestChangeSendTokenDenomProposal(t *testing.T) {
	tests := []struct {
		title       string
		description string
		metadata    banktypes.Metadata
		expectedErr bool
	}{
		{
			"test",
			"test",
			banktypes.Metadata{
				Name:        "Cosmos Hub Atom",
				Symbol:      "ATOM",
				Description: "The native staking token of the Cosmos Hub.",
				DenomUnits: []*banktypes.DenomUnit{
					{"uatom", uint32(0), []string{"microatom"}},
				},
				Base:    "uatom",
				Display: "atom",
			},
			false,
		},
		{
			"test",
			"test",
			banktypes.Metadata{
				Name: "",
			},
			true,
		},
		{
			"test",
			"test",
			banktypes.Metadata{
				Name:   "Cosmos Hub Atom",
				Symbol: "",
			},
			true,
		},
		{
			"test",
			"test",
			banktypes.Metadata{
				Name:   "Cosmos Hub Atom",
				Symbol: "ATOM",
				Base:   "",
			},
			true,
		},
		{
			"test",
			"test",
			banktypes.Metadata{
				Name:    "Cosmos Hub Atom",
				Symbol:  "ATOM",
				Base:    "uatom",
				Display: "",
			},
			true,
		},
	}

	for _, tc := range tests {
		msg := types.NewRegisterChangeSendTokenDenomProposal(tc.title, tc.description, tc.metadata)
		err := msg.ValidateBasic()
		if tc.expectedErr {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
	}
}
