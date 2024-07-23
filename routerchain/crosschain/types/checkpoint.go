package types

import (
	"encoding/binary"
	"encoding/hex"
	fmt "fmt"
	"math"
	"math/big"
	"strings"

	"github.com/NethermindEth/juno/core/felt"
	starknetgo "github.com/NethermindEth/starknet.go/curve"
	"github.com/NethermindEth/starknet.go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
	"github.com/router-protocol/sdk-go/routerchain/util"
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
	case multichainTypes.CHAIN_TYPE_NEAR, multichainTypes.CHAIN_TYPE_POLKADOT:
		return crosschainRequest.GetNearCheckpoint("")
	case multichainTypes.CHAIN_TYPE_COSMOS:
		return nil
	case multichainTypes.CHAIN_TYPE_STARKNET:
		return crosschainRequest.GetStarknetCheckpoint("")
	case multichainTypes.CHAIN_TYPE_ALEPH_ZERO:
		return crosschainRequest.GetAlephZeroCheckpoint("")
	case multichainTypes.CHAIN_TYPE_SOLANA:
		return crosschainRequest.GetSolanaCheckpoint("")
	case multichainTypes.CHAIN_TYPE_SUI:
		return crosschainRequest.GetSuiCheckpoint("")
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
	case multichainTypes.CHAIN_TYPE_NEAR, multichainTypes.CHAIN_TYPE_POLKADOT:
		return msg.GetNearCheckpoint("")
	case multichainTypes.CHAIN_TYPE_COSMOS:
		return nil
	case multichainTypes.CHAIN_TYPE_STARKNET:
		return msg.GetStarknetCheckpoint("")
	case multichainTypes.CHAIN_TYPE_ALEPH_ZERO:
		return msg.GetAlephZeroCheckpoint("")
	case multichainTypes.CHAIN_TYPE_SOLANA:
		return msg.GetSolanaCheckpoint("")
	case multichainTypes.CHAIN_TYPE_SUI:
		return msg.GetSuiCheckpoint("")
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
	routeRecipientByte, err := hex.DecodeString(strings.TrimPrefix(msg.RouteRecipient, "0x"))
	if err != nil {
		return nil
	}
	copy(routeRecipient[:], routeRecipientByte)

	var asmAddress [32]uint8
	asmAddressByte, err := hex.DecodeString(strings.TrimPrefix(metadata.AsmAddress, "0x"))
	if err != nil {
		return nil
	}
	copy(asmAddress[:], asmAddressByte)

	var handler [32]uint8
	handlerByte, err := hex.DecodeString(strings.TrimPrefix(requestPacket.Handler, "0x"))
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
	fmt.Println("Inside GetStarknetCheckpoint", "reqIdentifier", msg.RequestIdentifier)
	//SrcChainId
	srcChainId, err := utils.HexToFelt(msg.SrcChainId)
	if err != nil {
		fmt.Println(err)
	}

	//routeRecipient
	var routeRecipient *felt.Felt
	if msg.RouteRecipient == "" {
		// If RouteRecipient is empty, set it to 0x0
		var err error
		routeRecipient, err = utils.HexToFelt("0x0")
		if err != nil {
			fmt.Println("Error setting routeRecipient to 0x0:", err)
		}
	} else {
		// Convert msg.RouteRecipient to felt
		var err error
		routeRecipient, err = utils.HexToFelt(msg.RouteRecipient)
		if err != nil {
			fmt.Println("Error converting routeRecipient:", err)
		}
	}

	//routeAmount
	routeAmount := BigIntToFeltParts_newarray(msg.RouteAmount.BigInt())

	requestTimestamp := &big.Int{}
	requestTimestamp.SetUint64(uint64(msg.SrcTimestamp))

	//requestTimestamp
	requestTimestamp_ := BigIntToFeltParts_newarray(requestTimestamp)

	metadata := DecodeEvmContractMetadata(&msg)
	//AsmAddress
	asmAddress, err := utils.HexToFelt(metadata.AsmAddress)
	if err != nil {
		fmt.Println(err)
	}

	//DestchainId
	destChainId, err := utils.HexToFelt(msg.DestChainId)
	if err != nil {
		fmt.Println(err)
	}

	//RequestSender
	requestSender_arr := splitStringIntoHalves(msg.RequestSender)
	requestSender_felt, err := utils.HexArrToFelt(requestSender_arr)
	if err != nil {
		fmt.Println(err)
	}

	requestPacket := DecodeRouterCrosschainPacket(&msg)
	//HandlerAddress
	handler_address, err := utils.HexToFelt(requestPacket.Handler)
	if err != nil {
		fmt.Println(err)
	}

	//Packet
	processed_packet, err := processHexStrings(bytesToHexStrings(requestPacket.Payload))
	if err != nil {
		fmt.Println(err)
	}
	packet_felt, err := utils.HexArrToFelt(processed_packet)
	if err != nil {
		fmt.Println(err)
	}

	//IsReadCall
	isReadCall, err := utils.HexToFelt(boolToHex(metadata.IsReadCall))
	if err != nil {
		fmt.Println(err)
	}

	requestIdentifier := &big.Int{}
	requestIdentifier.SetUint64(msg.RequestIdentifier)

	requestIdentifier_felt := BigIntToFeltParts_newarray(requestIdentifier)

	// Run through the elements of the batch and serialize them
	method_name := "0x695265636569766500000000000000"
	method_name_felt, err := utils.HexToFelt(method_name)
	if err != nil {
		fmt.Println(err)
	}

	var serialized_data []*felt.Felt

	empty_string := "0x0"
	empty_string_felt, err := utils.HexToFelt(empty_string)
	if err != nil {
		fmt.Println(err)
	}
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

	hashed_val, err := hashSlice(new_big_int)
	if err != nil {
		fmt.Println(err)
	}

	ensure_bytes := Ensure32Bytes(hashed_val)
	return ensure_bytes
}

