package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
)

// GetQueryCmd returns the cli query commands for this module.
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group bandoracle queries under a subcommand.
	cmd := &cobra.Command{
		Use:                        "ojooracle",
		Short:                      fmt.Sprintf("Querying commands for the %s module", "ojooracle"),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())

	cmd.AddCommand(CmdFetchPriceResult())
	cmd.AddCommand(CmdLastFetchPriceID())
	cmd.AddCommand(CmdFetchPriceData())
	cmd.AddCommand(CmdDiscardData())

	return cmd
}

// GetTxCmd returns the transaction commands for this module.
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "ojooracle",
		Short:                      fmt.Sprintf("%s transactions subcommands", "ojooracle"),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand()

	return cmd
}
