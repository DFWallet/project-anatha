package utils

import (
	"io/ioutil"

	"github.com/DFWallet/anatha/codec"
	sdk "github.com/DFWallet/anatha/types"
)

type GovernorProposalJSON struct {
	Title       string         `json:"title" yaml:"title"`
	Description string         `json:"description" yaml:"description"`
	Governor    sdk.AccAddress `json:"governor" yaml:"governor"`
}

func ParseGovernorProposalJSON(cdc *codec.Codec, proposalFile string) (GovernorProposalJSON, error) {
	proposal := GovernorProposalJSON{}

	contents, err := ioutil.ReadFile(proposalFile)
	if err != nil {
		return proposal, err
	}

	if err := cdc.UnmarshalJSON(contents, &proposal); err != nil {
		return proposal, err
	}

	return proposal, nil
}
