package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	attestationTypes "github.com/router-protocol/sdk-go/routerchain/attestation/types"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
	"github.com/tendermint/tendermint/crypto/tmhash"
)

const TypeMsgInboundRequest = "inbound_request"

var _ sdk.Msg = &MsgInboundRequest{}

func NewMsgInboundRequest(orchestrator string, chainType multichainTypes.ChainType, chainId string, eventNonce uint64, blockHeight uint64, sourceSender string, sourceTxHash string, routerBridgeContract string, payload string) *MsgInboundRequest {
	return &MsgInboundRequest{
		Orchestrator:         orchestrator,
		ChainType:            chainType,
		ChainId:              chainId,
		EventNonce:           eventNonce,
		BlockHeight:          blockHeight,
		SourceSender:         sourceSender,
		SourceTxHash:         sourceTxHash,
		RouterBridgeContract: routerBridgeContract,
		Payload:              payload,
	}
}

func (msg *MsgInboundRequest) Route() string {
	return RouterKey
}

func (msg *MsgInboundRequest) Type() string {
	return TypeMsgInboundRequest
}

func (msg *MsgInboundRequest) GetSigners() []sdk.AccAddress {
	orchestrator, err := sdk.AccAddressFromBech32(msg.Orchestrator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{orchestrator}
}

func (msg *MsgInboundRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgInboundRequest) ValidateBasic() error {
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
func (msg *MsgInboundRequest) GetType() attestationTypes.ClaimType {
	return attestationTypes.CLAIM_TYPE_INBOUND_REQUEST
}

// Hash implements MsgInboundRequest.Hash
// modify this with care as it is security sensitive. If an element of the claim is not in this hash a single hostile validator
// could engineer a hash collision and execute a version of the claim with any unhashed data changed to benefit them.
// note that the Orchestrator is the only field excluded from this hash, this is because that value is used higher up in the store
// structure for who has made what claim and is verified by the msg ante-handler for signatures
func (msg *MsgInboundRequest) ClaimHash() ([]byte, error) {
	path := fmt.Sprintf("%d/%s/%d/%d/%s/%s/%s/%s", msg.ChainType, msg.ChainId, msg.EventNonce, msg.BlockHeight, msg.SourceSender, msg.SourceTxHash, msg.RouterBridgeContract, msg.Payload)
	return tmhash.Sum([]byte(path)), nil
}

func (msg MsgInboundRequest) GetClaimer() sdk.AccAddress {
	err := msg.ValidateBasic()
	if err != nil {
		panic("MsgInboundRequest failed ValidateBasic! Should have been handled earlier")
	}

	val, _ := sdk.AccAddressFromBech32(msg.Orchestrator)
	return val
}
