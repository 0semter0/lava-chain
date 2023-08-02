package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/lavanet/lava/x/pairing/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group pairing queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdProviders())
	cmd.AddCommand(CmdGetPairing())
	cmd.AddCommand(CmdVerifyPairing())
	cmd.AddCommand(CmdListUniquePaymentStorageClientProvider())
	cmd.AddCommand(CmdShowUniquePaymentStorageClientProvider())
	cmd.AddCommand(CmdListProviderPaymentStorage())
	cmd.AddCommand(CmdShowProviderPaymentStorage())
	cmd.AddCommand(CmdListEpochPayments())
	cmd.AddCommand(CmdShowEpochPayments())
	cmd.AddCommand(CmdUserMaxCu())

	cmd.AddCommand(CmdStaticProvidersList())
	cmd.AddCommand(CmdAccountInfo())
	cmd.AddCommand(CmdEffectivePolicy())

	cmd.AddCommand(CmdSdkPairing())

// this line is used by starport scaffolding # 1

	return cmd
}
