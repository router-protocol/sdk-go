package types

import (
	"encoding/hex"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	attestationTypes "github.com/router-protocol/sdk-go/routerchain/attestation/types"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
)

const (
	TypeMsgOutgoingBatchConfirm = "create_outgoing_batch_confirm"
)

var _ sdk.Msg = &MsgOutgoingBatchConfirm{}

func NewMsgOutgoingBatchConfirm(
	orchestrator string,
	destinationChainType multichainTypes.ChainType,
	destinationChainId string,
	outgoingBatchSender string,
	outgoingBatchNonce uint64,
	ethSigner string,
	signature string,

) *MsgOutgoingBatchConfirm {
	return &MsgOutgoingBatchConfirm{
		Orchestrator:         orchestrator,
		DestinationChainType: destinationChainType,
		DestinationChainId:   destinationChainId,
		OutgoingBatchSender:  outgoingBatchSender,
		OutgoingBatchNonce:   outgoingBatchNonce,
		EthSigner:            ethSigner,
		Signature:            signature,
	}
}

func (msg *MsgOutgoingBatchConfirm) Route() string {
	return RouterKey
}

func (msg *MsgOutgoingBatchConfirm) Type() string {
	return TypeMsgOutgoingBatchConfirm
}

func (msg *MsgOutgoingBatchConfirm) GetSigners() []sdk.AccAddress {
	orchestrator, err := sdk.AccAddressFromBech32(msg.Orchestrator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{orchestrator}
}

func (msg *MsgOutgoingBatchConfirm) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgOutgoingBatchConfirm) ValidateBasic() error {

	if _, err := sdk.AccAddressFromBech32(msg.Orchestrator); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid orchestrator address (%s)", err)
	}

	if err := attestationTypes.ValidateEthAddress(msg.EthSigner); err != nil {
		return sdkerrors.Wrap(err, "eth signer")
	}

	if _, err := hex.DecodeString(msg.Signature); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "Could not decode hex string %s", msg.Signature)
	}
	return nil

}
