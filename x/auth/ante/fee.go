package ante

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"
	"github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank"

	assets "github.com/csanti/terrona/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	cleank "github.com/csanti/terrona/x/policy/keeper"
)

var (
	_ FeeTx = (*types.StdTx)(nil) // assert StdTx implements FeeTx
)

// FeeTx defines the interface to be implemented by Tx to use the FeeDecorators
type FeeTx interface {
	sdk.Tx
	GetGas() uint64
	GetFee() sdk.Coins
	FeePayer() sdk.AccAddress
}

// MempoolFeeDecorator will check if the transaction's fee is at least as large
// as the local validator's minimum gasFee (defined in validator config).
// If fee is too low, decorator returns error and tx is rejected from mempool.
// Note this only applies when ctx.CheckTx = true
// If fee is high enough or not CheckTx, then call next AnteHandler
// CONTRACT: Tx must implement FeeTx to use MempoolFeeDecorator
type MempoolFeeDecorator struct{
	policyKeeper cleank.Keeper
}

func NewMempoolFeeDecorator(policyKeeper cleank.Keeper) MempoolFeeDecorator {
	return MempoolFeeDecorator{
		policyKeeper: policyKeeper,
	}
}

func (mfd MempoolFeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	feeTx, ok := tx.(FeeTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
	}
	feeCoins := feeTx.GetFee()
	
	gas := feeTx.GetGas()
	totalDue := sdk.ZeroInt()
	if !simulate {
		// gas check
		if ctx.IsCheckTx() {

			minGasPrices := ctx.MinGasPrices()
			if !minGasPrices.IsZero() {
				requiredFees := make(sdk.Coins, len(minGasPrices))

				// Determine the required fees by multiplying each required minimum gas
				// price by the gas limit, where fee = ceil(minGasPrice * gasLimit).
				glDec := sdk.NewDec(int64(gas))

				for i, gp := range minGasPrices {
					// mandatory denom to pay gas
					if gp.Denom != assets.MicroTpaperDenom {
						continue
					}
					fee := gp.Amount.Mul(glDec)
					requiredFees[i] = sdk.NewCoin(gp.Denom, fee.Ceil().RoundInt())
				}
				//fmt.Println(requiredFees)
				if !feeCoins.IsAnyGTE(requiredFees) {
					return ctx, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFee, "insufficient fees; got: %s required: %s", feeCoins, requiredFees)
				}
				//totalDue = totalDue.Add(requiredFees)
			}
		}

		// tax check
		// Get the amount of coins being sent, and calculate the tax accordingly to check if the fee is enough
		// txValue contains the toal amount of any coins in the transaction
		txValue := sdk.ZeroInt()
		for _, msg := range feeTx.GetMsgs() {
			switch msg := msg.(type) {
			case bank.MsgSend:
				// contains 
				for _, am := range msg.Amount {
					txValue = txValue.Add(am.Amount)
				}		
			// should add msgmultisend too..
			}
		}
		fmt.Println("Total value: ",txValue.Int64())
		// if transaction contains value transfer, need to check if it provided enough fees
		if !txValue.IsZero() {
			taxRate := mfd.policyKeeper.GetTaxRate(ctx)
			// fees can only be paid with tpaper
			// ignore other fees
			if len(feeCoins) == 0 || feeCoins[0].Denom != assets.MicroTpaperDenom {
				fmt.Println("Not enough to pay for tax")
				return ctx, sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "not valid coins")
			}

			taxDue := sdk.NewDecFromInt(txValue).Mul(taxRate).TruncateInt()
			fmt.Println("Tax: ",taxDue)
			totalDue = totalDue.Add(taxDue)
			fmt.Println("Total: ", totalDue)
			if totalDue.GT(feeCoins[0].Amount) {
				return ctx, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFee, "insufficient fees")
			}
		}

	}


	return next(ctx, tx, simulate)
}

// DeductFeeDecorator deducts fees from the first signer of the tx
// If the first signer does not have the funds to pay for the fees, return with InsufficientFunds error
// Call next AnteHandler if fees successfully deducted
// CONTRACT: Tx must implement FeeTx interface to use DeductFeeDecorator
type DeductFeeDecorator struct {
	ak           keeper.AccountKeeper
	supplyKeeper types.SupplyKeeper
}

func NewDeductFeeDecorator(ak keeper.AccountKeeper, sk types.SupplyKeeper) DeductFeeDecorator {
	return DeductFeeDecorator{
		ak:           ak,
		supplyKeeper: sk,
	}
}

func (dfd DeductFeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	feeTx, ok := tx.(FeeTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
	}

	if addr := dfd.supplyKeeper.GetModuleAddress(types.FeeCollectorName); addr == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.FeeCollectorName))
	}

	feePayer := feeTx.FeePayer()
	feePayerAcc := dfd.ak.GetAccount(ctx, feePayer)

	if feePayerAcc == nil {
		return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "fee payer address: %s does not exist", feePayer)
	}

	// deduct the fees
	if !feeTx.GetFee().IsZero() {
		err = DeductFees(dfd.supplyKeeper, ctx, feePayerAcc, feeTx.GetFee())
		if err != nil {
			return ctx, err
		}
	}

	return next(ctx, tx, simulate)
}

// DeductFees deducts fees from the given account.
//
// NOTE: We could use the BankKeeper (in addition to the AccountKeeper, because
// the BankKeeper doesn't give us accounts), but it seems easier to do this.
func DeductFees(supplyKeeper types.SupplyKeeper, ctx sdk.Context, acc exported.Account, fees sdk.Coins) error {
	blockTime := ctx.BlockHeader().Time
	coins := acc.GetCoins()

	if !fees.IsValid() {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFee, "invalid fee amount: %s", fees)
	}

	// verify the account has enough funds to pay for fees
	_, hasNeg := coins.SafeSub(fees)
	if hasNeg {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds,
			"insufficient funds to pay for fees; %s < %s", coins, fees)
	}

	// Validate the account has enough "spendable" coins as this will cover cases
	// such as vesting accounts.
	spendableCoins := acc.SpendableCoins(blockTime)
	if _, hasNeg := spendableCoins.SafeSub(fees); hasNeg {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds,
			"insufficient funds to pay for fees; %s < %s", spendableCoins, fees)
	}

	err := supplyKeeper.SendCoinsFromAccountToModule(ctx, acc.GetAddress(), types.FeeCollectorName, fees)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, err.Error())
	}

	return nil
}
