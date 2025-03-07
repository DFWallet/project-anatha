package cli

import (
	"bufio"
	"github.com/DFWallet/anatha/client/context"
	"github.com/DFWallet/anatha/codec"
	sdk "github.com/DFWallet/anatha/types"
	"github.com/DFWallet/anatha/x/auth"
	"github.com/DFWallet/anatha/x/auth/client/utils"
	distributionutils "github.com/DFWallet/project-anatha/x/distribution/client/utils"
	"github.com/DFWallet/project-anatha/x/distribution/internal/types"
	"github.com/DFWallet/project-anatha/x/governance"
	"github.com/spf13/cobra"
)

func GetCmdSubmitDevelopmentFundDistributionProposal(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "development-fund-distribution [proposal-file]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a Development Fund Distribution proposal",
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			proposal, err := distributionutils.ParseDevelopmentFundDistributionProposalJSON(cdc, args[0])
			if err != nil {
				return err
			}

			from := cliCtx.GetFromAddress()
			content := types.NewDevelopmentFundDistributionProposal(
				proposal.Title,
				proposal.Description,
				proposal.Amount,
				proposal.Recipient,
			)

			msg := governance.NewMsgSubmitProposal(content, from)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}

func GetCmdSubmitSecurityTokenFundDistributionProposal(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "security-token-fund-distribution [proposal-file]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a Security Token Fund Distribution proposal",
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			proposal, err := distributionutils.ParseSecurityTokenFundDistributionProposalJSON(cdc, args[0])
			if err != nil {
				return err
			}

			from := cliCtx.GetFromAddress()
			content := types.NewSecurityTokenFundDistributionProposal(
				proposal.Title,
				proposal.Description,
				proposal.Recipients,
			)

			msg := governance.NewMsgSubmitProposal(content, from)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}