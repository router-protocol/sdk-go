package types

import (
	"encoding/binary"
	fmt "fmt"

	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// CrossTalkAckReceiptKeyPrefix is the prefix to retrieve all CrossTalkAckReceipt
	CrossTalkAckReceiptKeyPrefix = "CrossTalkAckReceipt/value/"
)

// CrossTalkAckReceiptKey returns the store key to retrieve a CrossTalkAckReceipt from the index fields
func CrossTalkAckReceiptKey(
	chainType multichainTypes.ChainType,
	chainId string,
	eventNonce uint64,
	claimHash []byte,
) []byte {
	crossTalkAckReceiptKey := util.AppendBytes(util.UInt64Bytes(uint64(chainType)), []byte(chainId), util.UInt64Bytes(eventNonce), claimHash)
	fmt.Println("crossTalkAckReceiptKey", crossTalkAckReceiptKey)
	return crossTalkAckReceiptKey
}
