package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
)

func NewOutgoingBatchTx(destinationChainType multichainTypes.ChainType,
	destinationChainId string, sourceAddress sdk.AccAddress, batchNonce uint64, isAtomic bool,
	contractCalls []ContractCall, relayerFee sdk.Coin, outgoingTxFee sdk.Coin, expiryTimestamp int64) *OutgoingBatchTx {
	return &OutgoingBatchTx{
		DestinationChainType: destinationChainType,
		DestinationChainId:   destinationChainId,
		SourceAddress:        sourceAddress.String(),
		Nonce:                batchNonce,
		IsAtomic:             isAtomic,
		ContractCalls:        contractCalls,
		RelayerFee:           relayerFee,
		OutgoingTxFee:        outgoingTxFee,
		ExpiryTimestamp:      expiryTimestamp,
		Status:               OUTGOING_TX_CREATED,
	}
}

func NewOutgoingBatchConfirm(destinationChainType multichainTypes.ChainType,
	destinationChainId string, batchSender string, batchNonce uint64, ethSigner string, signature string, orchestrator string) *OutgoingBatchConfirm {
	return &OutgoingBatchConfirm{
		DestinationChainType: destinationChainType,
		DestinationChainId:   destinationChainId,
		OutgoingBatchSender:  batchSender,
		OutgoingBatchNonce:   batchNonce,
		EthSigner:            ethSigner,
		Signature:            signature,
		Orchestrator:         orchestrator,
	}
}

func NewOutboundAck(chainType multichainTypes.ChainType, chainId string, eventNonce uint64, blockHeight uint64,
	outboundTxNonce uint64, outboundTxRequestedBy string, relayerRouterAddress string, destinationTxHash string, contractAckResponses []bool, feeConsumed uint64,
) *OutboundAck {
	return &OutboundAck{
		ChainType:             chainType,
		ChainId:               chainId,
		EventNonce:            eventNonce,
		BlockHeight:           blockHeight,
		OutboundTxNonce:       outboundTxNonce,
		OutboundTxRequestedBy: outboundTxRequestedBy,
		RelayerRouterAddress:  relayerRouterAddress,
		DestinationTxHash:     destinationTxHash,
		ContractAckResponses:  contractAckResponses,
		FeeConsumed:           feeConsumed,
		Status:                OUTGOING_TX_ACK_RECEIVED,
	}
}
