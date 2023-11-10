package types

import (
	fmt "fmt"
	"strings"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

// constants
const (
	ProposalTypeRegisterAddNewDenomProposal string = "RegisterAddNewDenomProposal"
)

// Implements Proposal Interface
var (
	_ govtypes.Content = &RegisterAddNewDenomProposal{}
)

func init() {
	govtypes.RegisterProposalType(ProposalTypeRegisterAddNewDenomProposal)
	govtypes.RegisterProposalTypeCodec(&RegisterAddNewDenomProposal{}, "oracle/RegisterAddNewDenomProposal")
}

/*
PairProposal
*/

func NewRegisterAddNewDenomProposal(title, description string, denom string) govtypes.Content {
	return &RegisterAddNewDenomProposal{
		Title:       title,
		Description: description,
		Denom:       denom,
	}
}

func (*RegisterAddNewDenomProposal) ProposalRoute() string { return RouterKey }

func (*RegisterAddNewDenomProposal) ProposalType() string {
	return ProposalTypeRegisterAddNewDenomProposal
}

func (rtbp *RegisterAddNewDenomProposal) ValidateBasic() error {
	{
		if strings.TrimSpace(rtbp.Denom) == "" {
			return fmt.Errorf("Denom field cannot be blank")
		}
	}

	return nil
}
