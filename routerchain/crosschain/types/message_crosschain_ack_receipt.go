package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	attestationTypes "github.com/router-protocol/sdk-go/routerchain/attestation/types"
)

const TypeMsgCrosschainAckReceipt = "crosschain_ack_receipt"

var _ sdk.Msg = &MsgCrosschainAckReceipt{}

func NewMsgCrosschainAckReceipt(orchestrator string, ackReceiptSrcChainId string, ackReceiptIdentifier uint64, ackReceiptBlockHeight uint64, ackReceiptTxHash string, relayerRouterAddress string, requestIdentifier uint64, ackSrcChainId string, ackRequestIdentifier uint64, feeConsumed uint64) *MsgCrosschainAckReceipt {
	return &MsgCrosschainAckReceipt{
		Orchestrator:          orchestrator,
		AckReceiptSrcChainId:  ackReceiptSrcChainId,
		AckReceiptIdentifier:  ackReceiptIdentifier,
		AckReceiptBlockHeight: ackReceiptBlockHeight,
		AckReceiptTxHash:      ackReceiptTxHash,
		RelayerRouterAddress:  relayerRouterAddress,
		RequestIdentifier:     requestIdentifier,
		AckSrcChainId:         ackSrcChainId,
		AckRequestIdentifier:  ackRequestIdentifier,
		FeeConsumed:           feeConsumed,
	}
}

func (msg *MsgCrosschainAckReceipt) Route() string {
	return RouterKey
}

func (msg *MsgCrosschainAckReceipt) Type() string {
	return TypeMsgCrosschainAckReceipt
}

func (msg *MsgCrosschainAckReceipt) GetSigners() []sdk.AccAddress {
	orchestrator, err := sdk.AccAddressFromBech32(msg.Orchestrator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{orchestrator}
}

func (msg *MsgCrosschainAckReceipt) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCrosschainAckReceipt) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Orchestrator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid orchestrator address (%s)", err)
	}
	return nil
}

func (msg *MsgCrosschainAckReceipt) GetChainId() string {
	return msg.AckReceiptSrcChainId
}

func (msg *MsgCrosschainAckReceipt) GetBlockHeight() uint64 {
	return msg.AckReceiptBlockHeight
}

func (msg *MsgCrosschainAckReceipt) GetEventNonce() uint64 {
	return msg.AckReceiptIdentifier
}

/////////////////////////////
//     Implement Claim     //
/////////////////////////////

// GetType returns the type of the claim
func (msg *MsgCrosschainAckReceipt) GetType() attestationTypes.ClaimType {
	return attestationTypes.CLAIM_TYPE_CROSSCHAIN_ACK_RECEIPT
}

// Hash implements MsgCrosschainRequest.Hash
// modify this with care as it is security sensitive. If an element of the claim is not in this hash a single hostile validator
// could engineer a hash collision and execute a version of the claim with any unhashed data changed to benefit them.
// note that the Orchestrator is the only field excluded from this hash, this is because that value is used higher up in the store
// structure for who has made what claim and is verified by the msg ante-handler for signatures
func (msg *MsgCrosschainAckReceipt) ClaimHash() ([]byte, error) {

	// out, err := proto.Marshal(crosstalkRequestClaimHash)
	// return tmhash.Sum([]byte(out)), err

	return []byte{10}, nil
}

func (msg MsgCrosschainAckReceipt) GetClaimer() sdk.AccAddress {
	err := msg.ValidateBasic()
	if err != nil {
		panic("MsgCrosschainRequest failed ValidateBasic! Should have been handled earlier")
	}

	val, _ := sdk.AccAddressFromBech32(msg.Orchestrator)
	return val
}
