package cli

import (
	"context"

	"github.com/VelaChain/vela/x/market/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListLiqProv() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-liq-prov",
		Short: "list all liqProv",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllLiqProvRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.LiqProvAll(context.Background(), params)
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

func CmdShowLiqProv() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-liq-prov [pool-name] [address]",
		Short: "shows a liqProv",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argPoolName := args[0]
			argAddress := args[1]

			params := &types.QueryGetLiqProvRequest{
				PoolName: argPoolName,
				Address:  argAddress,
			}

			res, err := queryClient.LiqProv(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
