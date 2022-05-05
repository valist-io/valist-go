package contracts

import (
	"github.com/ethereum/go-ethereum/common"
)

//go:generate abigen -abi registry/registry.abi -bin registry/registry.bin -out registry/registry.go -pkg registry -type Registry
//go:generate abigen -abi license/license.abi -bin license/license.bin -out license/license.go -pkg license -type License
//go:generate abigen -abi erc20/erc20.abi -out erc20/erc20.go -pkg erc20 -type ERC20

var RegistryAddresses = map[string]common.Address{
	// Mumbai testnet
	"80001": common.HexToAddress("0xD504d012D78B81fA27288628f3fC89B0e2f56e24"),
	// Polygon mainnet
	"137": common.HexToAddress("0xD504d012D78B81fA27288628f3fC89B0e2f56e24"),
}

var LicenseAddresses = map[string]common.Address{
	// Mumbai testnet
	"80001": common.HexToAddress("0x3cE643dc61bb40bB0557316539f4A93016051b81"),
	// Polygon mainnet
	"137": common.HexToAddress("0x3cE643dc61bb40bB0557316539f4A93016051b81"),
}
