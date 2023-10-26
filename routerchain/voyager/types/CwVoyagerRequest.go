package types

import sdk "github.com/cosmos/cosmos-sdk/types"

type SudoFundDepositMsg struct {
	HandleFundDeposit HandleFundDeposit `json:"handle_fund_deposit"`
}

type HandleFundDeposit struct {
	SourceChainId          string  `json:"src_chain_id"`
	VoyagerContractAddress string  `json:"voyager_contract_address"`
	DepositId              uint64  `json:"deposit_id"`
	DestChainIdBytes       string  `json:"dest_chain_id_bytes"`
	SrcToken               string  `json:"src_token"`
	Depositor              string  `json:"depositor"`
	DestAmount             sdk.Int `json:"dest_amount"`
	Amount                 sdk.Int `json:"amount"`
	Recipient              string  `json:"recipient"`
	PartnerId              sdk.Int `json:"partner_id"`
	Message                string  `json:"message"`
	DepositWithMessage     bool    `json:"deposit_with_message"`
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
	SourceChainId          string  `json:"src_chain_id"`
	VoyagerContractAddress string  `json:"voyager_contract_address"`
	SrcToken               string  `json:"src_token"`
	FeeAmount              sdk.Int `json:"fee_amount"`
	DepositId              uint64  `json:"deposit_id"`
	EventNonce             uint64  `json:"event_nonce"`
	InitiateWithdrawal     bool    `json:"initiate_withdrawal"`
	Depositor              string  `json:"depositor"`
}
