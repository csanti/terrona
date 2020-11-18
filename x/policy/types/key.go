package types

const (
	// ModuleName is the name of the module
	ModuleName = "policy"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// RouterKey to be used for routing msgs
	RouterKey = ModuleName

	// QuerierRoute to be used for querier msgs
	QuerierRoute = ModuleName
)

var (
	// Keys for store prefixes
	TaxRateKey              = []byte{0x01} 
	MinStakeKey         	= []byte{0x02} 
)
