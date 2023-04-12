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
func (msg MsgCrosschainRequest) GetCheckpoint(routerIDstring string) []byte {

	/**
	     Always get checkpoint from crosschain request only.
	 	Crosschain request is dynamic while msg is static.
	**/
	crosschainRequest := NewCrosschainRequestFromMsg(&msg)

	switch crosschainRequest.DestChainType {
	case multichainTypes.CHAIN_TYPE_NEAR:
		return crosschainRequest.GetNearCheckpoint("")
	case multichainTypes.CHAIN_TYPE_COSMOS:
		return nil
	default:
		return crosschainRequest.GetEvmCheckpoint("")
	}
}

func (msg CrosschainRequest) GetNearCheckpoint(routerIDstring string) []byte {
	return msg.GetEvmCheckpoint(routerIDstring)
}

func (msg CrosschainRequest) GetEvmCheckpoint(routerIDstring string) []byte {
	//////////////////////////////////////////////////////////////////////
	/////  Build data with types required for iReceive gateway call  /////
	//////////////////////////////////////////////////////////////////////
	metadata := DecodeEvmMetadata(&msg)

	methodNameBytes := []uint8("iReceive")
	var crosschainMethodName [32]uint8
	copy(crosschainMethodName[:], methodNameBytes)

	routeAmount := msg.RouteAmount.BigInt()

	requestIdentifier := &big.Int{}
	requestIdentifier.SetUint64(msg.EventNonce)

	requestTimestamp := &big.Int{}
	requestTimestamp.SetUint64(uint64(msg.SrcTimestamp))

	routeRecipient := common.BytesToAddress(msg.RouteRecipient)
	asmAddress := common.BytesToAddress(metadata.AsmAddress)

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
		routeRecipient,
		asmAddress,
		msg.SrcChainId,
		msg.DestChainId,
		msg.RequestSender,
		msg.RequestPacket,
		metadata.IsReadCall,
	)

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
