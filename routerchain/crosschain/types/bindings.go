package types

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
	TokenPrice   string `json:"token_price"`
	TokenDecimal uint64 `json:"token_decimal"`
}
