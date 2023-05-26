package types

import (
	"encoding/binary"

	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// GasPriceKeyPrefix is the prefix to retrieve all GasPrice
	GasPriceKeyPrefix = "GasPrice/value/"
)

// GasPriceKey returns the store key to retrieve a GasPrice from the index fields
func GasPriceKey(
	priceFeederName string,
	chainId string,
) []byte {
	return util.AppendBytes([]byte(priceFeederName), []byte(chainId))
}
