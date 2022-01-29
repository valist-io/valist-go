package client

import (
	"context"

	valist "github.com/valist-io/valist-go"
)

type Client struct {
	Contract valist.ContractAPI
	Storage  valist.StorageAPI
}

// NewClient returns a client using the given contract and storage providers.
func NewClient(contract valist.ContractAPI, storage valist.StorageAPI) *Client {
	return &Client{contract, storage}
}

// GetTeam returns the team with the given name.
func (c *Client) GetTeam(ctx context.Context, teamName string) (*valist.Team, error) {
	metaURI, err := c.Contract.GetTeamMetaURI(ctx, teamName)
	if err != nil {
		return nil, err
	}
	return c.Storage.ReadTeamMeta(ctx, metaURI)
}

// GetProject returns the project with the given name.
func (c *Client) GetProject(ctx context.Context, teamName, projectName string) (*valist.Project, error) {
	metaURI, err := c.Contract.GetProjectMetaURI(ctx, teamName, projectName)
	if err != nil {
		return nil, err
	}
	return c.Storage.ReadProjectMeta(ctx, metaURI)
}

// GetRelease returns the release with the given name.
func (c *Client) GetRelease(ctx context.Context, teamName, projectName, releaseName string) (*valist.Release, error) {
	metaURI, err := c.Contract.GetReleaseMetaURI(ctx, teamName, projectName, releaseName)
	if err != nil {
		return nil, err
	}
	return c.Storage.ReadReleaseMeta(ctx, metaURI)
}

// GetLatestRelease returns the latest release.
func (c *Client) GetLatestRelease(ctx context.Context, teamName, projectName string) (*valist.Release, error) {
	releaseName, err := c.Contract.GetLatestReleaseName(ctx, teamName, projectName)
	if err != nil {
		return nil, err
	}
	return c.GetRelease(ctx, teamName, projectName, releaseName)
}

// CreateTeam creates a team with the given name, metadata, and members.
func (c *Client) CreateTeam(ctx context.Context, teamName string, team *valist.Team, members []string) error {
	metaURI, err := c.Storage.WriteTeamMeta(ctx, team)
	if err != nil {
		return err
	}
	return c.Contract.CreateTeam(ctx, teamName, metaURI, members)
}

// CreateProject creates a project under the team with the given name, metadata, and members.
func (c *Client) CreateProject(ctx context.Context, teamName, projectName string, project *valist.Project, members []string) error {
	metaURI, err := c.Storage.WriteProjectMeta(ctx, project)
	if err != nil {
		return err
	}
	return c.Contract.CreateProject(ctx, teamName, projectName, metaURI, members)
}

// CreateRelease creates a release under the project with the given name and metadata.
func (c *Client) CreateRelease(ctx context.Context, teamName, projectName, releaseName string, release *valist.Release) error {
	metaURI, err := c.Storage.WriteReleaseMeta(ctx, release)
	if err != nil {
		return err
	}
	return c.Contract.CreateRelease(ctx, teamName, projectName, releaseName, metaURI)
}
