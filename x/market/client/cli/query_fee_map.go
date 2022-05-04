package cli

import (
	"context"

	"github.com/VelaChain/vela/x/market/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListFeeMap() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-fee-map",
		Short: "list all feeMap",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllFeeMapRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.FeeMapAll(context.Background(), params)
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

func CmdShowFeeMap() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-fee-map [pool-name]",
		Short: "shows a feeMap",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argPoolName := args[0]

			params := &types.QueryGetFeeMapRequest{
				PoolName: argPoolName,
			}

			res, err := queryClient.FeeMap(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
