package cli

import (
	"os"
	"path/filepath"

	"github.com/QuadrateOrg/core/x/grow/types"
	"github.com/cosmos/cosmos-sdk/codec"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// ParseRegisterCoinProposal reads and parses a ParseRegisterCoinProposal from a file.
func ParseMetadataForLendAssetProposal(cdc codec.JSONCodec, metadataFile string) (banktypes.Metadata, string, error) {
	proposalMetadata := types.ProposalMetadataForRegisterLendAssetProposal{}

	contents, err := os.ReadFile(filepath.Clean(metadataFile))
	if err != nil {
		return banktypes.Metadata{}, "", err
	}

	if err = cdc.UnmarshalJSON(contents, &proposalMetadata); err != nil {
		return banktypes.Metadata{}, "", err
	}

	return proposalMetadata.AssetMetadata, proposalMetadata.OracleAssetId, nil
}

func ParseMetadataForGTokenPairProposal(cdc codec.JSONCodec, metadataFile string) (banktypes.Metadata, string, string, string, error) {
	proposalMetadata := types.ProposalMetadataForRegisterGTokenPairProposal{}

	contents, err := os.ReadFile(filepath.Clean(metadataFile))
	if err != nil {
		return banktypes.Metadata{}, "", "", "", err
	}

	if err = cdc.UnmarshalJSON(contents, &proposalMetadata); err != nil {
		return banktypes.Metadata{}, "", "", "", err
	}

	return proposalMetadata.GTokenMetadata, proposalMetadata.QStablePairId, proposalMetadata.MinAmountIn, proposalMetadata.MinAmountOut, nil
}
