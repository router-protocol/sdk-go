package types

import (
	"encoding/hex"
	fmt "fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/starknet.go/utils"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
	"github.com/router-protocol/sdk-go/routerchain/util"
)

///////////////////////////////////
//     Implement Ccheckpoint     //
///////////////////////////////////

// GetCheckpoint gets the checkpoint signature from the given MsgCrosschainAckRequest
func (msg MsgCrosschainAckRequest) GetCheckpoint(routerIDstring string) ([]byte, error) {
	/**
	     Always get checkpoint from crosschain request only.
	 	Crosschain request is dynamic while msg is static.
	**/
	crosschainAckRequest := NewCrosschainAckRequestFromMsg(&msg)

	switch crosschainAckRequest.AckDestChainType {
	case multichainTypes.CHAIN_TYPE_NEAR, multichainTypes.CHAIN_TYPE_POLKADOT:
		return crosschainAckRequest.GetNearCheckpoint("")
	case multichainTypes.CHAIN_TYPE_ALEPH_ZERO:
		return crosschainAckRequest.GetAlephZeroCheckpoint("")
	case multichainTypes.CHAIN_TYPE_COSMOS:
		return nil, nil
	case multichainTypes.CHAIN_TYPE_STARKNET:
		return crosschainAckRequest.GetStarknetCheckpoint("")
	default:
		return crosschainAckRequest.GetEvmCheckpoint("")
	}
}

// GetCheckpoint gets the checkpoint signature from the given MsgCrosschainAckRequest
func (msg CrosschainAckRequest) GetCheckpoint(routerIDstring string) ([]byte, error) {
	switch msg.AckDestChainType {
	case multichainTypes.CHAIN_TYPE_NEAR, multichainTypes.CHAIN_TYPE_POLKADOT:
		return msg.GetNearCheckpoint("")
	case multichainTypes.CHAIN_TYPE_ALEPH_ZERO:
		return msg.GetAlephZeroCheckpoint("")
	case multichainTypes.CHAIN_TYPE_COSMOS:
		return nil, nil
	default:
		return msg.GetEvmCheckpoint("")
	}
}

func (msg CrosschainAckRequest) GetAlephZeroCheckpoint(routerIDstring string) ([]byte, error) {
	//////////////////////////////////////////////////////////////////////
	/////  Build data with types required for iAck gateway call  /////
	//////////////////////////////////////////////////////////////////////

	methodNameBytes := []uint8("iAck")
	var crosschainAckMethodName [32]uint8
	copy(crosschainAckMethodName[:], methodNameBytes)

	requestIdentifier := &big.Int{}
	requestIdentifier.SetUint64(msg.RequestIdentifier)

	ackRequestIdentifier := &big.Int{}
	ackRequestIdentifier.SetUint64(msg.AckRequestIdentifier)

	var requestSender [32]uint8
	requestSenderByte, err := hex.DecodeString(strings.TrimPrefix(msg.RequestSender, "0x"))
	if err != nil {
		return nil, err
	}
	copy(requestSender[:], requestSenderByte)

	/////////////////////////////////////////////////
	/////  pack abi for iReceive function  //////////
	/////////////////////////////////////////////////
	abiDef, err := abi.JSON(strings.NewReader(util.CrosschainAckRequestAlephZeroCheckpointABIJSON))
	if err != nil {
		return nil, err
	}

	// the methodName needs to be the same as the 'name' above in the checkpointAbiJson
	// but other than that it's a constant that has no impact on the output. This is because
	// it gets encoded as a function name which we must then discard.
	abiEncodedBatch, err := abiDef.Pack("checkpoint",
		crosschainAckMethodName,
		msg.AckDestChainId,
		requestIdentifier,
		ackRequestIdentifier,
		msg.AckSrcChainId,
		requestSender,
		msg.ExecData,
		msg.ExecStatus,
	)

	// this should never happen outside of test since any case that could crash on encoding
	// should be filtered above.
	if err != nil {
		return nil, err
	}

	// we hash the resulting encoded bytes discarding the first 4 bytes these 4 bytes are the constant
	// method name 'checkpoint'. If you were to replace the checkpoint constant in this code you would
	// then need to adjust how many bytes you truncate off the front to get the output of abi.encode()
	return crypto.Keccak256Hash(abiEncodedBatch[4:]).Bytes(), nil
}

func (msg CrosschainAckRequest) GetNearCheckpoint(routerIDstring string) ([]byte, error) {
	//////////////////////////////////////////////////////////////////////
	/////  Build data with types required for iAck gateway call  /////
	//////////////////////////////////////////////////////////////////////

	methodNameBytes := []uint8("iAck")
	var crosschainAckMethodName [32]uint8
	copy(crosschainAckMethodName[:], methodNameBytes)

	requestIdentifier := &big.Int{}
	requestIdentifier.SetUint64(msg.RequestIdentifier)

	ackRequestIdentifier := &big.Int{}
	ackRequestIdentifier.SetUint64(msg.AckRequestIdentifier)

	requestSender := msg.RequestSender
	/////////////////////////////////////////////////
	/////  pack abi for iReceive function  //////////
	/////////////////////////////////////////////////
	abiDef, err := abi.JSON(strings.NewReader(util.CrosschainAckRequestNearCheckpointABIJSON))
	if err != nil {
		return nil, err
	}

	// the methodName needs to be the same as the 'name' above in the checkpointAbiJson
	// but other than that it's a constant that has no impact on the output. This is because
	// it gets encoded as a function name which we must then discard.
	abiEncodedBatch, err := abiDef.Pack("checkpoint",
		crosschainAckMethodName,
		msg.AckDestChainId,
		requestIdentifier,
		ackRequestIdentifier,
		msg.AckSrcChainId,
		requestSender,
		msg.ExecData,
		msg.ExecStatus,
	)

	// this should never happen outside of test since any case that could crash on encoding
	// should be filtered above.
	if err != nil {
		return nil, err
	}

	// we hash the resulting encoded bytes discarding the first 4 bytes these 4 bytes are the constant
	// method name 'checkpoint'. If you were to replace the checkpoint constant in this code you would
	// then need to adjust how many bytes you truncate off the front to get the output of abi.encode()
	return crypto.Keccak256Hash(abiEncodedBatch[4:]).Bytes(), nil
}

