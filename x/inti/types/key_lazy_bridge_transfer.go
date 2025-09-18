package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// LazyBridgeTransferKeyPrefix is the prefix to retrieve all LazyBridgeTransfer
	LazyBridgeTransferKeyPrefix = "LazyBridgeTransfer/value/"
)

// LazyBridgeTransferKey returns the store key to retrieve a LazyBridgeTransfer from the index fields
func LazyBridgeTransferKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
