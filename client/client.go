package client

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	files "github.com/ipfs/go-ipfs-files"
	httpapi "github.com/ipfs/go-ipfs-http-client"
	coreiface "github.com/ipfs/interface-go-ipfs-core"
	"github.com/ipfs/interface-go-ipfs-core/options"

	valist "github.com/valist-io/valist-go"
	"github.com/valist-io/valist-go/contract"
	license_contract "github.com/valist-io/valist-go/contract/license"
	valist_contract "github.com/valist-io/valist-go/contract/valist"
	"github.com/valist-io/valist-go/util"
)

var DefaultOptions = &Options{
	EthereumRpc: "https://rpc.valist.io",
	IPFSUrl:     "https://gateway.valist.io",
}

type Options struct {
	// EthereumRpc is the url to use for ethereum rpc calls and transactions.
	EthereumRpc string
	// IPFSUrl is the url to use for ipfs storage writes.
	IPFSUrl string
	// PrivateKey is the key to use for signing transactions.
	PrivateKey *ecdsa.PrivateKey
}

type Client struct {
	rpc     bind.DeployBackend
	ipfs    coreiface.CoreAPI
	gateway string
	private *ecdsa.PrivateKey
	chainID *big.Int
	valist  *valist_contract.Valist
	license *license_contract.License
}

// NewClient returns a client using the given contract and storage providers.
func NewClient(ctx context.Context, opts *Options) (*Client, error) {
	if opts == nil {
		opts = DefaultOptions
	}
	ipfs, err := httpapi.NewURLApiWithClient(opts.IPFSUrl, &http.Client{})
	if err != nil {
		return nil, err
	}
	rpc, err := ethclient.DialContext(ctx, opts.EthereumRpc)
	if err != nil {
		return nil, err
	}
	chainID, err := rpc.ChainID(ctx)
	if err != nil {
		return nil, err
	}
	valistAddress, ok := contract.ValistAddresses[chainID.String()]
	if !ok {
		return nil, fmt.Errorf("chain with id=%d not supported", chainID)
	}
	licenseAddress, ok := contract.LicenseAddresses[chainID.String()]
	if !ok {
		return nil, fmt.Errorf("chain with id=%d not supported", chainID)
	}
	valistContract, err := valist_contract.NewValist(valistAddress, rpc)
	if err != nil {
		return nil, err
	}
	licenseContract, err := license_contract.NewLicense(licenseAddress, rpc)
	if err != nil {
		return nil, err
	}
	return &Client{
		rpc:     rpc,
		ipfs:    ipfs,
		gateway: opts.IPFSUrl,
		private: opts.PrivateKey,
		chainID: chainID,
		valist:  valistContract,
		license: licenseContract,
	}, nil
}

// GetTeamMeta returns the metadata of the team with the given name.
func (c *Client) GetTeamMeta(ctx context.Context, teamName string) (*valist.TeamMeta, error) {
	metaURI, err := c.GetTeamMetaURI(ctx, teamName)
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
	metaURI, err := c.GetProjectMetaURI(ctx, teamName, projectName)
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
	metaURI, err := c.GetReleaseMetaURI(ctx, teamName, projectName, releaseName)
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
	releaseName, err := c.GetLatestReleaseName(ctx, teamName, projectName)
	if err != nil {
		return nil, err
	}
	return c.GetReleaseMeta(ctx, teamName, projectName, releaseName)
}

