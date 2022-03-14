package valist

import (
	"context"
	"math/big"
)

type TeamMeta struct {
	// Image is the URI of the team image
	Image string `json:"image"`
	// Name is the team friendly name.
	Name string `json:"name"`
	// Description is a short description of the team.
	Description string `json:"description"`
	// ExternalURL is a link to the team website.
	ExternalURL string `json:"external_url"`
}

type ProjectMeta struct {
	// Image is the URI of the project image
	Image string `json:"image"`
	// Name is the project friendly name.
	Name string `json:"name"`
	// ShortDescription is a short description of the project.
	ShortDescription string `json:"short_description"`
	// Description is a description of the project.
	Description string `json:"description"`
	// ExternalURL is a link to the project website.
	ExternalURL string `json:"external_url"`
}

type ReleaseMeta struct {
	// Image is the URI of the release image
	Image string `json:"image"`
	// Name is the unique release name.
	Name string `json:"name"`
	// Description is a description of the release.
	Description string `json:"description"`
	// ExternalURL is a link to the release assets.
	ExternalURL string `json:"external_url"`
	// Licenses contains a list of licenses for the release.
	Licenses []string `json:"licenses"`
}

type LicenseMeta struct {
	// Image is the URI of the license image
	Image string `json:"image"`
	// Name is the unique license name.
	Name string `json:"name"`
	// Description is a description of the license.
	Description string `json:"description"`
	// ExternalURL is a link to the license website.
	ExternalURL string `json:"external_url"`
}

