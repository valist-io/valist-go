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

	"github.com/valist-io/valist-go/contract/valist"
)

//go:generate abigen -abi valist/valist.abi -bin valist/valist.bin -out valist/valist.go -pkg valist -type Valist

type Backend interface {
	bind.DeployBackend
	bind.ContractBackend
}

type EVM struct {
	contract *valist.Valist
	backend  bind.DeployBackend
	private  *ecdsa.PrivateKey
	address  common.Address
	chainID  *big.Int
}

func NewEVM(address common.Address, backend Backend) (*EVM, error) {
	contract, err := valist.NewValist(address, backend)
	if err != nil {
		return nil, err
	}
	return &EVM{
		contract: contract,
		backend:  backend,
	}, nil
}

// Signer sets the transaction signing options
func (c *EVM) Signer(private *ecdsa.PrivateKey, chainID *big.Int) {
	c.private = private
	c.chainID = chainID
	c.address = crypto.PubkeyToAddress(private.PublicKey)
}

// CreateTeam creates a new team with the given members.
func (c *EVM) CreateTeam(ctx context.Context, teamName, metaURI string, members []string) error {
	var addresses []common.Address
	for _, hex := range members {
		addresses = append(addresses, common.HexToAddress(hex))
	}
	txopts, err := c.txopts(ctx)
	if err != nil {
		return err
	}
	tx, err := c.contract.CreateTeam(txopts, teamName, metaURI, addresses)
	if err != nil {
		return err
	}
	return c.waitMined(ctx, tx)
}

// CreateProject creates a new project. Requires the sender to be a member of the team.
func (c *EVM) CreateProject(ctx context.Context, teamName, projectName, metaURI string, members []string) error {
	var addresses []common.Address
	for _, hex := range members {
		addresses = append(addresses, common.HexToAddress(hex))
	}
	txopts, err := c.txopts(ctx)
	if err != nil {
		return err
	}
	tx, err := c.contract.CreateProject(txopts, teamName, projectName, metaURI, addresses)
	if err != nil {
		return err
	}
	return c.waitMined(ctx, tx)
}

// CreateRelease creates a new release. Requires the sender to be a member of the project.
func (c *EVM) CreateRelease(ctx context.Context, teamName, projectName, releaseName, metaURI string) error {
	txopts, err := c.txopts(ctx)
	if err != nil {
		return err
	}
	tx, err := c.contract.CreateRelease(txopts, teamName, projectName, releaseName, metaURI)
	if err != nil {
		return err
	}
	return c.waitMined(ctx, tx)
}

// AddTeamMember adds a member to the team. Requires the sender to be a member of the team.
func (c *EVM) AddTeamMember(ctx context.Context, teamName string, address string) error {
	txopts, err := c.txopts(ctx)
	if err != nil {
		return err
	}
	tx, err := c.contract.AddTeamMember(txopts, teamName, common.HexToAddress(address))
	if err != nil {
		return err
	}
	return c.waitMined(ctx, tx)
}

// RemoveTeamMember removes a member from the team. Requires the sender to be a member of the team.
func (c *EVM) RemoveTeamMember(ctx context.Context, teamName string, address string) error {
	txopts, err := c.txopts(ctx)
	if err != nil {
		return err
	}
	tx, err := c.contract.RemoveTeamMember(txopts, teamName, common.HexToAddress(address))
	if err != nil {
		return err
	}
	return c.waitMined(ctx, tx)
}

// AddProjectMemeber adds a member to the project. Requires the sender to be a member of the team.
func (c *EVM) AddProjectMember(ctx context.Context, teamName, projectName string, address string) error {
	txopts, err := c.txopts(ctx)
	if err != nil {
		return err
	}
	tx, err := c.contract.AddProjectMember(txopts, teamName, projectName, common.HexToAddress(address))
	if err != nil {
		return err
	}
	return c.waitMined(ctx, tx)
}

// RemoveProjectMember removes a member from the project. Requires the sender to be a member of the team.
func (c *EVM) RemoveProjectMember(ctx context.Context, teamName, projectName string, address string) error {
	txopts, err := c.txopts(ctx)
	if err != nil {
		return err
	}
	tx, err := c.contract.RemoveProjectMember(txopts, teamName, projectName, common.HexToAddress(address))
	if err != nil {
		return err
	}
	return c.waitMined(ctx, tx)
}

// SetTeamMetaURI sets the team metadata content ID. Requires the sender to be a member of the team.
func (c *EVM) SetTeamMetaURI(ctx context.Context, teamName, metaURI string) error {
	txopts, err := c.txopts(ctx)
	if err != nil {
		return err
	}
	tx, err := c.contract.SetTeamMetaCID(txopts, teamName, metaURI)
	if err != nil {
		return err
	}
	return c.waitMined(ctx, tx)
}

