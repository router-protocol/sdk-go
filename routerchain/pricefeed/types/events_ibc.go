package types

// IBC events
const (
	EventTypeTimeout = "timeout"
	// this line is used by starport scaffolding # ibc/packet/event

	AttributeKeyAckSuccess = "success"
	AttributeKeyAck        = "acknowledgement"
	AttributeKeyAckError   = "error"

	EventTypeRequestPacket  = "oracle_request_packet"
	EventTypeResponsePacket = "oracle_response_packet"
	EventTypePriceUpdate    = "price_update"

	AttributeKeyRequestID = "request_id"
	AttributeKeySymbol    = "symbol"
	AttributeKeyPrice     = "price"
	AttributeKeyTimestamp = "timestamp"
)
