package types

import (
	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateAdhocRequest = "create_adhoc_request"

var _ sdk.Msg = &MsgCreateAdhocRequest{}

func NewMsgCreateAdhocRequest(from string, chainIds []string, txHashes []string, metaData [][]byte, middlewareContractAddress string, fee sdk.Coin, gasLimit uint64) *MsgCreateAdhocRequest {
	return &MsgCreateAdhocRequest{
		From:                      from,
		ChainIds:                  chainIds,
		TxHashes:                  txHashes,
		MetaData:                  metaData,
		MiddlewareContractAddress: middlewareContractAddress,
		Fee:                       fee,
		CwGasLimit:                gasLimit,
	}
}

func (msg *MsgCreateAdhocRequest) Route() string {
	return RouterKey
}

func (msg *MsgCreateAdhocRequest) Type() string {
	return TypeMsgCreateAdhocRequest
}

func (msg *MsgCreateAdhocRequest) GetSigners() []sdk.AccAddress {
	acc, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sdk.AccAddress(acc)}
}

func (msg *MsgCreateAdhocRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateAdhocRequest) ValidateBasic() (err error) {
	if _, err = sdk.AccAddressFromBech32(msg.From); err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "the client address %s is invalid", msg.From)
	}

	numberOfAdhocRequests := len(msg.ChainIds)
	if len(msg.TxHashes) != numberOfAdhocRequests || len(msg.MetaData) != numberOfAdhocRequests {
		return errorsmod.Wrap(ErrInvalidInput, "error in counting the number of adhoc requests")
	}

	if msg.MiddlewareContractAddress != "" {
		if _, err = sdk.AccAddressFromBech32(msg.MiddlewareContractAddress); err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "the middleware contract address %s is invalid", msg.MiddlewareContractAddress)
		}
	}

	return nil
}
