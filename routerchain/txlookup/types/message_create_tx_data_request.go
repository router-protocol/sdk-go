package types

import (
	"github.com/cometbft/cometbft/crypto/tmhash"
	proto "github.com/cosmos/gogoproto/proto"
	attestationTypes "github.com/router-protocol/sdk-go/routerchain/attestation/types"

	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateTxDataRequest = "create_tx_data_request"

var _ sdk.Msg = &MsgCreateTxDataRequest{}

func NewMsgCreateTxDataRequest(orchestrator string, chainId string, txHash string, txData []byte, blockHeight uint64) *MsgCreateTxDataRequest {
	return &MsgCreateTxDataRequest{
		Orchestrator: orchestrator,
		ChainId:      chainId,
		TxHash:       txHash,
		TxData:       txData,
		BlockHeight:  blockHeight,
	}
}

func (msg *MsgCreateTxDataRequest) Route() string {
	return RouterKey
}

func (msg *MsgCreateTxDataRequest) Type() string {
	return TypeMsgCreateTxDataRequest
}

func (msg *MsgCreateTxDataRequest) GetSigners() []sdk.AccAddress {
	orchestrator, err := sdk.AccAddressFromBech32(msg.Orchestrator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{orchestrator}
}

func (msg *MsgCreateTxDataRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateTxDataRequest) ValidateBasic() error {
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
func (msg *MsgCreateTxDataRequest) GetType() attestationTypes.ClaimType {
	return attestationTypes.CLAIM_TYPE_ADHOC_REQUEST
}

// Hash implements MsgCreateTxDataRequest.Hash
// modify this with care as it is security sensitive. If an element of the claim is not in this hash a single hostile validator
// could engineer a hash collision and execute a version of the claim with any unhashed data changed to benefit them.
// note that the Orchestrator is the only field excluded from this hash, this is because that value is used higher up in the store
// structure for who has made what claim and is verified by the msg ante-handler for signatures
func (msg *MsgCreateTxDataRequest) ClaimHash() ([]byte, error) {
	txDataRequestClaimHash := TxDataRequestClaimHash{
		ChainId:     msg.ChainId,
		TxHash:      msg.TxHash,
		TxData:      msg.TxData,
		BlockHeight: msg.BlockHeight,
	}

	out, err := proto.Marshal(&txDataRequestClaimHash)
	return tmhash.Sum([]byte(out)), err
}

func (msg MsgCreateTxDataRequest) GetClaimer() sdk.AccAddress {
	err := msg.ValidateBasic()
	if err != nil {
		panic("MsgInboundRequest failed ValidateBasic! Should have been handled earlier")
	}

	val, _ := sdk.AccAddressFromBech32(msg.Orchestrator)
	return val
}

func (msg MsgCreateTxDataRequest) GetEventNonce() uint64 {
	return 0
}

func (msg MsgCreateTxDataRequest) GetContract() string {
	return msg.TxHash
}
