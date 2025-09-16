package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// PriceReportKeyPrefix is the prefix to retrieve all PriceReport
	PriceReportKeyPrefix = "PriceReport/value/"
)

// PriceReportKey returns the store key to retrieve a PriceReport from the index fields
func PriceReportKey(
	source string,
	oracle string,
) []byte {
	var key []byte

	sourceBytes := []byte(source)
	key = append(key, sourceBytes...)
	key = append(key, []byte("/")...)

	oracleBytes := []byte(oracle)
	key = append(key, oracleBytes...)
	key = append(key, []byte("/")...)

	return key
}
