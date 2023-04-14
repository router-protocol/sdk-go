package types

type SudoReceiveMsg struct {
	HandleIReceive HandleIReceive `json:"handle_i_receive"`
}

type HandleIReceive struct {
	RequestSender     string `json:"request_sender"`
	SourceChainId     string `json:"src_chain_id"`
	RequestIdentifier uint64 `json:"request_identifier"`
	Payload           string `json:"payload"`
}
