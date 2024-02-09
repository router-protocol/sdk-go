package types

import (
	"encoding/binary"

	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// NonceObservedStatusKeyPrefix is the prefix to retrieve all NonceObservedStatus
	NonceObservedStatusKeyPrefix = "NonceObservedStatus/value/"
)

// NonceObservedStatusKey returns the store key to retrieve a NonceObservedStatus from the index fields
func NonceObservedStatusKey(
	chainId string,
	contract string,
	eventNonce uint64,
) []byte {
	return util.AppendBytes([]byte(chainId), []byte(contract), util.UInt64Bytes(eventNonce))
}
