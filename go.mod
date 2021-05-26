module github.com/allinbits/cosmos-ng-scaffold

go 1.16

require (
	github.com/fdymylja/tmos v0.0.0-00010101000000-000000000000
	github.com/tendermint/tendermint v0.34.10
	google.golang.org/protobuf v1.25.0
)

replace github.com/fdymylja/tmos => github.com/allinbits/cosmos-sdk-poc v0.0.0-20210525072322-2edc986c37f4

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
