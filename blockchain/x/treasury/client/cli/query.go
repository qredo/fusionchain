package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/qredo/fusionchain/x/treasury/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(_ string) *cobra.Command {
	// Group treasury queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdKeyRequests())
	cmd.AddCommand(CmdKeys())
	cmd.AddCommand(CmdSignatureRequests())
	cmd.AddCommand(CmdKeyRequestById())
	cmd.AddCommand(CmdSignatureRequestById())
	cmd.AddCommand(CmdSignTransactionRequests())
	cmd.AddCommand(CmdSignTransactionRequestById())
	// this line is used by starport scaffolding # 1

	return cmd
}
