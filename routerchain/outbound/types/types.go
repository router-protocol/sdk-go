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

func NewOutboundAck(chainType multichainTypes.ChainType, chainId string, eventNonce uint64, blockHeight uint64,
	outboundTxNonce uint64, outboundTxRequestedBy string, relayer string, destinationTxHash string, contractAckResponses []*ContractAckResponse, feeConsumed uint64,
) *OutboundAck {
	return &OutboundAck{
		ChainType:             chainType,
		ChainId:               chainId,
		EventNonce:            eventNonce,
		BlockHeight:           blockHeight,
		OutboundTxNonce:       outboundTxNonce,
		OutboundTxRequestedBy: outboundTxRequestedBy,
		Relayer:               relayer,
		DestinationTxHash:     destinationTxHash,
		ContractAckResponses:  contractAckResponses,
		FeeConsumed:           feeConsumed,
		Status:                OUTGOING_TX_ACK_RECEIVED,
	}
}