// ContractAPI defines methods for interacting with smart contracts.
type ContractAPI interface {
	// CreateTeam creates a new team with the given members.
	CreateTeam(ctx context.Context, teamName, metaURI, beneficiary string, members []string) (TransactionAPI, error)
	// CreateProject creates a new project. Requires the sender to be a member of the team.
	CreateProject(ctx context.Context, teamName, projectName, metaURI string, members []string) (TransactionAPI, error)
	// CreateRelease creates a new release. Requires the sender to be a member of the project.
	CreateRelease(ctx context.Context, teamName, projectName, releaseName, metaURI string) (TransactionAPI, error)
	// CreateLicense Creates a new License and establishes the mint price.
	CreateLicense(ctx context.Context, teamName, projectName, licenseName, metaURI string, mintPrice *big.Int) (TransactionAPI, error)
	// MintLicense mints a new license to a recipient.
	MintLicense(ctx context.Context, teamName, projectName, licenseName, recipient string) (TransactionAPI, error)
	// AddTeamMember adds a member to the team. Requires the sender to be a member of the team.
	AddTeamMember(ctx context.Context, teamName string, address string) (TransactionAPI, error)
	// RemoveTeamMember removes a member from the team. Requires the sender to be a member of the team.
	RemoveTeamMember(ctx context.Context, teamName string, address string) (TransactionAPI, error)
	// AddProjectMemeber adds a member to the project. Requires the sender to be a member of the team.
	AddProjectMember(ctx context.Context, teamName, projectName string, address string) (TransactionAPI, error)
	// RemoveProjectMember removes a member from the project. Requires the sender to be a member of the team.
	RemoveProjectMember(ctx context.Context, teamName, projectName string, address string) (TransactionAPI, error)
	// SetTeamMetaURI sets the team metadata content ID. Requires the sender to be a member of the team.
	SetTeamMetaURI(ctx context.Context, teamName, metaURI string) (TransactionAPI, error)
	// SetProjectMetaURI sets the project metadata content ID. Requires the sender to be a member of the team.
	SetProjectMetaURI(ctx context.Context, teamName, projectName, metaURI string) (TransactionAPI, error)
	// SetTeamBeneficiary sets the team beneficiary to the new address.
	SetTeamBeneficiary(ctx context.Context, teamName, beneficiary string) (TransactionAPI, error)
	// ApproveRelease approves the release by adding the sender's address to the approvers list.
	// The sender's address will be removed from the rejectors list if it exists.
	ApproveRelease(ctx context.Context, teamName, projectName, releaseName string) (TransactionAPI, error)
	// RejectRelease rejects the release by adding the sender's address to the rejectors list.
	// The sender's address will be removed from the approvers list if it exists.
	RejectRelease(ctx context.Context, teamName, projectName, releaseName string) (TransactionAPI, error)
	// GetLatestReleaseName returns the latest release name.
	GetLatestReleaseName(ctx context.Context, teamName, projectName string) (string, error)
	// GetTeamMetaURI returns the team metadata URI.
	GetTeamMetaURI(ctx context.Context, teamName string) (string, error)
	// GetProjectMetaURI returns the project metadata URI.
	GetProjectMetaURI(ctx context.Context, teamName, projectName string) (string, error)
	// GetReleaseMetaURI returns the release metadata URI.
	GetReleaseMetaURI(ctx context.Context, teamName, projectName, releaseName string) (string, error)
	// GetLicenseMetaURI returns the license metadata URI.
	GetLicenseMetaURI(ctx context.Context, teamName, projectName, licenseName string) (string, error)
	// GetTeamNames returns a paginated list of team names.
	GetTeamNames(ctx context.Context, page *big.Int, size *big.Int) ([]string, error)
	// GetProjectNames returns a paginated list of project names.
	GetProjectNames(ctx context.Context, teamName string, page *big.Int, size *big.Int) ([]string, error)
	// GetReleaseNames returns a paginated list of release names.
	GetReleaseNames(ctx context.Context, teamName, projectName string, page *big.Int, size *big.Int) ([]string, error)
	// GetLicenseNames returns a paginated list of license names.
	GetLicenseNames(ctx context.Context, teamName, projectName string, page *big.Int, size *big.Int) ([]string, error)
	// GetTeamMembers returns a paginated list of team members.
	GetTeamMembers(ctx context.Context, teamName string, page *big.Int, size *big.Int) ([]string, error)
	// GetProjectMembers returns a paginated list of project members.
	GetProjectMembers(ctx context.Context, teamName, projectName string, page *big.Int, size *big.Int) ([]string, error)
	// GetReleaseApprovers returns a paginated list of release approvers.
	GetReleaseApprovers(ctx context.Context, teamName string, projectName, releaseName string, page *big.Int, size *big.Int) ([]string, error)
	// GetReleaseRejectors returns a paginated list of release rejectors.
	GetReleaseRejectors(ctx context.Context, teamName string, projectName, releaseName string, page *big.Int, size *big.Int) ([]string, error)
	// GetLicense price returns the mint price of the license.
	GetLicensePrice(ctx context.Context, teamName, projectName, licenseName string) (*big.Int, error)
	// GetTeamBeneficiary returns the team beneficiary address.
	GetTeamBeneficiary(ctx context.Context, teamName string) (string, error)
	// GetTeamID generates teamID from name.
	GetTeamID(ctx context.Context, teamName string) (*big.Int, error)
	// GetProjectID generates projectID from team ID and name.
	GetProjectID(ctx context.Context, teamID *big.Int, projectName string) (*big.Int, error)
	// GetReleaseID generates releaseID from project ID and name.
	GetReleaseID(ctx context.Context, projectID *big.Int, releaseName string) (*big.Int, error)
	// GetLicenseID generates licenseID from project ID and name.
	GetLicenseID(ctx context.Context, projectID *big.Int, licenseName string) (*big.Int, error)
}

// StorageAPI defines methods for reading and writing data.
type StorageAPI interface {
	// WriteJSON writes JSON data to storage.
	WriteJSON(ctx context.Context, data []byte) (string, error)
}

type TransactionAPI interface {
	// Wait waits until the transaction is finalized.
	Wait(ctx context.Context) error
	// Hash returns the transaction hash.
	Hash() string
}
