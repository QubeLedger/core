package cli

import (
	"strconv"

	"github.com/0xknstntn/quadrate/x/converter/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdErc20Cw20() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "erc-20-cw-20 [amount, address]",
		Short: "Broadcast message erc20-cw20",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			//argAmount := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgErc20Cw20(
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
