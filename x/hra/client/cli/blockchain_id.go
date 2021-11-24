package cli

import (
	"bufio"
	"github.com/DFWallet/anatha/client/context"
	"github.com/DFWallet/anatha/codec"
	sdk "github.com/DFWallet/anatha/types"
	"github.com/DFWallet/anatha/x/auth"
	"github.com/DFWallet/anatha/x/auth/client/utils"
	govtypes "github.com/DFWallet/project-anatha/x/governance"
	govutils "github.com/DFWallet/project-anatha/x/hra/client/utils"
	"github.com/DFWallet/project-anatha/x/hra/internal/types"
	"github.com/spf13/cobra"
)

func GetCmdSubmitRegisterBlockchainIdProposal(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-blockchain-id [proposal-file]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a proposal to register a blockchain id",
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			proposal, err := govutils.ParseBlockchainIdProposalJSON(cdc, args[0])
			if err != nil {
				return err
			}

			from := cliCtx.GetFromAddress()

			content := types.NewRegisterBlockchainIdProposal(proposal.Title, proposal.Description, proposal.BlockchainId)

			msg := govtypes.NewMsgSubmitProposal(content, from)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}

func GetCmdSubmitRemoveBlockchainIdProposal(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-blockchain-id [proposal-file]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a proposal to remove a blockchain id",
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			proposal, err := govutils.ParseBlockchainIdProposalJSON(cdc, args[0])
			if err != nil {
				return err
			}

			from := cliCtx.GetFromAddress()

			content := types.NewRemoveBlockchainIdProposal(proposal.Title, proposal.Description, proposal.BlockchainId)

			msg := govtypes.NewMsgSubmitProposal(content, from)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}
