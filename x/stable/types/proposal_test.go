package types_test

import (
	"testing"

	"github.com/QuadrateOrg/core/x/stable/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/require"
)

func TestRegisterPairProposal(t *testing.T) {
	tests := []struct {
		title             string
		description       string
		amountInMetadata  banktypes.Metadata
		amountOutMetadata banktypes.Metadata
		minAmountInt      string
		minAmountOut      string
		expectedErr       bool
	}{
		{
			"test",
			"test",
			banktypes.Metadata{
				Description: "",
				DenomUnits: []*banktypes.DenomUnit{
					{Denom: "uatom", Exponent: uint32(0), Aliases: []string{"microatom"}},
				},
				Base:    "uatom",
				Display: "atom",
				Name:    "ATOM",
				Symbol:  "ATOM",
			},
			banktypes.Metadata{
				Description: "",
				DenomUnits: []*banktypes.DenomUnit{
					{Denom: "uusd", Exponent: uint32(0), Aliases: []string{"microusd"}},
				},
				Base:    "uusd",
				Display: "usd",
				Name:    "USQ",
				Symbol:  "USQ",
			},
			"20uatom",
			"20uusd",
			false,
		},
		{
			"test",
			"test",
			banktypes.Metadata{
				Name: "",
			},
			banktypes.Metadata{
				Name: "",
			},
			"",
			"",
			true,
		},
		{
			"test",
			"test",
			banktypes.Metadata{
				Name:   "Cosmos Hub Atom",
				Symbol: "",
			},
			banktypes.Metadata{
				Name:   "Cosmos Hub Atom",
				Symbol: "",
			},
			"",
			"",
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
			banktypes.Metadata{
				Name:   "Cosmos Hub Atom",
				Symbol: "ATOM",
				Base:   "",
			},
			"",
			"",
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
			banktypes.Metadata{
				Name:    "Cosmos Hub Atom",
				Symbol:  "ATOM",
				Base:    "uatom",
				Display: "",
			},
			"",
			"",
			true,
		},
	}

	for _, tc := range tests {
		msg := types.NewRegisterPairProposal(tc.title, tc.description, tc.amountInMetadata, tc.amountOutMetadata, tc.minAmountInt, tc.minAmountOut)
		err := msg.ValidateBasic()
		if tc.expectedErr {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
	}
}

func TestRegisterChangeBurningFundAddressProposal(t *testing.T) {
	tests := []struct {
		title       string
		description string
		address     string
		expectedErr bool
	}{
		{
			"test",
			"test",
			"qube17ca7p2gvf6qcg0n6ucnkjpe3estscfdhaj9ep9",
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
		msg := types.NewRegisterChangeBurningFundAddressProposal(tc.title, tc.description, tc.address)
		err := msg.ValidateBasic()
		if tc.expectedErr {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
	}
}

func TestRegisterChangeReserveFundAddressProposal(t *testing.T) {
	tests := []struct {
		title       string
		description string
		address     string
		expectedErr bool
	}{
		{
			"test",
			"test",
			"qube17ca7p2gvf6qcg0n6ucnkjpe3estscfdhaj9ep9",
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
		msg := types.NewRegisterChangeReserveFundAddressProposal(tc.title, tc.description, tc.address)
		err := msg.ValidateBasic()
		if tc.expectedErr {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
	}
}
