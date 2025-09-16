package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// AggregatedPriceKeyPrefix is the prefix to retrieve all AggregatedPrice
	AggregatedPriceKeyPrefix = "AggregatedPrice/value/"
)

// AggregatedPriceKey returns the store key to retrieve a AggregatedPrice from the index fields
func AggregatedPriceKey(
	source string,
) []byte {
	var key []byte

	sourceBytes := []byte(source)
	key = append(key, sourceBytes...)
	key = append(key, []byte("/")...)

	return key
}
