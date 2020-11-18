package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	//"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/gov"
)

const (
	ProposalTypeTaxRateUpdate = "TaxRateUpdate"
	ProposalTypeMinStakeUpdate = "MinStakeUpdate"
)

// Assert TaxRateUpdateProposal implements govtypes.Content at compile-time
var _ gov.Content = TaxRateUpdateProposal{}
var _ gov.Content = MinStakeUpdateProposal{}

func init() {
	gov.RegisterProposalType(ProposalTypeTaxRateUpdate)
	gov.RegisterProposalType(ProposalTypeMinStakeUpdate)
}

// TaxRateUpdateProposal updates treasury tax-rate
type TaxRateUpdateProposal struct {
	Title       string  `json:"title" yaml:"title"`             // Title of the Proposal
	Description string  `json:"description" yaml:"description"` // Description of the Proposal
	TaxRate     sdk.Dec `json:"tax_rate" yaml:"tax_rate"`       // target TaxRate
}

// NewTaxRateUpdateProposal creates an TaxRateUpdateProposal.
func NewTaxRateUpdateProposal(title, description string, taxRate sdk.Dec) TaxRateUpdateProposal {
	return TaxRateUpdateProposal{title, description, taxRate}
}

// GetTitle returns the title of an TaxRateUpdateProposal.
func (p TaxRateUpdateProposal) GetTitle() string { return p.Title }

// GetDescription returns the description of an TaxRateUpdateProposal.
func (p TaxRateUpdateProposal) GetDescription() string { return p.Description }

// ProposalRoute returns the routing key of an TaxRateUpdateProposal.
func (TaxRateUpdateProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns the type of an TaxRateUpdateProposal.
func (p TaxRateUpdateProposal) ProposalType() string { return ProposalTypeTaxRateUpdate }

// ValidateBasic runs basic stateless validity checks
func (p TaxRateUpdateProposal) ValidateBasic() error {
	// skip validity tests for prototyping
	/*
	err := gov.ValidateAbstract(DefaultCodespace, p)
	if err != nil {
		return err
	}

	if !p.TaxRate.IsPositive() || p.TaxRate.GT(sdk.OneDec()) {
		return sdk.ErrInvalidCoins("Invalid tax-rate: " + p.TaxRate.String())
	}*/

	return nil
}

// String implements the Stringer interface.
func (p TaxRateUpdateProposal) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf(`TaxRateUpdateProposal :
  		Title:        %s
  		Description:  %s
		TaxRate:      %s
		`, p.Title, p.Description, p.TaxRate))
	return b.String()
}


type MinStakeUpdateProposal struct {
	Title        string  	`json:"title" yaml:"title"`                 // Title of the Proposal
	Description  string  	`json:"description" yaml:"description"`     // Description of the Proposal
	MinStake     sdk.Coin   `json:"min_stake" yaml:"min_stake"` 
}

// NewRewardWeightUpdateProposal creates an RewardWeightUpdateProposal.
func NewMinStakeUpdateProposal(title, description string, minStake sdk.Coin) MinStakeUpdateProposal {
	return MinStakeUpdateProposal{title, description, minStake}
}


func (p MinStakeUpdateProposal) GetTitle() string { return p.Title }


func (p MinStakeUpdateProposal) GetDescription() string { return p.Description }


func (MinStakeUpdateProposal) ProposalRoute() string { return RouterKey }


func (p MinStakeUpdateProposal) ProposalType() string { return ProposalTypeMinStakeUpdate }


func (p MinStakeUpdateProposal) ValidateBasic() error {
	// skip validation for prototyping

	return nil
}

// String implements the Stringer interface.
func (p MinStakeUpdateProposal) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf(`ProposalTypeMinStakeUpdate:
  		Title:        %s
  		Description:  %s
		MinStake:      %s
	`, p.Title, p.Description, p.MinStake))
	return b.String()
}