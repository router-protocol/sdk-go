package types

import (
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"

	sdkmath "cosmossdk.io/math"
)

func NewCrosschainRequest(
	srcChainId string,
	contract string,
	requestIdentifier uint64,
	srcBlockHeight uint64,
	sourceTxHash string,
	srcTimestamp uint64,
	srcTxOrigin string,
	routeAmount sdkmath.Int,
	routeRecipient string,
	destChainId string,
	requestSender string,
	requestMetadata []byte,
	requestPacket []byte,
	srcChainType multichainTypes.ChainType,
	destChainType multichainTypes.ChainType,
) *CrosschainRequest {
	return &CrosschainRequest{
		SrcChainId:        srcChainId,
		Contract:          contract,
		RequestIdentifier: requestIdentifier,
		BlockHeight:       srcBlockHeight,
		SourceTxHash:      sourceTxHash,
		SrcTimestamp:      srcTimestamp,
		SrcTxOrigin:       srcTxOrigin,
		RouteAmount:       routeAmount,
		RouteRecipient:    routeRecipient,
		DestChainId:       destChainId,
		RequestSender:     requestSender,
		RequestMetadata:   requestMetadata,
		RequestPacket:     requestPacket,
		SrcChainType:      srcChainType,
		DestChainType:     destChainType,
		Status:            CROSSCHAIN_TX_CREATED,
	}
}

func NewCrosschainRequestFromMsg(
	msg *MsgCrosschainRequest,
) *CrosschainRequest {
	return &CrosschainRequest{
		SrcChainId:        msg.SrcChainId,
		Contract:          msg.Contract,
		RequestIdentifier: msg.RequestIdentifier,
		BlockHeight:       msg.BlockHeight,
		SourceTxHash:      msg.SourceTxHash,
		SrcTimestamp:      msg.SrcTimestamp,
		SrcTxOrigin:       msg.SrcTxOrigin,
		RouteAmount:       msg.RouteAmount,
		RouteRecipient:    msg.RouteRecipient,
		DestChainId:       msg.DestChainId,
		RequestSender:     msg.RequestSender,
		RequestMetadata:   msg.RequestMetadata,
		RequestPacket:     msg.RequestPacket,
		SrcChainType:      msg.SrcChainType,
		DestChainType:     msg.DestChainType,
		Status:            CROSSCHAIN_TX_CREATED,
	}
}

func NewCrosschainRequestClaimHash(
	srcChainId string,
	contract string,
	requestIdentifier uint64,
	blockHeight uint64,
	sourceTxHash string,
	srcTimestamp uint64,
	srcTxOrigin string,
	routeAmount sdkmath.Int,
	routeRecipient string,
	destChainId string,
	requestSender string,
	requestMetadata []byte,
	requestPacket []byte,
	srcChainType multichainTypes.ChainType,
	destChainType multichainTypes.ChainType,
) *CrosschainRequestClaimHash {
	return &CrosschainRequestClaimHash{
		SrcChainId:        srcChainId,
		Contract:          contract,
		RequestIdentifier: requestIdentifier,
		BlockHeight:       blockHeight,
		SourceTxHash:      sourceTxHash,
		SrcTimestamp:      srcTimestamp,
		SrcTxOrigin:       srcTxOrigin,
		RouteAmount:       routeAmount,
		RouteRecipient:    routeRecipient,
		DestChainId:       destChainId,
		RequestSender:     requestSender,
		RequestMetadata:   requestMetadata,
		RequestPacket:     requestPacket,
		SrcChainType:      srcChainType,
		DestChainType:     destChainType,
	}
}

func NewCrosschainRequestConfirm(sourceChainId string, requestIdentifier uint64, claimHash []byte, orchestrator string, ethSigner string, signature string) *CrosschainRequestConfirm {
	return &CrosschainRequestConfirm{
		SourceChainId:     sourceChainId,
		RequestIdentifier: requestIdentifier,
		ClaimHash:         claimHash,
		Orchestrator:      orchestrator,
		EthSigner:         ethSigner,
		Signature:         signature,
	}
}

