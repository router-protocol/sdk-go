package types

import (
	"encoding/hex"
	fmt "fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
	"github.com/router-protocol/sdk-go/routerchain/util"

	"github.com/NethermindEth/juno/core/felt"
	starknetgo "github.com/NethermindEth/starknet.go/curve"
	"github.com/NethermindEth/starknet.go/utils"
)

///////////////////////////////////
//     Implement Ccheckpoint     //
///////////////////////////////////

// GetCheckpoint gets the checkpoint signature from the given CrossTalkRequest
func (msg MsgCrosschainRequest) GetCheckpoint(routerIDstring string) []byte {
	/**
	     Always get checkpoint from crosschain request only.
	 	Crosschain request is dynamic while msg is static.
	**/
	crosschainRequest := NewCrosschainRequestFromMsg(&msg)

	switch crosschainRequest.DestChainType {
	case multichainTypes.CHAIN_TYPE_NEAR, multichainTypes.CHAIN_TYPE_POLKADOT, multichainTypes.CHAIN_TYPE_BITCOIN:
		return crosschainRequest.GetNearCheckpoint("")
	case multichainTypes.CHAIN_TYPE_COSMOS:
		return nil
	case multichainTypes.CHAIN_TYPE_STARKNET:
		return crosschainRequest.GetStarknetCheckpoint("")
	case multichainTypes.CHAIN_TYPE_ALEPH_ZERO:
		return crosschainRequest.GetAlephZeroCheckpoint("")
	default:
		return crosschainRequest.GetEvmCheckpoint("")
	}
}

// GetCheckpoint gets the checkpoint signature from the given CrossTalkRequest
func (msg CrosschainRequest) GetCheckpoint(routerIDstring string) []byte {
	/**
	     Always get checkpoint from crosschain request only.
	 	Crosschain request is dynamic while msg is static.
	**/

	switch msg.DestChainType {
	case multichainTypes.CHAIN_TYPE_NEAR, multichainTypes.CHAIN_TYPE_POLKADOT, multichainTypes.CHAIN_TYPE_BITCOIN:
		return msg.GetNearCheckpoint("")
	case multichainTypes.CHAIN_TYPE_COSMOS:
		return nil
	case multichainTypes.CHAIN_TYPE_STARKNET:
		return msg.GetStarknetCheckpoint("")
	case multichainTypes.CHAIN_TYPE_ALEPH_ZERO:
		return msg.GetAlephZeroCheckpoint("")
	default:
		return msg.GetEvmCheckpoint("")
	}
}

