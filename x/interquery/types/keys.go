package types

const (
	// ModuleName defines the module name.
	ModuleName = "interquery"

	// StoreKey defines the primary module store key.
	StoreKey = ModuleName

	// RouterKey is the message route for interquery.
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key.
	QuerierRoute = ModuleName
)

// prefix bytes for the interquery persistent store.
const (
	prefixData         = iota + 1
	prefixQuery        = iota + 1
	prefixLatestHeight = iota + 1
)

var (
	KeyPrefixData         = []byte{prefixData}
	KeyPrefixQuery        = []byte{prefixQuery}
	KeyPrefixLatestHeight = []byte{prefixLatestHeight}
)
