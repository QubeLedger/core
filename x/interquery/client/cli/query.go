package cli

import (
	"context"

	"github.com/QuadrateOrg/core/x/interquery/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the parent command for all erc20 CLI query commands
func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the erc20 module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetInterchainQueryCmd(),
	)
	return cmd
}

func GetInterchainQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query [chain-id]",
		Short: "Get a interchain query by chain-id",
		Long:  "Get a interchain query by chain-id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQuerySrvrClient(clientCtx)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			msg := &types.QueryRequestsRequest{
				Pagination: pageReq,
				ChainId:    args[0],
			}

			res, err := queryClient.Queries(context.Background(), msg)

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
