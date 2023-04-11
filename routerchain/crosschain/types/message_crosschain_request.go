package types

import (
	fmt "fmt"
	"math/big"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	attestationTypes "github.com/router-protocol/sdk-go/routerchain/attestation/types"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
	"github.com/router-protocol/sdk-go/routerchain/util"
)

const TypeMsgCrosschainRequest = "crosschain_request"

var _ sdk.Msg = &MsgCrosschainRequest{}

func NewMsgCrosschainRequest(
	orchestrator string,
	srcChainId string,
	srcEventNonce uint64,
	srcBlockHeight uint64,
	sourceTxHash string,
	srcTimestamp uint64,
	srcTxOrigin string,
	routeAmount sdk.Int,
	routeRecipient []byte,
	destChainId string,
	requestSender []byte,
	requestMetadata []byte,
	requestPacket []byte) *MsgCrosschainRequest {
	return &MsgCrosschainRequest{
		Orchestrator:    orchestrator,
		SrcChainId:      srcChainId,
		EventNonce:      srcEventNonce,
		BlockHeight:     srcBlockHeight,
		SourceTxHash:    sourceTxHash,
		SrcTimestamp:    srcTimestamp,
		SrcTxOrigin:     srcTxOrigin,
		RouteAmount:     routeAmount,
		RouteRecipient:  routeRecipient,
		DestChainId:     destChainId,
		RequestSender:   requestSender,
		RequestMetadata: requestMetadata,
		RequestPacket:   requestPacket,
	}
}

func (msg *MsgCrosschainRequest) Route() string {
	return RouterKey
}

func (msg *MsgCrosschainRequest) Type() string {
	return TypeMsgCrosschainRequest
}

func (msg *MsgCrosschainRequest) GetSigners() []sdk.AccAddress {
	orchestrator, err := sdk.AccAddressFromBech32(msg.Orchestrator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{orchestrator}
}

func (msg *MsgCrosschainRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCrosschainRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Orchestrator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid orchestrator address (%s)", err)
	}
	return nil
}

func (msg *MsgCrosschainRequest) GetChainType() multichainTypes.ChainType {
	return 1
}

func (msg *MsgCrosschainRequest) GetChainId() string {
	return msg.SrcChainId
}

/////////////////////////////
//     Implement Claim     //
/////////////////////////////

// GetType returns the type of the claim
func (msg *MsgCrosschainRequest) GetType() attestationTypes.ClaimType {
	return attestationTypes.CLAIM_TYPE_CROSSCHAIN_REQUEST
}

// Hash implements MsgCrosschainRequest.Hash
// modify this with care as it is security sensitive. If an element of the claim is not in this hash a single hostile validator
// could engineer a hash collision and execute a version of the claim with any unhashed data changed to benefit them.
// note that the Orchestrator is the only field excluded from this hash, this is because that value is used higher up in the store
// structure for who has made what claim and is verified by the msg ante-handler for signatures
func (msg *MsgCrosschainRequest) ClaimHash() ([]byte, error) {

	// out, err := proto.Marshal(crosstalkRequestClaimHash)
	// return tmhash.Sum([]byte(out)), err

	return []byte{10}, nil
}

func (msg MsgCrosschainRequest) GetClaimer() sdk.AccAddress {
	err := msg.ValidateBasic()
	if err != nil {
		panic("MsgCrosschainRequest failed ValidateBasic! Should have been handled earlier")
	}

	val, _ := sdk.AccAddressFromBech32(msg.Orchestrator)
	return val
}

///////////////////////////////////
//     Implement Ccheckpoint     //
///////////////////////////////////

// GetCheckpoint gets the checkpoint signature from the given CrossTalkRequest
func (msg MsgCrosschainRequest) GetCheckpoint(routerIDstring string) []byte {
	metadata := msg.DecodeMetadata()
	abiDef, err := abi.JSON(strings.NewReader(util.CrosschainRequestCheckpointABIJSON))
	if err != nil {
		panic("Bad ABI constant!")
	}

	// the contract argument is not a arbitrary length array but a fixed length 32 byte
	// array, therefore we have to utf8 encode the string (the default in this case) and
	// then copy the variable length encoded data into a fixed length array. This function
	// will panic if gravityId is too long to fit in 32 bytes
	// routerID, err := util.StrToFixByteArray(routerIDstring)
	// if err != nil {
	// 	panic(err)
	// }

	// Create the methodName argument which salts the signature
	methodNameBytes := []uint8("iReceive")
	var crosschainMethodName [32]uint8
	copy(crosschainMethodName[:], methodNameBytes)

	routeAmount := msg.RouteAmount.BigInt()

	requestIdentifier := &big.Int{}
	requestIdentifier.SetUint64(msg.EventNonce)

	requestTimestamp := &big.Int{}
	requestTimestamp.SetUint64(uint64(msg.SrcTimestamp))

	routeRecipient := common.BytesToAddress(msg.RouteRecipient)
	asmAddress := common.BytesToAddress(metadata.asmAddress)

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
		metadata.isReadCall,
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

type CrosschainEVMMetadata struct {
	destGasLimit big.Int
	destGasPrice big.Int
	ackGasLimit  big.Int
	ackGasPrice  big.Int
	relayerFees  big.Int
	ackType      uint8
	isReadCall   bool
	asmAddress   []byte
}

func (msg MsgCrosschainRequest) DecodeMetadata() CrosschainEVMMetadata {
	abiDef, err := abi.JSON(strings.NewReader(util.CrosschainRequestMetadataABIJSON))
	if err != nil {
		panic("Bad ABI constant!")
	}

	var crosschainMetadata CrosschainEVMMetadata

	err = abiDef.UnpackIntoInterface(crosschainMetadata, "metadata", msg.RequestMetadata)
	if err != nil {
		panic(err)
	}

	return crosschainMetadata
}

func (msg CrosschainRequest) DecodeMetadata() CrosschainEVMMetadata {
	abiDef, err := abi.JSON(strings.NewReader(util.CrosschainRequestMetadataABIJSON))
	if err != nil {
		panic("Bad ABI constant!")
	}

	var crosschainMetadata CrosschainEVMMetadata

	err = abiDef.UnpackIntoInterface(crosschainMetadata, "metadata", msg.RequestMetadata)
	if err != nil {
		panic(err)
	}

	return crosschainMetadata
}
