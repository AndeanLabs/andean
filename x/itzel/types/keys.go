package types

const (
	// ModuleName defines the module name
	ModuleName = "itzel"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_itzel"
)

var (
	ParamsKey = []byte("p_itzel")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
