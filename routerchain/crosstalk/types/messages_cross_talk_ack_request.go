package types

import (
	fmt "fmt"
	"math/big"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
	attestationTypes "github.com/router-protocol/sdk-go/routerchain/attestation/types"
	types "github.com/router-protocol/sdk-go/routerchain/multichain/types"
	"github.com/router-protocol/sdk-go/routerchain/util"
	"github.com/tendermint/tendermint/crypto/tmhash"
)

const (
	TypeMsgCrossTalkAckRequest = "cross_talk_ack_request"
)

var _ sdk.Msg = &MsgCrossTalkAckRequest{}

func NewMsgCrossTalkAckRequest(
	orchestrator string,
	eventNonce uint64,
	blockHeight uint64,
	relayerRouterAddress string,
	sourceChainType types.ChainType,
	sourceChainId string,
	chainType types.ChainType,
	chainId string,
	destinationTxHash string,
	eventIdentifier uint64,
	crosstalkRequestSender []byte,
	crosstalkNonce uint64,
	contractAckResponses []byte,
	exeCode uint64,
	status bool,
	ethSigner string,
	signature string,
	execFlags []bool,
	execData [][]byte,

) *MsgCrossTalkAckRequest {
	return &MsgCrossTalkAckRequest{
		Orchestrator:           orchestrator,
		EventNonce:             eventNonce,
		BlockHeight:            blockHeight,
		RelayerRouterAddress:   relayerRouterAddress,
		SourceChainType:        sourceChainType,
		SourceChainId:          sourceChainId,
		ChainType:              chainType,
		ChainId:                chainId,
		DestinationTxHash:      destinationTxHash,
		EventIdentifier:        eventIdentifier,
		CrosstalkRequestSender: crosstalkRequestSender,
		CrosstalkNonce:         crosstalkNonce,
		ContractAckResponses:   contractAckResponses,
		ExeCode:                exeCode,
		Status:                 status,
		ExecFlags:              execFlags,
		ExecData:               execData,
		EthSigner:              ethSigner,
		Signature:              signature,
	}
}

func (msg *MsgCrossTalkAckRequest) Route() string {
	return RouterKey
}

func (msg *MsgCrossTalkAckRequest) Type() string {
	return TypeMsgCrossTalkAckRequest
}

func (msg *MsgCrossTalkAckRequest) GetSigners() []sdk.AccAddress {
	orchestrator, err := sdk.AccAddressFromBech32(msg.Orchestrator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{orchestrator}
}

func (msg *MsgCrossTalkAckRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCrossTalkAckRequest) ValidateBasic() error {
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
func (msg *MsgCrossTalkAckRequest) GetType() attestationTypes.ClaimType {
	return attestationTypes.CLAIM_TYPE_CROSSTALK_ACK_REQUEST
}

// Hash implements MsgCrossTalkAckRequest.Hash
// modify this with care as it is security sensitive. If an element of the claim is not in this hash a single hostile validator
// could engineer a hash collision and execute a version of the claim with any unhashed data changed to benefit them.
// note that the Orchestrator is the only field excluded from this hash, this is because that value is used higher up in the store
// structure for who has made what claim and is verified by the msg ante-handler for signatures
func (msg *MsgCrossTalkAckRequest) ClaimHash() ([]byte, error) {
	path := fmt.Sprintf("%d/%d/%s/%d/%s/%d/%s/%s/%d/%s/%d/%s/%d/%t", msg.EventNonce, msg.BlockHeight, msg.RelayerRouterAddress, msg.SourceChainType, msg.SourceChainId, msg.ChainType, msg.ChainId, msg.DestinationTxHash, msg.EventIdentifier, msg.CrosstalkRequestSender, msg.CrosstalkNonce, msg.ContractAckResponses, msg.ExeCode, msg.Status)
	return tmhash.Sum([]byte(path)), nil
}

func (msg MsgCrossTalkAckRequest) GetClaimer() sdk.AccAddress {
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
func (msg MsgCrossTalkAckRequest) GetCheckpoint(routerIDstring string) []byte {

	abi, err := abi.JSON(strings.NewReader(util.CrossTalkAckRequestCheckpointABIJSON))
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
	methodNameBytes := []uint8("crossTalkAck")
	var crosstalkMethodName [32]uint8
	copy(crosstalkMethodName[:], methodNameBytes)

	eventIdentifier := &big.Int{}
	eventIdentifier.SetUint64(msg.EventIdentifier)

	crossTalkNonce := &big.Int{}
	crossTalkNonce.SetUint64(msg.CrosstalkNonce)

	chainType := &big.Int{}
	chainType.SetUint64(uint64(msg.SourceChainType))

	destChainType := &big.Int{}
	destChainType.SetUint64(uint64(msg.ChainType))

	caller := []byte(msg.CrosstalkRequestSender)

	execFlags := make([]bool, 1)
	execFlags = append(execFlags, msg.Status)

	execData := make([][]byte, 1)
	execData = append(execData, msg.ContractAckResponses)

	// the methodName needs to be the same as the 'name' above in the checkpointAbiJson
	// but other than that it's a constant that has no impact on the output. This is because
	// it gets encoded as a function name which we must then discard.
	abiEncodedBatch, err := abi.Pack("checkpoint",
		crosstalkMethodName,
		eventIdentifier,
		crossTalkNonce,
		chainType,
		msg.SourceChainId,
		destChainType,
		msg.ChainId,
		caller,
		execFlags,
		execData,
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
