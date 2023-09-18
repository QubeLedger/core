package types

const (
	// ModuleName defines the module name
	ModuleName = "grow"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_grow"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	LoanKey                    = "Loan/value/"
	LoanCountKey               = "Loan/count/"
	PairKey                    = "Pair/value/"
	PairCountKey               = "Pair/count/"
	LendAssetKey               = "LendAsset/value/"
	LendAssetCountKey          = "LendAsset/count/"
	PositionKey                = "Position/value/"
	PositionCountKey           = "Position/count/"
	LiquidatorPositionKey      = "LiquidatorPosition/value/"
	LiquidatorPositionCountKey = "LiquidatorPosition/count/"
)
