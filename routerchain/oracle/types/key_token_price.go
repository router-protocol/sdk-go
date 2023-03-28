package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// TokenPriceKeyPrefix is the prefix to retrieve all TokenPrice
	TokenPriceKeyPrefix = "TokenPrice/value/"
)

// TokenPriceKey returns the store key to retrieve a TokenPrice from the index fields
func TokenPriceKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
