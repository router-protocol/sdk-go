package types

import (
	"encoding/binary"

	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// ContractConfigKeyPrefix is the prefix to retrieve all ContractConfig
	ContractConfigKeyPrefix = "ContractConfig/value/"

	// LastObservedEventNonceKey indexes the latest event nonce
	// [0xa34e56ab6fab9ee91e82ba216bfeb759]
	LastObservedEventNonceKey = "LastObservedEventNonceKey"
)

// ContractConfigKey returns the store key to retrieve a ContractConfig from the index fields
func ContractConfigKey(
	chainId string,
	contract string,
) []byte {
	return util.AppendBytes([]byte(chainId), []byte(contract))
}

// GetLastObservedEventNonceKey returns the store key to retrieve a LastObservedEventNonce from the index fields
func GetLastObservedEventNonceKey(
	chainId string,
	contract string,
) []byte {
	return util.AppendBytes([]byte(chainId), []byte(contract))
}
