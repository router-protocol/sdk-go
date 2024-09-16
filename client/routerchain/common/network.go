package common

import (
	"context"
	"fmt"
	"net"
	"strings"

	"google.golang.org/grpc/credentials"
)

type Network struct {
	ApiEndpoint         string
	TmEndpoint          string
	ChainGrpcEndpoint   string
	ChainEvmRpcEndpoint string
	ChainTlsCert        credentials.TransportCredentials
	ExchangeTlsCert     credentials.TransportCredentials
	ChainId             string
	Fee_denom           string
	Name                string
}

func LoadNetwork(name string, node string) (Network, error) {

	// `node` is not used, remove in future builds
	// Set default fields
	network := Network{
		Fee_denom: "route",
		Name:      name,
	}

	const TESTNET_V2 = "testnet-v2"
	const DEVNET_ALPHA = "devnet-alpha"
	const MAINNET = "mainnet"
	const LOCAL = "local"

	if name == LOCAL {
		network.ChainId = "router_9605-1"
		network.ApiEndpoint = "https://localhost:1317"
		network.TmEndpoint = "http://localhost:26657"
		network.ChainEvmRpcEndpoint = "http://localhost:8545"
		network.ChainGrpcEndpoint = "tcp://localhost:9090"
	} else if name == DEVNET_ALPHA {
		network.ChainId = "router_9605-1"
		network.ApiEndpoint = "https://devnet-alpha.lcd.routerprotocol.com:443"
		network.TmEndpoint = "https://devnet-alpha.tm.routerprotocol.com:443"
		network.ChainEvmRpcEndpoint = "https://devnet-alpha.evm.rpc.routerprotocol.com/"
		network.ChainGrpcEndpoint = "tcp://devnet-alpha.grpc.routerprotocol.com:9090"
	} else if name == MAINNET {
		network.ChainId = "router_9600-1"
		network.ApiEndpoint = "https://sentry.lcd.routerprotocol.com:443"
		network.TmEndpoint = "https://sentry.tm.rpc.routerprotocol.com:443"
		network.ChainEvmRpcEndpoint = "https://sentry.evm.rpc.routerprotocol.com/"
		network.ChainGrpcEndpoint = "tcp://sentry.grpc.routerprotocol.com:9090"
	} else if name == TESTNET_V2 {
		network.ChainId = "router_9607-1"
		network.ApiEndpoint = "https://lcd.sentry.routerchain.dev:443"
		network.TmEndpoint = "https://tmrpc.sentry.routerchain.dev:443"
		network.ChainEvmRpcEndpoint = "https://evmrpc.sentry.routerchain.dev/"
		network.ChainGrpcEndpoint = "tcp://grpc.sentry.routerchain.dev:9090"
	} else {
		return Network{}, fmt.Errorf("network %s not supported", name)
	}

	//Fetch chain ID
	// chainId, err := FetchChainID(network.TmEndpoint)
	// if err != nil {
	// 	fmt.Println("Error while fetching chain ID from default TmEndpoint ", "rpc", network.TmEndpoint)
	// 	// panic(err)
	// 	return Network{}, err
	// }
	// network.ChainId = chainId

	return network, nil
}

func DialerFunc(ctx context.Context, addr string) (net.Conn, error) {
	return Connect(addr)
}

// Connect dials the given address and returns a net.Conn. The protoAddr argument should be prefixed with the protocol,
// eg. "tcp://127.0.0.1:8080" or "unix:///tmp/test.sock"
func Connect(protoAddr string) (net.Conn, error) {
	proto, address := ProtocolAndAddress(protoAddr)
	conn, err := net.Dial(proto, address)
	return conn, err
}

// ProtocolAndAddress splits an address into the protocol and address components.
// For instance, "tcp://127.0.0.1:8080" will be split into "tcp" and "127.0.0.1:8080".
// If the address has no protocol prefix, the default is "tcp".
func ProtocolAndAddress(listenAddr string) (string, string) {
	protocol, address := "tcp", listenAddr
	parts := strings.SplitN(address, "://", 2)
	if len(parts) == 2 {
		protocol, address = parts[0], parts[1]
	}
	return protocol, address
}
