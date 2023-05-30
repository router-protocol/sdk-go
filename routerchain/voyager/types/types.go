package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
)

func NewFundDepositRequest(
	srcChainId string,
	srcChainType multichainTypes.ChainType,
	srcTxHash string,
	srcTimestamp uint64,
	contract string,
	depositId uint64,
	blockHeight uint64,
	destChainId []byte,
	amount sdk.Int,
	relayerFees sdk.Int,
	srcToken string,
	recipient []byte,
	depositor string) *FundDepositRequest {
	return &FundDepositRequest{
		SrcChainId:   srcChainId,
		SrcChainType: srcChainType,
		SrcTxHash:    srcTxHash,
		SrcTimestamp: srcTimestamp,
		Contract:     contract,
		DepositId:    depositId,
		BlockHeight:  blockHeight,
		DestChainId:  destChainId,
		Amount:       amount,
		RelayerFees:  relayerFees,
		SrcToken:     srcToken,
		Recipient:    recipient,
		Depositor:    depositor,
		Status:       "fund_deposit_request_created",
	}
}

func NewFundsDepositedFromMsg(
	msg *MsgFundsDeposited) *FundDepositRequest {
	return &FundDepositRequest{
		SrcChainId:   msg.SrcChainId,
		SrcChainType: msg.SrcChainType,
		SrcTxHash:    msg.SrcTxHash,
		SrcTimestamp: msg.SrcTimestamp,
		Contract:     msg.Contract,
		DepositId:    msg.DepositId,
		BlockHeight:  msg.BlockHeight,
		DestChainId:  msg.DestChainId,
		Amount:       msg.Amount,
		RelayerFees:  msg.RelayerFees,
		SrcToken:     msg.SrcToken,
		Recipient:    msg.Recipient,
		Depositor:    msg.Depositor,
		Status:       "fund_deposit_request_created",
	}
}

func NewFundDepositRequestClaimHash(
	srcChainId string,
	srcChainType multichainTypes.ChainType,
	srcTxHash string,
	srcTimestamp uint64,
	contract string,
	depositId uint64,
	blockHeight uint64,
	destChainId []byte,
	amount sdk.Int,
	relayerFees sdk.Int,
	srcToken string,
	recipient []byte,
	depositor string) *FundDepositRequestClaimHash {
	return &FundDepositRequestClaimHash{
		SrcChainId:   srcChainId,
		SrcChainType: srcChainType,
		SrcTxHash:    srcTxHash,
		SrcTimestamp: srcTimestamp,
		Contract:     contract,
		DepositId:    depositId,
		BlockHeight:  blockHeight,
		DestChainId:  destChainId,
		Amount:       amount,
		RelayerFees:  relayerFees,
		SrcToken:     srcToken,
		Recipient:    recipient,
		Depositor:    depositor,
	}
}

func NewFundPaidRequest(
	srcChainId string,
	srcChainType multichainTypes.ChainType,
	srcTxHash string,
	srcTimestamp uint64,
	contract string,
	eventNonce uint64,
	blockHeight uint64,
	messageHash []byte,
	forwarder string,
	forwarderRouterAddr []byte) *FundsPaidRequest {
	return &FundsPaidRequest{
		SrcChainId:          srcChainId,
		SrcChainType:        srcChainType,
		SrcTxHash:           srcTxHash,
		SrcTimestamp:        srcTimestamp,
		Contract:            contract,
		EventNonce:          eventNonce,
		BlockHeight:         blockHeight,
		MessageHash:         messageHash,
		Forwarder:           forwarder,
		ForwarderRouterAddr: forwarderRouterAddr,
		Status:              "fund_paid_request_created",
	}
}

func NewFundsPaidFromMsg(
	msg *MsgFundsPaid) *FundsPaidRequest {
	return &FundsPaidRequest{
		SrcChainId:          msg.SrcChainId,
		SrcChainType:        msg.SrcChainType,
		SrcTxHash:           msg.SrcTxHash,
		SrcTimestamp:        msg.SrcTimestamp,
		Contract:            msg.Contract,
		EventNonce:          msg.EventNonce,
		BlockHeight:         msg.BlockHeight,
		MessageHash:         msg.MessageHash,
		Forwarder:           msg.Forwarder,
		ForwarderRouterAddr: msg.ForwarderRouterAddr,
		Status:              "fund_paid_request_created",
	}
}

func NewFundPaidRequestClaimHash(
	srcChainId string,
	srcChainType multichainTypes.ChainType,
	srcTxHash string,
	srcTimestamp uint64,
	contract string,
	eventNonce uint64,
	blockHeight uint64,
	messageHash []byte,
	forwarder string,
	forwarderRouterAddr []byte) *FundsPaidRequestClaimHash {
	return &FundsPaidRequestClaimHash{
		SrcChainId:          srcChainId,
		SrcChainType:        srcChainType,
		SrcTxHash:           srcTxHash,
		SrcTimestamp:        srcTimestamp,
		Contract:            contract,
		EventNonce:          eventNonce,
		BlockHeight:         blockHeight,
		MessageHash:         messageHash,
		Forwarder:           forwarder,
		ForwarderRouterAddr: forwarderRouterAddr,
	}
}
