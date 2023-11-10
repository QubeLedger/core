package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"github.com/QuadrateOrg/core/x/oracle/client/cli"
	"github.com/QuadrateOrg/core/x/oracle/client/rest"
)

var (
	RegisterAddNewDenomProposal = govclient.NewProposalHandler(cli.NewRegisterAddNewDenomProposalCmd, rest.RegisterAddNewDenomProposalRESTHandler)
)
