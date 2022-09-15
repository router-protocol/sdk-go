package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/router-protocol/sdk-go/routerchain/attestation/types"
	multichaintypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
	chaintypes "github.com/router-protocol/sdk-go/routerchain/types"
)

const TypeMsgOutboundBatchRequest = "outbound_batch_request"

var _ sdk.Msg = &MsgOutboundBatchRequest{}

func NewMsgOutboundBatchRequest(sender string, destinationChainType multichaintypes.ChainType, destinationChainId string, contractCalls []ContractCall, isAtomic bool, relayerFee sdk.Coin, outgoingTxFee sdk.Coin) *MsgOutboundBatchRequest {
	return &MsgOutboundBatchRequest{
		Sender:               sender,
		DestinationChainType: destinationChainType,
		DestinationChainId:   destinationChainId,
		ContractCalls:        contractCalls,
		IsAtomic:             isAtomic,
		RelayerFee:           relayerFee,
		OutgoingTxFee:        outgoingTxFee,
	}
}

func (msg *MsgOutboundBatchRequest) Route() string {
	return RouterKey
}

func (msg *MsgOutboundBatchRequest) Type() string {
	return TypeMsgOutboundBatchRequest
}

func (msg *MsgOutboundBatchRequest) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgOutboundBatchRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgOutboundBatchRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	if msg.OutgoingTxFee.IsNil() || msg.OutgoingTxFee.IsZero() || msg.OutgoingTxFee.IsNegative() {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFee, "invalid outgoung tx fee (%d)", msg.OutgoingTxFee.Amount)
	}
	if msg.OutgoingTxFee.Denom != chaintypes.RouterCoin {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid outgoung tx fee denom (%d)", msg.OutgoingTxFee.Denom)
	}

	// RelayerFee can be zero or nil.
	if msg.RelayerFee.IsNegative() {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFee, "invalid relayer fee (%d)", msg.RelayerFee.Amount)
	}
	if msg.RelayerFee.Denom != chaintypes.RouterCoin {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid relayer fee denom (%d)", msg.OutgoingTxFee.Denom)
	}

	// Validate contract calls data
	if len(msg.ContractCalls) == 0 {
		return sdkerrors.Wrap(types.ErrInvalid, "No contract call in outbound request")
	}

	for _, contractCall := range msg.ContractCalls {
		if contractCall.DestinationContractAddress == nil {
			return sdkerrors.Wrap(types.ErrInvalid, "Destination contract address cannot be nil")
		}

		if contractCall.Payload == nil {
			return sdkerrors.Wrap(types.ErrInvalid, "Destination contract Payload cannot be nil")
		}
	}

	return nil
}
