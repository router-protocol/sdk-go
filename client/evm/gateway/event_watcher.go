package gateway

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gatewayWrapper "github.com/router-protocol/router-gateway-contracts/evm/build/bindings/go/GatewayUpgradeable"
)

func (gatewayContractClient *GatewayContractClient) QueryGatewaySendToRouterEvents(startBlock uint64, endBlock uint64) []*gatewayWrapper.GatewayUpgradeableRequestToRouterEvent {
	// create auth and transaction package for deploying smart contract
	filterOpts := &bind.FilterOpts{
		Start: startBlock,
		End:   &endBlock,
	}
	iter, _ := gatewayContractClient.GatewayWrapper.FilterRequestToRouterEvent(filterOpts, nil, nil)

	var sendToRouterEvents []*gatewayWrapper.GatewayUpgradeableRequestToRouterEvent

	for iter.Next() {
		sendToRouterEvents = append(sendToRouterEvents, iter.Event)
	}

	return sendToRouterEvents

}
