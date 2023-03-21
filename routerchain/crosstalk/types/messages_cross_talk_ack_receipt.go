package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	proto "github.com/gogo/protobuf/proto"
	attestationTypes "github.com/router-protocol/sdk-go/routerchain/attestation/types"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
	"github.com/tendermint/tendermint/crypto/tmhash"
)

const (
	TypeMsgCrossTalkAckReceipt = "create_cross_talk_ack_receipt"
)

var _ sdk.Msg = &MsgCrossTalkAckReceipt{}

func NewMsgCrossTalkAckReceipt(
	orchestrator string,
	eventNonce uint64,
	blockHeight uint64,
	relayerRouterAddress string,
	chainType multichainTypes.ChainType,
	chainId string,
	txHash string,
	feeConsumed uint64,
	eventIdentifier uint64,

) *MsgCrossTalkAckReceipt {
	return &MsgCrossTalkAckReceipt{
		Orchestrator:         orchestrator,
		EventNonce:           eventNonce,
		BlockHeight:          blockHeight,
		RelayerRouterAddress: relayerRouterAddress,
		ChainType:            chainType,
		ChainId:              chainId,
		TxHash:               txHash,
		FeeConsumed:          feeConsumed,
		EventIdentifier:      eventIdentifier,
	}
}

func (msg *MsgCrossTalkAckReceipt) Route() string {
	return RouterKey
}

func (msg *MsgCrossTalkAckReceipt) Type() string {
	return TypeMsgCrossTalkAckReceipt
}

func (msg *MsgCrossTalkAckReceipt) GetSigners() []sdk.AccAddress {
	orchestrator, err := sdk.AccAddressFromBech32(msg.Orchestrator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{orchestrator}
}

func (msg *MsgCrossTalkAckReceipt) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCrossTalkAckReceipt) ValidateBasic() error {
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
func (msg *MsgCrossTalkAckReceipt) GetType() attestationTypes.ClaimType {
	return attestationTypes.CLAIM_TYPE_CROSSTALK_ACK_RECEIPT
}

// Hash implements MsgCrossTalkAckReceipt.Hash
// modify this with care as it is security sensitive. If an element of the claim is not in this hash a single hostile validator
// could engineer a hash collision and execute a version of the claim with any unhashed data changed to benefit them.
// note that the Orchestrator is the only field excluded from this hash, this is because that value is used higher up in the store
// structure for who has made what claim and is verified by the msg ante-handler for signatures
func (msg *MsgCrossTalkAckReceipt) ClaimHash() ([]byte, error) {
	crossTalkAckReceiptClaimHash := NewCrossTalkAckReceiptClaimHash(
		msg.ChainType,
		msg.ChainId,
		msg.EventNonce,
		msg.BlockHeight,
		msg.RelayerRouterAddress,
		msg.TxHash,
		msg.EventIdentifier,
	)
	out, err := proto.Marshal(crossTalkAckReceiptClaimHash)
	return tmhash.Sum([]byte(out)), err
}

func (msg MsgCrossTalkAckReceipt) GetClaimer() sdk.AccAddress {
	err := msg.ValidateBasic()
	if err != nil {
		panic("MsgCrossTalkAckReceipt failed ValidateBasic! Should have been handled earlier")
	}

	val, _ := sdk.AccAddressFromBech32(msg.Orchestrator)
	return val
}
