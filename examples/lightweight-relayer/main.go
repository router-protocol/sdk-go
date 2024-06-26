package main

import (
	ethcmn "github.com/ethereum/go-ethereum/common"
)

const (
	// Tx Agrs
	NETWORK_NAME              = "devnet"
	ORCHESTRATOR_KEYRING_FROM = "orchestrator"
	PASSPHRASE                = "12345678"

	RELAYER_PRIVATE_KEY    = "982E38580B1D24D1C7DE91DE112F89F200E09E31395837DFBAF5A62D9FBE44F7"
	RELAYER_ETH_ADDRESS    = "0x33B4A007EcC80Bc99578c18Da07da704c5403236"
	RELAYER_ROUTER_ADDRESS = "router1fcvd3vcer6h5n37cqtajgwy6uk32w0t0wptqg9"
	// ETH_RPC                = "https://rc-testnet1.routerprotocol.com"
	ETH_RPC = "https://polygon-mumbai.g.alchemy.com/v2/7PDGHtfRIuSxeCJwzBlDLqk_6I20AIvz"
	// ETH_RPC = "https://rpc.ankr.com/avalanche_fuji"
)

var (
	GATEWAY_CONTRACT_ADDRESS = ethcmn.HexToAddress("0xF1530419ef623c32c805f3fC61030000C783F98c")
)

func main() {
	// eth client
	// ethClient, err := ethclient.Dial(ETH_RPC)
	// if err != nil {
	// 	panic(err)
	// }

	// // INITIALIZE RELAYER
	// gatewayContractClient := gateway.NewGatewayContractClient(ETH_RPC, GATEWAY_CONTRACT_ADDRESS)
	// routerchainClient := chainclient.InitialiseChainClient(NETWORK_NAME, ORCHESTRATOR_KEYRING_FROM, PASSPHRASE, "")
	// relayer := relayer.NewRelayer(ethClient, gatewayContractClient, routerchainClient, RELAYER_PRIVATE_KEY, RELAYER_ETH_ADDRESS, RELAYER_ROUTER_ADDRESS)

	// // TRIGGER ACTIONS
	// ctx, _ := context.WithTimeout(context.Background(), time.Minute)
	// // relayer.SubmitBatchTxToGateway(ctx, routerchainClient)
	// relayer.SubmitCrosstalkTxToGateway(ctx, routerchainClient)
	// relayer.SubmitCrosstalkAckTxToGateway(ctx, routerchainClient)

}
