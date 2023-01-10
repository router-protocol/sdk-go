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
	ExecutionCode         uint64                    `json:"execution_code"`
	ExecutionStatus       bool                      `json:"execution_status"`
	ExecFlags             []bool                    `json:"exec_flags"`
	ExecData              [][]byte                  `json:"exec_data"`
}
