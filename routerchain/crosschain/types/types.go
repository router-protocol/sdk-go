package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
)

func NewCrosschainRequest(
	srcChainId string,
	srcEventNonce uint64,
	srcBlockHeight uint64,
	sourceTxHash string,
	srcTimestamp uint64,
	srcTxOrigin string,
	routeAmount sdk.Int,
	routeRecipient []byte,
	destChainId string,
	requestSender []byte,
	requestMetadata []byte,
	requestPacket []byte,
	srcChainType multichainTypes.ChainType,
	destChainType multichainTypes.ChainType,
) *CrosschainRequest {
	return &CrosschainRequest{
		SrcChainId:      srcChainId,
		EventNonce:      srcEventNonce,
		BlockHeight:     srcBlockHeight,
		SourceTxHash:    sourceTxHash,
		SrcTimestamp:    srcTimestamp,
		SrcTxOrigin:     srcTxOrigin,
		RouteAmount:     routeAmount,
		RouteRecipient:  routeRecipient,
		DestChainId:     destChainId,
		RequestSender:   requestSender,
		RequestMetadata: requestMetadata,
		RequestPacket:   requestPacket,
		SrcChainType:    srcChainType,
		DestChainType:   destChainType,
		Status:          "created",
	}
}

func NewCrosschainRequestFromMsg(
	msg *MsgCrosschainRequest,
) *CrosschainRequest {
	return &CrosschainRequest{
		SrcChainId:      msg.SrcChainId,
		EventNonce:      msg.EventNonce,
		BlockHeight:     msg.BlockHeight,
		SourceTxHash:    msg.SourceTxHash,
		SrcTimestamp:    msg.SrcTimestamp,
		SrcTxOrigin:     msg.SrcTxOrigin,
		RouteAmount:     msg.RouteAmount,
		RouteRecipient:  msg.RouteRecipient,
		DestChainId:     msg.DestChainId,
		RequestSender:   msg.RequestSender,
		RequestMetadata: msg.RequestMetadata,
		RequestPacket:   msg.RequestPacket,
		SrcChainType:    msg.SrcChainType,
		DestChainType:   msg.DestChainType,
		Status:          "created",
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
		ChainId:              msg.ChainId,
		EventNonce:           msg.EventNonce,
		BlockHeight:          msg.BlockHeight,
		DestTxHash:           msg.DestTxHash,
		RelayerRouterAddress: msg.RelayerRouterAddress,
		SourceChainId:        msg.SourceChainId,
		RequestSender:        msg.RequestSender,
		RequestIdentifier:    msg.RequestIdentifier,
		ExecData:             msg.ExecData,
		ExecStatus:           msg.ExecStatus,
		EthSigner:            msg.EthSigner,
		Signature:            msg.Signature,
		Status:               "ack_created",
	}
}

func NewCrosschainAckRequestConfirm(destChainId string, ackRequestIdentifier uint64, claimHash []byte, orchestrator string, ethSigner string, signature string) *CrosschainAckRequestConfirm {
	return &CrosschainAckRequestConfirm{
		DestChainId:          destChainId,
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
		ChainId:              msg.ChainId,
		EventNonce:           msg.EventNonce,
		BlockHeight:          msg.BlockHeight,
		TxHash:               msg.TxHash,
		RelayerRouterAddress: msg.RelayerRouterAddress,
		RequestIdentifier:    msg.RequestIdentifier,
	}
}
