package cli

import (
	"context"
	"fmt"

	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/QuadrateOrg/core/x/grow/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group grow queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdQueryAssetByAssetId())
	cmd.AddCommand(CmdQueryPositionById())
	cmd.AddCommand(CmdQueryPositionByCreator())
	cmd.AddCommand(CmdQueryAllPosition())
	cmd.AddCommand(CmdQueryAllLiquidatorPosition())
	cmd.AddCommand(CmdQueryLiquidatorPositionByCreator())
	cmd.AddCommand(CmdQueryLiquidatorPositionById())
	cmd.AddCommand(CmdQueryGetAllAsset())
	cmd.AddCommand(CmdQueryLoanById())
	cmd.AddCommand(CmdQueryYieldPercentage())

	return cmd
}

func CmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "shows the parameters of the module",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Params(context.Background(), &types.QueryParamsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQueryAssetByAssetId() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "asset-by-assetId [id]",
		Short: "Get asset by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.AssetByAssetId(context.Background(), &types.QueryAssetByAssetIdRequest{Id: args[0]})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQueryGetAllAsset() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-assets",
		Short: "Get all assets",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.GetAllAssets(context.Background(), &types.QueryGetAllAssetsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQueryPositionById() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "position-by-id [id]",
		Short: "Get position by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.PositionById(context.Background(), &types.QueryPositionByIdRequest{Id: args[0]})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQueryPositionByCreator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "position-by-creator [address]",
		Short: "Get position by address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.PositionByCreator(context.Background(), &types.QueryPositionByCreatorRequest{Creator: args[0]})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQueryAllPosition() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "position",
		Short: "Get all position",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.AllPosition(context.Background(), &types.QueryAllPositionRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQueryAllLiquidatorPosition() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "liq-position",
		Short: "Get all Liquidator position",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.AllLiquidatorPosition(context.Background(), &types.QueryAllLiquidatorPositionRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQueryLiquidatorPositionByCreator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "liq-position-by-creator [address]",
		Short: "Get liquidator position by address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.LiquidatorPositionByCreator(context.Background(), &types.QueryLiquidatorPositionByCreatorRequest{Creator: args[0]})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQueryLiquidatorPositionById() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "liq-position-by-id [id]",
		Short: "Get liquidator position by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.LiquidatorPositionById(context.Background(), &types.QueryLiquidatorPositionByIdRequest{Id: args[0]})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQueryLoanById() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "loan-by-id [id]",
		Short: "Get loan by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.LoanById(context.Background(), &types.QueryLoanByIdRequest{Id: args[0]})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQueryYieldPercentage() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "yield-percentage [id]",
		Short: "Get yield percentage by gTokenPair id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.YieldPercentage(context.Background(), &types.QueryYieldPercentageRequest{Id: args[0]})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
