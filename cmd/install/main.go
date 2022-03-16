package main

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/ethclient"
	httpapi "github.com/ipfs/go-ipfs-http-client"

	"github.com/valist-io/valist-go/client"
	"github.com/valist-io/valist-go/contract"
	"github.com/valist-io/valist-go/storage"
	"github.com/valist-io/valist-go/util"
)

const help = `
Valist can be used to download and install software
in a completely decentralized way. Versions are fetched
from Ethereum and downloaded from IPFS.

Usage:
	valist-install <path>

Examples:
	valist-install ipfs/go-ipfs/v0.11.0
`

// Feel free to swap these URLs for your provider of choice.
// Valist doesn't rely on any centralized infrastructure.
const (
	rpcURL  = "https://rpc.valist.io"
	ipfsURL = "https://gateway.valist.io"
)

func main() {
	if len(os.Args) != 2 {
		panic(help)
	}
	path := strings.SplitN(os.Args[1], "/", 3)
	if len(path) != 3 {
		panic(help)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rpc, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		panic(err)
	}
	chainID, err := rpc.ChainID(ctx)
	if err != nil {
		panic(err)
	}
	evm, err := contract.NewEVM(rpc, chainID, nil)
	if err != nil {
		panic(err)
	}
	api, err := httpapi.NewURLApiWithClient(ipfsURL, &http.Client{})
	if err != nil {
		panic(err)
	}

	ipfs := storage.NewIPFS(api, ipfsURL)
	valist := client.NewClient(evm, ipfs)

	release, err := valist.GetReleaseMeta(ctx, path[0], path[1], path[2])
	if err != nil {
		panic(err)
	}
	data, err := util.Fetch(release.ExternalURL)
	if err != nil {
		panic(err)
	}
	_, err = os.Stdout.Write(data)
	if err != nil {
		panic(err)
	}
}
