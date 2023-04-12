package types

import (
	fmt "fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

///////////////////////////////////
//     Implement Metadata     //
///////////////////////////////////

type CrosschainEVMMetadata struct {
	DestGasLimit *big.Int `json:"destGasLimit"`
	DestGasPrice *big.Int `json:"destGasPrice"`
	AckGasLimit  *big.Int `json:"ackGasLimit"`
	AckGasPrice  *big.Int `json:"ackGasPrice"`
	RelayerFees  *big.Int `json:"relayerFees"`
	AckType      uint8    `json:"ackType"`
	IsReadCall   bool     `json:"isReadCall"`
	AsmAddress   []byte   `json:"asmAddress"`
}

type EvmMetadata interface {
	GetRequestMetadata() []byte
}

func DecodeEvmMetadata(msg EvmMetadata) *CrosschainEVMMetadata {
	metadataAbiStruct, _ := abi.NewType("tuple", "struct thing", []abi.ArgumentMarshaling{
		{Name: "destGasLimit", Type: "uint256"},
		{Name: "destGasPrice", Type: "uint256"},
		{Name: "ackGasLimit", Type: "uint256"},
		{Name: "ackGasPrice", Type: "uint256"},
		{Name: "relayerFees", Type: "uint256"},
		{Name: "ackType", Type: "uint8"},
		{Name: "isReadCall", Type: "bool"},
		{Name: "asmAddress", Type: "bytes"},
	})

	args := abi.Arguments{
		{Type: metadataAbiStruct, Name: "param_one"},
	}

	// Unpack the packed data
	unpackedData, err := args.Unpack(msg.GetRequestMetadata())
	if err != nil {
		fmt.Println("failed to unpack data: ", err)
		return nil
	}

	// crosschainEVMMetadata := unpackedData[0].(struct {
	// 	DestGasLimit *big.Int `json:"destGasLimit"`
	// 	DestGasPrice *big.Int `json:"destGasPrice"`
	// 	AckGasLimit  *big.Int `json:"ackGasLimit"`
	// 	AckGasPrice  *big.Int `json:"ackGasPrice"`
	// 	RelayerFees  *big.Int `json:"relayerFees"`
	// 	AckType      uint8    `json:"ackType"`
	// 	IsReadCall   bool     `json:"isReadCall"`
	// 	AsmAddress   []byte   `json:"asmAddress"`
	// })

	crosschainMetadata := unpackedData[0].(CrosschainEVMMetadata)
	return &crosschainMetadata
}
