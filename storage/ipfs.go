package storage

import (
	"context"
	"encoding/json"
	"io"
	"os"

	files "github.com/ipfs/go-ipfs-files"
	coreiface "github.com/ipfs/interface-go-ipfs-core"
	"github.com/ipfs/interface-go-ipfs-core/options"
	"github.com/ipfs/interface-go-ipfs-core/path"

	valist "github.com/valist-io/valist-go"
)

type IPFS struct {
	ipfs coreiface.CoreAPI
}

func NewIPFS(ipfs coreiface.CoreAPI) *IPFS {
	return &IPFS{ipfs}
}

// ReadTeamMeta returns team metadata from storage.
func (s *IPFS) ReadTeamMeta(ctx context.Context, metaURI string) (*valist.Team, error) {
	data, err := s.Read(ctx, metaURI)
	if err != nil {
		return nil, err
	}
	var team valist.Team
	if err := json.Unmarshal(data, &team); err != nil {
		return nil, err
	}
	return &team, nil
}

// ReadProjectMeta returns project metadata from storage.
func (s *IPFS) ReadProjectMeta(ctx context.Context, metaURI string) (*valist.Project, error) {
	data, err := s.Read(ctx, metaURI)
	if err != nil {
		return nil, err
	}
	var project valist.Project
	if err := json.Unmarshal(data, &project); err != nil {
		return nil, err
	}
	return &project, nil
}

// ReadReleaseMeta returns release metadata from storage.
func (s *IPFS) ReadReleaseMeta(ctx context.Context, metaURI string) (*valist.Release, error) {
	data, err := s.Read(ctx, metaURI)
	if err != nil {
		return nil, err
	}
	var release valist.Release
	if err := json.Unmarshal(data, &release); err != nil {
		return nil, err
	}
	return &release, nil
}

// WriteTeamMeta writes team metadata to storage.
func (s *IPFS) WriteTeamMeta(ctx context.Context, team *valist.Team) (string, error) {
	data, err := json.Marshal(team)
	if err != nil {
		return "", err
	}
	return s.Write(ctx, data)
}

// WriteProjectMeta writes project metadata to storage.
func (s *IPFS) WriteProjectMeta(ctx context.Context, project *valist.Project) (string, error) {
	data, err := json.Marshal(project)
	if err != nil {
		return "", err
	}
	return s.Write(ctx, data)
}

// WriteReleaseMeta writes release metadata to storage.
func (s *IPFS) WriteReleaseMeta(ctx context.Context, release *valist.Release) (string, error) {
	data, err := json.Marshal(release)
	if err != nil {
		return "", err
	}
	return s.Write(ctx, data)
}

// Read returns the contents of the given path.
func (s *IPFS) Read(ctx context.Context, uri string) ([]byte, error) {
	node, err := s.ipfs.Unixfs().Get(ctx, path.New(uri))
	if err != nil {
		return nil, err
	}
	file, ok := node.(files.File)
	if !ok {
		return nil, os.ErrNotExist
	}
	return io.ReadAll(file)
}

// Write writes the contents to storage.
func (s *IPFS) Write(ctx context.Context, data []byte) (string, error) {
	p, err := s.ipfs.Unixfs().Add(ctx, files.NewBytesFile(data), options.Unixfs.Pin(true))
	if err != nil {
		return "", err
	}
	return p.String(), nil
}
