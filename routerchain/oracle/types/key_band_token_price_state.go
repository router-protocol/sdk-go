package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// BandTokenPriceStateKeyPrefix is the prefix to retrieve all BandTokenPriceState
	BandTokenPriceStateKeyPrefix = "BandTokenPriceState/value/"
)

// BandTokenPriceStateKey returns the store key to retrieve a BandTokenPriceState from the symbol fields
func BandTokenPriceStateKey(
	symbol string,
) []byte {
	var key []byte

	symbolBytes := []byte(symbol)
	key = append(key, symbolBytes...)
	key = append(key, []byte("/")...)

	return key
}
