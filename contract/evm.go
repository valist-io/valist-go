package contract

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	valist "github.com/valist-io/valist-go"
	license_contract "github.com/valist-io/valist-go/contract/license"
	valist_contract "github.com/valist-io/valist-go/contract/valist"
)

//go:generate abigen -abi valist/valist.abi -bin valist/valist.bin -out valist/valist.go -pkg valist -type Valist
//go:generate abigen -abi license/license.abi -bin license/license.bin -out license/license.go -pkg license -type License

var valistAddresses = map[string]common.Address{
	// Deterministic simulated
	"1337": common.HexToAddress("0x333c3310824b7c685133F2BeDb2CA4b8b4DF633d"),
	// Mumbai testnet
	"80001": common.HexToAddress("0x9569bEb0Eba900495cF58028DB094D824d0AE850"),
	// Polygon mainnet
	"137": common.HexToAddress("0xc70A069eC7F887a7497a4bdC7bE666C1e18c8DC3"),
}

var licenseAddresses = map[string]common.Address{
	// Deterministic simulated
	"1337": common.HexToAddress("0x8bDa78331C916A08481428e4b07C96D3e916D165"),
	// Mumbai testnet
	"80001": common.HexToAddress("0x597bfcE7F9363b6eBc229f2023F9EcD716C88120"),
	// Polygon mainnet
	"137": common.HexToAddress("0xb85ed41d49Eba25aE6186921Ea63b6055903e810"),
}

type EVM_Transaction struct {
	transaction *types.Transaction
	backend     bind.DeployBackend
}

// Wait waits until the transaction is finalized.
func (tx *EVM_Transaction) Wait(ctx context.Context) error {
	if sim, ok := tx.backend.(*backends.SimulatedBackend); ok {
		sim.Commit()
	}
	_, err := bind.WaitMined(ctx, tx.backend, tx.transaction)
	return err
}

// Hash returns the transaction hash.
func (tx *EVM_Transaction) Hash() string {
	return tx.transaction.Hash().Hex()
}

type Backend interface {
	bind.DeployBackend
	bind.ContractBackend
}

type EVM struct {
	valist  *valist_contract.Valist
	license *license_contract.License
	private *ecdsa.PrivateKey
	chainID *big.Int
	backend Backend
}

func NewEVM(backend Backend, chainID *big.Int, private *ecdsa.PrivateKey) (*EVM, error) {
	valistAddress, ok := valistAddresses[chainID.String()]
	if !ok {
		return nil, fmt.Errorf("chain with id=%d not supported", chainID)
	}
	licenseAddress, ok := licenseAddresses[chainID.String()]
	if !ok {
		return nil, fmt.Errorf("chain with id=%d not supported", chainID)
	}
	valist, err := valist_contract.NewValist(valistAddress, backend)
	if err != nil {
		return nil, err
	}
	license, err := license_contract.NewLicense(licenseAddress, backend)
	if err != nil {
		return nil, err
	}
	return &EVM{
		valist:  valist,
		license: license,
		private: private,
		chainID: chainID,
		backend: backend,
	}, nil
}

