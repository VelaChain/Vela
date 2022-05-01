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

func CmdSwap() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "swap [amount-in] [denom-in] [min-amount-out] [denom-out]",
		Short: "Broadcast message swap",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAmountIn := args[0]
			argDenomIn := args[1]
			argMinAmountOut := args[2]
			argDenomOut := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSwap(
				clientCtx.GetFromAddress().String(),
				argAmountIn,
				argDenomIn,
				argMinAmountOut,
				argDenomOut,
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
