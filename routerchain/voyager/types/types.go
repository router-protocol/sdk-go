package types

import (
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
	destChainId string,
	destChainType multichainTypes.ChainType,
	amount uint64,
	relayerFees uint64,
	srcToken string,
	recipient string,
	depositor string) *FundDepositRequest {
	return &FundDepositRequest{
		SrcChainId:    srcChainId,
		SrcChainType:  srcChainType,
		SrcTxHash:     srcTxHash,
		SrcTimestamp:  srcTimestamp,
		Contract:      contract,
		DepositId:     depositId,
		BlockHeight:   blockHeight,
		DestChainId:   destChainId,
		DestChainType: destChainType,
		Amount:        amount,
		RelayerFees:   relayerFees,
		SrcToken:      srcToken,
		Recipient:     recipient,
		Depositor:     depositor,
		Status:        "fund_deposit_request_created",
	}
}

func NewFundsDepositedFromMsg(
	msg *MsgFundsDeposited) *FundDepositRequest {
	return &FundDepositRequest{
		SrcChainId:    msg.SrcChainId,
		SrcChainType:  msg.SrcChainType,
		SrcTxHash:     msg.SrcTxHash,
		SrcTimestamp:  msg.SrcTimestamp,
		Contract:      msg.Contract,
		DepositId:     msg.DepositId,
		BlockHeight:   msg.BlockHeight,
		DestChainId:   msg.DestChainId,
		DestChainType: msg.DestChainType,
		Amount:        msg.Amount,
		RelayerFees:   msg.RelayerFees,
		SrcToken:      msg.SrcToken,
		Recipient:     msg.Recipient,
		Depositor:     msg.Depositor,
		Status:        "fund_deposit_request_created",
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
	destChainId string,
	destChainType multichainTypes.ChainType,
	amount uint64,
	relayerFees uint64,
	srcToken string,
	recipient string,
	depositor string) *FundDepositRequestClaimHash {
	return &FundDepositRequestClaimHash{
		SrcChainId:    srcChainId,
		SrcChainType:  srcChainType,
		SrcTxHash:     srcTxHash,
		SrcTimestamp:  srcTimestamp,
		Contract:      contract,
		DepositId:     depositId,
		BlockHeight:   blockHeight,
		DestChainId:   destChainId,
		DestChainType: destChainType,
		Amount:        amount,
		RelayerFees:   relayerFees,
		SrcToken:      srcToken,
		Recipient:     recipient,
		Depositor:     depositor,
	}
}
