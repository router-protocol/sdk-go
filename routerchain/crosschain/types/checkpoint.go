package types

import (
	fmt "fmt"
	"math/big"
	"strings"

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
func (msg MsgCrosschainRequest) GetCheckpoint(routerIDstring string) ([]byte, error) {
	/**
	     Always get checkpoint from crosschain request only.
	 	Crosschain request is dynamic while msg is static.
	**/
	crosschainRequest := NewCrosschainRequestFromMsg(&msg)

	switch crosschainRequest.DestChainType {
	case multichainTypes.CHAIN_TYPE_NEAR:
		return crosschainRequest.GetNearCheckpoint("")
	case multichainTypes.CHAIN_TYPE_COSMOS:
		return nil, nil
	default:
		return crosschainRequest.GetEvmCheckpoint("")
	}
}

// GetCheckpoint gets the checkpoint signature from the given CrossTalkRequest
func (msg CrosschainRequest) GetCheckpoint(routerIDstring string) ([]byte, error) {
	/**
	     Always get checkpoint from crosschain request only.
	 	Crosschain request is dynamic while msg is static.
	**/

	switch msg.DestChainType {
	case multichainTypes.CHAIN_TYPE_NEAR:
		return msg.GetNearCheckpoint("")
	case multichainTypes.CHAIN_TYPE_COSMOS:
		return nil, nil
	default:
		return msg.GetEvmCheckpoint("")
	}
}

func (msg CrosschainRequest) GetNearCheckpoint(routerIDstring string) ([]byte, error) {
	return msg.GetEvmCheckpoint(routerIDstring)
}

func (msg CrosschainRequest) GetEvmCheckpoint(routerIDstring string) ([]byte, error) {
	//////////////////////////////////////////////////////////////////////
	/////  Build data with types required for iReceive gateway call  /////
	//////////////////////////////////////////////////////////////////////
	fmt.Println("Get EvmCheckpoint")
	metadata := DecodeEvmContractMetadata(&msg)
	requestPacket, err := DecodeRouterCrosschainPacket(&msg)
	if err != nil {
		return nil, err
	}
	fmt.Println("Decoded Metadata", metadata)

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

	fmt.Println("Checkpoint- crosschainMethodName ", crosschainMethodName)
	fmt.Println("Checkpoint-routeAmount", routeAmount)
	fmt.Println("Checkpoint-requestIdentifier", requestIdentifier)
	fmt.Println("Checkpoint-requestTimestamp", requestTimestamp)
	fmt.Println("Checkpoint-msg.SrcChainId", msg.SrcChainId)
	fmt.Println("Checkpoint-routeRecipient", routeRecipient)
	fmt.Println("Checkpoint-msg.DestChainId", msg.DestChainId)
	fmt.Println("Checkpoint-asmAddress", asmAddress)
	fmt.Println("Checkpoint-msg.RequestSender", msg.RequestSender)

	fmt.Println("Checkpoint-handler", handler)
	fmt.Println("Checkpoint-Payload", requestPacket.Payload)
	fmt.Println("Checkpoint-abiEncodedBatch", abiEncodedBatch)

	// this should never happen outside of test since any case that could crash on encoding
	// should be filtered above.
	if err != nil {
		panic(fmt.Sprintf("Error packing checkpoint! %s/n", err))
	}

	// we hash the resulting encoded bytes discarding the first 4 bytes these 4 bytes are the constant
	// method name 'checkpoint'. If you were to replace the checkpoint constant in this code you would
	// then need to adjust how many bytes you truncate off the front to get the output of abi.encode()
	return crypto.Keccak256Hash(abiEncodedBatch[4:]).Bytes(), nil
}
