package storage

import (
	"context"

	files "github.com/ipfs/go-ipfs-files"
	coreiface "github.com/ipfs/interface-go-ipfs-core"
	"github.com/ipfs/interface-go-ipfs-core/options"
)

type IPFS struct {
	ipfs    coreiface.CoreAPI
	gateway string
}

func NewIPFS(ipfs coreiface.CoreAPI, gateway string) *IPFS {
	return &IPFS{ipfs, gateway}
}

// WriteJSON writes JSON data to storage.
func (s *IPFS) WriteJSON(ctx context.Context, data []byte) (string, error) {
	p, err := s.ipfs.Unixfs().Add(ctx, files.NewBytesFile(data), options.Unixfs.Pin(true))
	if err != nil {
		return "", err
	}
	return s.gateway + p.String(), nil
}
