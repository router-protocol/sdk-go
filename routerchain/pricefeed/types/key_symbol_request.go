package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// SymbolRequestKeyPrefix is the prefix to retrieve all SymbolRequest
	SymbolRequestKeyPrefix = "SymbolRequest/value/"
)

// SymbolRequestKey returns the store key to retrieve a SymbolRequest from the index fields
func SymbolRequestKey(
	symbol string,
) []byte {
	var key []byte = []byte(symbol)
	return key
}
