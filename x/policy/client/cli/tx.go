package cli

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/cosmos/cosmos-sdk/x/gov"

	"github.com/csanti/terrona/x/policy/types"
)

// GetCmdSubmitTaxRateUpdateProposal implements the command to submit a tax-rate-update proposal
func GetCmdSubmitTaxRateUpdateProposal(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tax-rate-update [proposal-file]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a tax rate update proposal",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Submit a tax rate update proposal along with an initial deposit.
The proposal details must be supplied via a JSON file.
Example:
$ %s tx gov submit-proposal tax-rate-update <path/to/proposal.json> --from=<key_or_address>
Where proposal.json contains:
{
  "title": "Update Tax Rate",
  "description": "Lets update tax rate to 1.5%%",
  "tax_rate": "0.015",
  "deposit": [
    {
      "denom": "stake",
      "amount": "10000"
    }
  ]
}
`,
				version.ClientName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			proposal, err := ParseTaxRateUpdateProposalJSON(cdc, args[0])
			if err != nil {
				return err
			}

			from := cliCtx.GetFromAddress()
			content := types.NewTaxRateUpdateProposal(proposal.Title, proposal.Description, proposal.TaxRate)

			msg := gov.NewMsgSubmitProposal(content, proposal.Deposit, from)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}

func GetCmdSubmitMinStakeUpdateProposal(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "min-stake-update [proposal-file]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a minimum stake update proposal",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Submit a minimum stake update proposal along with an initial deposit.
The proposal details must be supplied via a JSON file.
Example:
$ %s tx gov submit-proposal min-stake-update <path/to/proposal.json> --from=<key_or_address>
Where proposal.json contains:
{
  "title": "Update Tax Rate",
  "description": "Lets update tax rate to 1.5%%",
  "min_stake": {
  	"denom": "umasks",
  	"amount": "1000000"
  },
  "deposit": [
    {
      "denom": "utpaper",
      "amount": "10000"
    }
  ]
}
`,
				version.ClientName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			proposal, err := ParseMinStakeUpdateProposalJSON(cdc, args[0])
			if err != nil {
				return err
			}

			from := cliCtx.GetFromAddress()
			content := types.NewMinStakeUpdateProposal(proposal.Title, proposal.Description, proposal.MinStake)

			msg := gov.NewMsgSubmitProposal(content, proposal.Deposit, from)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}