// GetLicenseMeta returns the metadata of the license with the given name.
func (c *Client) GetLicenseMeta(ctx context.Context, teamName, projectName, licenseName string) (*valist.LicenseMeta, error) {
	metaURI, err := c.GetLicenseMetaURI(ctx, teamName, projectName, licenseName)
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
func (c *Client) CreateTeam(ctx context.Context, teamName string, team *valist.TeamMeta, beneficiary string, members []string) (*types.Transaction, error) {
	data, err := json.Marshal(team)
	if err != nil {
		return nil, err
	}
	metaURI, err := c.WriteJSON(ctx, data)
	if err != nil {
		return nil, err
	}
	var addresses []common.Address
	for _, hex := range members {
		addresses = append(addresses, common.HexToAddress(hex))
	}
	txopts, err := c.txopts(ctx)
	if err != nil {
		return nil, err
	}
	return c.valist.CreateTeam(txopts, teamName, metaURI, common.HexToAddress(beneficiary), addresses)
}

// CreateProject creates a project under the team with the given name, metadata, and members.
func (c *Client) CreateProject(ctx context.Context, teamName, projectName string, project *valist.ProjectMeta, members []string) (*types.Transaction, error) {
	data, err := json.Marshal(project)
	if err != nil {
		return nil, err
	}
	metaURI, err := c.WriteJSON(ctx, data)
	if err != nil {
		return nil, err
	}
	var addresses []common.Address
	for _, hex := range members {
		addresses = append(addresses, common.HexToAddress(hex))
	}
	txopts, err := c.txopts(ctx)
	if err != nil {
		return nil, err
	}
	return c.valist.CreateProject(txopts, teamName, projectName, metaURI, addresses)
}

// CreateRelease creates a release under the project with the given name and metadata.
func (c *Client) CreateRelease(ctx context.Context, teamName, projectName, releaseName string, release *valist.ReleaseMeta) (*types.Transaction, error) {
	data, err := json.Marshal(release)
	if err != nil {
		return nil, err
	}
	metaURI, err := c.WriteJSON(ctx, data)
	if err != nil {
		return nil, err
	}
	txopts, err := c.txopts(ctx)
	if err != nil {
		return nil, err
	}
	return c.valist.CreateRelease(txopts, teamName, projectName, releaseName, metaURI)
}

// CreateLicense creates a license under the project with the given name and metadata.
func (c *Client) CreateLicense(ctx context.Context, teamName, projectName, licenseName string, license *valist.LicenseMeta, mintPrice *big.Int) (*types.Transaction, error) {
	data, err := json.Marshal(license)
	if err != nil {
		return nil, err
	}
	metaURI, err := c.WriteJSON(ctx, data)
	if err != nil {
		return nil, err
	}
	txopts, err := c.txopts(ctx)
	if err != nil {
		return nil, err
	}
	return c.license.CreateLicense(txopts, teamName, projectName, licenseName, metaURI, mintPrice)
}

// MintLicense mints a new license to a recipient.
func (c *Client) MintLicense(ctx context.Context, teamName, projectName, licenseName, recipient string) (*types.Transaction, error) {
	txopts, err := c.txopts(ctx)
	if err != nil {
		return nil, err
	}
	txopts.Value, err = c.GetLicensePrice(ctx, teamName, projectName, licenseName)
	if err != nil {
		return nil, err
	}
	return c.license.MintLicense(txopts, teamName, projectName, licenseName, common.HexToAddress(recipient))
}

// AddTeamMember adds a member to the team. Requires the sender to be a member of the team.
func (c *Client) AddTeamMember(ctx context.Context, teamName string, address string) (*types.Transaction, error) {
	txopts, err := c.txopts(ctx)
	if err != nil {
		return nil, err
	}
	return c.valist.AddTeamMember(txopts, teamName, common.HexToAddress(address))
}

// RemoveTeamMember removes a member from the team. Requires the sender to be a member of the team.
func (c *Client) RemoveTeamMember(ctx context.Context, teamName string, address string) (*types.Transaction, error) {
	txopts, err := c.txopts(ctx)
	if err != nil {
		return nil, err
	}
	return c.valist.RemoveTeamMember(txopts, teamName, common.HexToAddress(address))
}

// AddProjectMemeber adds a member to the project. Requires the sender to be a member of the team.
func (c *Client) AddProjectMember(ctx context.Context, teamName, projectName string, address string) (*types.Transaction, error) {
	txopts, err := c.txopts(ctx)
	if err != nil {
		return nil, err
	}
	return c.valist.AddProjectMember(txopts, teamName, projectName, common.HexToAddress(address))
}

// RemoveProjectMember removes a member from the project. Requires the sender to be a member of the team.
func (c *Client) RemoveProjectMember(ctx context.Context, teamName, projectName string, address string) (*types.Transaction, error) {
	txopts, err := c.txopts(ctx)
	if err != nil {
		return nil, err
	}
	return c.valist.RemoveProjectMember(txopts, teamName, projectName, common.HexToAddress(address))
}

// SetTeamMetaURI sets the team metadata content ID. Requires the sender to be a member of the team.
func (c *Client) SetTeamMetaURI(ctx context.Context, teamName, metaURI string) (*types.Transaction, error) {
	txopts, err := c.txopts(ctx)
	if err != nil {
		return nil, err
	}
	return c.valist.SetTeamMetaURI(txopts, teamName, metaURI)
}

// SetProjectMetaURI sets the project metadata content ID. Requires the sender to be a member of the team.
func (c *Client) SetProjectMetaURI(ctx context.Context, teamName, projectName, metaURI string) (*types.Transaction, error) {
	txopts, err := c.txopts(ctx)
	if err != nil {
		return nil, err
	}
	return c.valist.SetProjectMetaURI(txopts, teamName, projectName, metaURI)
}

// SetTeamBeneficiary sets the team beneficiary to the new address.
func (c *Client) SetTeamBeneficiary(ctx context.Context, teamName, beneficiary string) (*types.Transaction, error) {
	txopts, err := c.txopts(ctx)
	if err != nil {
		return nil, err
	}
	teamID, err := c.GetTeamID(ctx, teamName)
	if err != nil {
		return nil, err
	}
	return c.valist.SetTeamBeneficiary(txopts, teamID, common.HexToAddress(beneficiary))
}

// ApproveRelease approves the release by adding the sender's address to the approvers list.
// The sender's address will be removed from the rejectors list if it exists.
func (c *Client) ApproveRelease(ctx context.Context, teamName, projectName, releaseName string) (*types.Transaction, error) {
	txopts, err := c.txopts(ctx)
	if err != nil {
		return nil, err
	}
	return c.valist.ApproveRelease(txopts, teamName, projectName, releaseName)
}

// RejectRelease rejects the release by adding the sender's address to the rejectors list.
// The sender's address will be removed from the approvers list if it exists.
func (c *Client) RejectRelease(ctx context.Context, teamName, projectName, releaseName string) (*types.Transaction, error) {
	txopts, err := c.txopts(ctx)
	if err != nil {
		return nil, err
	}
	return c.valist.RejectRelease(txopts, teamName, projectName, releaseName)
}

// GetLatestReleaseName returns the latest release name.
func (c *Client) GetLatestReleaseName(ctx context.Context, teamName, projectName string) (string, error) {
	return c.valist.GetLatestReleaseName(c.callopts(ctx), teamName, projectName)
}

// GetTeamMetaURI returns the team metadata URI.
func (c *Client) GetTeamMetaURI(ctx context.Context, teamName string) (string, error) {
	return c.valist.GetTeamMetaURI(c.callopts(ctx), teamName)
}

// GetProjectMetaURI returns the project metadata URI.
func (c *Client) GetProjectMetaURI(ctx context.Context, teamName, projectName string) (string, error) {
	return c.valist.GetProjectMetaURI(c.callopts(ctx), teamName, projectName)
}

// GetReleaseMetaURI returns the release metadata URI.
func (c *Client) GetReleaseMetaURI(ctx context.Context, teamName, projectName, releaseName string) (string, error) {
	return c.valist.GetReleaseMetaURI(c.callopts(ctx), teamName, projectName, releaseName)
}

// GetLicenseMetaURI returns the license metadata URI.
func (c *Client) GetLicenseMetaURI(ctx context.Context, teamName, projectName, licenseName string) (string, error) {
	return c.license.GetLicenseMetaURI(c.callopts(ctx), teamName, projectName, licenseName)
}

// GetTeamNames returns a paginated list of team names.
func (c *Client) GetTeamNames(ctx context.Context, page *big.Int, size *big.Int) ([]string, error) {
	return c.valist.GetTeamNames(c.callopts(ctx), page, size)
}

// GetProjectNames returns a paginated list of project names.
func (c *Client) GetProjectNames(ctx context.Context, teamName string, page *big.Int, size *big.Int) ([]string, error) {
	return c.valist.GetProjectNames(c.callopts(ctx), teamName, page, size)
}

// GetReleaseNames returns a paginated list of release names.
func (c *Client) GetReleaseNames(ctx context.Context, teamName, projectName string, page *big.Int, size *big.Int) ([]string, error) {
	return c.valist.GetReleaseNames(c.callopts(ctx), teamName, projectName, page, size)
}

// GetLicenseNames returns a paginated list of license names.
func (c *Client) GetLicenseNames(ctx context.Context, teamName, projectName string, page *big.Int, size *big.Int) ([]string, error) {
	return c.license.GetNamesByProjectID(c.callopts(ctx), teamName, projectName, page, size)
}

// GetTeamMembers returns a paginated list of team members.
func (c *Client) GetTeamMembers(ctx context.Context, teamName string, page *big.Int, size *big.Int) ([]string, error) {
	addresses, err := c.valist.GetTeamMembers(c.callopts(ctx), teamName, page, size)
	if err != nil {
		return nil, err
	}
	var members []string
	for _, address := range addresses {
		members = append(members, address.Hex())
	}
	return members, nil
}

// GetProjectMembers returns a paginated list of project members.
func (c *Client) GetProjectMembers(ctx context.Context, teamName, projectName string, page *big.Int, size *big.Int) ([]string, error) {
	addresses, err := c.valist.GetProjectMembers(c.callopts(ctx), teamName, projectName, page, size)
	if err != nil {
		return nil, err
	}
	var members []string
	for _, address := range addresses {
		members = append(members, address.Hex())
	}
	return members, nil
}

// GetReleaseApprovers returns a paginated list of release approvers.
func (c *Client) GetReleaseApprovers(ctx context.Context, teamName string, projectName, releaseName string, page *big.Int, size *big.Int) ([]string, error) {
	addresses, err := c.valist.GetReleaseApprovers(c.callopts(ctx), teamName, projectName, releaseName, page, size)
	if err != nil {
		return nil, err
	}
	var approvers []string
	for _, address := range addresses {
		approvers = append(approvers, address.Hex())
	}
	return approvers, nil
}

// GetReleaseRejectors returns a paginated list of release rejectors.
func (c *Client) GetReleaseRejectors(ctx context.Context, teamName string, projectName, releaseName string, page *big.Int, size *big.Int) ([]string, error) {
	addresses, err := c.valist.GetReleaseRejectors(c.callopts(ctx), teamName, projectName, releaseName, page, size)
	if err != nil {
		return nil, err
	}
	var rejectors []string
	for _, address := range addresses {
		rejectors = append(rejectors, address.Hex())
	}
	return rejectors, nil
}

// GetTeamID generates teamID from name.
func (c *Client) GetTeamID(ctx context.Context, teamName string) (*big.Int, error) {
	return c.valist.GetTeamID(c.callopts(ctx), teamName)
}

// GetProjectID generates projectID from team ID and name.
func (c *Client) GetProjectID(ctx context.Context, teamID *big.Int, projectName string) (*big.Int, error) {
	return c.valist.GetProjectID(c.callopts(ctx), teamID, projectName)
}

// GetReleaseID generates releaseID from project ID and name.
func (c *Client) GetReleaseID(ctx context.Context, projectID *big.Int, releaseName string) (*big.Int, error) {
	return c.valist.GetReleaseID(c.callopts(ctx), projectID, releaseName)
}

// GetLicenseID generates licenseID from project ID and name.
func (c *Client) GetLicenseID(ctx context.Context, projectID *big.Int, licenseName string) (*big.Int, error) {
	return c.license.GetLicenseID(c.callopts(ctx), projectID, licenseName)
}

// GetLicense price returns the mint price of the license.
func (c *Client) GetLicensePrice(ctx context.Context, teamName, projectName, licenseName string) (*big.Int, error) {
	teamID, err := c.GetTeamID(ctx, teamName)
	if err != nil {
		return nil, err
	}
	projectID, err := c.GetProjectID(ctx, teamID, projectName)
	if err != nil {
		return nil, err
	}
	licenseID, err := c.GetLicenseID(ctx, projectID, licenseName)
	if err != nil {
		return nil, err
	}
	return c.license.PriceByID(c.callopts(ctx), licenseID)
}

// GetTeamBeneficiary returns the team beneficiary address.
func (c *Client) GetTeamBeneficiary(ctx context.Context, teamName string) (string, error) {
	teamID, err := c.GetTeamID(ctx, teamName)
	if err != nil {
		return "", err
	}
	address, err := c.valist.GetTeamBeneficiary(c.callopts(ctx), teamID)
	if err != nil {
		return "", err
	}
	return address.Hex(), nil
}

// WriteJSON writes JSON data to storage.
func (c *Client) WriteJSON(ctx context.Context, data []byte) (string, error) {
	p, err := c.ipfs.Unixfs().Add(ctx, files.NewBytesFile(data), options.Unixfs.Pin(true))
	if err != nil {
		return "", err
	}
	return c.gateway + p.String(), nil
}

// callopts returns options for executing calls.
func (c *Client) callopts(ctx context.Context) *bind.CallOpts {
	opts := &bind.CallOpts{Context: ctx}
	if c.private != nil {
		opts.From = crypto.PubkeyToAddress(c.private.PublicKey)
	}
	return opts
}

// txopts returns options for executing transactions.
func (c *Client) txopts(ctx context.Context) (*bind.TransactOpts, error) {
	if c.private == nil {
		return nil, fmt.Errorf("signer key not set")
	}
	txopts, err := bind.NewKeyedTransactorWithChainID(c.private, c.chainID)
	if err != nil {
		return nil, err
	}
	txopts.Context = ctx
	return txopts, nil
}
