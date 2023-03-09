package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// RouterQuery contains router custom queries.
type RouterQuery struct {
	GasPrice   *GasPrice   `json:"gas_price,omitempty"`
	TokenPrice *TokenPrice `json:"token_price,omitempty"`
}

type GasPrice struct {
	ChainType uint32 `json:"chain_type"`
	ChainId   string `json:"chain_id"`
}

type GasPriceResponse struct {
	GasPrice uint64 `json:"gas_price"`
}

type TokenPrice struct {
	Symbol string `json:"symbol"`
}

type TokenPriceResponse struct {
	TokenPrice string `json:"token_price"`
}

type RouterMsg struct {
	/// Contracts can create and send Outbound batch requests to routerchain
	OutboundBatchRequests *OutboundBatchRequests `json:"outbound_batch_requests,omitempty"`
}

type OutboundBatchRequests struct {
	OutboundBatchRequests []OutboundBatchRequest `json:"outbound_batch_requests,omitempty"`
}

type OutboundBatchResponses struct {
	OutboundBatchResponses []OutboundBatchResponse `json:"outbound_batch_responses,omitempty"`
}

type OutboundBatchResponse struct {
	OutboundBatchNonce    uint64   `json:"outbound_batch_nonce,omitempty"`
	OutgoingTxFeeDeducted sdk.Coin `json:"outgoing_tx_fee_deducted"`
}

type OutboundBatchRequest struct {
	DestinationChainType int32               `json:"destination_chain_type,omitempty"`
	DestinationChainId   string              `json:"destination_chain_id,omitempty"`
	ContractCalls        []SmartContractCall `json:"contract_calls"`
	RelayerFee           sdk.Coin            `json:"relayer_fee"`
	OutgoingTxFee        OutgoingTxFee       `json:"outgoing_tx_fee"`
	IsAtomic             bool                `json:"is_atomic"`
	ExpiryTimestamp      uint64              `json:"exp_timestamp"`
	RouteAmount          sdk.Int             `json:"route_amount"`
	RouteRecipient       []byte              `json:"route_recipient"`
}

type OutgoingTxFee struct {
	GasLimit uint64 `json:"gas_limit"`
	GasPrice uint64 `json:"gas_price"`
}

type SmartContractCall struct {
	DestinationContractAddress []byte `json:"destination_contract_address,omitempty"`
	Payload                    []byte `json:"payload,omitempty"`
}
