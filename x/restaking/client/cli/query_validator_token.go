package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"lightmos/x/restaking/types"
)

func CmdListValidatorToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-validator-token",
		Short: "list all validatorToken",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllValidatorTokenRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ValidatorTokenAll(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowValidatorToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-validator-token [address]",
		Short: "shows a validatorToken",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			address := args[0]
			params := &types.QueryGetValidatorTokenRequest{
				Address: address,
			}

			res, err := queryClient.ValidatorToken(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
