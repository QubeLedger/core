package types

import (
	fmt "fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
)

// constants
const (
	ProposalTypeChangeBaseTokenDenom string = "ChangeBaseTokenDenom"
	ProposalTypeChangeSendTokenDenom string = "ChangeSendTokenDenom"
)

// Implements Proposal Interface
var (
	_ govtypes.Content = &ChangeBaseTokenDenom{}
	_ govtypes.Content = &ChangeSendTokenDenom{}
)

func init() {
	govtypes.RegisterProposalType(ProposalTypeChangeBaseTokenDenom)
	govtypes.RegisterProposalTypeCodec(&ChangeBaseTokenDenom{}, "stable/ChangeBaseTokenDenom")
	govtypes.RegisterProposalTypeCodec(&ChangeSendTokenDenom{}, "stable/ChangeSendTokenDenom")
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

func NewRegisterChangeSendTokenDenomProposal(title, description string, coinMetadata ...banktypes.Metadata) govtypes.Content {
	return &ChangeSendTokenDenom{
		Title:       title,
		Description: description,
		Metadata:    coinMetadata,
	}
}

func (*ChangeSendTokenDenom) ProposalRoute() string { return RouterKey }

func (*ChangeSendTokenDenom) ProposalType() string {
	return ProposalTypeChangeSendTokenDenom
}

func (rtbp *ChangeSendTokenDenom) ValidateBasic() error {
	for _, metadata := range rtbp.Metadata {
		if strings.TrimSpace(metadata.Name) == "" {
			return fmt.Errorf("name field cannot be blank")
		}

		if strings.TrimSpace(metadata.Symbol) == "" {
			return fmt.Errorf("symbol field cannot be blank")
		}

		if err := sdk.ValidateDenom(metadata.Base); err != nil {
			return fmt.Errorf("invalid metadata base denom: %w", err)
		}

		if err := sdk.ValidateDenom(metadata.Display); err != nil {
			return fmt.Errorf("invalid metadata display denom: %w", err)
		}
		if err := ibctransfertypes.ValidateIBCDenom(metadata.Base); err != nil {
			return err
		}

		if err := validateIBCVoucherMetadata(metadata); err != nil {
			return err
		}
	}

	return nil
}

func (rtbp *ChangeBaseTokenDenom) ValidateBasic() error {
	for _, metadata := range rtbp.Metadata {
		if strings.TrimSpace(metadata.Name) == "" {
			return fmt.Errorf("name field cannot be blank")
		}

		if strings.TrimSpace(metadata.Symbol) == "" {
			return fmt.Errorf("symbol field cannot be blank")
		}

		if err := sdk.ValidateDenom(metadata.Base); err != nil {
			return fmt.Errorf("invalid metadata base denom: %w", err)
		}

		if err := sdk.ValidateDenom(metadata.Display); err != nil {
			return fmt.Errorf("invalid metadata display denom: %w", err)
		}
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
