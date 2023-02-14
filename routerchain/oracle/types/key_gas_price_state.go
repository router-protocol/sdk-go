package types

import (
	"encoding/binary"

	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// GasPriceStateKeyPrefix is the prefix to retrieve all GasPriceState
	GasPriceStateKeyPrefix = "GasPriceState/value/"
)

// GasPriceStateKey returns the store key to retrieve a GasPriceState from the index fields
func GasPriceStateKey(
	chainType multichainTypes.ChainType,
	chainId string,
) []byte {

	gasPriceStateKey := util.AppendBytes(util.UInt64Bytes(uint64(chainType)), []byte(chainId))
	return gasPriceStateKey
}
