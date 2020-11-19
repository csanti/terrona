# Terrona

**Terrona** is a blockchain prototype built with [CosmosSDK](https://github.com/cosmos/cosmos-sdk)

## Native Currencies

Terrona is designed with the concept of managing scarce resources in a fictional **world where a virus has taken over**.

|  Currency   | Denomination  | Description  	|
| ----------- | ------------- |  ------------- |
| 🧻 TPaper   | utpaper | Native currenty, used to pay tax induced fees, so block rewards are also paid in tpaper.|
| 😷 Masks    | umasks | Staking currency. Must stake more than `MinStake` of this coin to become a validator.	|

## Characteristics

- Only Masks can be used as stake to become a validator or delegate
- All fees and rewards will be paid with *Tpaper*
- To become a validator, a node must be holding a minimum amount of Masks, detrmined by ``MinStake``
- Fees depend on a ``TaxRate`` managed by the blockchain
- ``TaxRate`` can only be paid with *Tpaper*
- ``TaxRate`` and ``MinStake`` can be modified using the governance module, creating and approving update proposals.
- Validators receive block rewards (*Tpaper* fees enforced by the tax), distributed with the distribution module

## How to test basic functionality

[Starport](https://github.com/tendermint/starport) used for bootstraping and genesis management, **required** for building the genesis file.

``config.yml`` defines user accounts and gentx to create the initial validator

Generate genesis based on ``config.yml`` and start the node. It will also build *terronad* and *terronacli*.
```
starport serve -v
```

Starport will generate a genesis containing a `gentx` that will attempt to create a validator. This is only possible if the staking amount is greater than `MinStake` (`DefaultMinStake=1000000umasks`). Only `umasks` is accepted as stake currency.
After Starport generates the initial configuration files, binaries can be recompiled with `make`.

### Useful commands:

Send currency between user accounts:
```
terronacli tx send $(terronacli keys show user2 -a) $(terronacli keys show user3 -a) 1000umasks --fees 10utpaper
```
If the tx does not contain enough fees to cover the tax (`DefaultTaxRate=0.01`) or the fees are in a different currency other than `utpaper`, the tx will not be added to the pool.

Generate governance proposal to update minimum stake:
```
terronacli tx gov submit-proposal min-stake-update min_stake_proposal.json --from $(terronacli keys show user3 -a)
```

Generate governance proposal to update the tax rate:
```
terronacli tx gov submit-proposal tax-rate-update tax_proposal.json --from $(terronacli keys show user3 -a)
```

Check the block rewards paid to the validator in ``utpaper``:
```
terronacli q distribution rewards $(terronacli keys show user1 -a)
```