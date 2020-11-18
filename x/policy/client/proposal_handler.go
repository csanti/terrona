package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	"github.com/csanti/terrona/x/policy/client/cli"
	"github.com/csanti/terrona/x/policy/client/rest"
)

// param change proposal handler
var (
	TaxRateUpdateProposalHandler      = govclient.NewProposalHandler(cli.GetCmdSubmitTaxRateUpdateProposal, rest.TaxRateUpdateProposalRESTHandler)
	MinStakeUpdateProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitMinStakeUpdateProposal, rest.MinStakeUpdateProposalRESTHandler)
)