package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/gov/client/cli"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/spf13/cobra"

	"github.com/comdex-official/comdex/x/ojooracle/types"
)

func NewCmdSubmitFetchPriceProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fetch-price [market-script-id] [requested-validator-count] [sufficient-validator-count] [twa-batch-size] [accepted-height-diff]",
		Args:  cobra.ExactArgs(5),
		Short: "Make a new FetchPrice query request via an existing BandChain market script",
		RunE: func(cmd *cobra.Command, args []string) error {
			// retrieve the market script id.

			// retrieve the requested validator count.
			twaBatch, err := strconv.ParseUint(args[3], 10, 64)
			if err != nil {
				return err
			}

			acceptedHeightDiff, err := strconv.ParseInt(args[4], 10, 64)
			if err != nil {
				return err
			}

			channel, err := cmd.Flags().GetString(flagChannel)
			if err != nil {
				return err
			}

			// retrieve the list of symbols for the requested market script.
			symbols, err := cmd.Flags().GetStringSlice(flagSymbols)
			if err != nil {
				return err
			}

			// retrieve the multiplier for the symbols' price.
			multiplier, err := cmd.Flags().GetUint64(flagMultiplier)
			if err != nil {
				return err
			}

			calldata := &types.FetchPriceCallData{
				Symbols:    symbols,
				Multiplier: multiplier,
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			pmsg := types.NewMsgFetchPriceData(
				clientCtx.GetFromAddress().String(),
				channel,
				calldata,
				twaBatch,
				acceptedHeightDiff,
			)

			title, err := cmd.Flags().GetString(cli.FlagTitle)
			if err != nil {
				return err
			}

			description, err := cmd.Flags().GetString(cli.FlagDescription)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()

			depositStr, err := cmd.Flags().GetString(cli.FlagDeposit)
			if err != nil {
				return err
			}
			deposit, err := sdk.ParseCoinsNormalized(depositStr)
			if err != nil {
				return err
			}

			content := types.NewFetchPriceProposal(title, description, *pmsg)

			msg, err := govtypes.NewMsgSubmitProposal(content, deposit, from)
			if err != nil {
				return err
			}

			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(cli.FlagTitle, "", "title of proposal")
	cmd.Flags().String(cli.FlagDescription, "", "description of proposal")
	cmd.Flags().String(cli.FlagDeposit, "", "deposit of proposal")
	_ = cmd.MarkFlagRequired(cli.FlagTitle)
	_ = cmd.MarkFlagRequired(cli.FlagDescription)
	cmd.Flags().String(flagChannel, "", "The channel id")
	err := cmd.MarkFlagRequired(flagChannel)
	if err != nil {
		return nil
	}
	cmd.Flags().StringSlice(flagSymbols, nil, "Symbols used in calling the market script")
	cmd.Flags().Uint64(flagMultiplier, 1000000, "Multiplier used in calling the market script")
	cmd.Flags().String(flagFeeLimit, "", "the maximum tokens that will be paid to all data source providers")
	cmd.Flags().Uint64(flagPrepareGas, 200000, "Prepare gas used in fee counting for prepare request")
	cmd.Flags().Uint64(flagExecuteGas, 200000, "Execute gas used in fee counting for execute request")

	return cmd
}
