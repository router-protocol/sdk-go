package types

import (
	"encoding/binary"

	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// IncomingTxKeyPrefix is the prefix to retrieve all IncomingTx
	IncomingTxKeyPrefix = "IncomingTx/value/"
)

// IncomingTxKey returns the store key to retrieve a IncomingTx from the index fields
func IncomingTxKey(
	chainType multichainTypes.ChainType,
	chainId string,
	eventNonce uint64,
	claimHash []byte,
) []byte {
	return util.AppendBytes(util.UInt64Bytes(uint64(chainType)), []byte(chainId), util.UInt64Bytes(eventNonce), claimHash)
}
