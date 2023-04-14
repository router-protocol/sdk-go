package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type SudoMsg struct {
	HandleIAck HandleIAck `json:"handle_i_ack"`
}

type HandleIAck struct {
	RequestIdentifier uint64   `json:"request_identifier"`
	ExecFlag          bool     `json:"exec_flag"`
	ExecData          string   `json:"exec_data"`
	RefundAmount      sdk.Coin `json:"refund_amount"`
}
