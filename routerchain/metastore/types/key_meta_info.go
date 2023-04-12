package types

import (
	"encoding/binary"

	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// MetaInfoKeyPrefix is the prefix to retrieve all MetaInfo
	MetaInfoKeyPrefix = "MetaInfo/value/"
)

// MetaInfoKey returns the store key to retrieve a MetaInfo from the index fields
func MetaInfoKey(
	chainId string,
	dappAddress []byte,
) []byte {

	return util.AppendBytes([]byte(chainId), dappAddress)
}