// CreateTeam creates a new team with the given members.
func (c *EVM) CreateTeam(ctx context.Context, teamName, metaURI, beneficiary string, members []string) (valist.TransactionAPI, error) {
	var addresses []common.Address
	for _, hex := range members {
		addresses = append(addresses, common.HexToAddress(hex))
	}
	txopts, err := c.txopts(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := c.valist.CreateTeam(txopts, teamName, metaURI, common.HexToAddress(beneficiary), addresses)
	if err != nil {
		return nil, err
	}
	return &EVM_Transaction{tx, c.backend}, nil
}

// CreateProject creates a new project. Requires the sender to be a member of the team.
func (c *EVM) CreateProject(ctx context.Context, teamName, projectName, metaURI string, members []string) (valist.TransactionAPI, error) {
	var addresses []common.Address
	for _, hex := range members {
		addresses = append(addresses, common.HexToAddress(hex))
	}
	txopts, err := c.txopts(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := c.valist.CreateProject(txopts, teamName, projectName, metaURI, addresses)
	if err != nil {
		return nil, err
	}
	return &EVM_Transaction{tx, c.backend}, nil
}

// CreateRelease creates a new release. Requires the sender to be a member of the project.
func (c *EVM) CreateRelease(ctx context.Context, teamName, projectName, releaseName, metaURI string) (valist.TransactionAPI, error) {
	txopts, err := c.txopts(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := c.valist.CreateRelease(txopts, teamName, projectName, releaseName, metaURI)
	if err != nil {
		return nil, err
	}
	return &EVM_Transaction{tx, c.backend}, nil
}

// CreateLicense Creates a new License and establishes the mint price.
func (c *EVM) CreateLicense(ctx context.Context, teamName, projectName, licenseName, metaURI string, mintPrice *big.Int) (valist.TransactionAPI, error) {
	txopts, err := c.txopts(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := c.license.CreateLicense(txopts, teamName, projectName, licenseName, metaURI, mintPrice)
	if err != nil {
		return nil, err
	}
	return &EVM_Transaction{tx, c.backend}, nil
}

// MintLicense mints a new license to a recipient.
func (c *EVM) MintLicense(ctx context.Context, teamName, projectName, licenseName, recipient string) (valist.TransactionAPI, error) {
	txopts, err := c.txopts(ctx)
	if err != nil {
		return nil, err
	}
	txopts.Value, err = c.GetLicensePrice(ctx, teamName, projectName, licenseName)
	if err != nil {
		return nil, err
	}
	tx, err := c.license.MintLicense(txopts, teamName, projectName, licenseName, common.HexToAddress(recipient))
	if err != nil {
		return nil, err
	}
	return &EVM_Transaction{tx, c.backend}, nil
}

// AddTeamMember adds a member to the team. Requires the sender to be a member of the team.
func (c *EVM) AddTeamMember(ctx context.Context, teamName string, address string) (valist.TransactionAPI, error) {
	txopts, err := c.txopts(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := c.valist.AddTeamMember(txopts, teamName, common.HexToAddress(address))
	if err != nil {
		return nil, err
	}
	return &EVM_Transaction{tx, c.backend}, nil
}

// RemoveTeamMember removes a member from the team. Requires the sender to be a member of the team.
func (c *EVM) RemoveTeamMember(ctx context.Context, teamName string, address string) (valist.TransactionAPI, error) {
	txopts, err := c.txopts(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := c.valist.RemoveTeamMember(txopts, teamName, common.HexToAddress(address))
	if err != nil {
		return nil, err
	}
	return &EVM_Transaction{tx, c.backend}, nil
}

// AddProjectMemeber adds a member to the project. Requires the sender to be a member of the team.
func (c *EVM) AddProjectMember(ctx context.Context, teamName, projectName string, address string) (valist.TransactionAPI, error) {
	txopts, err := c.txopts(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := c.valist.AddProjectMember(txopts, teamName, projectName, common.HexToAddress(address))
	if err != nil {
		return nil, err
	}
	return &EVM_Transaction{tx, c.backend}, nil
}

// RemoveProjectMember removes a member from the project. Requires the sender to be a member of the team.
func (c *EVM) RemoveProjectMember(ctx context.Context, teamName, projectName string, address string) (valist.TransactionAPI, error) {
	txopts, err := c.txopts(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := c.valist.RemoveProjectMember(txopts, teamName, projectName, common.HexToAddress(address))
	if err != nil {
		return nil, err
	}
	return &EVM_Transaction{tx, c.backend}, nil
}

// SetTeamMetaURI sets the team metadata content ID. Requires the sender to be a member of the team.
func (c *EVM) SetTeamMetaURI(ctx context.Context, teamName, metaURI string) (valist.TransactionAPI, error) {
	txopts, err := c.txopts(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := c.valist.SetTeamMetaURI(txopts, teamName, metaURI)
	if err != nil {
		return nil, err
	}
	return &EVM_Transaction{tx, c.backend}, nil
}

// SetProjectMetaURI sets the project metadata content ID. Requires the sender to be a member of the team.
func (c *EVM) SetProjectMetaURI(ctx context.Context, teamName, projectName, metaURI string) (valist.TransactionAPI, error) {
	txopts, err := c.txopts(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := c.valist.SetProjectMetaURI(txopts, teamName, projectName, metaURI)
	if err != nil {
		return nil, err
	}
	return &EVM_Transaction{tx, c.backend}, nil
}

// SetTeamBeneficiary sets the team beneficiary to the new address.
func (c *EVM) SetTeamBeneficiary(ctx context.Context, teamName, beneficiary string) (valist.TransactionAPI, error) {
	txopts, err := c.txopts(ctx)
	if err != nil {
		return nil, err
	}
	teamID, err := c.GetTeamID(ctx, teamName)
	if err != nil {
		return nil, err
	}
	tx, err := c.valist.SetTeamBeneficiary(txopts, teamID, common.HexToAddress(beneficiary))
	if err != nil {
		return nil, err
	}
	return &EVM_Transaction{tx, c.backend}, nil
}

// ApproveRelease approves the release by adding the sender's address to the approvers list.
// The sender's address will be removed from the rejectors list if it exists.
func (c *EVM) ApproveRelease(ctx context.Context, teamName, projectName, releaseName string) (valist.TransactionAPI, error) {
	txopts, err := c.txopts(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := c.valist.ApproveRelease(txopts, teamName, projectName, releaseName)
	if err != nil {
		return nil, err
	}
	return &EVM_Transaction{tx, c.backend}, nil
}

// RejectRelease rejects the release by adding the sender's address to the rejectors list.
// The sender's address will be removed from the approvers list if it exists.
func (c *EVM) RejectRelease(ctx context.Context, teamName, projectName, releaseName string) (valist.TransactionAPI, error) {
	txopts, err := c.txopts(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := c.valist.RejectRelease(txopts, teamName, projectName, releaseName)
	if err != nil {
		return nil, err
	}
	return &EVM_Transaction{tx, c.backend}, nil
}

// GetLatestReleaseName returns the latest release name.
func (c *EVM) GetLatestReleaseName(ctx context.Context, teamName, projectName string) (string, error) {
	return c.valist.GetLatestReleaseName(c.callopts(ctx), teamName, projectName)
}

// GetTeamMetaURI returns the team metadata URI.
func (c *EVM) GetTeamMetaURI(ctx context.Context, teamName string) (string, error) {
	return c.valist.GetTeamMetaURI(c.callopts(ctx), teamName)
}

// GetProjectMetaURI returns the project metadata URI.
func (c *EVM) GetProjectMetaURI(ctx context.Context, teamName, projectName string) (string, error) {
	return c.valist.GetProjectMetaURI(c.callopts(ctx), teamName, projectName)
}

// GetReleaseMetaURI returns the release metadata URI.
func (c *EVM) GetReleaseMetaURI(ctx context.Context, teamName, projectName, releaseName string) (string, error) {
	return c.valist.GetReleaseMetaURI(c.callopts(ctx), teamName, projectName, releaseName)
}

// GetLicenseMetaURI returns the license metadata URI.
func (c *EVM) GetLicenseMetaURI(ctx context.Context, teamName, projectName, licenseName string) (string, error) {
	return c.license.GetLicenseMetaURI(c.callopts(ctx), teamName, projectName, licenseName)
}

// GetTeamNames returns a paginated list of team names.
func (c *EVM) GetTeamNames(ctx context.Context, page *big.Int, size *big.Int) ([]string, error) {
	return c.valist.GetTeamNames(c.callopts(ctx), page, size)
}

// GetProjectNames returns a paginated list of project names.
func (c *EVM) GetProjectNames(ctx context.Context, teamName string, page *big.Int, size *big.Int) ([]string, error) {
	return c.valist.GetProjectNames(c.callopts(ctx), teamName, page, size)
}

// GetReleaseNames returns a paginated list of release names.
func (c *EVM) GetReleaseNames(ctx context.Context, teamName, projectName string, page *big.Int, size *big.Int) ([]string, error) {
	return c.valist.GetReleaseNames(c.callopts(ctx), teamName, projectName, page, size)
}

// GetLicenseNames returns a paginated list of license names.
func (c *EVM) GetLicenseNames(ctx context.Context, teamName, projectName string, page *big.Int, size *big.Int) ([]string, error) {
	return c.license.GetNamesByProjectID(c.callopts(ctx), teamName, projectName, page, size)
}

// GetTeamMembers returns a paginated list of team members.
func (c *EVM) GetTeamMembers(ctx context.Context, teamName string, page *big.Int, size *big.Int) ([]string, error) {
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
func (c *EVM) GetProjectMembers(ctx context.Context, teamName, projectName string, page *big.Int, size *big.Int) ([]string, error) {
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
func (c *EVM) GetReleaseApprovers(ctx context.Context, teamName string, projectName, releaseName string, page *big.Int, size *big.Int) ([]string, error) {
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
func (c *EVM) GetReleaseRejectors(ctx context.Context, teamName string, projectName, releaseName string, page *big.Int, size *big.Int) ([]string, error) {
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
func (c *EVM) GetTeamID(ctx context.Context, teamName string) (*big.Int, error) {
	return c.valist.GetTeamID(c.callopts(ctx), teamName)
}

// GetProjectID generates projectID from team ID and name.
func (c *EVM) GetProjectID(ctx context.Context, teamID *big.Int, projectName string) (*big.Int, error) {
	return c.valist.GetProjectID(c.callopts(ctx), teamID, projectName)
}

// GetReleaseID generates releaseID from project ID and name.
func (c *EVM) GetReleaseID(ctx context.Context, projectID *big.Int, releaseName string) (*big.Int, error) {
	return c.valist.GetReleaseID(c.callopts(ctx), projectID, releaseName)
}

// GetLicenseID generates licenseID from project ID and name.
func (c *EVM) GetLicenseID(ctx context.Context, projectID *big.Int, licenseName string) (*big.Int, error) {
	return c.license.GetLicenseID(c.callopts(ctx), projectID, licenseName)
}

// GetLicense price returns the mint price of the license.
func (c *EVM) GetLicensePrice(ctx context.Context, teamName, projectName, licenseName string) (*big.Int, error) {
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
func (c *EVM) GetTeamBeneficiary(ctx context.Context, teamName string) (string, error) {
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

// callopts returns options for executing calls.
func (c *EVM) callopts(ctx context.Context) *bind.CallOpts {
	opts := &bind.CallOpts{Context: ctx}
	if c.private != nil {
		opts.From = crypto.PubkeyToAddress(c.private.PublicKey)
	}
	return opts
}

// txopts returns options for executing transactions.
func (c *EVM) txopts(ctx context.Context) (*bind.TransactOpts, error) {
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
