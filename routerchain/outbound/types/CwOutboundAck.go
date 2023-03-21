package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
)

type SudoMsg struct {
	HandleOutboundAck HandleOutboundAck `json:"handle_outbound_ack"`
}

type HandleOutboundAck struct {
	OutboundTxRequestedBy string                    `json:"outbound_tx_requested_by"`
	DestinationChainId    string                    `json:"destination_chain_id"`
	DestinationChainType  multichainTypes.ChainType `json:"destination_chain_type"`
	OutboundBatchNonce    uint64                    `json:"outbound_batch_nonce"`
	ExecutionCode         uint64                    `json:"execution_code"`
	ExecutionStatus       bool                      `json:"execution_status"`
	ExecFlags             []bool                    `json:"exec_flags"`
	ExecData              []string                  `json:"exec_data"`
	RefundAmount          sdk.Coin                  `json:"refund_amount"`
}