// SetProjectMetaURI sets the project metadata content ID. Requires the sender to be a member of the team.
func (c *EVM) SetProjectMetaURI(ctx context.Context, teamName, projectName, metaURI string) error {
	txopts, err := c.txopts(ctx)
	if err != nil {
		return err
	}
	tx, err := c.contract.SetProjectMetaCID(txopts, teamName, projectName, metaURI)
	if err != nil {
		return err
	}
	return c.waitMined(ctx, tx)
}

// ApproveRelease approves the release by adding the sender's address to the approvers list.
// The sender's address will be removed from the rejectors list if it exists.
func (c *EVM) ApproveRelease(ctx context.Context, teamName, projectName, releaseName string) error {
	txopts, err := c.txopts(ctx)
	if err != nil {
		return err
	}
	tx, err := c.contract.ApproveRelease(txopts, teamName, projectName, releaseName)
	if err != nil {
		return err
	}
	return c.waitMined(ctx, tx)
}

// RejectRelease rejects the release by adding the sender's address to the rejectors list.
// The sender's address will be removed from the approvers list if it exists.
func (c *EVM) RejectRelease(ctx context.Context, teamName, projectName, releaseName string) error {
	txopts, err := c.txopts(ctx)
	if err != nil {
		return err
	}
	tx, err := c.contract.RejectRelease(txopts, teamName, projectName, releaseName)
	if err != nil {
		return err
	}
	return c.waitMined(ctx, tx)
}

// GetLatestReleaseName returns the latest release name.
func (c *EVM) GetLatestReleaseName(ctx context.Context, teamName, projectName string) (string, error) {
	return c.contract.GetLatestReleaseName(c.callopts(ctx), teamName, projectName)
}

// GetTeamMetaURI returns the team metadata URI.
func (c *EVM) GetTeamMetaURI(ctx context.Context, teamName string) (string, error) {
	return c.contract.GetTeamMetaCID(c.callopts(ctx), teamName)
}

// GetProjectMetaURI returns the project metadata URI.
func (c *EVM) GetProjectMetaURI(ctx context.Context, teamName, projectName string) (string, error) {
	return c.contract.GetProjectMetaCID(c.callopts(ctx), teamName, projectName)
}

// GetReleaseMetaURI returns the release metadata URI.
func (c *EVM) GetReleaseMetaURI(ctx context.Context, teamName, projectName, releaseName string) (string, error) {
	return c.contract.GetReleaseMetaCID(c.callopts(ctx), teamName, projectName, releaseName)
}

// GetTeamNames returns a paginated list of team names.
func (c *EVM) GetTeamNames(ctx context.Context, page *big.Int, size *big.Int) ([]string, error) {
	return c.contract.GetTeamNames(c.callopts(ctx), page, size)
}

// GetProjectNames returns a paginated list of project names.
func (c *EVM) GetProjectNames(ctx context.Context, teamName string, page *big.Int, size *big.Int) ([]string, error) {
	return c.contract.GetProjectNames(c.callopts(ctx), teamName, page, size)
}

// GetReleaseNames returns a paginated list of release names.
func (c *EVM) GetReleaseNames(ctx context.Context, teamName, projectName string, page *big.Int, size *big.Int) ([]string, error) {
	return c.contract.GetReleaseNames(c.callopts(ctx), teamName, projectName, page, size)
}

// GetTeamMembers returns a paginated list of team members.
func (c *EVM) GetTeamMembers(ctx context.Context, teamName string, page *big.Int, size *big.Int) ([]string, error) {
	addresses, err := c.contract.GetTeamMembers(c.callopts(ctx), teamName, page, size)
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
	addresses, err := c.contract.GetProjectMembers(c.callopts(ctx), teamName, projectName, page, size)
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
	addresses, err := c.contract.GetReleaseApprovers(c.callopts(ctx), teamName, projectName, releaseName, page, size)
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
	addresses, err := c.contract.GetReleaseRejectors(c.callopts(ctx), teamName, projectName, releaseName, page, size)
	if err != nil {
		return nil, err
	}
	var rejectors []string
	for _, address := range addresses {
		rejectors = append(rejectors, address.Hex())
	}
	return rejectors, nil
}

// callopts returns options for executing calls.
func (c *EVM) callopts(ctx context.Context) *bind.CallOpts {
	return &bind.CallOpts{
		Context: ctx,
		From:    c.address,
	}
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

// waitMined waits for a transaction to be mined.
func (c *EVM) waitMined(ctx context.Context, tx *types.Transaction) error {
	if sim, ok := c.backend.(*backends.SimulatedBackend); ok {
		sim.Commit()
	}
	_, err := bind.WaitMined(ctx, c.backend, tx)
	return err
}
