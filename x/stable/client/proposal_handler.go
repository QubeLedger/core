package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"github.com/QuadrateOrg/core/x/stable/client/cli"
	"github.com/QuadrateOrg/core/x/stable/client/rest"
)

var (
	RegisterChangeBaseTokenDenomHendler = govclient.NewProposalHandler(cli.NewRegisterChangeBaseTokenDenomProposalCmd, rest.RegisterChangeBaseTokenDenomRESTHandler)
)
