package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
	routerchaintypes "github.com/router-protocol/sdk-go/routerchain/types"
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
	contractCalls []ContractCall, relayerFee sdk.Coin, destinationGasLimit uint64, destinationGasPrice uint64, outgoingTxFeeInRoute sdk.Coin, chainTimeStamp int64, expiryTimestamp int64, routeAmount sdk.Int, routeRecipient []byte, asmAddress []byte, outboundAckGasLimit uint64, middlewareContractType routerchaintypes.MiddlewareContractType) *OutgoingBatchTx {
	return &OutgoingBatchTx{
		DestinationChainType:   destinationChainType,
		DestinationChainId:     destinationChainId,
		SourceAddress:          sourceAddress.String(),
		Nonce:                  batchNonce,
		IsAtomic:               isAtomic,
		ContractCalls:          contractCalls,
		RelayerFee:             relayerFee,
		DestinationGasLimit:    destinationGasLimit,
		DestinationGasPrice:    destinationGasPrice,
		OutgoingTxFeeInRoute:   outgoingTxFeeInRoute,
		OutboundAckGasLimit:    outboundAckGasLimit,
		ChainTimestamp:         uint64(chainTimeStamp),
		ExpiryTimestamp:        expiryTimestamp,
		RouteAmount:            routeAmount,
		RouteRecipient:         routeRecipient,
		AsmAddress:             asmAddress,
		Status:                 OUTGOING_TX_CREATED,
		MiddlewareContractType: middlewareContractType,
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
	feeConsumedInRoute *sdk.Coin,
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
		FeeConsumedInRoute:    feeConsumedInRoute,
		ContractAckResponses:  contractAckResponses,
		ExeCode:               exeCode,
		ExecStatus:            execStatus,
		ExecFlags:             execFlags,
		ExecData:              execData,
		Status:                OUTGOING_TX_ACK_RECEIVED,
	}
}

func NewOutboundAckClaimHash(
	chainType multichainTypes.ChainType,
	chainId string,
	eventNonce uint64,
	blockHeight uint64,
	outboundTxNonce uint64,
	outboundTxRequestedBy string,
	relayerRouterAddress string,
	destinationTxHash string,
	contractAckResponses []byte,
	exeCode uint64,
	execStatus bool,
	execFlags []bool,
	execData [][]byte,
) *OutboundAckClaimHash {
	return &OutboundAckClaimHash{
		ChainType:             chainType,
		ChainId:               chainId,
		EventNonce:            eventNonce,
		BlockHeight:           blockHeight,
		OutboundTxNonce:       outboundTxNonce,
		OutboundTxRequestedBy: outboundTxRequestedBy,
		RelayerRouterAddress:  relayerRouterAddress,
		DestinationTxHash:     destinationTxHash,
		ContractAckResponses:  contractAckResponses,
		ExeCode:               exeCode,
		ExecStatus:            execStatus,
		ExecFlags:             execFlags,
		ExecData:              execData,
	}
}
