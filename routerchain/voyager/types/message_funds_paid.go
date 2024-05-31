package types

import (
	"github.com/cometbft/cometbft/crypto/tmhash"
	proto "github.com/cosmos/gogoproto/proto"
	attestationTypes "github.com/router-protocol/sdk-go/routerchain/attestation/types"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"

	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgFundsPaid = "funds_paid"

var _ sdk.Msg = &MsgFundsPaid{}

func NewMsgFundsPaid(orchestrator string, srcChainId string, srcChainType multichainTypes.ChainType, srcTxHash string, srcTimestamp uint64, contract string, eventNonce uint64, blockHeight uint64, messageHash []byte, forwarder string, forwarderRouterAddr string, execData []byte, execStatus bool) *MsgFundsPaid {
	return &MsgFundsPaid{
		Orchestrator:        orchestrator,
		SrcChainId:          srcChainId,
		SrcChainType:        srcChainType,
		SrcTxHash:           srcTxHash,
		SrcTimestamp:        srcTimestamp,
		Contract:            contract,
		EventNonce:          eventNonce,
		BlockHeight:         blockHeight,
		MessageHash:         messageHash,
		Forwarder:           forwarder,
		ForwarderRouterAddr: forwarderRouterAddr,
		ExecData:            execData,
		ExecStatus:          execStatus,
	}
}

func (msg *MsgFundsPaid) Route() string {
	return RouterKey
}

func (msg *MsgFundsPaid) Type() string {
	return TypeMsgFundsPaid
}

func (msg *MsgFundsPaid) GetSigners() []sdk.AccAddress {
	orchestrator, err := sdk.AccAddressFromBech32(msg.Orchestrator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{orchestrator}
}

func (msg *MsgFundsPaid) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgFundsPaid) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Orchestrator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid orchestrator address (%s)", err)
	}
	return nil
}

/////////////////////////////
//     Implement Claim     //
/////////////////////////////

// GetType returns the type of the claim
func (msg *MsgFundsPaid) GetType() attestationTypes.ClaimType {
	return attestationTypes.CLAIM_TYPE_FUND_PAID_REQUEST
}

func (msg *MsgFundsPaid) GetChainId() string {
	return msg.SrcChainId
}

// Hash implements MsgFundsPaid.Hash
// modify this with care as it is security sensitive. If an element of the claim is not in this hash a single hostile validator
// could engineer a hash collision and execute a version of the claim with any unhashed data changed to benefit them.
// note that the Orchestrator is the only field excluded from this hash, this is because that value is used higher up in the store
// structure for who has made what claim and is verified by the msg ante-handler for signatures
func (msg *MsgFundsPaid) ClaimHash() ([]byte, error) {
	fundPaidRequestClaimHash := NewFundPaidRequestClaimHash(
		msg.SrcChainId,
		msg.SrcChainType,
		msg.SrcTxHash,
		msg.SrcTimestamp,
		msg.Contract,
		msg.EventNonce,
		msg.BlockHeight,
		msg.MessageHash,
		msg.Forwarder,
		msg.ForwarderRouterAddr,
		msg.ExecData,
		msg.ExecStatus,
	)

	out, err := proto.Marshal(fundPaidRequestClaimHash)
	return tmhash.Sum([]byte(out)), err
}

func (msg MsgFundsPaid) GetClaimer() sdk.AccAddress {
	err := msg.ValidateBasic()
	if err != nil {
		panic("MsgInboundRequest failed ValidateBasic! Should have been handled earlier")
	}

	val, _ := sdk.AccAddressFromBech32(msg.Orchestrator)
	return val
}
