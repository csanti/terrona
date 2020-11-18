package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	assets "github.com/csanti/terrona/types"
)

// GenesisState - all clean state that must be provided at genesis
type GenesisState struct {
	TaxRate              sdk.Dec           `json:"tax_rate" yaml:"tax_rate"`
	MinStake			 sdk.Coin          `json:"min_stake" yaml:"min_stake"`
}

var (
	DefaultTaxRate	= sdk.NewDecWithPrec(1,3)
	DefaultMinStake = sdk.NewCoin(assets.MicroMaskDenom, sdk.ZeroInt())
)

// NewGenesisState creates a new GenesisState object
func NewGenesisState(taxRate sdk.Dec, minStake sdk.Coin) GenesisState {
	return GenesisState {
		TaxRate:	taxRate,
		MinStake:	minStake,
	}
}

// DefaultGenesisState - default GenesisState used by Cosmos Hub
func DefaultGenesisState() GenesisState {
	return GenesisState{
		TaxRate:	DefaultTaxRate,
		MinStake:	DefaultMinStake,
	}
}

// ValidateGenesis validates the clean genesis parameters
func ValidateGenesis(data GenesisState) error {
	// should implement validation, skip for prototype
	return nil
}
