package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/crypto/tmhash"

	attestationTypes "github.com/router-protocol/sdk-go/routerchain/attestation/types"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
)

const (
	TypeMsgCrossTalkRequest = "cross_talk_request"
)

var _ sdk.Msg = &MsgCrossTalkRequest{}

func NewMsgCrossTalkRequest(
	orchestrator string,
	index string,
	eventNonce uint64,
	blockHeight uint64,
	sourceChainType multichainTypes.ChainType,
	sourceChainId string,
	destinationChainType multichainTypes.ChainType,
	destinationChainId string,
	requestSender string,
	requestNonce uint64,
	isAtomic bool,
	gasLimit uint64,
	gasPrice uint64,
	expiryTimestamp uint64,
	ethSigner string,
	signature string,

) *MsgCrossTalkRequest {
	return &MsgCrossTalkRequest{
		Orchestrator:         orchestrator,
		EventNonce:           eventNonce,
		BlockHeight:          blockHeight,
		ChainType:            sourceChainType,
		ChainId:              sourceChainId,
		DestinationChainType: destinationChainType,
		DestinationChainId:   destinationChainId,
		RequestSender:        requestSender,
		RequestNonce:         requestNonce,
		IsAtomic:             isAtomic,
		GasLimit:             gasLimit,
		GasPrice:             gasPrice,
		ExpiryTimestamp:      expiryTimestamp,
		EthSigner:            ethSigner,
		Signature:            signature,
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
	path := fmt.Sprintf("%d/%d/%d/%d/%d/%d/%s/%d/%s/%s/%t/%d/%d/%d/%s/%s", msg.EventNonce, msg.BlockHeight, msg.ChainType, msg.ChainId, msg.DestinationChainType, msg.DestinationChainId, msg.RequestSender, msg.RequestNonce, msg.IsAtomic, msg.GasLimit, msg.GasPrice, msg.ExpiryTimestamp, msg.EthSigner, msg.Signature)
	return tmhash.Sum([]byte(path)), nil
}

func (msg MsgCrossTalkRequest) GetClaimer() sdk.AccAddress {
	err := msg.ValidateBasic()
	if err != nil {
		panic("MsgCrossTalkRequest failed ValidateBasic! Should have been handled earlier")
	}

	val, _ := sdk.AccAddressFromBech32(msg.Orchestrator)
	return val
}
