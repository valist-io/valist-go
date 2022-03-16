package contract

import (
	"github.com/ethereum/go-ethereum/common"
)

//go:generate abigen -abi valist/valist.abi -bin valist/valist.bin -out valist/valist.go -pkg valist -type Valist
//go:generate abigen -abi license/license.abi -bin license/license.bin -out license/license.go -pkg license -type License

var ValistAddresses = map[string]common.Address{
	// Deterministic simulated
	"1337": common.HexToAddress("0x333c3310824b7c685133F2BeDb2CA4b8b4DF633d"),
	// Mumbai testnet
	"80001": common.HexToAddress("0x9569bEb0Eba900495cF58028DB094D824d0AE850"),
	// Polygon mainnet
	"137": common.HexToAddress("0xc70A069eC7F887a7497a4bdC7bE666C1e18c8DC3"),
}

var LicenseAddresses = map[string]common.Address{
	// Deterministic simulated
	"1337": common.HexToAddress("0x8bDa78331C916A08481428e4b07C96D3e916D165"),
	// Mumbai testnet
	"80001": common.HexToAddress("0x597bfcE7F9363b6eBc229f2023F9EcD716C88120"),
	// Polygon mainnet
	"137": common.HexToAddress("0xb85ed41d49Eba25aE6186921Ea63b6055903e810"),
}
