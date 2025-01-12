package types

type SudoReceiveMsg struct {
	HandleAdhocRequest HandleAdhocRequest `json:"handle_adhoc_request"`
}

type HandleAdhocRequest struct {
	ChainIds    []string `json:"chain_ids"`
	TxHashes    []string `json:"tx_hashes"`
	TxResponses [][]byte `json:"tx_responses"`
	MetaData    [][]byte `json:"meta_data"`
}
