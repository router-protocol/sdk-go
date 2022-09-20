package gateway

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gatewayWrapper "github.com/router-protocol/router-gateway-contracts/evm/wrappers"
)

func (gatewayContractClient *GatewayContractClient) QueryGatewaySendToRouterEvents(startBlock uint64, endBlock uint64) []*gatewayWrapper.GatewayRequestToRouterEvent {
	// create auth and transaction package for deploying smart contract
	filterOpts := &bind.FilterOpts{
		Start: startBlock,
		End:   &endBlock,
	}
	iter, _ := gatewayContractClient.GatewayWrapper.FilterRequestToRouterEvent(filterOpts, nil)

	var sendToRouterEvents []*gatewayWrapper.GatewayRequestToRouterEvent

	for iter.Next() {
		sendToRouterEvents = append(sendToRouterEvents, iter.Event)
	}

	return sendToRouterEvents

}