func (msg CrosschainRequest) GetSuiCheckpoint(routerIDstring string) []byte {
    //////////////////////////////////////////////////////////////////////
    /////  Build data with types required for iReceive gateway call  /////
    //////////////////////////////////////////////////////////////////////
    metadata := DecodeEvmContractMetadata(&msg)
    requestPacket := DecodeRouterCrosschainPacket(&msg)

    parsedABI, err := abi.JSON(strings.NewReader(util.CrosschainRequestSuiCheckpointABIJSON))
    if err != nil {
        panic("Bad ABI constant!")
    }

    crosschainMethodName := make([]byte, 32)
    copy(crosschainMethodName, "iReceive")

    routeAmount := msg.RouteAmount.BigInt()

    requestIdentifier := new(big.Int).SetUint64(msg.RequestIdentifier)

    requestTimestamp := new(big.Int).SetUint64(uint64(msg.SrcTimestamp))

    srcChainID := msg.SrcChainId

    destChainID := msg.DestChainId

    routeRecipient := make([]byte, 32)
    if msg.RouteRecipient != "" {
        routeRecipient = common.HexToAddress(msg.RouteRecipient).Bytes()
    }

	asmAddress := []byte{}
    if metadata.AsmAddress != "" {
        asmAddress = common.HexToAddress(metadata.AsmAddress).Bytes()
    }

    requestSender := common.HexToAddress(msg.RequestSender).Hex()

    handlerAddress, err := hex.DecodeString(strings.TrimPrefix(requestPacket.Handler, "0x"))
    if err != nil {
        return nil
    }

    packet := requestPacket.Payload
    isReadCall := metadata.IsReadCall

    // Pack the arguments to generate the ABI-encoded bytes
    abiEncodedBatch, err := parsedABI.Pack("checkpoint",
        crosschainMethodName,
        routeAmount,
        requestIdentifier,
        requestTimestamp,
        srcChainID,
        routeRecipient,
        destChainID,
        asmAddress,
        requestSender,
        handlerAddress,
        packet,
        isReadCall,
    )
    if err != nil {
        panic(fmt.Sprintf("Error packing checkpoint! %s/n", err))
    }

    return crypto.Keccak256Hash(abiEncodedBatch[4:]).Bytes()
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

	hash, _ = starknetgo.Curve.PedersenHash([]*big.Int{hash, big.NewInt(int64(112))})
	hash, _ = starknetgo.Curve.PedersenHash([]*big.Int{hash, big.NewInt(int64(11))})
	fmt.Println("Final Hash : ", hash)

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

func (msg CrosschainRequest) GetSolanaCheckpoint(routerIDstring string) []byte {
	//////////////////////////////////////////////////////////////////////
	/////  Build data with types required for iReceive gateway call  /////
	//////////////////////////////////////////////////////////////////////
	const CHUNK_LIMIT = 800
	metadata := DecodeEvmContractMetadata(&msg)
	requestPacket := DecodeRouterCrosschainPacket(&msg)
	srcChainIdBytes := []byte(msg.SrcChainId)
	destChainIdBytes := []byte(msg.DestChainId)
	requestSenderBytes := []byte(msg.RequestSender)
	routeRecipientBytes, err := utils.HexToBytes(msg.RouteRecipient)
	if err != nil {
		panic(fmt.Sprintf("Error decoding route recipient: %s", err.Error()))
	}
	handlerAddressBytes, err := utils.HexToBytes(requestPacket.Handler)
	if err != nil {
		panic(fmt.Sprintf("Error decoding handler address: %s", err.Error()))
	}
	asmAddressBytes, err := utils.HexToBytes(metadata.AsmAddress)
	if err != nil {
		panic(fmt.Sprintf("Error decoding asm address: %s", err.Error()))
	}
	// Calculate hash before attesting packet
	length := 32 + 8 + 16 + 16 + len(srcChainIdBytes) + len(routeRecipientBytes) +
		len(destChainIdBytes) + len(asmAddressBytes) + len(requestSenderBytes) +
		len(handlerAddressBytes) + 1
	data := make([]byte, length)
	offset := 0

	// add method name
	copy(data[offset:], []byte{105, 82, 101, 99, 101, 105, 118, 101, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	offset += 32

	// add route amount
	binary.LittleEndian.PutUint64(data[offset:], msg.RouteAmount.Uint64())
	offset += 8

	// Add request identifier
	requestIdentifierBytes := make([]byte, 16)
	binary.LittleEndian.PutUint64(requestIdentifierBytes, msg.RequestIdentifier)
	copy(data[offset:], requestIdentifierBytes)
	offset += 16

	// add request timestamp
	requestTimestampBytes := make([]byte, 16)
	binary.LittleEndian.PutUint64(requestTimestampBytes, msg.SrcTimestamp)
	copy(data[offset:], requestTimestampBytes)
	offset += 16

	// add chin id
	copy(data[offset:], srcChainIdBytes)
	offset += len(srcChainIdBytes)

	// add dest chain
	copy(data[offset:], destChainIdBytes)
	offset += len(destChainIdBytes)

	// add asm address
	copy(data[offset:], asmAddressBytes)
	offset += len(asmAddressBytes)

	// add request sender
	copy(data[offset:], requestSenderBytes)
	offset += len(requestSenderBytes)

	// add handler address
	copy(data[offset:], handlerAddressBytes)
	offset += len(handlerAddressBytes)

	if metadata.IsReadCall {
		data[offset] = 1
	} else {
		data[offset] = 0
	}

	// Calculate hash after attesting packet
	packet := requestPacket.Payload
	packetLength := len(packet)
	result := crypto.Keccak256Hash(data).Bytes()
	times := int(math.Floor(float64(packetLength)/float64(CHUNK_LIMIT))) + func() int {
		if packetLength%CHUNK_LIMIT == 0 {
			return 0
		}
		return 1
	}()

	for idx := 0; idx < times; idx++ {
		from := idx * CHUNK_LIMIT
		to := (idx + 1) * CHUNK_LIMIT
		if to > packetLength {
			to = packetLength
		}
		buf := make([]byte, 32+to-from)
		copy(buf[:32], result)
		copy(buf[32:], packet[from:to])
		result = crypto.Keccak256Hash(buf).Bytes()
	}

	if packetLength == 0 {
		buf := make([]byte, 32)
		copy(buf[:32], result)
		result = crypto.Keccak256Hash(buf).Bytes()
	}

	return result
}

func boolToHex(b bool) string {
	if b {
		return "0x01"
	}
	return "0x00"
}
