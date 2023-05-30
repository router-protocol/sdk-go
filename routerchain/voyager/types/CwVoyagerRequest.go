package types

import sdk "github.com/cosmos/cosmos-sdk/types"

type SudoFundDepositMsg struct {
	HandleFundDeposit HandleFundDeposit `json:"handle_fund_deposit"`
}

type HandleFundDeposit struct {
	SourceChainId    string  `json:"src_chain_id"`
	DepositId        uint64  `json:"deposit_id"`
	DestChainIdBytes string  `json:"dest_chain_id_bytes"`
	SrcToken         string  `json:"src_token"`
	Depositor        string  `json:"depositor"`
	RelayerFeePct    sdk.Int `json:"relayer_fee_pct"`
	Amount           sdk.Int `json:"amount"`
	Recipient        string  `json:"recipient"`
}

type HandleFundDepositOld struct {
	// SourceChainId string `json:"src_chain_id"`
	DepositId uint64 `json:"deposit_id"`
	// DestChainIdBytes string `json:"dest_chain_id_bytes"`
	// SrcToken         string `json:"src_token"`
	// Depositor        string `json:"depositor"`
	// RelayerFeePct    uint64 `json:"relayer_fee_pct"`
	// Amount           uint64 `json:"amount"`
	// Recipient        string `json:"recipient"`
}

type SudoFundsPaidMsg struct {
	HandleFundsPaid HandleFundsPaid `json:"handle_funds_paid"`
}

type HandleFundsPaid struct {
	SourceChainId          string `json:"src_chain_id"`
	RequestIdentifier      uint64 `json:"request_identifier"`
	MessageHash            string `json:"message_hash"`
	ForwarderAddress       string `json:"forwarder_address"`
	ForwarderRouterAddress string `json:"forwarder_router_address"`
}

type SudoCounterMsg struct {
	HandleFundDeposit HandleFundDeposit `json:"handle_fund_deposit"`
}

type SetCounter struct {
	Counter uint64 `json:"counter"`
}
