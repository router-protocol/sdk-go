package types

import (
	"encoding/binary"
	"encoding/hex"
	fmt "fmt"
	"math/big"
)

///////////////////////////////////
//     Implement Metadata     //
///////////////////////////////////

type ContractMetadata struct {
	DestGasLimit uint64   `json:"destGasLimit"`
	DestGasPrice uint64   `json:"destGasPrice"`
	AckGasLimit  uint64   `json:"ackGasLimit"`
	AckGasPrice  uint64   `json:"ackGasPrice"`
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
	requestMetadataStr := hex.EncodeToString(msg.GetRequestMetadata())
	// Slice the string into required chunks and store them in separate variables
	var (
		chunk1 = requestMetadataStr[:16]
		chunk2 = requestMetadataStr[16:32]
		chunk3 = requestMetadataStr[32:48]
		chunk4 = requestMetadataStr[48:64]
		chunk5 = requestMetadataStr[64:96]
		chunk6 = requestMetadataStr[96:98]
		chunk7 = requestMetadataStr[98:100]
		chunk8 = requestMetadataStr[100:]
	)

	// Convert hexadecimal string to uint64
	destGasLimit, _ := hex.DecodeString(chunk1)
	destGasPrice, _ := hex.DecodeString(chunk2)
	ackGasLimit, _ := hex.DecodeString(chunk3)
	ackGasPrice, _ := hex.DecodeString(chunk4)

	// Convert hexadecimal string to *big.Int for 128-bit integer
	relayerFees := new(big.Int)
	relayerFees.SetString(chunk5, 16)

	// Convert hexadecimal string to uint8
	var ackType uint8
	fmt.Sscanf(chunk6, "%x", &ackType)

	// Convert hexadecimal string to bool
	var isReadCall bool
	if chunk7 == "01" {
		isReadCall = true
	}

	// Convert hexadecimal string to bytes
	asmAddress, _ := hex.DecodeString(chunk8)

	metadata := ContractMetadata{
		DestGasLimit: binary.BigEndian.Uint64(destGasLimit),
		DestGasPrice: binary.BigEndian.Uint64(destGasPrice),
		AckGasLimit:  binary.BigEndian.Uint64(ackGasLimit),
		AckGasPrice:  binary.BigEndian.Uint64(ackGasPrice),
		RelayerFees:  relayerFees,
		AckType:      ackType,
		IsReadCall:   isReadCall,
		AsmAddress:   asmAddress,
	}

	fmt.Println("sdk-go GetCheckpoint", "crosschainMetadata", "metadata", metadata)
	return &metadata
}
