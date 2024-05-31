package types

import (
	"github.com/cometbft/cometbft/crypto/tmhash"
	proto "github.com/cosmos/gogoproto/proto"
	attestationTypes "github.com/router-protocol/sdk-go/routerchain/attestation/types"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDepositInfoUpdated = "deposit_info_updated"

var _ sdk.Msg = &MsgDepositInfoUpdated{}

func NewMsgDepositInfoUpdated(orchestrator string, srcChainId string, srcChainType multichainTypes.ChainType, srcTxHash string, srcTimestamp uint64, depositId uint64, contract string, eventNonce uint64, blockHeight uint64, feeAmount sdkmath.Int, initiatewithdrawal bool, srcToken string, depositor string) *MsgDepositInfoUpdated {
	return &MsgDepositInfoUpdated{
		Orchestrator:       orchestrator,
		SrcChainId:         srcChainId,
		SrcChainType:       srcChainType,
		SrcTxHash:          srcTxHash,
		SrcTimestamp:       srcTimestamp,
		DepositId:          depositId,
		Contract:           contract,
		EventNonce:         eventNonce,
		BlockHeight:        blockHeight,
		FeeAmount:          feeAmount,
		Initiatewithdrawal: initiatewithdrawal,
		SrcToken:           srcToken,
		Depositor:          depositor,
	}
}

func (msg *MsgDepositInfoUpdated) Route() string {
	return RouterKey
}

func (msg *MsgDepositInfoUpdated) Type() string {
	return TypeMsgDepositInfoUpdated
}

func (msg *MsgDepositInfoUpdated) GetSigners() []sdk.AccAddress {
	orchestrator, err := sdk.AccAddressFromBech32(msg.Orchestrator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{orchestrator}
}

func (msg *MsgDepositInfoUpdated) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDepositInfoUpdated) ValidateBasic() error {
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
func (msg *MsgDepositInfoUpdated) GetType() attestationTypes.ClaimType {
	return attestationTypes.CLAIM_DEPOSIT_INFO_UPDATED_REQUEST
}

func (msg *MsgDepositInfoUpdated) GetChainId() string {
	return msg.SrcChainId
}

// Hash implements MsgDepositInfoUpdated.Hash
// modify this with care as it is security sensitive. If an element of the claim is not in this hash a single hostile validator
// could engineer a hash collision and execute a version of the claim with any unhashed data changed to benefit them.
// note that the Orchestrator is the only field excluded from this hash, this is because that value is used higher up in the store
// structure for who has made what claim and is verified by the msg ante-handler for signatures
func (msg *MsgDepositInfoUpdated) ClaimHash() ([]byte, error) {
	fundDepositRequestClaimHash := NewDepositInfoUpdatedRequestClaimHash(
		msg.SrcChainId,
		msg.SrcChainType,
		msg.SrcTxHash,
		msg.SrcTimestamp,
		msg.DepositId,
		msg.Contract,
		msg.EventNonce,
		msg.BlockHeight,
		msg.FeeAmount,
		msg.Initiatewithdrawal,
		msg.SrcToken,
		msg.Depositor,
	)

	out, err := proto.Marshal(fundDepositRequestClaimHash)
	return tmhash.Sum([]byte(out)), err
}

func (msg MsgDepositInfoUpdated) GetClaimer() sdk.AccAddress {
	err := msg.ValidateBasic()
	if err != nil {
		panic("MsgInboundRequest failed ValidateBasic! Should have been handled earlier")
	}

	val, _ := sdk.AccAddressFromBech32(msg.Orchestrator)
	return val
}
