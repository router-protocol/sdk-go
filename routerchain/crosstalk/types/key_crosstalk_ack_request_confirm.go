package types

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// CrosstalkAckRequestConfirmKeyPrefix is the prefix to retrieve all CrosstalkAckRequestConfirm
	CrosstalkAckRequestConfirmKeyPrefix = "CrosstalkAckRequestConfirm/value/"
)

// CrosstalkAckRequestConfirmKey returns the store key to retrieve a CrosstalkAckRequestConfirm from the index fields
func CrosstalkAckRequestConfirmKey(
	chainType multichainTypes.ChainType,
	chainId string,
	eventNonce uint64,
	claimHash []byte,
	orchestrator sdk.AccAddress,
) []byte {
	return util.AppendBytes(util.UInt64Bytes(uint64(chainType)), []byte(chainId), util.UInt64Bytes(eventNonce), claimHash, orchestrator.Bytes())
}
