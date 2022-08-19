package types

import (
	"encoding/binary"

	"github.com/router-protocol/router-chain/util"
)

var _ binary.ByteOrder

const (
	// ChainConfigKeyPrefix is the prefix to retrieve all ChainConfig
	ChainConfigKeyPrefix = "ChainConfig/value/"

	// LastObservedEventNonceKey indexes the latest event nonce
	// [0xa34e56ab6fab9ee91e82ba216bfeb759]
	LastObservedEventNonceKey = "LastObservedEventNonceKey"
)

// ChainConfigKey returns the store key to retrieve a ChainConfig from the index fields
func ChainConfigKey(
	chainType ChainType,
	chainId string,
) []byte {
	return util.AppendBytes(util.UInt64Bytes(uint64(chainType)), []byte(chainId))
}

// GetLastObservedEventNonceKey returns the store key to retrieve a LastObservedEventNonce from the index fields
func GetLastObservedEventNonceKey(
	chainType ChainType,
	chainId string,
) []byte {
	return util.AppendBytes(util.UInt64Bytes(uint64(chainType)), []byte(chainId))
}
