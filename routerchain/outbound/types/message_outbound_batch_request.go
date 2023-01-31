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

func NewMsgOutboundBatchRequest(sender string, destinationChainType multichaintypes.ChainType, destinationChainId string, contractCalls []ContractCall, isAtomic bool, relayerFee sdk.Coin, destinationGasLimit uint64, destinationGasPrice uint64) *MsgOutboundBatchRequest {
	return &MsgOutboundBatchRequest{
		Sender:               sender,
		DestinationChainType: destinationChainType,
		DestinationChainId:   destinationChainId,
		ContractCalls:        contractCalls,
		IsAtomic:             isAtomic,
		RelayerFee:           relayerFee,
		DestinationGasLimit:  destinationGasLimit,
		DestinationGasPrice:  destinationGasPrice,
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

	if msg.DestinationGasLimit <= 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFee, "invalid outgoung tx gas limit (%d)", msg.DestinationGasLimit)
	}
	if msg.DestinationGasPrice <= 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid outgoung tx gas price (%d)", msg.DestinationGasPrice)
	}

	// RelayerFee can be zero or nil.
	if msg.RelayerFee.IsNegative() {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFee, "invalid relayer fee (%d)", msg.RelayerFee.Amount)
	}
	if msg.RelayerFee.Denom != chaintypes.RouterCoin {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid relayer fee denom (%d)", msg.RelayerFee.Denom)
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
