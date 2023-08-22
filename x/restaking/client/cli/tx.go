package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"lightmos/x/restaking/types"
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

	cmd.AddCommand(CmdSendBuyOrder())
	cmd.AddCommand(CmdCancelSellOrder())
	cmd.AddCommand(CmdCancelBuyOrder())
	cmd.AddCommand(CmdCreateValidatorCmd())
	cmd.AddCommand(CmdSendRetireShare())
	cmd.AddCommand(CmdSendUndelegate())
	cmd.AddCommand(CmdWithdrawToken())
	// this line is used by starport scaffolding # 1

	return cmd
}
