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
	ProposalTypeRegisterPairProposal                       string = "RegisterPairProposal"
	ProposalTypeRegisterChangeStabilityFundAddressProposal string = "RegisterChangeStabilityFundAddressProposal"
)

// Implements Proposal Interface
var (
	_ govtypes.Content = &RegisterPairProposal{}
	_ govtypes.Content = &RegisterChangeStabilityFundAddressProposal{}
)

func init() {
	govtypes.RegisterProposalType(ProposalTypeRegisterPairProposal)
	govtypes.RegisterProposalTypeCodec(&RegisterPairProposal{}, "stable/RegisterPairProposal")

	govtypes.RegisterProposalType(ProposalTypeRegisterChangeStabilityFundAddressProposal)
	govtypes.RegisterProposalTypeCodec(&RegisterChangeStabilityFundAddressProposal{}, "stable/RegisterChangeStabilityFundAddressProposal")
}

func NewRegisterPairProposal(title, description string, amountInDenom banktypes.Metadata, amountOutDenom banktypes.Metadata, minAmount string) govtypes.Content {
	return &RegisterPairProposal{
		Title:             title,
		Description:       description,
		AmountInMetadata:  amountInDenom,
		AmountOutMetadata: amountOutDenom,
		MinAmountIn:       minAmount,
	}
}

func NewRegisterChangeStabilityFundAddressProposal(title, description string, address string) govtypes.Content {
	return &RegisterChangeStabilityFundAddressProposal{
		Title:       title,
		Description: description,
		Address:     address,
	}
}

func (*RegisterChangeStabilityFundAddressProposal) ProposalRoute() string { return RouterKey }

func (*RegisterChangeStabilityFundAddressProposal) ProposalType() string {
	return ProposalTypeRegisterChangeStabilityFundAddressProposal
}

func (rtbp *RegisterChangeStabilityFundAddressProposal) ValidateBasic() error {
	if len(rtbp.Address) == 0 {
		return ErrInvalidLength
	}
	_, err := sdk.AccAddressFromBech32(rtbp.Address)
	if err != nil {
		return nil
	}
	return nil
}

func (*RegisterPairProposal) ProposalRoute() string { return RouterKey }

func (*RegisterPairProposal) ProposalType() string {
	return ProposalTypeRegisterPairProposal
}

func (rtbp *RegisterPairProposal) ValidateBasic() error {
	{
		if strings.TrimSpace(rtbp.AmountInMetadata.Name) == "" {
			return fmt.Errorf("name field cannot be blank")
		}

		if strings.TrimSpace(rtbp.AmountInMetadata.Symbol) == "" {
			return fmt.Errorf("symbol field cannot be blank")
		}

		if err := sdk.ValidateDenom(rtbp.AmountInMetadata.Base); err != nil {
			return fmt.Errorf("invalid metadata base denom: %w", err)
		}

		if err := sdk.ValidateDenom(rtbp.AmountInMetadata.Display); err != nil {
			return fmt.Errorf("invalid metadata display denom: %w", err)
		}
		if err := ibctransfertypes.ValidateIBCDenom(rtbp.AmountInMetadata.Base); err != nil {
			return err
		}

		if err := validateIBCVoucherMetadata(rtbp.AmountInMetadata); err != nil {
			return err
		}
	}
	{
		if strings.TrimSpace(rtbp.AmountOutMetadata.Name) == "" {
			return fmt.Errorf("name field cannot be blank")
		}

		if strings.TrimSpace(rtbp.AmountOutMetadata.Symbol) == "" {
			return fmt.Errorf("symbol field cannot be blank")
		}

		if err := sdk.ValidateDenom(rtbp.AmountOutMetadata.Base); err != nil {
			return fmt.Errorf("invalid metadata base denom: %w", err)
		}

		if err := sdk.ValidateDenom(rtbp.AmountOutMetadata.Display); err != nil {
			return fmt.Errorf("invalid metadata display denom: %w", err)
		}
		if err := ibctransfertypes.ValidateIBCDenom(rtbp.AmountOutMetadata.Base); err != nil {
			return err
		}

		if err := validateIBCVoucherMetadata(rtbp.AmountOutMetadata); err != nil {
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
