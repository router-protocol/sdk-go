package types

import (
	"encoding/binary"

	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// OutgoingBatchTxKeyPrefix is the prefix to retrieve all OutgoingBatchTx
	OutgoingBatchTxKeyPrefix = "OutgoingBatchTx/value/"

	// LatestOutgoingBatchNonce indexes the latest outgoing batch nonce of source address
	LatestOutgoingBatchNoncePrefix = "OutgoingBatchTx/LatestOutgoingBatchNonce"
)

var ()

// OutgoingBatchTxKey returns the store key to retrieve a OutgoingBatchTx from the index fields
func OutgoingBatchTxKey(
	chainType multichainTypes.ChainType,
	chainId string,
	sourceAddress string,
	batchNonce uint64,
) []byte {
	return util.AppendBytes(util.UInt64Bytes(uint64(chainType)), []byte(chainId), []byte(sourceAddress), util.UInt64Bytes(batchNonce))
}

// LatestOutgoingBatchNonceKey returns the store key to retrieve a LatestOutgoingBatchNonceKey from the index fields
func LatestOutgoingBatchNonceKey(
	chainType multichainTypes.ChainType,
	chainId string,
	sourceAddress string,
) []byte {
	return util.AppendBytes(util.UInt64Bytes(uint64(chainType)), []byte(chainId), []byte(sourceAddress))
}
