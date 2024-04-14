package cli

import (
	"fmt"
	"time"

	"github.com/QuadrateOrg/core/x/perpetual/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

/* #nosec */
var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

//nolint:all
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

	// this line is used by starport scaffolding # 1

	cmd.AddCommand(CmdOpenPosition())
	cmd.AddCommand(CmdClosePosition())

	return cmd
}

func CmdOpenPosition() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "open-position [trade-type] [leverage] [trading-asset] [collateral-coins]",
		Short: "Broadcast message open-position",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTradeType := args[0]
			argLeverage := args[1]
			argTradingAsset := args[2]
			argCollateralCoins := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			var trade_type types.PerpetualTradeType

			switch argTradeType {
			case "Long":
				trade_type = types.PerpetualTradeType_PERPETUAL_LONG_POSITION
			case "Short":
				trade_type = types.PerpetualTradeType_PERPETUAL_SHORT_POSITION
			}

			msg := types.NewMsgOpen(
				clientCtx.GetFromAddress().String(),
				trade_type,
				sdk.MustNewDecFromStr(argLeverage),
				argTradingAsset,
				argCollateralCoins,
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

func CmdClosePosition() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "close-position [amount] [id]",
		Short: "Broadcast message close-position",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAmount := args[0]
			argId := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			v, ok := sdk.NewIntFromString(argAmount)
			if !ok {
				return types.ErrNotSdkInt
			}

			msg := types.NewMsgClose(
				clientCtx.GetFromAddress().String(),
				argId,
				v,
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
