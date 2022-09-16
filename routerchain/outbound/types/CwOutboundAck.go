package types

import (
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
)

type SudoMsg struct {
	HandleOutboundAck HandleOutboundAck `json:"handle_outbound_ack"`
}

type HandleOutboundAck struct {
	OutboundTxRequestedBy string                    `json:"outbound_tx_requested_by"`
	DestinationChainType  multichainTypes.ChainType `json:"destination_chain_type"`
	DestinationChainId    string                    `json:"destination_chain_id"`
	OutboundBatchNonce    uint64                    `json:"outbound_batch_nonce"`
	ContractAckResponses  []ContractAck             `json:"contract_ack_responses"`
}

type ContractAck struct {
	DestinationContractAddress []byte `json:"destination_contract_address"`
	AckStatus                  bool   `json:"ack_status"`
	ResponsePayload            []byte `json:"response_payload"`
}
