package types

import (
	"encoding/binary"

	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// IncomingTxKeyPrefix is the prefix to retrieve all IncomingTx
	IncomingTxKeyPrefix = "IncomingTx/value/"
)

// IncomingTxKey returns the store key to retrieve a IncomingTx from the index fields
func IncomingTxKey(
	chainId uint64,
	eventNonce uint64,
	claimHash []byte,
) []byte {
	return util.AppendBytes(util.UInt64Bytes(chainId), util.UInt64Bytes(eventNonce), claimHash)
}
