package cli

import (
	"errors"
	channelutils "github.com/cosmos/ibc-go/v7/modules/core/04-channel/client/utils"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"lightmos/x/restaking/types"
)

var _ = strconv.Itoa(0)

func CmdChangePairState() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "change-pair-state [port] [channel-id] [source-denom] [target-denom] [price] [author]",
		Short: "Broadcast message changePairState",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			srcPort := args[0]
			srcChannel := args[1]
			argSourceDenom := args[2]
			argTargetDenom := args[3]
			argPrice, err := cast.ToInt32E(args[4])
			if err != nil {
				return err
			}

			author := args[5]
			if author != "buyer" && author != "seller" {
				return errors.New("author is buyer or seller")
			}
			// Get the relative timeout timestamp
			timeoutTimestamp, err := cmd.Flags().GetUint64(flagPacketTimeoutTimestamp)
			if err != nil {
				return err
			}
			consensusState, _, _, err := channelutils.QueryLatestConsensusState(clientCtx, srcPort, srcChannel)
			if err != nil {
				return err
			}
			if timeoutTimestamp != 0 {
				timeoutTimestamp = consensusState.GetTimestamp() + timeoutTimestamp
			}

			msg := types.NewMsgChangePairState(
				clientCtx.GetFromAddress().String(),
				srcPort,
				srcChannel,
				argSourceDenom,
				argTargetDenom,
				argPrice,
				author,
				timeoutTimestamp,
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
