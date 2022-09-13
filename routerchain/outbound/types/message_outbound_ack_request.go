package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	attestationTypes "github.com/router-protocol/sdk-go/routerchain/attestation/types"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
	"github.com/tendermint/tendermint/crypto/tmhash"
)

const TypeMsgOutboundAckRequest = "outbound_ack_request"

var _ sdk.Msg = &MsgOutboundAckRequest{}

func NewMsgOutboundAckRequest(orchestrator string, chainType multichainTypes.ChainType, chainId string, outboundTxNonce uint64, outboundTxRequestedBy string, relayer string, destinationTxHash string, feeConsumed uint64, contractAckResponse []*ContractAckResponse, eventNonce uint64, blockHeight uint64) *MsgOutboundAckRequest {
	return &MsgOutboundAckRequest{
		Orchestrator:          orchestrator,
		ChainType:             chainType,
		ChainId:               chainId,
		OutboundTxNonce:       outboundTxNonce,
		OutboundTxRequestedBy: outboundTxRequestedBy,
		Relayer:               relayer,
		DestinationTxHash:     destinationTxHash,
		FeeConsumed:           feeConsumed,
		ContractAckResponses:  contractAckResponse,
		EventNonce:            eventNonce,
		BlockHeight:           blockHeight,
	}
}

func (msg *MsgOutboundAckRequest) Route() string {
	return RouterKey
}

func (msg *MsgOutboundAckRequest) Type() string {
	return TypeMsgOutboundAckRequest
}

func (msg *MsgOutboundAckRequest) GetSigners() []sdk.AccAddress {
	orchestrator, err := sdk.AccAddressFromBech32(msg.Orchestrator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{orchestrator}
}

func (msg *MsgOutboundAckRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgOutboundAckRequest) ValidateBasic() error {
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
func (msg *MsgOutboundAckRequest) GetType() attestationTypes.ClaimType {
	return attestationTypes.CLAIM_TYPE_INBOUND_REQUEST
}

// Hash implements MsgOutboundAckRequest.Hash
// modify this with care as it is security sensitive. If an element of the claim is not in this hash a single hostile validator
// could engineer a hash collision and execute a version of the claim with any unhashed data changed to benefit them.
// note that the Orchestrator is the only field excluded from this hash, this is because that value is used higher up in the store
// structure for who has made what claim and is verified by the msg ante-handler for signatures
func (msg *MsgOutboundAckRequest) ClaimHash() ([]byte, error) {
	// TODO: @venky check the type for ContractAckResponses
	path := fmt.Sprintf("%d/%s/%d/%s/%s/%s/%d/%v/%d/%d", msg.ChainType, msg.ChainId, msg.OutboundTxNonce, msg.OutboundTxRequestedBy, msg.Relayer, msg.DestinationTxHash, msg.FeeConsumed, msg.ContractAckResponses, msg.EventNonce, msg.BlockHeight)
	return tmhash.Sum([]byte(path)), nil
}

func (msg MsgOutboundAckRequest) GetClaimer() sdk.AccAddress {
	err := msg.ValidateBasic()
	if err != nil {
		panic("MsgOutboundAckRequest failed ValidateBasic! Should have been handled earlier")
	}

	val, _ := sdk.AccAddressFromBech32(msg.Orchestrator)
	return val
}
