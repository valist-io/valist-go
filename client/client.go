package client

import (
	"context"
	"encoding/json"
	"math/big"

	valist "github.com/valist-io/valist-go"
	"github.com/valist-io/valist-go/util"
)

type Client struct {
	Contract valist.ContractAPI
	Storage  valist.StorageAPI
}

// NewClient returns a client using the given contract and storage providers.
func NewClient(contract valist.ContractAPI, storage valist.StorageAPI) *Client {
	return &Client{contract, storage}
}

// GetTeamMeta returns the metadata of the team with the given name.
func (c *Client) GetTeamMeta(ctx context.Context, teamName string) (*valist.TeamMeta, error) {
	metaURI, err := c.Contract.GetTeamMetaURI(ctx, teamName)
	if err != nil {
		return nil, err
	}
	data, err := util.Fetch(metaURI)
	if err != nil {
		return nil, err
	}
	var team valist.TeamMeta
	if err := json.Unmarshal(data, &team); err != nil {
		return nil, err
	}
	return &team, nil
}

// GetProjectMeta returns the metadata of the project with the given name.
func (c *Client) GetProjectMeta(ctx context.Context, teamName, projectName string) (*valist.ProjectMeta, error) {
	metaURI, err := c.Contract.GetProjectMetaURI(ctx, teamName, projectName)
	if err != nil {
		return nil, err
	}
	data, err := util.Fetch(metaURI)
	if err != nil {
		return nil, err
	}
	var project valist.ProjectMeta
	if err := json.Unmarshal(data, &project); err != nil {
		return nil, err
	}
	return &project, nil
}

// GetReleaseMeta returns the metadata of the release with the given name.
func (c *Client) GetReleaseMeta(ctx context.Context, teamName, projectName, releaseName string) (*valist.ReleaseMeta, error) {
	metaURI, err := c.Contract.GetReleaseMetaURI(ctx, teamName, projectName, releaseName)
	if err != nil {
		return nil, err
	}
	data, err := util.Fetch(metaURI)
	if err != nil {
		return nil, err
	}
	var release valist.ReleaseMeta
	if err := json.Unmarshal(data, &release); err != nil {
		return nil, err
	}
	return &release, nil
}

// GetLatestReleaseMeta returns the latest release.
func (c *Client) GetLatestReleaseMeta(ctx context.Context, teamName, projectName string) (*valist.ReleaseMeta, error) {
	releaseName, err := c.Contract.GetLatestReleaseName(ctx, teamName, projectName)
	if err != nil {
		return nil, err
	}
	return c.GetReleaseMeta(ctx, teamName, projectName, releaseName)
}

// GetLicenseMeta returns the metadata of the license with the given name.
func (c *Client) GetLicenseMeta(ctx context.Context, teamName, projectName, licenseName string) (*valist.LicenseMeta, error) {
	metaURI, err := c.Contract.GetLicenseMetaURI(ctx, teamName, projectName, licenseName)
	if err != nil {
		return nil, err
	}
	data, err := util.Fetch(metaURI)
	if err != nil {
		return nil, err
	}
	var release valist.LicenseMeta
	if err := json.Unmarshal(data, &release); err != nil {
		return nil, err
	}
	return &release, nil
}

// CreateTeam creates a team with the given name, metadata, and members.
func (c *Client) CreateTeam(ctx context.Context, teamName string, team *valist.TeamMeta, beneficiary string, members []string) (valist.TransactionAPI, error) {
	data, err := json.Marshal(team)
	if err != nil {
		return nil, err
	}
	metaURI, err := c.Storage.WriteJSON(ctx, data)
	if err != nil {
		return nil, err
	}
	return c.Contract.CreateTeam(ctx, teamName, metaURI, beneficiary, members)
}

// CreateProject creates a project under the team with the given name, metadata, and members.
func (c *Client) CreateProject(ctx context.Context, teamName, projectName string, project *valist.ProjectMeta, members []string) (valist.TransactionAPI, error) {
	data, err := json.Marshal(project)
	if err != nil {
		return nil, err
	}
	metaURI, err := c.Storage.WriteJSON(ctx, data)
	if err != nil {
		return nil, err
	}
	return c.Contract.CreateProject(ctx, teamName, projectName, metaURI, members)
}

// CreateRelease creates a release under the project with the given name and metadata.
func (c *Client) CreateRelease(ctx context.Context, teamName, projectName, releaseName string, release *valist.ReleaseMeta) (valist.TransactionAPI, error) {
	data, err := json.Marshal(release)
	if err != nil {
		return nil, err
	}
	metaURI, err := c.Storage.WriteJSON(ctx, data)
	if err != nil {
		return nil, err
	}
	return c.Contract.CreateRelease(ctx, teamName, projectName, releaseName, metaURI)
}

// CreateLicense creates a license under the project with the given name and metadata.
func (c *Client) CreateLicense(ctx context.Context, teamName, projectName, licenseName string, license *valist.LicenseMeta, mintPrice *big.Int) (valist.TransactionAPI, error) {
	data, err := json.Marshal(license)
	if err != nil {
		return nil, err
	}
	metaURI, err := c.Storage.WriteJSON(ctx, data)
	if err != nil {
		return nil, err
	}
	return c.Contract.CreateLicense(ctx, teamName, projectName, licenseName, metaURI, mintPrice)
}
