package types

import (
	fmt "fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

///////////////////////////////////
//     Implement Metadata     //
///////////////////////////////////

type ContractMetadata struct {
	DestGasLimit *big.Int `json:"destGasLimit"`
	DestGasPrice *big.Int `json:"destGasPrice"`
	AckGasLimit  *big.Int `json:"ackGasLimit"`
	AckGasPrice  *big.Int `json:"ackGasPrice"`
	RelayerFees  *big.Int `json:"relayerFees"`
	AckType      uint8    `json:"ackType"`
	IsReadCall   bool     `json:"isReadCall"`
	AsmAddress   []byte   `json:"asmAddress"`
}

type IContractMetadata interface {
	GetRequestMetadata() []byte
}

func DecodeContractMetadata(msg IContractMetadata) *ContractMetadata {
	fmt.Println("sdk-go GetCheckpoint", "Decode Evm checkpoint")

	// Define the ABI types for the function argument and return value
	metadataType, _ := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{Name: "destGasLimit", Type: "uint256"},
		{Name: "destGasPrice", Type: "uint256"},
		{Name: "ackGasLimit", Type: "uint256"},
		{Name: "ackGasPrice", Type: "uint256"},
		{Name: "relayerFees", Type: "uint256"},
		{Name: "ackType", Type: "uint8"},
		{Name: "isReadCall", Type: "bool"},
		{Name: "asmAddress", Type: "bytes"},
	})

	args := abi.Arguments{{Type: metadataType}}

	fmt.Println("sdk-go GetCheckpoint", "metadataAbiStruct", metadataType)
	fmt.Println("sdk-go GetCheckpoint", "args", args)

	// Unpack the function argument
	unpackedData, err := args.Unpack(msg.GetRequestMetadata())
	if err != nil {
		fmt.Println("failed to unpack data:", err)
		return nil
	}

	fmt.Println("sdk-go GetCheckpoint", "unpackedData", unpackedData)

	crosschainEVMMetadataArgs := unpackedData[0].(struct {
		DestGasLimit *big.Int `json:"destGasLimit"`
		DestGasPrice *big.Int `json:"destGasPrice"`
		AckGasLimit  *big.Int `json:"ackGasLimit"`
		AckGasPrice  *big.Int `json:"ackGasPrice"`
		RelayerFees  *big.Int `json:"relayerFees"`
		AckType      uint8    `json:"ackType"`
		IsReadCall   bool     `json:"isReadCall"`
		AsmAddress   []byte   `json:"asmAddress"`
	})

	fmt.Println("sdk-go GetCheckpoint", "crosschainEVMMetadataArgs", "DestGasLimit", crosschainEVMMetadataArgs.DestGasLimit)
	fmt.Println("sdk-go GetCheckpoint", "crosschainEVMMetadataArgs", "DestGasPrice", crosschainEVMMetadataArgs.DestGasPrice)
	fmt.Println("sdk-go GetCheckpoint", "crosschainEVMMetadataArgs", "DestGasLimit", crosschainEVMMetadataArgs.DestGasLimit)
	fmt.Println("sdk-go GetCheckpoint", "crosschainEVMMetadataArgs", "AckGasLimit", crosschainEVMMetadataArgs.AckGasLimit)
	fmt.Println("sdk-go GetCheckpoint", "crosschainEVMMetadataArgs", "AckGasPrice", crosschainEVMMetadataArgs.AckGasPrice)
	fmt.Println("sdk-go GetCheckpoint", "crosschainEVMMetadataArgs", "RelayerFees", crosschainEVMMetadataArgs.RelayerFees)
	fmt.Println("sdk-go GetCheckpoint", "crosschainEVMMetadataArgs", "AckType", crosschainEVMMetadataArgs.AckType)
	fmt.Println("sdk-go GetCheckpoint", "crosschainEVMMetadataArgs", "IsReadCall", crosschainEVMMetadataArgs.IsReadCall)
	fmt.Println("sdk-go GetCheckpoint", "crosschainEVMMetadataArgs", "AsmAddress", crosschainEVMMetadataArgs.AsmAddress)

	// Cast the unpacked data to the Metadata struct
	// metadata, ok := unpackedData[0].(Metadata)
	// if !ok {
	// 	fmt.Println("failed to cast data to Metadata struct")
	// 	return nil
	// }

	metadata := ContractMetadata{
		DestGasLimit: crosschainEVMMetadataArgs.DestGasLimit,
		DestGasPrice: crosschainEVMMetadataArgs.DestGasPrice,
		AckGasLimit:  crosschainEVMMetadataArgs.AckGasLimit,
		AckGasPrice:  crosschainEVMMetadataArgs.AckGasPrice,
		RelayerFees:  crosschainEVMMetadataArgs.RelayerFees,
		AckType:      crosschainEVMMetadataArgs.AckType,
		IsReadCall:   crosschainEVMMetadataArgs.IsReadCall,
		AsmAddress:   crosschainEVMMetadataArgs.AsmAddress,
	}

	fmt.Println("sdk-go GetCheckpoint", "crosschainMetadata", "metadata", metadata)
	return &metadata
}
