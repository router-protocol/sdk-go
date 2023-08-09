package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	proto "github.com/gogo/protobuf/proto"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
	"github.com/tendermint/tendermint/crypto/tmhash"
)

func (c CrosschainAckRequest) ValidateBasic() error {
	//TODO: Validate id?
	//TODO: Validate cosmos sender?

	return nil
}

// Hash implements IncomingTx.Hash
// modify this with care as it is security sensitive. If an element of the claim is not in this hash a single hostile validator
// could engineer a hash collision and execute a version of the claim with any unhashed data changed to benefit them.
// note that the Orchestrator is the only field excluded from this hash, this is because that value is used higher up in the store
// structure for who has made what claim and is verified by the msg ante-handler for signatures
func (msg *CrosschainAckRequest) ClaimHash() ([]byte, error) {
	crosschainAckRequestClaimHash := NewCrosschainAckRequestClaimHash(
		msg.AckSrcChainId,
		msg.Contract,
		msg.AckRequestIdentifier,
		msg.BlockHeight,
		msg.DestTxHash,
		msg.RelayerRouterAddress,
		msg.AckDestChainId,
		msg.RequestSender,
		msg.RequestIdentifier,
		msg.AckSrcChainType,
		msg.AckDestChainType,
		msg.FeeConsumed,
		msg.ExecData,
		msg.ExecStatus,
	)

	out, err := proto.Marshal(crosschainAckRequestClaimHash)
	return tmhash.Sum([]byte(out)), err
}

func (msg *CrosschainAckRequest) GetChainId() string {
	return msg.AckDestChainId
}

func (msg *CrosschainAckRequest) RelayerType() RelayerType {
	switch msg.AckDestChainType {
	case multichainTypes.CHAIN_TYPE_COSMOS:
		return IBC_RELAYER

	case multichainTypes.CHAIN_TYPE_ROUTER:
		return RELAYER_NONE

	default:
		return ROUTER_RELAYER
	}
}

func (msg *CrosschainAckRequest) ValidationType() ValidationType {
	switch msg.AckSrcChainType {
	case multichainTypes.CHAIN_TYPE_COSMOS:
		return IBC_VALIDATION

	case multichainTypes.CHAIN_TYPE_ROUTER:
		return DEFAULT_VALIDATION

	default:
		return ORCHESTRATOR_VALIDATION
	}
}

func (msg *CrosschainAckRequest) WorkflowType() WorkflowType {
	// If ack request is to_router, then it's Inbound_ack
	if msg.AckDestChainType == multichainTypes.CHAIN_TYPE_ROUTER {
		return INBOUND_ACK
	}

	// If ack request is from_router, then it's Outbound_ack
	if msg.AckSrcChainType == multichainTypes.CHAIN_TYPE_ROUTER {
		return OUTBOUND_ACK
	}

	// If ack request is via_router, then it's Crosstalk_ack
	return CROSSTALK_ACK
}

func (msg *CrosschainAckRequest) GetBytes() []byte {
	cdc := codec.NewProtoCodec(codectypes.NewInterfaceRegistry())
	return sdk.MustSortJSON(cdc.MustMarshalJSON(msg))
}
