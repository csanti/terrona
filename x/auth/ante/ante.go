package ante

import (
	//"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	cante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	cleank "github.com/csanti/terrona/x/policy/keeper"
)

// NewAnteHandler returns an AnteHandler that checks and increments sequence
// numbers, checks signatures & account numbers, and deducts fees from the first
// signer.
func NewAnteHandler(ak keeper.AccountKeeper, supplyKeeper types.SupplyKeeper, sigGasConsumer cante.SignatureVerificationGasConsumer, policyKeeper cleank.Keeper) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		cante.NewSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first
		NewMempoolFeeDecorator(policyKeeper),
		cante.NewValidateBasicDecorator(),
		cante.NewValidateMemoDecorator(ak),
		cante.NewConsumeGasForTxSizeDecorator(ak),
		cante.NewSetPubKeyDecorator(ak), // SetPubKeyDecorator must be called before all signature verification decorators
		cante.NewValidateSigCountDecorator(ak),
		cante.NewDeductFeeDecorator(ak, supplyKeeper),
		cante.NewSigGasConsumeDecorator(ak, sigGasConsumer),
		cante.NewSigVerificationDecorator(ak),
		cante.NewIncrementSequenceDecorator(ak), // innermost AnteDecorator
	)
}
