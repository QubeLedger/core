package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/gov/client/cli"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/QuadrateOrg/core/x/stable/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds()) // #nosec
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

	cmd.AddCommand(CmdMintUsq())
	cmd.AddCommand(CmdBurnUsq())
	// this line is used by starport scaffolding # 1

	return cmd
}

func CmdMintUsq() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint [amount] [denom]",
		Short: "Broadcast message mint",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgMint(
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

func CmdBurnUsq() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn [amount] [denom]",
		Short: "Broadcast message burn",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
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

func NewRegisterPairProposalCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-pair [metadata]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a register pair proposal",
		Long:  `Submit a proposal for couple's registration to x/stable along with the down payment. The offer data should be submitted as a JSON file.`,
		Example: fmt.Sprintf(`qubed tx gov submit-proposal register-pair metadata.json --from=<key_or_address>

		Where metadata.json contains (example):
		{
			"amountInMetadata": {
				"description": "The native staking and governance token of the Cosmos chain",
				"denom_units": [
					{
						"denom": "ibc/<HASH>",
						"exponent": 0,
						"aliases": ["ibcuatom"]
					},
					{
						"denom": "ATOM",
						"exponent": 6
					}
				],
				"base": "ibc/<HASH>",
				"display": "ATOM",
				"name": "Atom",
				"symbol": "ATOM"
			},
			"amountOutMetadata": {
				"description": "First algorithmic stablecoin backed by ATOM",
				"denom_units": [
					{
						"denom": "uusd",
						"exponent": 0,
						"aliases": ["uusd"]
					},
					{
						"denom": "USQ",
						"exponent": 6
					}
				],
				"base": "uusd",
				"display": "USQ",
				"name": "USQ",
				"symbol": "USQ"
			},
			"minAmountIn": "20ibc/<HASH>"
		}`,
		),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			title, err := cmd.Flags().GetString(cli.FlagTitle)
			if err != nil {
				return err
			}

			description, err := cmd.Flags().GetString(cli.FlagDescription)
			if err != nil {
				return err
			}

			depositStr, err := cmd.Flags().GetString(cli.FlagDeposit)
			if err != nil {
				return err
			}

			deposit, err := sdk.ParseCoinsNormalized(depositStr)
			if err != nil {
				return err
			}
			amountInMetadata, amountoutMetadata, minAmount, err := ParseMetadata(clientCtx.Codec, args[0])
			if err != nil {
				return err
			}
			from := clientCtx.GetFromAddress()

			content := types.NewRegisterPairProposal(title, description, amountInMetadata, amountoutMetadata, minAmount)

			msg, err := govtypes.NewMsgSubmitProposal(content, deposit, from)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	cmd.Flags().String(cli.FlagTitle, "", "title of proposal")
	cmd.Flags().String(cli.FlagDescription, "", "description of proposal")
	cmd.Flags().String(cli.FlagDeposit, "1uqube", "deposit of proposal")
	if err := cmd.MarkFlagRequired(cli.FlagTitle); err != nil {
		panic(err)
	}
	if err := cmd.MarkFlagRequired(cli.FlagDescription); err != nil {
		panic(err)
	}
	if err := cmd.MarkFlagRequired(cli.FlagDeposit); err != nil {
		panic(err)
	}

	return cmd
}
