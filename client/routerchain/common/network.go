package common

import (
	"context"
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
	if name == "local" {
		return Network{
			ApiEndpoint:         "https://localhost:1317",
			TmEndpoint:          "http://localhost:26657",
			ChainEvmRpcEndpoint: "http://localhost:8545",
			ChainGrpcEndpoint:   "tcp://localhost:9090",
			ChainId:             "router_9000-1",
			Fee_denom:           "route",
			Name:                "local",
		}
	} else if name == "local-docker" {
		return Network{
			ApiEndpoint:         "http://router-chain-container:1317",
			TmEndpoint:          "http://router-chain-container:26657",
			ChainEvmRpcEndpoint: "http://router-chain-container:8545",
			ChainGrpcEndpoint:   "tcp://router-chain-container:9090",
			ChainId:             "router_9000-1",
			Fee_denom:           "route",
			Name:                "local-docker",
		}
	} else if name == "devnet-alpha" {
		return Network{
			ApiEndpoint:         "https://devnet-alpha.lcd.routerprotocol.com:443",
			TmEndpoint:          "https://devnet-alpha.tm.routerprotocol.com:443",
			ChainEvmRpcEndpoint: "https://devnet-alpha.evm.rpc.routerprotocol.com/",
			ChainGrpcEndpoint:   "tcp://devnet-alpha.grpc.routerprotocol.com:9090",
			ChainId:             "router_9000-1",
			Fee_denom:           "route",
			Name:                "devnet-alpha",
		}
	} else if name == "devnet-internal" {
		return Network{
			ApiEndpoint:       "https://devnet-internal.lcd.routerprotocol.com:443",
			TmEndpoint:        "https://devnet-internal.tm.routerprotocol.com:443",
			ChainGrpcEndpoint: "tcp://devnet-internal.grpc.routerprotocol.com:9090",
			ChainId:           "router_9000-1",
			Fee_denom:         "route",
			Name:              "devnet-internal",
		}
	} else if name == "devnet-alpha-nondocker" {
		return Network{
			ApiEndpoint:         "http://3.110.19.140:1317",
			TmEndpoint:          "http://3.110.19.140:26657",
			ChainEvmRpcEndpoint: "http://3.110.19.140:8545",
			ChainGrpcEndpoint:   "tcp://3.110.19.140:9090",
			ChainId:             "router_9000-1",
			Fee_denom:           "route",
			Name:                "devnet-alpha-nondocker",
		}
	} else if name == "devnet" {
		return Network{
			ApiEndpoint:         "https://devnet.lcd.routerprotocol.com:443",
			TmEndpoint:          "https://devnet.tm.routerprotocol.com:443",
			ChainEvmRpcEndpoint: "https://devnet.evm.rpc.routerprotocol.com/",
			ChainGrpcEndpoint:   "tcp://devnet.grpc.routerprotocol.com:9090",
			ChainId:             "router_9000-1",
			Fee_denom:           "route",
			Name:                "devnet",
		}
	} else if name == "testnet" {
		return Network{
			ApiEndpoint:         "https://lcd.testnet.routerchain.dev:443",
			TmEndpoint:          "https://tm.rpc.testnet.routerchain.dev:443",
			ChainEvmRpcEndpoint: "https://evm.rpc.testnet.routerchain.dev/",
			ChainGrpcEndpoint:   "tcp://grpc.testnet.routerchain.dev:9090",
			ChainId:             "router_9000-1",
			Fee_denom:           "route",
			Name:                "testnet",
		}
	} else if name == "load-test" {
		return Network{

			ApiEndpoint:         "http://13.233.149.217:1317",
			TmEndpoint:          "http://13.233.149.217:26657",
			ChainEvmRpcEndpoint: "http://13.233.149.217:8545",
			ChainGrpcEndpoint:   "tcp://13.233.149.217:9090",
			ChainId:             "router_9000-1",
			Fee_denom:           "route",
			Name:                "local-docker",
		}
	}
	return Network{}
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
