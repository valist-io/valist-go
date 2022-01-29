package valist

import (
	"context"
	"math/big"
)

type Team struct {
	// Name is the organization friendly name.
	Name string `json:"name"`
	// Description is a short description of the organization.
	Description string `json:"description"`
	// Homepage is a link to the organization website.
	Homepage string `json:"homepage"`
}

type Project struct {
	// Name is the repository friendly name.
	Name string `json:"name"`
	// Description is a short description of the repository.
	Description string `json:"description"`
	// Homepage is the website for the repository.
	Homepage string `json:"homepage"`
	// Repository is the source code url for the repository.
	Repository string `json:"repository"`
}

type Release struct {
	// Name is the full release path.
	Name string `json:"name"`
	// Version is the release version.
	Version string `json:"version"`
	// Readme contains the readme contents.
	Readme string `json:"readme"`
	// License contains the license type.
	License string `json:"license"`
	// Dependencies contains a list of all dependencies.
	Dependencies []string `json:"dependencies"`
	// Artifacts is a mapping of names to artifacts.
	Artifacts map[string]Artifact `json:"artifacts"`
}

// Artifact is file contained in a release.
type Artifact struct {
	// SHA256 is the sha256 of the file.
	SHA256 string `json:"sha256"`
	// Provider is the path to the artifact file.
	Provider string `json:"provider"`
}

// ContractAPI defines methods for interacting with smart contracts.
type ContractAPI interface {
	// CreateTeam creates a new team with the given members.
	CreateTeam(ctx context.Context, teamName, metaURI string, members []string) error
	// CreateProject creates a new project. Requires the sender to be a member of the team.
	CreateProject(ctx context.Context, teamName, projectName, metaURI string, members []string) error
	// CreateRelease creates a new release. Requires the sender to be a member of the project.
	CreateRelease(ctx context.Context, teamName, projectName, releaseName, metaURI string) error
	// AddTeamMember adds a member to the team. Requires the sender to be a member of the team.
	AddTeamMember(ctx context.Context, teamName string, address string) error
	// RemoveTeamMember removes a member from the team. Requires the sender to be a member of the team.
	RemoveTeamMember(ctx context.Context, teamName string, address string) error
	// AddProjectMemeber adds a member to the project. Requires the sender to be a member of the team.
	AddProjectMember(ctx context.Context, teamName, projectName string, address string) error
	// RemoveProjectMember removes a member from the project. Requires the sender to be a member of the team.
	RemoveProjectMember(ctx context.Context, teamName, projectName string, address string) error
	// SetTeamMetaURI sets the team metadata content ID. Requires the sender to be a member of the team.
	SetTeamMetaURI(ctx context.Context, teamName, metaURI string) error
	// SetProjectMetaURI sets the project metadata content ID. Requires the sender to be a member of the team.
	SetProjectMetaURI(ctx context.Context, teamName, projectName, metaURI string) error
	// ApproveRelease approves the release by adding the sender's address to the approvers list.
	// The sender's address will be removed from the rejectors list if it exists.
	ApproveRelease(ctx context.Context, teamName, projectName, releaseName string) error
	// RejectRelease rejects the release by adding the sender's address to the rejectors list.
	// The sender's address will be removed from the approvers list if it exists.
	RejectRelease(ctx context.Context, teamName, projectName, releaseName string) error
	// GetLatestReleaseName returns the latest release name.
	GetLatestReleaseName(ctx context.Context, teamName, projectName string) (string, error)
	// GetTeamMetaURI returns the team metadata URI.
	GetTeamMetaURI(ctx context.Context, teamName string) (string, error)
	// GetProjectMetaURI returns the project metadata URI.
	GetProjectMetaURI(ctx context.Context, teamName, projectName string) (string, error)
	// GetReleaseMetaURI returns the release metadata URI.
	GetReleaseMetaURI(ctx context.Context, teamName, projectName, releaseName string) (string, error)
	// GetTeamNames returns a paginated list of team names.
	GetTeamNames(ctx context.Context, page *big.Int, size *big.Int) ([]string, error)
	// GetProjectNames returns a paginated list of project names.
	GetProjectNames(ctx context.Context, teamName string, page *big.Int, size *big.Int) ([]string, error)
	// GetReleaseNames returns a paginated list of release names.
	GetReleaseNames(ctx context.Context, teamName, projectName string, page *big.Int, size *big.Int) ([]string, error)
	// GetTeamMembers returns a paginated list of team members.
	GetTeamMembers(ctx context.Context, teamName string, page *big.Int, size *big.Int) ([]string, error)
	// GetProjectMembers returns a paginated list of project members.
	GetProjectMembers(ctx context.Context, teamName, projectName string, page *big.Int, size *big.Int) ([]string, error)
	// GetReleaseApprovers returns a paginated list of release approvers.
	GetReleaseApprovers(ctx context.Context, teamName string, projectName, releaseName string, page *big.Int, size *big.Int) ([]string, error)
	// GetReleaseRejectors returns a paginated list of release rejectors.
	GetReleaseRejectors(ctx context.Context, teamName string, projectName, releaseName string, page *big.Int, size *big.Int) ([]string, error)
}

// StorageAPI defines methods for reading and writing data.
type StorageAPI interface {
	// ReadTeamMeta returns team metadata from storage.
	ReadTeamMeta(ctx context.Context, metaURI string) (*Team, error)
	// ReadProjectMeta returns project metadata from storage.
	ReadProjectMeta(ctx context.Context, metaURI string) (*Project, error)
	// ReadReleaseMeta returns release metadata from storage.
	ReadReleaseMeta(ctx context.Context, metaURI string) (*Release, error)
	// WriteTeamMeta writes team metadata to storage.
	WriteTeamMeta(ctx context.Context, team *Team) (string, error)
	// WriteProjectMeta writes project metadata to storage.
	WriteProjectMeta(ctx context.Context, project *Project) (string, error)
	// WriteReleaseMeta writes release metadata to storage.
	WriteReleaseMeta(ctx context.Context, release *Release) (string, error)
	// Read returns the contents of the given path.
	Read(ctx context.Context, uri string) ([]byte, error)
	// Write writes the contents to storage.
	Write(ctx context.Context, data []byte) (string, error)
}
