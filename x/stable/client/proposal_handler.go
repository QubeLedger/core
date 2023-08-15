package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"github.com/QubeLedger/core/x/stable/client/cli"
	"github.com/QubeLedger/core/x/stable/client/rest"
)

var (
	RegisterPairHandler                     = govclient.NewProposalHandler(cli.NewRegisterPairProposalCmd, rest.RegisterPairRESTHandler)
	RegisterChangeBurningFundAddressHandler = govclient.NewProposalHandler(cli.NewRegisterChangeBurningFundAddressProposalCmd, rest.RegisterChangeBurningFundAddressProposalRESTHandler)
)
