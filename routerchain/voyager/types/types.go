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
	partnerId sdk.Int,
	message []byte,
	depositWithMessage bool,
	destToken []byte) *FundDepositRequest {
	return &FundDepositRequest{
		SrcChainId:         srcChainId,
		SrcChainType:       srcChainType,
		SrcTxHash:          srcTxHash,
		SrcTimestamp:       srcTimestamp,
		Contract:           contract,
		DepositId:          depositId,
		BlockHeight:        blockHeight,
		DestChainId:        destChainId,
		Amount:             amount,
		DestAmount:         destAmount,
		SrcToken:           srcToken,
		Recipient:          recipient,
		Depositor:          depositor,
		PartnerId:          partnerId,
		Message:            message,
		DepositWithMessage: depositWithMessage,
		DestToken:          destToken,
		Status:             VOYAGER_FUND_DEPOSIT_REQUEST_CREATED,
	}
}

func NewFundsDepositedFromMsg(
	msg *MsgFundsDeposited) *FundDepositRequest {
	return &FundDepositRequest{
		SrcChainId:         msg.SrcChainId,
		SrcChainType:       msg.SrcChainType,
		SrcTxHash:          msg.SrcTxHash,
		SrcTimestamp:       msg.SrcTimestamp,
		Contract:           msg.Contract,
		DepositId:          msg.DepositId,
		BlockHeight:        msg.BlockHeight,
		DestChainId:        msg.DestChainId,
		Amount:             msg.Amount,
		DestAmount:         msg.DestAmount,
		SrcToken:           msg.SrcToken,
		Recipient:          msg.Recipient,
		Depositor:          msg.Depositor,
		PartnerId:          msg.PartnerId,
		Message:            msg.Message,
		DepositWithMessage: msg.DepositWithMessage,
		DestToken:          msg.DestToken,
		Status:             VOYAGER_FUND_DEPOSIT_REQUEST_CREATED,
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
	partnerID sdk.Int,
	message []byte,
	depositWithMessage bool,
	destToken []byte) *FundDepositRequestClaimHash {
	return &FundDepositRequestClaimHash{
		SrcChainId:         srcChainId,
		SrcChainType:       srcChainType,
		SrcTxHash:          srcTxHash,
		SrcTimestamp:       srcTimestamp,
		Contract:           contract,
		DepositId:          depositId,
		BlockHeight:        blockHeight,
		DestChainId:        destChainId,
		Amount:             amount,
		DestAmount:         destAmount,
		SrcToken:           srcToken,
		Recipient:          recipient,
		Depositor:          depositor,
		PartnerId:          partnerID,
		Message:            message,
		DepositWithMessage: depositWithMessage,
		DestToken:          destToken,
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
	forwarderRouterAddr string,
	execData []byte,
	execStatus bool) *FundsPaidRequest {
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
		ExecData:            execData,
		ExecStatus:          execStatus,
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
	forwarderRouterAddr string,
	execData []byte,
	execStatus bool) *FundsPaidRequestClaimHash {
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
		ExecData:            execData,
		ExecStatus:          execStatus,
		ForwarderRouterAddr: forwarderRouterAddr,
	}
}

func NewDepositInfoUpdatedRequest(
	srcChainId string,
	srcChainType multichainTypes.ChainType,
	srcTxHash string,
	srcTimestamp uint64,
	depositId uint64,
	contract string,
	eventNonce uint64,
	blockHeight uint64,
	feeAmount sdk.Int,
	initiatewithdrawal bool,
	srcToken string,
	depositor string) *DepositUpdateInfoRequest {

	return &DepositUpdateInfoRequest{
		DepositId:          depositId,
		SrcChainId:         srcChainId,
		SrcChainType:       srcChainType,
		SrcTxHash:          srcTxHash,
		SrcTimestamp:       srcTimestamp,
		Contract:           contract,
		EventNonce:         eventNonce,
		BlockHeight:        blockHeight,
		FeeAmount:          feeAmount,
		Initiatewithdrawal: initiatewithdrawal,
		SrcToken:           srcToken,
		Depositor:          depositor,
		Status:             VOYAGER_DEPOSIT_UPDATE_INFO_REQUEST_CREATED,
	}
}

func NewDepositInfoUpdatedRequestFromMsg(
	msg *MsgDepositInfoUpdated) *DepositUpdateInfoRequest {
	return &DepositUpdateInfoRequest{
		SrcChainId:         msg.SrcChainId,
		SrcChainType:       msg.SrcChainType,
		SrcTxHash:          msg.SrcTxHash,
		SrcTimestamp:       msg.SrcTimestamp,
		DepositId:          msg.DepositId,
		Contract:           msg.Contract,
		EventNonce:         msg.EventNonce,
		BlockHeight:        msg.BlockHeight,
		FeeAmount:          msg.FeeAmount,
		Initiatewithdrawal: msg.Initiatewithdrawal,
		SrcToken:           msg.SrcToken,
		Depositor:          msg.Depositor,
		Status:             VOYAGER_DEPOSIT_UPDATE_INFO_REQUEST_CREATED,
	}
}

func NewDepositInfoUpdatedRequestClaimHash(
	srcChainId string,
	srcChainType multichainTypes.ChainType,
	srcTxHash string,
	srcTimestamp uint64,
	depositId uint64,
	contract string,
	eventNonce uint64,
	blockHeight uint64,
	feeAmount sdk.Int,
	initiatewithdrawal bool,
	srcToken string,
	depositor string) *DepositUpdateInfoRequestClaimHash {
	return &DepositUpdateInfoRequestClaimHash{
		SrcChainId:         srcChainId,
		SrcChainType:       srcChainType,
		SrcTxHash:          srcTxHash,
		SrcTimestamp:       srcTimestamp,
		DepositId:          depositId,
		Contract:           contract,
		EventNonce:         eventNonce,
		BlockHeight:        blockHeight,
		FeeAmount:          feeAmount,
		SrcToken:           srcToken,
		Depositor:          depositor,
		Initiatewithdrawal: initiatewithdrawal,
	}
}
