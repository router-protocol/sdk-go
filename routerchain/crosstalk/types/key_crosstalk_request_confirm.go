package types

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// CrosstalkRequestConfirmKeyPrefix is the prefix to retrieve all CrosstalkRequestConfirm
	CrosstalkRequestConfirmKeyPrefix = "CrosstalkRequestConfirm/value/"
)

// CrosstalkRequestConfirmKey returns the store key to retrieve a CrosstalkRequestConfirm from the index fields
func CrosstalkRequestConfirmKey(
	sourceChainType multichainTypes.ChainType,
	sourceChainId string,
	eventNonce uint64,
	claimHash []byte,
	orchestrator sdk.AccAddress,
) []byte {
	return util.AppendBytes(util.UInt64Bytes(uint64(sourceChainType)), []byte(sourceChainId), util.UInt64Bytes(eventNonce), claimHash, orchestrator.Bytes())
}
