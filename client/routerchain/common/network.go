package common

import (
	"context"
	"fmt"
	"net"
	"path"
	"runtime"
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

func getFileAbsPath(relativePath string) string {
	_, filename, _, _ := runtime.Caller(0)
	return path.Join(path.Dir(filename), relativePath)
}

func LoadNetwork(name string, node string) Network {

	// Set default fields
	network := Network{
		Fee_denom: "route",
		Name:      name,
	}

	// Set endpoints
	if name == "local" {
		network.ApiEndpoint = "https://localhost:1317"
		network.TmEndpoint = "http://localhost:26657"
		network.ChainEvmRpcEndpoint = "http://localhost:8545"
		network.ChainGrpcEndpoint = "tcp://localhost:9090"
	} else if name == "devnet-alpha" {
		network.ApiEndpoint = "https://devnet-alpha.lcd.routerprotocol.com:443"
		network.TmEndpoint = "https://devnet-alpha.tm.routerprotocol.com:443"
		network.ChainEvmRpcEndpoint = "https://devnet-alpha.evm.rpc.routerprotocol.com/"
		network.ChainGrpcEndpoint = "tcp://devnet-alpha.grpc.routerprotocol.com:9090"
	} else if name == "devnet" {
		network.ApiEndpoint = "https://devnet.lcd.routerprotocol.com:443"
		network.TmEndpoint = "https://devnet.tm.routerprotocol.com:443"
		network.ChainEvmRpcEndpoint = "https://devnet.evm.rpc.routerprotocol.com/"
		network.ChainGrpcEndpoint = "tcp://devnet.grpc.routerprotocol.com:9090"
	} else if name == "testnet" {
		network.ApiEndpoint = "https://lcd.testnet.routerchain.dev:443"
		network.TmEndpoint = "https://tm.rpc.testnet.routerchain.dev:443"
		network.ChainEvmRpcEndpoint = "https://evm.rpc.testnet.routerchain.dev/"
		network.ChainGrpcEndpoint = "tcp://grpc.testnet.routerchain.dev:9090"
	} else if name == "load-test" {
		network.ApiEndpoint = "http://13.235.246.63:1317"
		network.TmEndpoint = "http://13.235.246.63:26657"
		network.ChainEvmRpcEndpoint = "http://13.235.246.63:8545"
		network.ChainGrpcEndpoint = "tcp://13.235.246.63:9090"
	}

	//Fetch chain ID
	chainId, err := FetchChainID(network.TmEndpoint)
	if err != nil {
		fmt.Println("Error while fetching chain ID", "rpc", network.TmEndpoint)
		panic(err)
	}

	network.ChainId = chainId
	return network
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
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
