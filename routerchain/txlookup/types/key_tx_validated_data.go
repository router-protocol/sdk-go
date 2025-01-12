package types

import (
	"encoding/binary"

	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// ValidatedTxDataRequestKeyPrefix is the prefix to retrieve all validated tx data requests
	ValidatedTxDataRequestKeyPrefix = "ValidatedTxDataRequest/value/"
)

// ValidatedTxDataRequestKey returns the store key to retrieve a transaction data request from the index fields
func ValidatedTxDataRequestKey(
	chainId string,
	txHash string,
) []byte {
	return util.AppendBytes([]byte(chainId), []byte(txHash))
}
