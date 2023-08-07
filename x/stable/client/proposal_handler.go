package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"github.com/QuadrateOrg/core/x/stable/client/cli"
	"github.com/QuadrateOrg/core/x/stable/client/rest"
)

var (
	RegisterPairHandler                       = govclient.NewProposalHandler(cli.NewRegisterPairProposalCmd, rest.RegisterPairRESTHandler)
	RegisterChangeStabilityFundAddressHandler = govclient.NewProposalHandler(cli.NewRegisterChangeStabilityFundAddressProposalCmd, rest.RegisterChangeStabilityFundAddressProposalRESTHandler)
)
