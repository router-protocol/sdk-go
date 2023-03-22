package types

import (
	"encoding/binary"

	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// MetaInfoKeyPrefix is the prefix to retrieve all MetaInfo
	MetaInfoKeyPrefix = "MetaInfo/value/"
)

// MetaInfoKey returns the store key to retrieve a MetaInfo from the index fields
func MetaInfoKey(
	chainType multichainTypes.ChainType,
	chainId string,
	dappAddress []byte,
) []byte {

	return util.AppendBytes(util.UInt64Bytes(uint64(chainType)), []byte(chainId), dappAddress)
}
