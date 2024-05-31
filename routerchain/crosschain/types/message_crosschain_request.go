package types

import (
	"github.com/cometbft/cometbft/crypto/tmhash"
	proto "github.com/cosmos/gogoproto/proto"
	attestationTypes "github.com/router-protocol/sdk-go/routerchain/attestation/types"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCrosschainRequest = "crosschain_request"

var _ sdk.Msg = &MsgCrosschainRequest{}

func NewMsgCrosschainRequest(
	orchestrator string,
	srcChainId string,
	contract string,
	requestIdentifier uint64,
	srcBlockHeight uint64,
	sourceTxHash string,
	srcTimestamp uint64,
	srcTxOrigin string,
	routeAmount sdkmath.Int,
	routeRecipient string,
	destChainId string,
	requestSender string,
	requestMetadata []byte,
	requestPacket []byte,
	srcChainType multichainTypes.ChainType,
	destChainType multichainTypes.ChainType,
) *MsgCrosschainRequest {
	return &MsgCrosschainRequest{
		Orchestrator:      orchestrator,
		SrcChainId:        srcChainId,
		Contract:          contract,
		SrcChainType:      srcChainType,
		DestChainType:     destChainType,
		RequestIdentifier: requestIdentifier,
		BlockHeight:       srcBlockHeight,
		SourceTxHash:      sourceTxHash,
		SrcTimestamp:      srcTimestamp,
		SrcTxOrigin:       srcTxOrigin,
		RouteAmount:       routeAmount,
		RouteRecipient:    routeRecipient,
		DestChainId:       destChainId,
		RequestSender:     requestSender,
		RequestMetadata:   requestMetadata,
		RequestPacket:     requestPacket,
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
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid orchestrator address (%s)", err)
	}
	return nil
}

func (msg *MsgCrosschainRequest) GetChainType() multichainTypes.ChainType {
	return 1
}

func (msg *MsgCrosschainRequest) GetChainId() string {
	return msg.SrcChainId
}
func (msg *MsgCrosschainRequest) GetEventNonce() uint64 {
	return msg.RequestIdentifier
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
	crosschainRequestClaimHash := NewCrosschainRequestClaimHash(
		msg.SrcChainId,
		msg.Contract,
		msg.RequestIdentifier,
		msg.BlockHeight,
		msg.SourceTxHash,
		msg.SrcTimestamp,
		msg.SrcTxOrigin,
		msg.RouteAmount,
		msg.RouteRecipient,
		msg.DestChainId,
		msg.RequestSender,
		msg.RequestMetadata,
		msg.RequestPacket,
		msg.SrcChainType,
		msg.DestChainType)

	out, err := proto.Marshal(crosschainRequestClaimHash)
	return tmhash.Sum([]byte(out)), err
}

func (msg MsgCrosschainRequest) GetClaimer() sdk.AccAddress {
	err := msg.ValidateBasic()
	if err != nil {
		panic("MsgCrosschainRequest failed ValidateBasic! Should have been handled earlier")
	}

	val, _ := sdk.AccAddressFromBech32(msg.Orchestrator)
	return val
}
