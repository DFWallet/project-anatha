package client

import (
	"github.com/DFWallet/project-anatha/x/fee/client/cli"
	govclient "github.com/DFWallet/project-anatha/x/governance/client"
)

var (
	AddFeeExcludedMessageProposalHandler    = govclient.NewProposalHandler(cli.GetCmdSubmitAddFeeExcludedMessageProposal)
	RemoveFeeExcludedMessageProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitRemoveFeeExcludedMessageProposal)
)
