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

type EvmContractMetadata struct {
	DestGasLimit uint64   `json:"destGasLimit"`
	DestGasPrice uint64   `json:"destGasPrice"`
	AckGasLimit  uint64   `json:"ackGasLimit"`
	AckGasPrice  uint64   `json:"ackGasPrice"`
	RelayerFees  *big.Int `json:"relayerFees"`
	AckType      uint8    `json:"ackType"`
	IsReadCall   bool     `json:"isReadCall"`
	AsmAddress   string   `json:"asmAddress"`
}

type IContractMetadata interface {
	GetRequestMetadata() []byte
}

func DecodeEvmContractMetadata(msg IContractMetadata) *EvmContractMetadata {
	metadata := &EvmContractMetadata{
		RelayerFees: big.NewInt(0),
	}
	requestMetadataStr := hex.EncodeToString(msg.GetRequestMetadata())

	if len(requestMetadataStr) < 100 {
		return metadata
	}

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

	destGasLimit, err := hex.DecodeString(chunk1)
	if err != nil {
		destGasLimit = make([]byte, 8)
	}
	destGasPrice, err := hex.DecodeString(chunk2)
	if err != nil {
		destGasPrice = make([]byte, 8)
	}
	ackGasLimit, err := hex.DecodeString(chunk3)
	if err != nil {
		ackGasLimit = make([]byte, 8)
	}
	ackGasPrice, err := hex.DecodeString(chunk4)
	if err != nil {
		ackGasPrice = make([]byte, 8)
	}

	if _, ok := metadata.RelayerFees.SetString(chunk5, 16); !ok {
		metadata.RelayerFees.SetInt64(0) // Set to zero if there's an issue
	}

	var ackType uint8
	if _, err := fmt.Sscanf(chunk6, "%x", &ackType); err != nil {
		ackType = 0
	}

	var isReadCall bool
	if chunk7 == "01" {
		isReadCall = true
	}

	asmAddressByte, err := hex.DecodeString(chunk8)
	if err != nil {
		asmAddressByte = make([]byte, 0)
	}
	asmAddress := string(asmAddressByte)

	metadata.DestGasLimit = binary.BigEndian.Uint64(destGasLimit)
	metadata.DestGasPrice = binary.BigEndian.Uint64(destGasPrice)
	metadata.AckGasLimit = binary.BigEndian.Uint64(ackGasLimit)
	metadata.AckGasPrice = binary.BigEndian.Uint64(ackGasPrice)
	metadata.AckType = ackType
	metadata.IsReadCall = isReadCall
	metadata.AsmAddress = asmAddress

	return metadata
}
