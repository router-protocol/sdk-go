package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
)

const (
	TypeOutgoingBatchTxCreated    = "outgoing_batch_tx_created"
	AttributeOutgoingTxNonce      = "outgoing_tx_nonce"
	AttributeDestinationChainType = "destinationChainType"
	AttributeDestinationChainId   = "destinationChainId"
	AttributeContractCalls        = "contractCalls"
	AttributeRelayerFee           = "relayerFee"
	AttributeOutgoingTxFee        = "outgoingTxFee"
	AttributeIsAtomic             = "isAtomic"
	AttributeSourceAddress        = "sourceAddress"
	AttributeExpiryTimestamp      = "expiryTimestamp"
	AttributeStatus               = "status"
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

func NewOutboundAck(chainType multichainTypes.ChainType,
	chainId string,
	eventNonce uint64,
	blockHeight uint64,
	outboundTxNonce uint64,
	outboundTxRequestedBy string,
	relayerRouterAddress string,
	destinationTxHash string,
	feeConsumed uint64,
	contractAckResponses []byte,
	exeCode uint64,
	execStatus bool,
	execFlags []bool,
	execData [][]byte,
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
		FeeConsumed:           feeConsumed,
		ContractAckResponses:  contractAckResponses,
		ExeCode:               exeCode,
		ExecStatus:            execStatus,
		ExecFlags:             execFlags,
		ExecData:              execData,
		Status:                OUTGOING_TX_ACK_RECEIVED,
	}
}
