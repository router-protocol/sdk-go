module github.com/router-protocol/sdk-go

go 1.16

require (
	github.com/CosmWasm/wasmd v0.27.0
	github.com/StackExchange/wmi v1.2.1 // indirect
	github.com/allegro/bigcache v1.2.1 // indirect
	github.com/armon/go-metrics v0.4.0 // indirect
	github.com/bitly/go-simplejson v0.5.0 // indirect
	github.com/btcsuite/btcd v0.22.1
	github.com/btcsuite/btcutil v1.0.3-0.20201208143702-a53e38424cce
	github.com/bugsnag/osext v0.0.0-20130617224835-0dd3f918b21b // indirect
	github.com/bugsnag/panicwrap v0.0.0-20151223152923-e2c28503fcd0 // indirect
	github.com/cenkalti/backoff/v4 v4.1.2 // indirect
	github.com/cosmos/cosmos-sdk v0.45.5-0.20220523154235-2921a1c3c918
	github.com/cosmos/ibc-go/v3 v3.1.0
	github.com/dustin/go-humanize v1.0.1-0.20200219035652-afde56e7acac // indirect
	github.com/ethereum/go-ethereum v1.10.19
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/google/btree v1.0.1 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/improbable-eng/grpc-web v0.15.0 // indirect
	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0 // indirect
	github.com/mattn/go-runewidth v0.0.10 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.2-0.20181231171920-c182affec369 // indirect
	github.com/onsi/gomega v1.19.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/regen-network/cosmos-proto v0.3.1
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/rogpeppe/go-internal v1.6.2 // indirect
	github.com/rs/zerolog v1.26.1 // indirect
	github.com/spf13/cobra v1.5.0 // indirect
	github.com/spf13/viper v1.12.0 // indirect
	github.com/stretchr/testify v1.8.0
	github.com/tendermint/tendermint v0.34.20-0.20220517115723-e6f071164839
	github.com/tklauser/go-sysconf v0.3.7 // indirect
	github.com/tyler-smith/go-bip39 v1.1.0
	github.com/xlab/suplog v1.3.1
	golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa // indirect
	golang.org/x/net v0.0.0-20220624214902-1bab6f366d9e // indirect
	golang.org/x/sys v0.0.0-20220610221304-9f5ed59c137d // indirect
	google.golang.org/genproto v0.0.0-20220810155839-1856144b1d9c
	google.golang.org/grpc v1.48.0
	gopkg.in/yaml.v2 v2.4.0
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4
	google.golang.org/grpc => google.golang.org/grpc v1.33.2
)
