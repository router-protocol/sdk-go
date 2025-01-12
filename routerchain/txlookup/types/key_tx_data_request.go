package types

import (
	"encoding/binary"

	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// TxDataRequestKeyPrefix is the prefix to retrieve all tx data requests
	TxDataRequestKeyPrefix = "TxDataRequest/value/"
)

// TxDataRequestKey returns the store key to retrieve a transaction data request from the index fields
func TxDataRequestKey(
	chainId string,
	txHash string,
	claimHash []byte,
) []byte {
	return util.AppendBytes([]byte(chainId), []byte(txHash), claimHash)
}
