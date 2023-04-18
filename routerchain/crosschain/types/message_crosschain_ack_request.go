package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	attestationTypes "github.com/router-protocol/sdk-go/routerchain/attestation/types"
)

const TypeMsgCrosschainAckRequest = "crosschain_ack_request"

var _ sdk.Msg = &MsgCrosschainAckRequest{}

func NewMsgCrosschainAckRequest(orchestrator string, ackSrcChainId string, eventNonce uint64, blockHeight uint64, destTxHash string, relayerRouterAddress string, ackDestChainId string, requestSender []byte, requestIdentifier uint64, execData []byte, execStatus bool) *MsgCrosschainAckRequest {
	return &MsgCrosschainAckRequest{
		Orchestrator:         orchestrator,
		AckSrcChainId:        ackSrcChainId,
		EventNonce:           eventNonce,
		BlockHeight:          blockHeight,
		DestTxHash:           destTxHash,
		RelayerRouterAddress: relayerRouterAddress,
		AckDestChainId:       ackDestChainId,
		RequestIdentifier:    requestIdentifier,
		RequestSender:        requestSender,
		ExecData:             execData,
		ExecStatus:           execStatus,
	}
}

func (msg *MsgCrosschainAckRequest) Route() string {
	return RouterKey
}

func (msg *MsgCrosschainAckRequest) Type() string {
	return TypeMsgCrosschainAckRequest
}

func (msg *MsgCrosschainAckRequest) GetSigners() []sdk.AccAddress {
	orchestrator, err := sdk.AccAddressFromBech32(msg.Orchestrator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{orchestrator}
}

func (msg *MsgCrosschainAckRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCrosschainAckRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Orchestrator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid orchestrator address (%s)", err)
	}
	return nil
}

func (msg *MsgCrosschainAckRequest) GetChainId() string {
	return msg.AckSrcChainId
}

/////////////////////////////
//     Implement Claim     //
/////////////////////////////

// GetType returns the type of the claim
func (msg *MsgCrosschainAckRequest) GetType() attestationTypes.ClaimType {
	return attestationTypes.CLAIM_TYPE_CROSSCHAIN_ACK_REQUEST
}

// Hash implements MsgCrosschainRequest.Hash
// modify this with care as it is security sensitive. If an element of the claim is not in this hash a single hostile validator
// could engineer a hash collision and execute a version of the claim with any unhashed data changed to benefit them.
// note that the Orchestrator is the only field excluded from this hash, this is because that value is used higher up in the store
// structure for who has made what claim and is verified by the msg ante-handler for signatures
func (msg *MsgCrosschainAckRequest) ClaimHash() ([]byte, error) {

	// out, err := proto.Marshal(crosstalkRequestClaimHash)
	// return tmhash.Sum([]byte(out)), err

	return []byte{10}, nil
}

func (msg MsgCrosschainAckRequest) GetClaimer() sdk.AccAddress {
	err := msg.ValidateBasic()
	if err != nil {
		panic("MsgCrosschainRequest failed ValidateBasic! Should have been handled earlier")
	}

	val, _ := sdk.AccAddressFromBech32(msg.Orchestrator)
	return val
}
