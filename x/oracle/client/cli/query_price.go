package cli

import (
	"context"

	"github.com/QubeLedger/core/x/oracle/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListPrice() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-price",
		Short: "list all price",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllPriceRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.PriceAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowPrice() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-price",
		Short: "shows a price",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			id := uint64(0) //strconv.ParseUint(0, 10, 64)

			params := &types.QueryGetPriceRequest{
				Id: id,
			}

			res, err := queryClient.Price(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
