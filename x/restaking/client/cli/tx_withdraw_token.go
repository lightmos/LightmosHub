package cli

import (
	"strconv"

	lightmos "lightmos/types"
	"lightmos/x/restaking/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdWithdrawToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraw-token [amount]",
		Short: "take away the vouncher used for restaking",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAmount := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			amount, err := lightmos.ParseCoinNormalized(argAmount)
			if err != nil {
				return err
			}

			msg := types.NewMsgWithdrawToken(
				clientCtx.GetFromAddress().String(),
				amount,
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
