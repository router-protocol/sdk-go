package types

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
	multichainTypes "github.com/router-protocol/router-chain/x/multichain/types"
	"github.com/router-protocol/sdk-go/routerchain/util"
)

var _ binary.ByteOrder

const (
	// OutgoingBatchConfirmKeyPrefix is the prefix to retrieve all OutgoingBatchConfirm
	OutgoingBatchConfirmKeyPrefix = "OutgoingBatchConfirm/value/"
)

// OutgoingBatchConfirmKey returns the store key to retrieve a OutgoingBatchConfirm from the index fields
func OutgoingBatchConfirmKey(
	chainType multichainTypes.ChainType,
	chainId string,
	sourceAddress string,
	nonce uint64,
	orchestrator sdk.AccAddress,
) []byte {
	return util.AppendBytes(util.UInt64Bytes(uint64(chainType)), []byte(chainId), []byte(sourceAddress), util.UInt64Bytes(nonce), orchestrator.Bytes())
}
