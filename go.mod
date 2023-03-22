module github.com/router-protocol/sdk-go

go 1.16

require (
	cosmossdk.io/math v1.0.0-beta.4
	github.com/CosmWasm/wasmd v0.27.0
	github.com/cosmos/cosmos-sdk v0.46.8
	github.com/cosmos/gogoproto v1.4.4 // indirect
	github.com/cosmos/ibc-go/v6 v6.1.0
	github.com/ethereum/go-ethereum v1.10.26
	github.com/evmos/ethermint v0.21.0
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/router-protocol/router-gateway-contracts v1.2.2
	github.com/tendermint/tendermint v0.34.24
	github.com/xlab/suplog v1.3.1
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f
	google.golang.org/grpc v1.53.0
	gopkg.in/yaml.v2 v2.4.0
)

replace (
	github.com/CosmWasm/wasmd => github.com/notional-labs/wasmd v0.30.1-0.20230130020453-795c06e7bc42
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
)
