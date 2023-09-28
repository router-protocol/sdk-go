package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	proto "github.com/gogo/protobuf/proto"
	"github.com/tendermint/tendermint/crypto/tmhash"
)

const TypeMsgValsetUpdatedClaim = "valset_updated_claim"

var _ sdk.Msg = &MsgValsetUpdatedClaim{}

func NewMsgValsetUpdatedClaim(orchestrator string, chainId string, contract string, eventNonce uint64, valsetNonce uint64, blockHeight uint64, srcTxHash string, members []BridgeValidator) *MsgValsetUpdatedClaim {
	return &MsgValsetUpdatedClaim{
		Orchestrator: orchestrator,
		ChainId:      chainId,
		Contract:     contract,
		EventNonce:   eventNonce,
		ValsetNonce:  valsetNonce,
		BlockHeight:  blockHeight,
		SourceTxHash: srcTxHash,
		Members:      members,
	}
}

func (msg *MsgValsetUpdatedClaim) Route() string {
	return RouterKey
}

func (msg *MsgValsetUpdatedClaim) Type() string {
	return TypeMsgValsetUpdatedClaim
}

func (msg *MsgValsetUpdatedClaim) GetSigners() []sdk.AccAddress {
	orchestrator, err := sdk.AccAddressFromBech32(msg.Orchestrator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{orchestrator}
}

func (msg *MsgValsetUpdatedClaim) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgValsetUpdatedClaim) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Orchestrator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid orchestrator address (%s)", err)
	}

	if msg.ChainId == "" || len(msg.ChainId) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidChainID, "invalid chain id (%s)", msg.ChainId)
	}
	return nil
}

/////////////////////////////
//     Implement Claim     //
/////////////////////////////

// GetType returns the type of the claim
func (e *MsgValsetUpdatedClaim) GetType() ClaimType {
	return CLAIM_TYPE_VALSET_UPDATED
}

// Hash implements BridgeDeposit.Hash
// modify this with care as it is security sensitive. If an element of the claim is not in this hash a single hostile validator
// could engineer a hash collision and execute a version of the claim with any unhashed data changed to benefit them.
// note that the Orchestrator is the only field excluded from this hash, this is because that value is used higher up in the store
// structure for who has made what claim and is verified by the msg ante-handler for signatures
func (msg *MsgValsetUpdatedClaim) ClaimHash() ([]byte, error) {
	valsetUpdatedClaimHash := NewValsetUpdatedClaimHash(
		msg.GetChainId(),
		msg.GetContract(),
		msg.GetEventNonce(),
		msg.GetValsetNonce(),
		msg.GetBlockHeight(),
		msg.GetSourceTxHash(),
		msg.GetMembers(),
	)

	out, err := proto.Marshal(valsetUpdatedClaimHash)
	return tmhash.Sum([]byte(out)), err
}

func (msg MsgValsetUpdatedClaim) GetClaimer() sdk.AccAddress {
	err := msg.ValidateBasic()
	if err != nil {
		panic("MsgValsetUpdatedClaim failed ValidateBasic! Should have been handled earlier")
	}

	val, _ := sdk.AccAddressFromBech32(msg.Orchestrator)
	return val
}
