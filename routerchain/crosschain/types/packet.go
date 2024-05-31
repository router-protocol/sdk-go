package types

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
)

///////////////////////////////////
//     Implement Packet Decoding   //
///////////////////////////////////

type CrosschainRouterPacket struct {
	Handler string `json:"handler"`
	Payload []byte `json:"payload"`
}

type RouterCrosschainPacket interface {
	GetRequestPacket() []byte
}

func DecodeRouterCrosschainPacket(msg RouterCrosschainPacket) *CrosschainRouterPacket {
	packet := &CrosschainRouterPacket{}

	if len(msg.GetRequestPacket()) < 100 {
		return packet
	}

	Bytes, err := abi.NewType("bytes", "", nil)
	if err != nil {
		return packet
	}

	String, err := abi.NewType("string", "", nil)
	if err != nil {
		return packet
	}

	arguments := abi.Arguments{
		{
			Type: String,
		},
		{
			Type: Bytes,
		},
	}

	data, err := arguments.Unpack(msg.GetRequestPacket())
	if err != nil {
		return packet
	}

	// Assign the decoded values
	handlerContractAddress := data[0].(string)
	payload := data[1].([]byte)

	packet.Handler = handlerContractAddress
	packet.Payload = payload

	return packet
}
