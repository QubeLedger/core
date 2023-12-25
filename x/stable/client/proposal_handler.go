package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"github.com/QuadrateOrg/core/x/stable/client/cli"
	"github.com/QuadrateOrg/core/x/stable/client/rest"
)

var (
	RegisterPairHandler                     = govclient.NewProposalHandler(cli.NewRegisterPairProposalCmd, rest.RegisterPairRESTHandler)
	RegisterChangeBurningFundAddressHandler = govclient.NewProposalHandler(cli.NewRegisterChangeBurningFundAddressProposalCmd, rest.RegisterChangeBurningFundAddressProposalRESTHandler)
	RegisterChangeReserveFundAddressHandler = govclient.NewProposalHandler(cli.NewRegisterChangeReserveFundAddressProposalCmd, rest.RegisterChangeReserveFundAddressProposalRESTHandler)
	RegisterDeletePairHandler               = govclient.NewProposalHandler(cli.NewRegisterDeletePairProposalCmd, rest.RegisterDeletePairProposalRESTHandler)
)
