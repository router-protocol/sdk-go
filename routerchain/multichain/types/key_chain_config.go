package types

import (
	"encoding/binary"
)

var _ binary.ByteOrder

const (
	// ChainConfigKeyPrefix is the prefix to retrieve all ChainConfig
	ChainConfigKeyPrefix = "ChainConfig/value/"
)

// ChainConfigKey returns the store key to retrieve a ChainConfig from the index fields
func ChainConfigKey(
	chainId string,
) []byte {
	return []byte(chainId)
}
