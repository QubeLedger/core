package types

import (
	fmt "fmt"
	"strings"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
)

// constants
const (
	ProposalTypeChangeBaseTokenDenom string = "ChangeBaseTokenDenom"
)

// Implements Proposal Interface
var (
	_ govtypes.Content = &ChangeBaseTokenDenom{}
)

func init() {
	govtypes.RegisterProposalType(ProposalTypeChangeBaseTokenDenom)
	govtypes.RegisterProposalTypeCodec(&ChangeBaseTokenDenom{}, "stable/ChangeBaseTokenDenom")
}

func NewRegisterChangeBaseTokenDenomProposal(title, description string, coinMetadata ...banktypes.Metadata) govtypes.Content {
	return &ChangeBaseTokenDenom{
		Title:       title,
		Description: description,
		Metadata:    coinMetadata,
	}
}

func (*ChangeBaseTokenDenom) ProposalRoute() string { return RouterKey }

func (*ChangeBaseTokenDenom) ProposalType() string {
	return ProposalTypeChangeBaseTokenDenom
}

func (rtbp *ChangeBaseTokenDenom) ValidateBasic() error {
	for _, metadata := range rtbp.Metadata {
		if err := ibctransfertypes.ValidateIBCDenom(metadata.Base); err != nil {
			return err
		}

		if err := validateIBCVoucherMetadata(metadata); err != nil {
			return err
		}
	}

	return nil
}

func validateIBCVoucherMetadata(metadata banktypes.Metadata) error {
	// Check ibc/ denom
	denomSplit := strings.SplitN(metadata.Base, "/", 2)

	if denomSplit[0] == metadata.Base && strings.TrimSpace(metadata.Base) != "" {
		// Not IBC
		return nil
	}

	if len(denomSplit) != 2 || denomSplit[0] != ibctransfertypes.DenomPrefix {
		// NOTE: should be unaccessible (covered on ValidateIBCDenom)
		return fmt.Errorf("invalid metadata. %s denomination should be prefixed with the format 'ibc/", metadata.Base)
	}

	return nil
}
