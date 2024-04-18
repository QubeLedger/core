package types

const (
	// ModuleName defines the module name
	ModuleName = "stable"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_stable"

	// Version defines the current version the IBC module supports
	Version = "stable-1"

	// PortID is the default port id that module binds to
	PortID = "stable"

	SystemModuleAccount = ModuleName + "_system_account"
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("stable-port-")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	PairKey      = "Pair/value/"
	PairCountKey = "Pair/count/"
)
