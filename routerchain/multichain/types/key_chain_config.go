package types

import (
	"encoding/binary"

	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// ChainConfigKeyPrefix is the prefix to retrieve all ChainConfig
	ChainConfigKeyPrefix = "ChainConfig/value/"
)

// ChainConfigKey returns the store key to retrieve a ChainConfig from the index fields
func ChainConfigKey(
	chainId uint64,
) []byte {
	return util.AppendBytes(util.HashString(ChainConfigKeyPrefix), util.UInt64Bytes(chainId))
}
