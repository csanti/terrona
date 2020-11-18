package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/gov"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
  // this line is used by starport scaffolding # 1
  cdc.RegisterConcrete(TaxRateUpdateProposal{}, "clean/TaxRateUpdateProposal", nil)
  cdc.RegisterConcrete(MinStakeUpdateProposal{}, "clean/MinStakeUpdateProposal", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()

	gov.RegisterProposalTypeCodec(MinStakeUpdateProposal{}, "clean/MinStakeUpdateProposal")
	gov.RegisterProposalTypeCodec(TaxRateUpdateProposal{}, "clean/TaxRateUpdateProposal")
}
