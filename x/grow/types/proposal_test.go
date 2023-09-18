package types_test

import (
	"testing"

	apptypes "github.com/QuadrateOrg/core/types"
	"github.com/QuadrateOrg/core/x/grow/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/require"
)

func TestRegisterLendAssetProposal(t *testing.T) {
	apptypes.SetConfig()
	tests := []struct {
		title         string
		description   string
		assetMetadata banktypes.Metadata
		OracleAssetId string
		expectedErr   bool
	}{
		{
			"test",
			"test",
			banktypes.Metadata{
				Description: "Grow USQ",
				DenomUnits: []*banktypes.DenomUnit{
					{Denom: "ugusd", Exponent: uint32(0), Aliases: []string{"microusd"}},
				},
				Base:    "ugusd",
				Display: "gusd",
				Name:    "gUSQ",
				Symbol:  "gUSQ",
			},
			"ATOM",
			false,
		},
		{
			"test",
			"test",
			banktypes.Metadata{
				Description: "Grow USQ",
				DenomUnits: []*banktypes.DenomUnit{
					{Denom: "ugusd", Exponent: uint32(0), Aliases: []string{"microusd"}},
				},
				Base:    "ugusd",
				Display: "gusd",
				Name:    "gUSQ",
				Symbol:  "gUSQ",
			},
			"",
			true,
		},
		{
			"test",
			"test",
			banktypes.Metadata{
				Name: "",
			},
			"ATOM",
			true,
		},
		{
			"test",
			"test",
			banktypes.Metadata{
				Name:   "gUSQ",
				Symbol: "",
			},
			"ATOM",
			true,
		},
		{
			"test",
			"test",
			banktypes.Metadata{
				Name:   "gUSQ",
				Symbol: "gUSQ",
				Base:   "",
			},
			"ATOM",
			true,
		},
		{
			"test",
			"test",
			banktypes.Metadata{
				Name:    "gUSQ",
				Symbol:  "gUSQ",
				Base:    "ugusd",
				Display: "",
			},
			"ATOM",
			true,
		},
	}

	for _, tc := range tests {
		msg := types.NewRegisterLendAssetProposal(tc.title, tc.description, tc.assetMetadata, tc.OracleAssetId)
		err := msg.ValidateBasic()
		if tc.expectedErr {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
	}
}

func TestRegisterGTokenPairProposal(t *testing.T) {
	apptypes.SetConfig()
	tests := []struct {
		title          string
		description    string
		gTokenMetadata banktypes.Metadata
		qStablePairId  string
		minAmountInt   string
		minAmountOut   string
		expectedErr    bool
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
			"testqstableid",
			"20uatom",
			"20uusd",
			false,
		},
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
			"",
			"20uatom",
			"20uusd",
			true,
		},
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
			"testqstableid",
			"",
			"20uusd",
			true,
		},
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
			"testqstableid",
			"20uatom",
			"",
			true,
		},
		{
			"test",
			"test",
			banktypes.Metadata{
				Name: "",
			},
			"testqstableid",
			"20uatom",
			"20uusd",
			true,
		},
		{
			"test",
			"test",
			banktypes.Metadata{
				Name:   "Cosmos Hub Atom",
				Symbol: "",
			},
			"testqstableid",
			"20uatom",
			"20uusd",
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
			"testqstableid",
			"20uatom",
			"20uusd",
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
			"testqstableid",
			"20uatom",
			"20uusd",
			true,
		},
	}

	for _, tc := range tests {
		msg := types.NewRegisterGTokenPairProposal(tc.title, tc.description, tc.gTokenMetadata, tc.qStablePairId, tc.minAmountInt, tc.minAmountOut)
		err := msg.ValidateBasic()
		if tc.expectedErr {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
	}
}

func TestRegisterChangeGrowYieldReserveAddressProposal(t *testing.T) {
	apptypes.SetConfig()
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
		msg := types.NewRegisterChangeGrowYieldReserveAddressProposal(tc.title, tc.description, tc.address)
		err := msg.ValidateBasic()
		if tc.expectedErr {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
	}
}

func TestRegisterChangeUSQReserveAddressProposal(t *testing.T) {
	apptypes.SetConfig()
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
		msg := types.NewRegisterChangeUSQReserveAddressProposal(tc.title, tc.description, tc.address)
		err := msg.ValidateBasic()
		if tc.expectedErr {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
	}
}

func TestRegisterChangeGrowStakingReserveAddressProposal(t *testing.T) {
	apptypes.SetConfig()
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
		msg := types.NewRegisterChangeGrowStakingReserveAddressProposal(tc.title, tc.description, tc.address)
		err := msg.ValidateBasic()
		if tc.expectedErr {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
	}
}

func TestRegisterChangeRealRateProposal(t *testing.T) {
	apptypes.SetConfig()
	tests := []struct {
		title       string
		description string
		value       int64
		expectedErr bool
	}{
		{
			"test",
			"test",
			10,
			false,
		},
		{
			"test",
			"test",
			0,
			true,
		},
		{
			"test",
			"test",
			-25,
			true,
		},
	}

	for _, tc := range tests {
		msg := types.NewRegisterChangeRealRateProposal(tc.title, tc.description, uint64(tc.value))
		err := msg.ValidateBasic()
		if tc.expectedErr {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
	}
}

func TestRegisterChangeRealBorrowProposal(t *testing.T) {
	apptypes.SetConfig()
	tests := []struct {
		title       string
		description string
		value       int64
		expectedErr bool
	}{
		{
			"test",
			"test",
			10,
			false,
		},
		{
			"test",
			"test",
			0,
			true,
		},
		{
			"test",
			"test",
			-25,
			true,
		},
	}

	for _, tc := range tests {
		msg := types.NewRegisterChangeBorrowRateProposal(tc.title, tc.description, uint64(tc.value))
		err := msg.ValidateBasic()
		if tc.expectedErr {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
	}
}

func TestRegisterRemoveLendAssetProposal(t *testing.T) {
	apptypes.SetConfig()
	tests := []struct {
		title       string
		description string
		value       string
		expectedErr bool
	}{
		{
			"test",
			"test",
			"testid",
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
		msg := types.NewRegisterRemoveLendAssetProposal(tc.title, tc.description, tc.value)
		err := msg.ValidateBasic()
		if tc.expectedErr {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
	}
}

func TestRegisterRemoveGTokenPairProposal(t *testing.T) {
	apptypes.SetConfig()
	tests := []struct {
		title       string
		description string
		value       string
		expectedErr bool
	}{
		{
			"test",
			"test",
			"testid",
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
		msg := types.NewRegisterRemoveGTokenPairProposal(tc.title, tc.description, tc.value)
		err := msg.ValidateBasic()
		if tc.expectedErr {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
	}
}
