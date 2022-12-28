package types

import (
	"encoding/binary"

	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// CrossTalkAckRequestKeyPrefix is the prefix to retrieve all CrossTalkAckRequest
	CrossTalkAckRequestKeyPrefix = "CrossTalkAckRequest/value/"
)

// CrossTalkAckRequestKey returns the store key to retrieve a CrossTalkAckRequest from the index fields
func CrossTalkAckRequestKey(
	chainType multichainTypes.ChainType,
	chainId string,
	eventNonce uint64,
	claimHash []byte,
) []byte {
	crossTalkAckRequestKey := util.AppendBytes(util.UInt64Bytes(uint64(chainType)), []byte(chainId), util.UInt64Bytes(eventNonce), claimHash)
	return crossTalkAckRequestKey
}
