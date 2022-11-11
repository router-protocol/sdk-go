package gateway

import (
	"log"

	ethcmn "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	gatewayWrapper "github.com/router-protocol/router-gateway-contracts/evm/build/bindings/go/GatewayUpgradeable"
)

type GatewayContractClient struct {
	ethClient      *ethclient.Client
	GatewayWrapper *gatewayWrapper.GatewayUpgradeable

	gatewayContractAddress ethcmn.Address
	ethRpc                 string
}

func NewGatewayContractClient(ethRpc string, gatewayContractAddress ethcmn.Address) *GatewayContractClient {
	// address of etherum env
	client, err := ethclient.Dial(ethRpc)
	if err != nil {
		panic(err)
	}

	gatewayWrapperInstance, err := gatewayWrapper.NewGatewayUpgradeable(gatewayContractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	return &GatewayContractClient{
		ethClient:              client,
		GatewayWrapper:         gatewayWrapperInstance,
		ethRpc:                 ethRpc,
		gatewayContractAddress: gatewayContractAddress,
	}

}
