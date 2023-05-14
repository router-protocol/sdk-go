package types

import (
	proto "github.com/gogo/protobuf/proto"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
	"github.com/tendermint/tendermint/crypto/tmhash"
)

func (c CrosschainRequest) ValidateBasic() error {
	//TODO: Validate id?
	//TODO: Validate cosmos sender?

	return nil
}

// Hash implements IncomingTx.Hash
// modify this with care as it is security sensitive. If an element of the claim is not in this hash a single hostile validator
// could engineer a hash collision and execute a version of the claim with any unhashed data changed to benefit them.
// note that the Orchestrator is the only field excluded from this hash, this is because that value is used higher up in the store
// structure for who has made what claim and is verified by the msg ante-handler for signatures
func (msg *CrosschainRequest) ClaimHash() ([]byte, error) {
	crosschainRequestClaimHash := NewCrosschainRequestClaimHash(
		msg.SrcChainId,
		msg.RequestIdentifier,
		msg.BlockHeight,
		msg.SourceTxHash,
		msg.SrcTimestamp,
		msg.SrcTxOrigin,
		msg.RouteAmount,
		msg.RouteRecipient,
		msg.DestChainId,
		msg.RequestSender,
		msg.RequestMetadata,
		msg.RequestPacket,
		msg.SrcChainType,
		msg.DestChainType)

	out, err := proto.Marshal(crosschainRequestClaimHash)
	return tmhash.Sum([]byte(out)), err
}

func (msg *CrosschainRequest) ExecuteMsgType() ExecuteMsgType {
	// Routerchain <-> Cosmos communication is only using IBC
	if (msg.DestChainType == multichainTypes.CHAIN_TYPE_COSMOS) || (msg.DestChainType == multichainTypes.CHAIN_TYPE_ROUTER && msg.SrcChainType == multichainTypes.CHAIN_TYPE_COSMOS) {
		return EXECUTE_IBC_MSG
	} else { // Routerchain <-> EVM/Near communication is only through Gateway contract call
		return EXECUTE_CONTRACT_CALL
	}
}

func (msg *CrosschainRequest) RelayerType() RelayerType {
	switch msg.DestChainType {
	case multichainTypes.CHAIN_TYPE_COSMOS:
		return IBC_RELAYER

	case multichainTypes.CHAIN_TYPE_ROUTER:
		return RELAYER_NONE

	default:
		return ROUTER_RELAYER
	}
}

func (msg *CrosschainRequest) AckRelayerType() RelayerType {
	switch msg.SrcChainType {
	case multichainTypes.CHAIN_TYPE_COSMOS:
		return IBC_RELAYER

	case multichainTypes.CHAIN_TYPE_ROUTER:
		return RELAYER_NONE

	default:
		return ROUTER_RELAYER
	}
}

func (msg *CrosschainRequest) IsConfirmationRequired() bool {
	// Skip Crosstalk requests as it's already confirmed by Orchestrator when MsgCrosschainRequest is submitted
	if msg.RelayerType() == ROUTER_RELAYER && msg.WorkflowType() != CROSSTALK {
		return true
	}
	return false
}

func (msg *CrosschainRequest) WorkflowType() WorkflowType {
	// If request is to_router, then it's Inbound
	if msg.DestChainType == multichainTypes.CHAIN_TYPE_ROUTER {
		return INBOUND
	}

	// If request is from_router, then it's Outbound
	if msg.SrcChainType == multichainTypes.CHAIN_TYPE_ROUTER {
		return OUTBOUND
	}

	// If request is via_router, then it's Crosstalk
	return CROSSTALK
}

func (msg *CrosschainRequest) ValidationType() ValidationType {
	switch msg.SrcChainType {
	case multichainTypes.CHAIN_TYPE_COSMOS:
		return IBC_VALIDATION

	case multichainTypes.CHAIN_TYPE_ROUTER:
		return DEFAULT_VALIDATION

	default:
		return ORCHESTRATOR_VALIDATION
	}
}

func (msg *CrosschainRequest) AckValidationType() ValidationType {
	switch msg.DestChainType {
	case multichainTypes.CHAIN_TYPE_COSMOS:
		return IBC_VALIDATION

	case multichainTypes.CHAIN_TYPE_ROUTER:
		return DEFAULT_VALIDATION

	default:
		return ORCHESTRATOR_VALIDATION
	}
}
