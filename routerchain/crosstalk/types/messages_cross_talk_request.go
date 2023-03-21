package types

import (
	fmt "fmt"
	"math/big"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
	proto "github.com/gogo/protobuf/proto"
	"github.com/tendermint/tendermint/crypto/tmhash"

	attestationTypes "github.com/router-protocol/sdk-go/routerchain/attestation/types"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
	"github.com/router-protocol/sdk-go/routerchain/util"
)

const (
	TypeMsgCrossTalkRequest = "cross_talk_request"
)

var _ sdk.Msg = &MsgCrossTalkRequest{}

func NewMsgCrossTalkRequest(
	orchestrator string,
	eventNonce uint64,
	blockHeight uint64,
	sourceChainType multichainTypes.ChainType,
	sourceChainId string,
	sourceTxHash string,
	sourceTimeStamp uint64,
	destinationChainType multichainTypes.ChainType,
	destinationChainId string,
	destinationGasLimit uint64,
	destinationGasPrice uint64,
	requestSender []byte,
	requestNonce uint64,
	isAtomic bool,
	expiryTimestamp uint64,
	destContractAddresses [][]byte,
	destContractPayloads [][]byte,
	ackType CrossTalkRequestAckType,
	ackGasLimit uint64,
	ackGasPrice uint64,
	ethSigner string,
	signature string,
	requestTxOrigin string,
	isReadCall bool,
	feePayer []byte,
	asmAddress []byte,

) *MsgCrossTalkRequest {
	return &MsgCrossTalkRequest{
		Orchestrator:          orchestrator,
		EventNonce:            eventNonce,
		BlockHeight:           blockHeight,
		ChainType:             sourceChainType,
		ChainId:               sourceChainId,
		SourceTxHash:          sourceTxHash,
		SourceTimestamp:       sourceTimeStamp,
		DestinationChainType:  destinationChainType,
		DestinationChainId:    destinationChainId,
		DestinationGasLimit:   destinationGasLimit,
		DestinationGasPrice:   destinationGasPrice,
		RequestSender:         requestSender,
		RequestTxOrigin:       requestTxOrigin,
		IsReadCall:            isReadCall,
		RequestNonce:          requestNonce,
		IsAtomic:              isAtomic,
		ExpiryTimestamp:       expiryTimestamp,
		DestContractAddresses: destContractAddresses,
		DestContractPayloads:  destContractPayloads,
		AckType:               ackType,
		AckGasLimit:           ackGasLimit,
		AckGasPrice:           ackGasPrice,
		EthSigner:             ethSigner,
		Signature:             signature,
		FeePayer:              feePayer,
		AsmAddress:            asmAddress,
	}
}

func (msg *MsgCrossTalkRequest) Route() string {
	return RouterKey
}

func (msg *MsgCrossTalkRequest) Type() string {
	return TypeMsgCrossTalkRequest
}

func (msg *MsgCrossTalkRequest) GetSigners() []sdk.AccAddress {
	orchestrator, err := sdk.AccAddressFromBech32(msg.Orchestrator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{orchestrator}
}

func (msg *MsgCrossTalkRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCrossTalkRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Orchestrator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid orchestrator address (%s)", err)
	}
	return nil
}

/////////////////////////////
//     Implement Claim     //
/////////////////////////////

// GetType returns the type of the claim
func (msg *MsgCrossTalkRequest) GetType() attestationTypes.ClaimType {
	return attestationTypes.CLAIM_TYPE_CROSSTALK_REQUEST
}

// Hash implements MsgCrossTalkRequest.Hash
// modify this with care as it is security sensitive. If an element of the claim is not in this hash a single hostile validator
// could engineer a hash collision and execute a version of the claim with any unhashed data changed to benefit them.
// note that the Orchestrator is the only field excluded from this hash, this is because that value is used higher up in the store
// structure for who has made what claim and is verified by the msg ante-handler for signatures
func (msg *MsgCrossTalkRequest) ClaimHash() ([]byte, error) {
	crosstalkRequestClaimHash := NewCrossTalkRequestClaimHash(
		msg.EventNonce,
		msg.BlockHeight,
		msg.ChainType,
		msg.ChainId,
		msg.SourceTxHash,
		msg.SourceTimestamp,
		msg.DestinationChainType,
		msg.DestinationChainId,
		msg.DestinationGasLimit,
		msg.DestinationGasPrice,
		msg.RequestSender,
		msg.RequestNonce,
		msg.IsAtomic,
		msg.ExpiryTimestamp,
		msg.DestContractAddresses,
		msg.DestContractPayloads,
		msg.AckType,
		msg.AckGasLimit,
		msg.AckGasPrice,
		msg.RequestTxOrigin,
		msg.IsReadCall,
		msg.FeePayer,
		msg.AsmAddress)

	out, err := proto.Marshal(crosstalkRequestClaimHash)
	return tmhash.Sum([]byte(out)), err
}

func (msg MsgCrossTalkRequest) GetClaimer() sdk.AccAddress {
	err := msg.ValidateBasic()
	if err != nil {
		panic("MsgCrossTalkRequest failed ValidateBasic! Should have been handled earlier")
	}

	val, _ := sdk.AccAddressFromBech32(msg.Orchestrator)
	return val
}

///////////////////////////////////
//     Implement Ccheckpoint     //
///////////////////////////////////

// GetCheckpoint gets the checkpoint signature from the given CrossTalkRequest
func (msg MsgCrossTalkRequest) GetCheckpoint(routerIDstring string) []byte {

	abi, err := abi.JSON(strings.NewReader(util.CrossTalkRequestCheckpointABIJSON))
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
	methodNameBytes := []uint8("requestFromSource")
	var crosstalkMethodName [32]uint8
	copy(crosstalkMethodName[:], methodNameBytes)

	eventIdentifier := &big.Int{}
	eventIdentifier.SetUint64(msg.EventNonce)

	crossTalkNonce := &big.Int{}
	crossTalkNonce.SetUint64(msg.RequestNonce)

	destChainType := &big.Int{}
	destChainType.SetUint64(uint64(msg.DestinationChainType))

	srcChainType := &big.Int{}
	srcChainType.SetUint64(uint64(msg.ChainType))

	caller := []byte(msg.RequestSender)

	expTimestamp := &big.Int{}
	expTimestamp.SetUint64(uint64(msg.ExpiryTimestamp))

	chainTimestamp := &big.Int{}
	chainTimestamp.SetUint64(uint64(msg.SourceTimestamp))

	// the methodName needs to be the same as the 'name' above in the checkpointAbiJson
	// but other than that it's a constant that has no impact on the output. This is because
	// it gets encoded as a function name which we must then discard.
	abiEncodedBatch, err := abi.Pack("checkpoint",
		crosstalkMethodName,
		eventIdentifier,
		crossTalkNonce,
		destChainType,
		msg.DestinationChainId,
		msg.ChainId,
		srcChainType,
		caller,
		msg.IsAtomic,
		chainTimestamp,
		expTimestamp,
		msg.AsmAddress,
		msg.DestContractAddresses,
		msg.DestContractPayloads,
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
