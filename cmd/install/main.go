package main

import (
	"context"
	"os"
	"strings"

	"github.com/valist-io/valist-go/client"
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

	valist, err := client.NewClient(ctx, nil)
	if err != nil {
		panic(err)
	}
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
