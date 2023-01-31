package types

import (
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
)

type SudoMsg struct {
	HandleInboundReq HandleInboundReq `json:"handle_inbound_req"`
}

type HandleInboundReq struct {
	Sender        string                    `json:"sender"`
	ChainType     multichainTypes.ChainType `json:"chain_type"`
	SourceChainId string                    `json:"source_chain_id"`
	EventNonce    uint64                    `json:"event_nonce"`
	Payload       string                    `json:"payload"`
}
