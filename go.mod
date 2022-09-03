module github.com/theneverse/neverse-core

go 1.13

require (
	github.com/binance-chain/tss-lib v1.3.3-0.20210411025750-fffb56b30511
	github.com/btcsuite/btcd v0.21.0-beta
	github.com/deckarep/golang-set v0.0.0-20180603214616-504e848d77ea
	github.com/ethereum/go-ethereum v1.10.8
	github.com/gogo/protobuf v1.3.2
	github.com/golang/mock v1.5.0
	github.com/hyperledger/fabric v2.1.1+incompatible
	github.com/hyperledger/fabric-protos-go v0.0.0-20201028172056-a3136dde2354
	github.com/iancoleman/orderedmap v0.2.0
	github.com/libp2p/go-libp2p-core v0.5.6
	github.com/looplab/fsm v0.2.0
	github.com/sirupsen/logrus v1.8.1
	github.com/stretchr/testify v1.7.2
	github.com/theneverse/go-lightp2p v0.0.0-20220903065817-d780be311b10
	github.com/theneverse/neverse-kit v0.0.0-20220903070058-cf115e048a59
	github.com/theneverse/neverse-model v0.0.0-20220903075443-51d5a15e0456
	go.uber.org/atomic v1.7.0
	gopkg.in/check.v1 v1.0.0-20200902074654-038fdea0a05b
)

replace github.com/binance-chain/tss-lib => github.com/dawn-to-dusk/tss-lib v1.3.2-0.20220422023240-5ddc16a330ed

replace github.com/agl/ed25519 => github.com/binance-chain/edwards25519 v0.0.0-20200305024217-f36fc4b53d43
