package client

import (
	govclient "github.com/DFWallet/project-anatha/x/governance/client"
	"github.com/DFWallet/project-anatha/x/upgrade/client/cli"
)

var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitSoftwareUpgradeProposal)
