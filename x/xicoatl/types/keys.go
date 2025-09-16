package types

const (
	// ModuleName defines the module name
	ModuleName = "xicoatl"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_xicoatl"
)

var (
	ParamsKey = []byte("p_xicoatl")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
