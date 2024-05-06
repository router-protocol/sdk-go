package types

import (
	"encoding/hex"
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

// GetCheckpoint gets the checkpoint signature from the given MsgCrosschainAckRequest
func (msg MsgCrosschainAckRequest) GetCheckpoint(routerIDstring string) ([]byte, error) {
	/**
	     Always get checkpoint from crosschain request only.
	 	Crosschain request is dynamic while msg is static.
	**/
	crosschainAckRequest := NewCrosschainAckRequestFromMsg(&msg)

	switch crosschainAckRequest.AckDestChainType {
	case multichainTypes.CHAIN_TYPE_NEAR, multichainTypes.CHAIN_TYPE_POLKADOT, multichainTypes.CHAIN_TYPE_BITCOIN:
		return crosschainAckRequest.GetNearCheckpoint("")
	case multichainTypes.CHAIN_TYPE_ALEPH_ZERO:
		return crosschainAckRequest.GetAlephZeroCheckpoint("")
	case multichainTypes.CHAIN_TYPE_COSMOS:
		return nil, nil
	default:
		return crosschainAckRequest.GetEvmCheckpoint("")
	}
}

// GetCheckpoint gets the checkpoint signature from the given MsgCrosschainAckRequest
func (msg CrosschainAckRequest) GetCheckpoint(routerIDstring string) ([]byte, error) {
	switch msg.AckDestChainType {
	case multichainTypes.CHAIN_TYPE_NEAR, multichainTypes.CHAIN_TYPE_POLKADOT, multichainTypes.CHAIN_TYPE_BITCOIN:
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
	requestSenderByte, err := hex.DecodeString(msg.RequestSender)
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