func NewCrosschainAckRequestFromMsg(
	msg *MsgCrosschainAckRequest,
) *CrosschainAckRequest {
	return &CrosschainAckRequest{
		AckSrcChainId:        msg.AckSrcChainId,
		Contract:             msg.Contract,
		AckRequestIdentifier: msg.AckRequestIdentifier,
		BlockHeight:          msg.BlockHeight,
		DestTxHash:           msg.DestTxHash,
		RelayerRouterAddress: msg.RelayerRouterAddress,
		AckDestChainId:       msg.AckDestChainId,
		RequestSender:        msg.RequestSender,
		RequestIdentifier:    msg.RequestIdentifier,
		ExecData:             msg.ExecData,
		ExecStatus:           msg.ExecStatus,
		EthSigner:            msg.EthSigner,
		Signature:            msg.Signature,
		FeeConsumed:          msg.FeeConsumed,
		AckSrcChainType:      msg.AckSrcChainType,
		AckDestChainType:     msg.AckDestChainType,

		Status: CROSSCHAIN_ACK_TX_CREATED,
	}
}

func NewCrosschainAckRequestClaimHash(
	ackSrcChainId string,
	contract string,
	ackRequestIdentifier uint64,
	blockHeight uint64,
	destTxHash string,
	relayerRouterAddress string,
	ackDestChainId string,
	requestSender string,
	requestIdentifier uint64,
	ackSrcChainType multichainTypes.ChainType,
	ackDestChainType multichainTypes.ChainType,
	feeConsumed uint64,
	execData []byte,
	execStatus bool,
) *CrosschainAckRequestClaimHash {
	return &CrosschainAckRequestClaimHash{
		AckSrcChainId:        ackSrcChainId,
		Contract:             contract,
		AckRequestIdentifier: ackRequestIdentifier,
		BlockHeight:          blockHeight,
		DestTxHash:           destTxHash,
		RelayerRouterAddress: relayerRouterAddress,
		AckDestChainId:       ackDestChainId,
		RequestSender:        requestSender,
		RequestIdentifier:    requestIdentifier,
		AckSrcChainType:      ackSrcChainType,
		AckDestChainType:     ackDestChainType,
		FeeConsumed:          feeConsumed,
		ExecData:             execData,
		ExecStatus:           execStatus,
	}
}

func NewCrosschainAckRequestConfirm(ackSrcChainId string, ackRequestIdentifier uint64, claimHash []byte, orchestrator string, ethSigner string, signature string) *CrosschainAckRequestConfirm {
	return &CrosschainAckRequestConfirm{
		AckSrcChainId:        ackSrcChainId,
		AckRequestIdentifier: ackRequestIdentifier,
		ClaimHash:            claimHash,
		Orchestrator:         orchestrator,
		EthSigner:            ethSigner,
		Signature:            signature,
	}
}

func NewCrosschainAckReceiptFromMsg(
	msg *MsgCrosschainAckReceipt,
) *CrosschainAckReceipt {
	return &CrosschainAckReceipt{
		AckReceiptSrcChainId:  msg.AckReceiptSrcChainId,
		Contract:              msg.Contract,
		AckReceiptIdentifier:  msg.AckReceiptIdentifier,
		AckReceiptBlockHeight: msg.AckReceiptBlockHeight,
		AckReceiptTxHash:      msg.AckReceiptTxHash,
		RelayerRouterAddress:  msg.RelayerRouterAddress,
		RequestIdentifier:     msg.RequestIdentifier,
		AckSrcChainId:         msg.AckSrcChainId,
		AckRequestIdentifier:  msg.AckRequestIdentifier,
		AckExecData:           msg.AckExecData,
		AckExecStatus:         msg.AckExecStatus,
		FeeConsumed:           msg.FeeConsumed,
		Status:                CROSSCHAIN_ACK_RECEIPT_TX_CREATED,
	}
}

func NewCrosschainAckReceiptClaimHash(
	ackReceiptSrcChainId string,
	contract string,
	ackReceiptIdentifier uint64,
	ackReceiptBlockHeight uint64,
	ackReceiptTxHash string,
	relayerRouterAddress string,
	requestIdentifier uint64,
	ackSrcChainId string,
	ackRequestIdentifier uint64,
	feeConsumed uint64,
	ackExecData []byte,
	ackExecStatus bool,
) *CrosschainAckReceiptClaimHash {
	return &CrosschainAckReceiptClaimHash{
		AckReceiptSrcChainId:  ackReceiptSrcChainId,
		Contract:              contract,
		AckReceiptIdentifier:  ackReceiptIdentifier,
		AckReceiptBlockHeight: ackReceiptBlockHeight,
		AckReceiptTxHash:      ackReceiptTxHash,
		RelayerRouterAddress:  relayerRouterAddress,
		RequestIdentifier:     requestIdentifier,
		AckSrcChainId:         ackSrcChainId,
		AckRequestIdentifier:  ackRequestIdentifier,
		FeeConsumed:           feeConsumed,
		AckExecData:           ackExecData,
		AckExecStatus:         ackExecStatus,
	}
}
