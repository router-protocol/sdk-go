package types

import (
	"encoding/binary"

	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// PriceKeyPrefix is the prefix to retrieve all Price
	PriceKeyPrefix = "Price/value/"
)

// PriceKey returns the store key to retrieve a Price from the index fields
func PriceKey(
	priceFeederName string,
	symbol string,
) []byte {
	return util.AppendBytes([]byte(priceFeederName), []byte(symbol))
}
