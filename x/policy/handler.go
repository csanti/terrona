package policy

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/csanti/terrona/x/policy/keeper"
	"github.com/csanti/terrona/x/policy/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

func NewPolicyUpdateHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case types.TaxRateUpdateProposal:
			return handleTaxRateUpdateProposal(ctx, k, c)
		case types.MinStakeUpdateProposal:
			return handleMinStakeProposal(ctx, k, c)
		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized proposal content type: %T", c)
		}
	}
}

// handleTaxRateUpdateProposal is a handler for updating tax rate
func handleTaxRateUpdateProposal(ctx sdk.Context, k keeper.Keeper, p types.TaxRateUpdateProposal) error {
	fmt.Println("handleTaxRateUpdateProposal")
	newTaxRate := p.TaxRate
	k.SetTaxRate(ctx, newTaxRate)
	fmt.Println(k.GetTaxRate(ctx))
	return nil
}

// handleRewardWeightUpdateProposal is a handler for updating reward-weight
func handleMinStakeProposal(ctx sdk.Context, k keeper.Keeper, p types.MinStakeUpdateProposal) error {
	fmt.Println("handleMinStakeProposal")
	newMinStake := p.MinStake
	k.SetMinStake(ctx, newMinStake)
	return nil
}