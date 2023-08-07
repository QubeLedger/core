package cli

import (
	"os"
	"path/filepath"

	"github.com/QuadrateOrg/core/x/stable/types"
	"github.com/cosmos/cosmos-sdk/codec"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// ParseRegisterCoinProposal reads and parses a ParseRegisterCoinProposal from a file.
func ParseMetadata(cdc codec.JSONCodec, metadataFile string) (banktypes.Metadata, banktypes.Metadata, string, error) {
	proposalMetadata := types.ProposalMetadata{}

	contents, err := os.ReadFile(filepath.Clean(metadataFile))
	if err != nil {
		return banktypes.Metadata{}, banktypes.Metadata{}, "", err
	}

	if err = cdc.UnmarshalJSON(contents, &proposalMetadata); err != nil {
		return banktypes.Metadata{}, banktypes.Metadata{}, "", err
	}

	return proposalMetadata.AmountInMetadata, proposalMetadata.AmountOutMetadata, proposalMetadata.MinAmountIn, nil
}
