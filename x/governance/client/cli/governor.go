package cli

import (
	"bufio"
	"github.com/DFWallet/anatha/client/context"
	"github.com/DFWallet/anatha/codec"
	sdk "github.com/DFWallet/anatha/types"
	"github.com/DFWallet/anatha/x/auth"
	"github.com/DFWallet/anatha/x/auth/client/utils"
	govutils "github.com/DFWallet/project-anatha/x/governance/client/utils"
	govtypes "github.com/DFWallet/project-anatha/x/governance/internal/types"
	"github.com/spf13/cobra"
)

func GetCmdSubmitAddGovernorProposal(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-governor [proposal-file]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a proposal to add a new governor",
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			proposal, err := govutils.ParseGovernorProposalJSON(cdc, args[0])
			if err != nil {
				return err
			}

			from := cliCtx.GetFromAddress()

			content := govtypes.NewAddGovernorProposal(proposal.Title, proposal.Description, proposal.Governor)

			msg := govtypes.NewMsgSubmitProposal(content, from)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}

func GetCmdSubmitRemoveGovernorProposal(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-governor [proposal-file]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a proposal to remove a governor",
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			proposal, err := govutils.ParseGovernorProposalJSON(cdc, args[0])
			if err != nil {
				return err
			}

			from := cliCtx.GetFromAddress()

			content := govtypes.NewRemoveGovernorProposal(proposal.Title, proposal.Description, proposal.Governor)

			msg := govtypes.NewMsgSubmitProposal(content, from)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}