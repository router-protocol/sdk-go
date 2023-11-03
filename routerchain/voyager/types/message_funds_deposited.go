package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	proto "github.com/gogo/protobuf/proto"
	attestationTypes "github.com/router-protocol/sdk-go/routerchain/attestation/types"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
	"github.com/tendermint/tendermint/crypto/tmhash"
)

const TypeMsgFundsDeposited = "funds_deposited"

var _ sdk.Msg = &MsgFundsDeposited{}

func NewMsgFundsDeposited(orchestrator string, srcChainId string, srcChainType multichainTypes.ChainType, srcTxHash string, srcTimestamp uint64, contract string, depositId uint64, blockHeight uint64, destChainId []byte, amount sdk.Int, destAmount sdk.Int, srcToken string, recipient []byte, depositor string, partnerID sdk.Int, message []byte, depositWithMessage bool, destToken []byte) *MsgFundsDeposited {
	return &MsgFundsDeposited{
		Orchestrator:       orchestrator,
		SrcChainId:         srcChainId,
		SrcChainType:       srcChainType,
		SrcTxHash:          srcTxHash,
		SrcTimestamp:       srcTimestamp,
		Contract:           contract,
		DepositId:          depositId,
		BlockHeight:        blockHeight,
		DestChainId:        destChainId,
		Amount:             amount,
		DestAmount:         destAmount,
		SrcToken:           srcToken,
		Recipient:          recipient,
		Depositor:          depositor,
		PartnerId:          partnerID,
		Message:            message,
		DestToken:          destToken,
		DepositWithMessage: depositWithMessage,
	}
}

func (msg *MsgFundsDeposited) Route() string {
	return RouterKey
}

func (msg *MsgFundsDeposited) Type() string {
	return TypeMsgFundsDeposited
}

func (msg *MsgFundsDeposited) GetSigners() []sdk.AccAddress {
	orchestrator, err := sdk.AccAddressFromBech32(msg.Orchestrator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{orchestrator}
}

func (msg *MsgFundsDeposited) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgFundsDeposited) ValidateBasic() error {
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
func (msg *MsgFundsDeposited) GetType() attestationTypes.ClaimType {
	return attestationTypes.CLAIM_TYPE_FUND_DEPOSITED_REQUEST
}

func (msg *MsgFundsDeposited) GetChainId() string {
	return msg.SrcChainId
}

func (msg *MsgFundsDeposited) GetEventNonce() uint64 {
	return msg.DepositId
}

// Hash implements MsgFundsDeposited.Hash
// modify this with care as it is security sensitive. If an element of the claim is not in this hash a single hostile validator
// could engineer a hash collision and execute a version of the claim with any unhashed data changed to benefit them.
// note that the Orchestrator is the only field excluded from this hash, this is because that value is used higher up in the store
// structure for who has made what claim and is verified by the msg ante-handler for signatures
func (msg *MsgFundsDeposited) ClaimHash() ([]byte, error) {
	fundDepositRequestClaimHash := NewFundDepositRequestClaimHash(
		msg.SrcChainId,
		msg.SrcChainType,
		msg.SrcTxHash,
		msg.SrcTimestamp,
		msg.Contract,
		msg.DepositId,
		msg.BlockHeight,
		msg.DestChainId,
		msg.Amount,
		msg.DestAmount,
		msg.SrcToken,
		msg.Recipient,
		msg.Depositor,
		msg.PartnerId,
		msg.Message,
		msg.DepositWithMessage,
		msg.DestToken,
	)

	out, err := proto.Marshal(fundDepositRequestClaimHash)
	return tmhash.Sum([]byte(out)), err
}

func (msg MsgFundsDeposited) GetClaimer() sdk.AccAddress {
	err := msg.ValidateBasic()
	if err != nil {
		panic("MsgInboundRequest failed ValidateBasic! Should have been handled earlier")
	}

	val, _ := sdk.AccAddressFromBech32(msg.Orchestrator)
	return val
}
