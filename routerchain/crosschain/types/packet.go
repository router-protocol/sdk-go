package types

import (
	fmt "fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

///////////////////////////////////
//     Implement Packet     //
///////////////////////////////////

type CrosschainRouterPacket struct {
	DestContractAddress []byte `json:"destContractAddress"`
	Payload             []byte `json:"payload"`
}

type RouterCrosschainPacket interface {
	GetRequestPacket() []byte
}

func DecodeRouterCrosschainPacket(msg RouterCrosschainPacket) *CrosschainRouterPacket {
	fmt.Println("sdk-go Get packet", "Decode router packet")

	Bytes, _ := abi.NewType("bytes", "", nil)

	arguments := abi.Arguments{
		{
			Type: Bytes,
		},
		{
			Type: Bytes,
		},
	}

	data, err := arguments.Unpack(msg.GetRequestPacket())
	if err != nil {
		panic(err)
	}

	// // Print the decoded values
	contractAddress := data[0].([]byte)
	payload := data[1].([]byte)
	fmt.Printf("Data 1: %v\n", string(contractAddress))
	fmt.Printf("Data 2: %v\n", (payload))

	packet := CrosschainRouterPacket{
		DestContractAddress: contractAddress,
		Payload:             payload,
	}

	fmt.Println("sdk-go GetCheckpoint", "crosschainpacket", "packet", packet)
	return &packet
}
