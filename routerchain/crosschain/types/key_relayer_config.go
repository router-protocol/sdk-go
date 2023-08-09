package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// RelayerConfigKeyPrefix is the prefix to retrieve all RelayerConfig
	RelayerConfigKeyPrefix = "RelayerConfig/value/"
)

// RelayerConfigKey returns the store key to retrieve a RelayerConfig from the index fields
func RelayerConfigKey(
	chainId string,
) []byte {
	var key []byte

	chainIdBytes := []byte(chainId)
	key = append(key, chainIdBytes...)
	key = append(key, []byte("/")...)

	return key
}
