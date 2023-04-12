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
	chainId string,
	eventNonce uint64,
	claimHash []byte,
) []byte {
	crosschainAckReceiptKey := util.AppendBytes([]byte(chainId), util.UInt64Bytes(eventNonce), claimHash)

	return crosschainAckReceiptKey
}