func (msg CrosschainAckRequest) GetEvmCheckpoint(routerIDstring string) ([]byte, error) {
	//////////////////////////////////////////////////////////////////////
	/////  Build data with types required for iAck gateway call  /////
	//////////////////////////////////////////////////////////////////////

	methodNameBytes := []uint8("iAck")
	var crosschainAckMethodName [32]uint8
	copy(crosschainAckMethodName[:], methodNameBytes)

	requestIdentifier := &big.Int{}
	requestIdentifier.SetUint64(msg.RequestIdentifier)

	ackRequestIdentifier := &big.Int{}
	ackRequestIdentifier.SetUint64(msg.AckRequestIdentifier)

	requestSender := common.HexToAddress(msg.RequestSender)
	/////////////////////////////////////////////////
	/////  pack abi for iReceive function  //////////
	/////////////////////////////////////////////////
	abiDef, err := abi.JSON(strings.NewReader(util.CrosschainAckRequestCheckpointABIJSON))
	if err != nil {
		return nil, err
	}

	// the methodName needs to be the same as the 'name' above in the checkpointAbiJson
	// but other than that it's a constant that has no impact on the output. This is because
	// it gets encoded as a function name which we must then discard.
	abiEncodedBatch, err := abiDef.Pack("checkpoint",
		crosschainAckMethodName,
		msg.AckDestChainId,
		requestIdentifier,
		ackRequestIdentifier,
		msg.AckSrcChainId,
		requestSender,
		msg.ExecData,
		msg.ExecStatus,
	)

	// this should never happen outside of test since any case that could crash on encoding
	// should be filtered above.
	if err != nil {
		return nil, err
	}

	// we hash the resulting encoded bytes discarding the first 4 bytes these 4 bytes are the constant
	// method name 'checkpoint'. If you were to replace the checkpoint constant in this code you would
	// then need to adjust how many bytes you truncate off the front to get the output of abi.encode()
	return crypto.Keccak256Hash(abiEncodedBatch[4:]).Bytes(), nil
}

func (msg CrosschainAckRequest) GetStarknetCheckpoint(routerIDstring string) ([]byte, error) {
	method_name := "0x6941636b0000000000000000000000"
	method_name_felt, err := utils.HexToFelt(method_name)
	if err != nil {
		fmt.Println(err)
	}

	requestIdentifier_big := &big.Int{}
	requestIdentifier_big.SetUint64(msg.RequestIdentifier)

	requestIdentifier_felt := BigIntToFeltParts_newarray(requestIdentifier_big)

	ackRequestIdentifier_big := &big.Int{}
	ackRequestIdentifier_big.SetUint64(msg.AckRequestIdentifier)

	ackRequestIdentifier_felt := BigIntToFeltParts_newarray(ackRequestIdentifier_big)

	//RequestSender
	requestSender_arr := splitStringIntoHalves(msg.RequestSender)
	requestSender_felt, err := utils.HexArrToFelt(requestSender_arr)
	if err != nil {
		fmt.Println(err)
	}

	//DestchainId
	destChainId, err := utils.HexToFelt(msg.AckDestChainId)
	if err != nil {
		fmt.Println(err)
	}

	//SoruceChainId
	sourceChainId, err := utils.HexToFelt(msg.AckSrcChainId)
	if err != nil {
		fmt.Println(err)
	}

	//ExecData
	execData_big := utils.BytesToBig(msg.ExecData)
	execData_felt := BigIntToFeltParts_newarray(execData_big)

	//ExecStatus
	execstatus, err := utils.HexToFelt(boolToHex(msg.ExecStatus))
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
	serialized_data = append(serialized_data, sourceChainId)
	serialized_data = append(serialized_data, requestIdentifier_felt...)
	serialized_data = append(serialized_data, ackRequestIdentifier_felt...)
	serialized_data = append(serialized_data, destChainId)
	serialized_data = generic_append(requestSender_felt, serialized_data)
	serialized_data = append(serialized_data, execData_felt...)
	serialized_data = append(serialized_data, execstatus)

	new_big_int := utils.FeltArrToBigIntArr(serialized_data)

	hashed_val, err := hashSlice(new_big_int)
	if err != nil {
		fmt.Println(err)
	}
	ensure_bytes := Ensure32Bytes(hashed_val)
	return ensure_bytes, nil

}

func Ensure32Bytes(hash *big.Int) []byte {
	hashBytes := hash.Bytes()
	if len(hashBytes) < 32 {
		// If hashBytes is less than 32 bytes, create a new slice with leading zeros
		paddedBytes := make([]byte, 32-len(hashBytes), 32)
		paddedBytes = append(paddedBytes, hashBytes...)
		return paddedBytes
	}
	return hashBytes
}
