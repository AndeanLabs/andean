package types

const (
	// ModuleName defines the module name
	ModuleName = "inti"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_inti"
)

var (
	ParamsKey = []byte("p_inti")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
