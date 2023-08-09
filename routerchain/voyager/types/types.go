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
	destAmount sdk.Int,
	srcToken string,
	recipient []byte,
	depositor string,
	partnerId sdk.Int) *FundDepositRequest {
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
		DestAmount:   destAmount,
		SrcToken:     srcToken,
		Recipient:    recipient,
		Depositor:    depositor,
		PartnerId:    partnerId,
		Status:       VOYAGER_FUND_DEPOSIT_REQUEST_CREATED,
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
		DestAmount:   msg.DestAmount,
		SrcToken:     msg.SrcToken,
		Recipient:    msg.Recipient,
		Depositor:    msg.Depositor,
		PartnerId:    msg.PartnerId,
		Status:       VOYAGER_FUND_DEPOSIT_REQUEST_CREATED,
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
	destAmount sdk.Int,
	srcToken string,
	recipient []byte,
	depositor string,
	partnerID sdk.Int) *FundDepositRequestClaimHash {
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
		DestAmount:   destAmount,
		SrcToken:     srcToken,
		Recipient:    recipient,
		Depositor:    depositor,
		PartnerId:    partnerID,
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
	forwarderRouterAddr string) *FundsPaidRequest {
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
		Status:              VOYAGER_FUND_PAID_REQUEST_CREATED,
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
		Status:              VOYAGER_FUND_PAID_REQUEST_CREATED,
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
	forwarderRouterAddr string) *FundsPaidRequestClaimHash {
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
