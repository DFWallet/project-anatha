package client

import (
	"github.com/DFWallet/project-anatha/x/distribution/client/cli"
	govclient "github.com/DFWallet/project-anatha/x/governance/client"
)

var DevelopmentFundDistributionProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitDevelopmentFundDistributionProposal)
var SecurityTokenFundDistributionProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitSecurityTokenFundDistributionProposal)
