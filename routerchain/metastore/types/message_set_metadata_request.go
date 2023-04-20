package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	proto "github.com/gogo/protobuf/proto"
	attestationTypes "github.com/router-protocol/sdk-go/routerchain/attestation/types"
	"github.com/tendermint/tendermint/crypto/tmhash"
)

const TypeMsgCreateMetadataRequest = "set_metadata_request"

var _ sdk.Msg = &MsgCreateMetadataRequest{}

func NewMsgCreateMetadataRequest(orchestrator string, chainId string, eventNonce uint64, blockHeight uint64, daapAddress string, feePayer string) *MsgCreateMetadataRequest {
	return &MsgCreateMetadataRequest{
		Orchestrator: orchestrator,
		ChainId:      chainId,
		EventNonce:   eventNonce,
		BlockHeight:  blockHeight,
		DaapAddress:  daapAddress,
		FeePayer:     feePayer,
	}
}

func (msg *MsgCreateMetadataRequest) Route() string {
	return RouterKey
}

func (msg *MsgCreateMetadataRequest) Type() string {
	return TypeMsgCreateMetadataRequest
}

func (msg *MsgCreateMetadataRequest) GetSigners() []sdk.AccAddress {
	orchestrator, err := sdk.AccAddressFromBech32(msg.Orchestrator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{orchestrator}
}

func (msg *MsgCreateMetadataRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateMetadataRequest) ValidateBasic() error {
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
func (msg *MsgCreateMetadataRequest) GetType() attestationTypes.ClaimType {
	return attestationTypes.CLAIM_TYPE_SET_METADATA_REQUEST
}

// Hash implements MsgCreateMetadataRequest.Hash
// modify this with care as it is security sensitive. If an element of the claim is not in this hash a single hostile validator
// could engineer a hash collision and execute a version of the claim with any unhashed data changed to benefit them.
// note that the Orchestrator is the only field excluded from this hash, this is because that value is used higher up in the store
// structure for who has made what claim and is verified by the msg ante-handler for signatures
func (msg *MsgCreateMetadataRequest) ClaimHash() ([]byte, error) {
	metadataRequestClaimHash := NewMetadataRequestClaimHash(
		msg.ChainId,
		msg.EventNonce,
		msg.BlockHeight,
		msg.DaapAddress,
		msg.FeePayer,
	)

	out, err := proto.Marshal(metadataRequestClaimHash)
	return tmhash.Sum([]byte(out)), err
}

func (msg MsgCreateMetadataRequest) GetClaimer() sdk.AccAddress {
	err := msg.ValidateBasic()
	if err != nil {
		panic("MsgInboundRequest failed ValidateBasic! Should have been handled earlier")
	}

	val, _ := sdk.AccAddressFromBech32(msg.Orchestrator)
	return val
}
