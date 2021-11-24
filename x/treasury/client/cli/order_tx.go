package cli

import (
	"bufio"
	"fmt"
	"github.com/DFWallet/anatha/client/context"
	"github.com/DFWallet/anatha/x/auth"
	"github.com/DFWallet/anatha/x/auth/client/utils"
	denom "github.com/DFWallet/project-anatha/utils"

	"github.com/spf13/cobra"

	"github.com/DFWallet/anatha/client"
	"github.com/DFWallet/anatha/client/flags"
	"github.com/DFWallet/anatha/codec"
	"github.com/DFWallet/project-anatha/x/treasury/internal/types"

	sdk "github.com/DFWallet/anatha/types"
)

func GetCmdOrder(cdc *codec.Codec) *cobra.Command {
	orderTxCmd := &cobra.Command{
		Use:                        "order",
		Short:                      fmt.Sprintf("%s transactions subcommands", "order"),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	orderTxCmd.AddCommand(flags.PostCommands(
		GetCmdAddSellOrder(cdc),
		GetCmdAddBuyOrder(cdc),
	)...)

	return orderTxCmd
}

func GetCmdAddSellOrder(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "sell [anatha-amount]",
		Short: "Create a sell order",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			amount, err := denom.ParseAndConvertCoins(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateSellOrder(cliCtx.GetFromAddress(), amount)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdAddBuyOrder(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "buy [ast-amount]",
		Short: "Create a buy order",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			amount, err := denom.ParseAndConvertCoins(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateBuyOrder(cliCtx.GetFromAddress(), amount)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
