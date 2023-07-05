package types

import (
	"errors"

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

func DecodeRouterCrosschainPacket(msg RouterCrosschainPacket) (*CrosschainRouterPacket, error) {

	packet := CrosschainRouterPacket{}

	if len(msg.GetRequestPacket()) < 100 {
		return nil, errors.New("Insufficient request packet length")
	}

	Bytes, _ := abi.NewType("bytes", "", nil)
	String, _ := abi.NewType("string", "", nil)

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
		return nil, err
	}

	// // Print the decoded values
	handlerContractAddress := data[0].(string)
	payload := data[1].([]byte)

	packet = CrosschainRouterPacket{
		Handler: handlerContractAddress,
		Payload: payload,
	}

	return &packet, nil
}
