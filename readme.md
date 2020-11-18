# Terrona

**Terrona** is a blockchain prototype built with [CosmosSDK](https://github.com/cosmos/cosmos-sdk)

## Native Currencies

Terrona is designed with the concept of managing the most important resources in a world where a virus has taken over.

| Currency | Denomination  | Description  |
| -------- | ------------- |  ------------- |
| TPaper   | utpaper       | x |
| Masks    | umasks        | x |

## Characteristics

- Only Masks can be used as staking to become a validator or delegate
- All fees and rewards will be paid with Tpaper
- To becoma a validator, a node must be holding a minimum amount of Masks
- Fees depend on a taxrate managed by the blockchain
- Tax rate and minimum staking amount can be modified using the governance module

## How to test

[Starport](https://github.com/tendermint/starport) used for bootstraping and genesis management