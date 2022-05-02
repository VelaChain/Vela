package cli

import (
	"strconv"

	"github.com/VelaChain/vela/x/market/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSendShares() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send-shares [shares] [denom-a] [denom-b] [address]",
		Short: "Broadcast message send-shares",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argShares := args[0]
			argDenomA := args[1]
			argDenomB := args[2]
			argAddress := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSendShares(
				clientCtx.GetFromAddress().String(),
				argShares,
				argDenomA,
				argDenomB,
				argAddress,
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
