package types

import (
	sdkmath "cosmossdk.io/math"
)

type SudoFundDepositMsg struct {
	HandleFundDeposit HandleFundDeposit `json:"handle_fund_deposit"`
}

type HandleFundDeposit struct {
	SourceChainId          string      `json:"src_chain_id"`
	VoyagerContractAddress string      `json:"voyager_contract_address"`
	DepositId              uint64      `json:"deposit_id"`
	DestChainIdBytes       string      `json:"dest_chain_id_bytes"`
	SrcToken               string      `json:"src_token"`
	Depositor              string      `json:"depositor"`
	DestAmount             sdkmath.Int `json:"dest_amount"`
	Amount                 sdkmath.Int `json:"amount"`
	Recipient              string      `json:"recipient"`
	PartnerId              sdkmath.Int `json:"partner_id"`
	Message                string      `json:"message"`
	DepositWithMessage     bool        `json:"deposit_with_message"`
	DestToken              string      `json:"dest_token"`
}

type SudoFundsPaidMsg struct {
	HandleFundsPaid HandleFundsPaid `json:"handle_funds_paid"`
}

type HandleFundsPaid struct {
	SourceChainId          string `json:"src_chain_id"`
	VoyagerContractAddress string `json:"voyager_contract_address"`
	RequestIdentifier      uint64 `json:"request_identifier"`
	MessageHash            string `json:"message_hash"`
	ForwarderAddress       string `json:"forwarder_address"`
	ForwarderRouterAddress string `json:"forwarder_router_address"`
	SrcTimestamp           uint64 `json:"src_timestamp"`
}

type SudoHandleDepositInfoUpdateMsg struct {
	HandleDepositInfoUpdate HandleDepositInfoUpdate `json:"handle_deposit_info_update"`
}

type HandleDepositInfoUpdate struct {
	SourceChainId          string      `json:"src_chain_id"`
	VoyagerContractAddress string      `json:"voyager_contract_address"`
	SrcToken               string      `json:"src_token"`
	FeeAmount              sdkmath.Int `json:"fee_amount"`
	DepositId              uint64      `json:"deposit_id"`
	EventNonce             uint64      `json:"event_nonce"`
	InitiateWithdrawal     bool        `json:"initiate_withdrawal"`
	Depositor              string      `json:"depositor"`
}
