module github.com/router-protocol/sdk-go

go 1.16

require (
	cosmossdk.io/math v1.0.0-beta.4
	github.com/CosmWasm/wasmd v0.27.0
	github.com/StackExchange/wmi v1.2.1 // indirect
	github.com/btcsuite/btcd v0.22.2
	github.com/btcsuite/btcutil v1.0.3-0.20201208143702-a53e38424cce
	github.com/cosmos/cosmos-sdk v0.46.8
	github.com/cosmos/ibc-go/v6 v6.1.0
	github.com/dustin/go-humanize v1.0.1-0.20200219035652-afde56e7acac // indirect
	github.com/ethereum/go-ethereum v1.10.26
	github.com/evmos/ethermint v0.21.0
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0 // indirect
	github.com/mattn/go-runewidth v0.0.10 // indirect
	github.com/pkg/errors v0.9.1
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/router-protocol/router-gateway-contracts v1.2.2
	github.com/stretchr/testify v1.8.1
	github.com/tendermint/tendermint v0.34.24
	github.com/tyler-smith/go-bip39 v1.1.0
	github.com/xlab/suplog v1.3.1
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f
	google.golang.org/grpc v1.53.0
	gopkg.in/yaml.v2 v2.4.0
)

replace (
	github.com/CosmWasm/wasmd => github.com/router-protocol/wasmd v0.31.0-rc0-router
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4
)
