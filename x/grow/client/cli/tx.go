package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/QuadrateOrg/core/x/grow/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdDeposit())
	cmd.AddCommand(CmdWithdrawal())
	cmd.AddCommand(CmdCreateLend())
	cmd.AddCommand(CmdDeleteLend())
	// this line is used by starport scaffolding # 1

	return cmd
}

func CmdCreateLend() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-lend [amount] [pair-id]",
		Short: "Broadcast message create-lend",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAmount := args[0]
			argPairId := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateLend(
				clientCtx.GetFromAddress().String(),
				argAmount,
				argPairId,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteLend() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-lend [amount] [pair-id]",
		Short: "Broadcast message delete-lend",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAmount := args[0]
			argPairId := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteLend(
				clientCtx.GetFromAddress().String(),
				argAmount,
				argPairId,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeposit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposit [amount] [pair-id]",
		Short: "Broadcast message deposit",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAmount := args[0]
			argPairId := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeposit(
				clientCtx.GetFromAddress().String(),
				argAmount,
				argPairId,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdWithdrawal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdrawal [amount] [pair-id]",
		Short: "Broadcast message withdrawal",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAmount := args[0]
			argPairId := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgWithdrawal(
				clientCtx.GetFromAddress().String(),
				argAmount,
				argPairId,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