func (msg CrosschainRequest) GetAlephZeroCheckpoint(routerIDstring string) []byte {
	//////////////////////////////////////////////////////////////////////
	/////  Build data with types required for iReceive gateway call  /////
	//////////////////////////////////////////////////////////////////////
	metadata := DecodeEvmContractMetadata(&msg)
	requestPacket := DecodeRouterCrosschainPacket(&msg)

	methodNameBytes := []uint8("iReceive")
	var crosschainMethodName [32]uint8
	copy(crosschainMethodName[:], methodNameBytes)

	routeAmount := msg.RouteAmount.BigInt()

	requestIdentifier := &big.Int{}
	requestIdentifier.SetUint64(msg.RequestIdentifier)

	requestTimestamp := &big.Int{}
	requestTimestamp.SetUint64(uint64(msg.SrcTimestamp))

	var routeRecipient [32]uint8
	routeRecipientByte, err := hex.DecodeString(msg.RouteRecipient)
	if err != nil {
		return nil
	}
	copy(routeRecipient[:], routeRecipientByte)

	var asmAddress [32]uint8
	asmAddressByte, err := hex.DecodeString(metadata.AsmAddress)
	if err != nil {
		return nil
	}
	copy(asmAddress[:], asmAddressByte)

	var handler [32]uint8
	handlerByte, err := hex.DecodeString(requestPacket.Handler)
	if err != nil {
		return nil
	}
	copy(handler[:], handlerByte)

	/////////////////////////////////////////////////
	/////  pack abi for iReceive function  //////////
	/////////////////////////////////////////////////
	abiDef, err := abi.JSON(strings.NewReader(util.CrosschainRequestAlephZeroCheckpointABIJSON))
	if err != nil {
		panic("Bad ABI constant!")
	}

	// the methodName needs to be the same as the 'name' above in the checkpointAbiJson
	// but other than that it's a constant that has no impact on the output. This is because
	// it gets encoded as a function name which we must then discard.
	abiEncodedBatch, err := abiDef.Pack("checkpoint",
		crosschainMethodName,
		routeAmount,
		requestIdentifier,
		requestTimestamp,
		msg.SrcChainId,
		routeRecipient,
		msg.DestChainId,
		asmAddress,
		msg.RequestSender,
		handler,
		requestPacket.Payload,
		metadata.IsReadCall,
	)

	// fmt.Println("Checkpoint- crosschainMethodName ", crosschainMethodName)
	// fmt.Println("Checkpoint-routeAmount", routeAmount)
	// fmt.Println("Checkpoint-requestIdentifier", requestIdentifier)
	// fmt.Println("Checkpoint-requestTimestamp", requestTimestamp)
	// fmt.Println("Checkpoint-msg.SrcChainId", msg.SrcChainId)
	// fmt.Println("Checkpoint-routeRecipient", routeRecipient)
	// fmt.Println("Checkpoint-msg.DestChainId", msg.DestChainId)
	// fmt.Println("Checkpoint-asmAddress", asmAddress)
	// fmt.Println("Checkpoint-msg.RequestSender", msg.RequestSender)

	// fmt.Println("Checkpoint-handler", handler)
	// fmt.Println("Checkpoint-Payload", requestPacket.Payload)
	// fmt.Println("Checkpoint-abiEncodedBatch", abiEncodedBatch)

	// this should never happen outside of test since any case that could crash on encoding
	// should be filtered above.
	if err != nil {
		panic(fmt.Sprintf("Error packing checkpoint! %s/n", err))
	}

	// we hash the resulting encoded bytes discarding the first 4 bytes these 4 bytes are the constant
	// method name 'checkpoint'. If you were to replace the checkpoint constant in this code you would
	// then need to adjust how many bytes you truncate off the front to get the output of abi.encode()
	return crypto.Keccak256Hash(abiEncodedBatch[4:]).Bytes()
}

func (msg CrosschainRequest) GetNearCheckpoint(routerIDstring string) []byte {
	//////////////////////////////////////////////////////////////////////
	/////  Build data with types required for iReceive gateway call  /////
	//////////////////////////////////////////////////////////////////////
	metadata := DecodeEvmContractMetadata(&msg)
	requestPacket := DecodeRouterCrosschainPacket(&msg)

	methodNameBytes := []uint8("iReceive")
	var crosschainMethodName [32]uint8
	copy(crosschainMethodName[:], methodNameBytes)

	routeAmount := msg.RouteAmount.BigInt()

	requestIdentifier := &big.Int{}
	requestIdentifier.SetUint64(msg.RequestIdentifier)

	requestTimestamp := &big.Int{}
	requestTimestamp.SetUint64(uint64(msg.SrcTimestamp))

	routeRecipient := msg.RouteRecipient
	asmAddress := metadata.AsmAddress
	handler := requestPacket.Handler

	/////////////////////////////////////////////////
	/////  pack abi for iReceive function  //////////
	/////////////////////////////////////////////////
	abiDef, err := abi.JSON(strings.NewReader(util.CrosschainRequestNearCheckpointABIJSON))
	if err != nil {
		panic("Bad ABI constant!")
	}

	// the methodName needs to be the same as the 'name' above in the checkpointAbiJson
	// but other than that it's a constant that has no impact on the output. This is because
	// it gets encoded as a function name which we must then discard.
	abiEncodedBatch, err := abiDef.Pack("checkpoint",
		crosschainMethodName,
		routeAmount,
		requestIdentifier,
		requestTimestamp,
		msg.SrcChainId,
		routeRecipient,
		msg.DestChainId,
		asmAddress,
		msg.RequestSender,
		handler,
		requestPacket.Payload,
		metadata.IsReadCall,
	)

	// fmt.Println("Checkpoint- crosschainMethodName ", crosschainMethodName)
	// fmt.Println("Checkpoint-routeAmount", routeAmount)
	// fmt.Println("Checkpoint-requestIdentifier", requestIdentifier)
	// fmt.Println("Checkpoint-requestTimestamp", requestTimestamp)
	// fmt.Println("Checkpoint-msg.SrcChainId", msg.SrcChainId)
	// fmt.Println("Checkpoint-routeRecipient", routeRecipient)
	// fmt.Println("Checkpoint-msg.DestChainId", msg.DestChainId)
	// fmt.Println("Checkpoint-asmAddress", asmAddress)
	// fmt.Println("Checkpoint-msg.RequestSender", msg.RequestSender)

	// fmt.Println("Checkpoint-handler", handler)
	// fmt.Println("Checkpoint-Payload", requestPacket.Payload)
	// fmt.Println("Checkpoint-abiEncodedBatch", abiEncodedBatch)

	// this should never happen outside of test since any case that could crash on encoding
	// should be filtered above.
	if err != nil {
		panic(fmt.Sprintf("Error packing checkpoint! %s/n", err))
	}

	// we hash the resulting encoded bytes discarding the first 4 bytes these 4 bytes are the constant
	// method name 'checkpoint'. If you were to replace the checkpoint constant in this code you would
	// then need to adjust how many bytes you truncate off the front to get the output of abi.encode()
	return crypto.Keccak256Hash(abiEncodedBatch[4:]).Bytes()
}

