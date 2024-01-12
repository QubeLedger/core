package client

import (
	"github.com/QuadrateOrg/core/x/grow/client/cli"
	"github.com/QuadrateOrg/core/x/grow/client/rest"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
)

var (
	RegisterLendAssetProposalHandler                       = govclient.NewProposalHandler(cli.NewRegisterLendAssetProposalCmd, rest.RegisterLendAssetProposalRESTHandler)
	RegisterGTokenPairProposalHandler                      = govclient.NewProposalHandler(cli.NewRegisterGTokenPairProposalCmd, rest.RegisterGTokenPairProposalRESTHandler)
	RegisterChangeGrowYieldReserveAddressProposalHandler   = govclient.NewProposalHandler(cli.NewRegisterChangeGrowYieldReserveAddressProposalCmd, rest.RegisterChangeGrowYieldReserveAddressProposalRESTHandler)
	RegisterChangeUSQReserveAddressProposalHandler         = govclient.NewProposalHandler(cli.NewRegisterChangeUSQReserveAddressProposalCmd, rest.RegisterChangeUSQReserveAddressProposalRESTHandler)
	RegisterChangeGrowStakingReserveAddressProposalHandler = govclient.NewProposalHandler(cli.NewRegisterChangeGrowStakingReserveAddressProposalCmd, rest.RegisterChangeGrowStakingReserveAddressProposalRESTHandler)
	RegisterChangeRealRateProposalHandler                  = govclient.NewProposalHandler(cli.NewRegisterChangeRealRateProposalCmd, rest.RegisterChangeRealRateProposalRESTHandler)
	RegisterChangeBorrowRateProposalHandler                = govclient.NewProposalHandler(cli.NewRegisterChangeBorrowRateProposalCmd, rest.RegisterChangeBorrowRateProposalRESTHandler)
	RegisterChangeLendRateProposalHandler                  = govclient.NewProposalHandler(cli.NewRegisterChangeLendRateProposalCmd, rest.RegisterChangeLendRateProposalRESTHandler)
	RegisterActivateGrowModuleProposalHandler              = govclient.NewProposalHandler(cli.NewRegisterActivateGrowModuleProposalCmd, rest.RegisterActivateGrowModuleProposalRESTHandler)
	RegisterRemoveLendAssetProposalHandler                 = govclient.NewProposalHandler(cli.NewRegisterRemoveLendAssetProposalCmd, rest.RegisterRemoveLendAssetProposalRESTHandler)
	RegisterRemoveGTokenPairProposalHandler                = govclient.NewProposalHandler(cli.NewRegisterRemoveGTokenPairProposalCmd, rest.RegisterRemoveGTokenPairProposalRESTHandler)
)
