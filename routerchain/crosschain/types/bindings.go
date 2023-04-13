package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// RouterQuery contains router custom queries.
type RouterQuery struct {
	GasPrice   *GasPrice   `json:"gas_price,omitempty"`
	TokenPrice *TokenPrice `json:"token_price,omitempty"`
}

type GasPrice struct {
	ChainId string `json:"chain_id"`
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
	/// Contracts can create and send crosschain calls to routerchain
	CrosschainCall *CrosschainCall `json:"crosschain_call,omitempty"`
}

type CrosschainResponse struct {
	RequestIdentifier     uint64   `json:"request_identifier,omitempty"`
	CrosschainFeeDeducted sdk.Coin `json:"crosschain_fee_deducted"`
}

type CrosschainCall struct {
	Version         uint64  `json:"version"`
	RouteAmount     sdk.Int `json:"route_amount"`
	RouteRecipient  []byte  `json:"route_recipient"`
	DestChainId     string  `json:"dest_chain_id,omitempty"`
	RequestMetadata []byte  `json:"request_metadata"`
	RequestPacket   []byte  `json:"request_packet"`
}