func (msg CrosschainRequest) GetEvmCheckpoint(routerIDstring string) []byte {
	//////////////////////////////////////////////////////////////////////
	/////  Build data with types required for iReceive gateway call  /////
	//////////////////////////////////////////////////////////////////////
	metadata := DecodeEvmContractMetadata(&msg)
	requestPacket := DecodeRouterCrosschainPacket(&msg)

	methodNameBytes := []uint8("iReceive")
	var crosschainMethodName [32]uint8
	copy(crosschainMethodName[:], methodNameBytes)

	routeAmount := msg.RouteAmount.BigInt()

	requestIdentifier := &big.Int{}
	requestIdentifier.SetUint64(msg.RequestIdentifier)

	requestTimestamp := &big.Int{}
	requestTimestamp.SetUint64(uint64(msg.SrcTimestamp))

	routeRecipient := common.HexToAddress(msg.RouteRecipient)
	asmAddress := common.HexToAddress(metadata.AsmAddress)
	handler := common.HexToAddress(requestPacket.Handler)

	/////////////////////////////////////////////////
	/////  pack abi for iReceive function  //////////
	/////////////////////////////////////////////////
	abiDef, err := abi.JSON(strings.NewReader(util.CrosschainRequestCheckpointABIJSON))
	if err != nil {
		panic("Bad ABI constant!")
	}

	// the methodName needs to be the same as the 'name' above in the checkpointAbiJson
	// but other than that it's a constant that has no impact on the output. This is because
	// it gets encoded as a function name which we must then discard.
	abiEncodedBatch, err := abiDef.Pack("checkpoint",
		crosschainMethodName,
		routeAmount,
		requestIdentifier,
		requestTimestamp,
		msg.SrcChainId,
		routeRecipient,
		msg.DestChainId,
		asmAddress,
		msg.RequestSender,
		handler,
		requestPacket.Payload,
		metadata.IsReadCall,
	)

	// fmt.Println("Checkpoint- crosschainMethodName ", crosschainMethodName)
	// fmt.Println("Checkpoint-routeAmount", routeAmount)
	// fmt.Println("Checkpoint-requestIdentifier", requestIdentifier)
	// fmt.Println("Checkpoint-requestTimestamp", requestTimestamp)
	// fmt.Println("Checkpoint-msg.SrcChainId", msg.SrcChainId)
	// fmt.Println("Checkpoint-routeRecipient", routeRecipient)
	// fmt.Println("Checkpoint-msg.DestChainId", msg.DestChainId)
	// fmt.Println("Checkpoint-asmAddress", asmAddress)
	// fmt.Println("Checkpoint-msg.RequestSender", msg.RequestSender)

	// fmt.Println("Checkpoint-handler", handler)
	// fmt.Println("Checkpoint-Payload", requestPacket.Payload)
	// fmt.Println("Checkpoint-abiEncodedBatch", abiEncodedBatch)

	// this should never happen outside of test since any case that could crash on encoding
	// should be filtered above.
	if err != nil {
		panic(fmt.Sprintf("Error packing checkpoint! %s/n", err))
	}

	// we hash the resulting encoded bytes discarding the first 4 bytes these 4 bytes are the constant
	// method name 'checkpoint'. If you were to replace the checkpoint constant in this code you would
	// then need to adjust how many bytes you truncate off the front to get the output of abi.encode()
	return crypto.Keccak256Hash(abiEncodedBatch[4:]).Bytes()
}

