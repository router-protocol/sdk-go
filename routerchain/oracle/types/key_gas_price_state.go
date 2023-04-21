package types

import (
	"encoding/binary"
)

var _ binary.ByteOrder

const (
	// GasPriceStateKeyPrefix is the prefix to retrieve all GasPriceState
	GasPriceStateKeyPrefix = "GasPriceState/value/"
)

// GasPriceStateKey returns the store key to retrieve a GasPriceState from the index fields
func GasPriceStateKey(
	chainId string,
) []byte {

	gasPriceStateKey := []byte(chainId)
	return gasPriceStateKey
}
