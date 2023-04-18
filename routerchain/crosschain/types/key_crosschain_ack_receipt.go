package types

import (
	"encoding/binary"

	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// CrosschainAckReceiptKeyPrefix is the prefix to retrieve all CrosschainAckReceipt
	CrosschainAckReceiptKeyPrefix = "CrosschainAckReceipt/value/"
)

// CrosschainAckReceiptKey returns the store key to retrieve a CrosschainAckReceipt from the index fields
func CrosschainAckReceiptKey(
	ackReceiptSrcChainId string,
	ackReceiptIdentifier uint64,
	claimHash []byte,
) []byte {
	crosschainAckReceiptKey := util.AppendBytes([]byte(ackReceiptSrcChainId), util.UInt64Bytes(ackReceiptIdentifier), claimHash)

	return crosschainAckReceiptKey
}
