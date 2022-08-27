package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
)

func NewOutgoingBatchTx(destinationChainType multichainTypes.ChainType,
	destinationChainId string, sourceAddress string, batchNonce uint64, isAtomic bool,
	contractCalls []ContractCall, relayerFee sdk.Coin, outgoingTxFee sdk.Coin) *OutgoingBatchTx {
	return &OutgoingBatchTx{
		DestinationChainType: destinationChainType,
		DestinationChainId:   destinationChainId,
		SourceAddress:        sourceAddress,
		Nonce:                batchNonce,
		IsAtomic:             isAtomic,
		ContractCalls:        contractCalls,
		RelayerFee:           relayerFee,
		OutgoingTxFee:        outgoingTxFee,
		Status:               OUTGOING_TX_CREATED,
	}
}
