package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
	"github.com/tendermint/tendermint/crypto/tmhash"
)

const TypeMsgValsetUpdatedClaim = "valset_updated_claim"

var _ sdk.Msg = &MsgValsetUpdatedClaim{}

func NewMsgValsetUpdatedClaim(orchestrator string, chainType multichainTypes.ChainType, chainId string, eventNonce uint64, valsetNonce uint64, blockHeight uint64, srcTxHash string, members []BridgeValidator) *MsgValsetUpdatedClaim {
	return &MsgValsetUpdatedClaim{
		Orchestrator: orchestrator,
		ChainType:    chainType,
		ChainId:      chainId,
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
func (b *MsgValsetUpdatedClaim) ClaimHash() ([]byte, error) {
	var members BridgeValidators = b.Members
	internalMembers, err := members.ToInternal()
	if err != nil {
		return nil, sdkerrors.Wrap(err, "invalid members")
	}
	internalMembers.Sort()
	path := fmt.Sprintf("%d/%s/%d/%d/%d/%s/%x/", b.ChainType, b.ChainId, b.EventNonce, b.ValsetNonce, b.BlockHeight, b.SourceTxHash, internalMembers.ToExternal())
	return tmhash.Sum([]byte(path)), nil
}

func (msg MsgValsetUpdatedClaim) GetClaimer() sdk.AccAddress {
	err := msg.ValidateBasic()
	if err != nil {
		panic("MsgValsetUpdatedClaim failed ValidateBasic! Should have been handled earlier")
	}

	val, _ := sdk.AccAddressFromBech32(msg.Orchestrator)
	return val
}