func (msg CrosschainRequest) GetStarknetCheckpoint(routerIDstring string) []byte {
	//SrcChainId
	srcChainId, _ := utils.HexToFelt(msg.SrcChainId)

	//routeRecipient
	routeRecipient, _ := utils.HexToFelt(msg.RouteRecipient)

	//routeAmount
	routeAmount := BigIntToFeltParts_newarray(msg.RouteAmount.BigInt())

	requestTimestamp := &big.Int{}
	requestTimestamp.SetUint64(uint64(msg.SrcTimestamp))

	//requestTimestamp
	requestTimestamp_ := BigIntToFeltParts_newarray(requestTimestamp)

	metadata := DecodeEvmContractMetadata(&msg)
	//AsmAddress
	asmAddress, _ := utils.HexToFelt(metadata.AsmAddress)

	//DestchainId
	destChainId, _ := utils.HexToFelt(msg.DestChainId)

	//RequestSender
	requestSender_arr := splitStringIntoHalves(msg.RequestSender)
	requestSender_felt, _ := utils.HexArrToFelt(requestSender_arr)

	requestPacket := DecodeRouterCrosschainPacket(&msg)
	//HandlerAddress
	handler_address, _ := utils.HexToFelt(requestPacket.Handler)

	//Packet
	processed_packet, _ := processHexStrings(bytesToHexStrings(requestPacket.Payload))
	packet_felt, err := utils.HexArrToFelt(processed_packet)
	if err != nil {
		fmt.Println(err)
	}

	//IsReadCall
	isReadCall, _ := utils.HexToFelt(boolToHex(metadata.IsReadCall))

	requestIdentifier := &big.Int{}
	requestIdentifier.SetUint64(msg.RequestIdentifier)

	requestIdentifier_felt := BigIntToFeltParts_newarray(requestIdentifier)

	// Run through the elements of the batch and serialize them
	method_name := "0x695265636569766500000000000000"
	method_name_felt, _ := utils.HexToFelt(method_name)

	var serialized_data []*felt.Felt

	empty_string := "0x0"
	empty_string_felt, _ := utils.HexToFelt(empty_string)
	serialized_data = append(serialized_data, empty_string_felt)
	serialized_data = append(serialized_data, method_name_felt)
	serialized_data = append(serialized_data, routeAmount...)
	serialized_data = append(serialized_data, requestIdentifier_felt...)
	serialized_data = append(serialized_data, requestTimestamp_...)
	serialized_data = append(serialized_data, srcChainId)
	serialized_data = append(serialized_data, routeRecipient)
	serialized_data = append(serialized_data, destChainId)
	serialized_data = append(serialized_data, asmAddress)
	serialized_data = generic_append(requestSender_felt, serialized_data)
	serialized_data = append(serialized_data, handler_address)
	serialized_data = generic_append(packet_felt, serialized_data)
	serialized_data = append(serialized_data, isReadCall)

	new_big_int := utils.FeltArrToBigIntArr(serialized_data)

	hashed_val, _ := hashSlice(new_big_int)
	return hashed_val.Bytes()
}

