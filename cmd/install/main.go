package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	httpapi "github.com/ipfs/go-ipfs-http-client"

	"github.com/valist-io/valist-go/client"
	"github.com/valist-io/valist-go/contract"
	"github.com/valist-io/valist-go/storage"
)

const help = `
Valist can be used to download and install software
in a completely decentralized way. Versions are fetched
from Ethereum and downloaded from IPFS.

Usage:
	valist-install <path>

Examples:
	valist-install ipfs/go-ipfs/v0.11.0/darwin-arm64 > /usr/local/bin/ipfs
`

// Feel free to swap these URLs for your provider of choice.
// Valist doesn't rely on any centralized infrastructure.
const (
	rpcURL  = "https://rpc.valist.io"
	ipfsURL = "https://gateway.valist.io"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, help)
		os.Exit(1)
	}
	path := strings.SplitN(os.Args[1], "/", 4)
	if len(path) != 4 {
		fmt.Fprintln(os.Stderr, help)
		os.Exit(1)
	}

	rpc, err := ethclient.Dial(rpcURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
	evm, err := contract.NewEVM(common.HexToAddress("0x0"), rpc)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
	api, err := httpapi.NewURLApiWithClient(ipfsURL, &http.Client{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ipfs := storage.NewIPFS(api)
	valist := client.NewClient(evm, ipfs)

	release, err := valist.GetRelease(ctx, path[0], path[1], path[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
	artifact, ok := release.Artifacts[path[3]]
	if !ok {
		fmt.Printf("artifact not found: %s", path[3])
		os.Exit(1)
	}
	data, err := valist.Storage.Read(ctx, artifact.Provider)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
	_, err = os.Stdout.Write(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}
