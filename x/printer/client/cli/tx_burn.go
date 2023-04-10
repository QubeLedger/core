package cli

import (
	"strconv"
	"strings"

	"github.com/QuadrateOrg/core/x/printer/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdBurn() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn [amount] [validator]",
		Short: "USQ combustion to obtain Qube",
		Args:  cobra.ExactArgs(2),
		Long:  strings.TrimSpace(`$ quadrated printer burn 10000usq qubevaloperkgb366x4hqa4d3457cjjizwsnhp9m4hhw7fby72`),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			//argAmount := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgBurn(
				clientCtx.GetFromAddress().String(),
				args[0],
				args[1],
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