func hashSlice(input []*big.Int) (*big.Int, error) {
	// Check if the input slice is empty and return an error or a specific hash value
	if len(input) == 0 {
		return nil, fmt.Errorf("input slice is empty")
	}

	// Initialize hash with the first element, if specific initialization is needed, adjust accordingly
	hash, _ := starknetgo.Curve.PedersenHash([]*big.Int{big.NewInt(int64(0)), input[0]})

	// Start from the second element since we used the first element to initialize the hash
	for _, val := range input[1:] {

		// Use the current hash and the next element to compute the new hash
		newHash, err := starknetgo.Curve.PedersenHash([]*big.Int{hash, val})
		if err != nil {
			return nil, fmt.Errorf("error computing PedersenHash: %w", err)
		}
		fmt.Println("New Hash : ", newHash)
		hash = newHash // Update hash with the new hash value for the next iteration
	}

	hash, _ = starknetgo.Curve.PedersenHash([]*big.Int{hash, big.NewInt(int64(10))})
	hash, _ = starknetgo.Curve.PedersenHash([]*big.Int{hash, big.NewInt(int64(11))})

	return hash, nil
}

func generic_append(val []*felt.Felt, to_append []*felt.Felt) []*felt.Felt {
	size := utils.Uint64ToFelt(uint64(len(val)))
	to_append = append(to_append, size)
	to_append = append(to_append, val...)
	return to_append
}

func processHexStrings(hexStrings []string) ([]string, error) {
	var processedStrings []string

	for _, hexStr := range hexStrings {
		// Check if the string starts with 0x and has a valid digit as the next character
		if strings.HasPrefix(hexStr, "0x") && len(hexStr) > 2 && hexStr[2] >= '1' && hexStr[2] <= '9' {
			// Remove the 0x prefix to simplify processing
			trimmedStr := hexStr[2:]

			// Find the index of the first occurrence of 00
			zeroIndex := strings.Index(trimmedStr, "00")
			if zeroIndex == -1 {
				// If 00 is not found, append the original string
				processedStrings = append(processedStrings, hexStr)
			} else {
				// Extract the substring up to and including the first occurrence of 00 and prepend 0x
				result := "0x" + trimmedStr[:zeroIndex+2]
				processedStrings = append(processedStrings, result)
			}
		} else {
			// If the string does not meet the criteria, append it as is
			processedStrings = append(processedStrings, hexStr)
		}
	}

	return processedStrings, nil
}

func bytesToHexStrings(bytes []byte) []string {
	var result []string

	for len(bytes) > 0 {
		// Take the first 32 bytes as a chunk
		chunkSize := 32
		if len(bytes) < chunkSize {
			chunkSize = len(bytes)
		}

		// Convert the chunk to a hexadecimal string
		hexStr := hex.EncodeToString(bytes[:chunkSize])

		// Append the hex string to the result
		result = append(result, "0x"+hexStr)

		// Remove the processed bytes from concatenatedBytes
		bytes = bytes[chunkSize:]
	}

	return result
}

func splitStringIntoHalves(input string) []string {
	length := len(input)

	// Calculate the midpoint index
	midpoint := length / 2

	// Split the string into two halves
	firstHalf := input[:midpoint]
	secondHalf := input[midpoint:]

	// Create a string slice to store the halves
	halves := []string{firstHalf, secondHalf}

	return halves
}

func BigIntToFeltParts_newarray(val *big.Int) []*felt.Felt {
	var feltArr []*felt.Felt
	lowHex, highHex := BigIntToHexU128Parts(val)
	low_hex, _ := utils.HexToFelt(lowHex)
	high_hex, _ := utils.HexToFelt(highHex)
	feltArr = append(feltArr, low_hex)
	feltArr = append(feltArr, high_hex)
	return feltArr
}

func BigIntToHexU128Parts(num *big.Int) (lowHex, highHex string) {
	// Create a big.Int with a value of 2^128
	u128 := new(big.Int)
	u128.Exp(big.NewInt(2), big.NewInt(128), nil)

	// Calculate the low part
	low := new(big.Int).And(num, new(big.Int).Sub(u128, big.NewInt(1)))
	lowHex = fmt.Sprintf("0x%032x", low)

	// Calculate the high part
	high := new(big.Int).Rsh(num, 128)
	highHex = fmt.Sprintf("0x%04x", high)

	return lowHex, highHex

}

func boolToHex(b bool) string {
	if b {
		return "0x01"
	}
	return "0x00"
}
