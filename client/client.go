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
	"github.com/valist-io/valist-go/contracts"
	erc20_contract "github.com/valist-io/valist-go/contracts/erc20"
	license_contract "github.com/valist-io/valist-go/contracts/license"
	registry_contract "github.com/valist-io/valist-go/contracts/registry"
	"github.com/valist-io/valist-go/utils"
)

var DefaultOptions = &Options{
	EthereumRpc: "https://rpc.valist.io",
	IpfsHost:    "https://pin.valist.io",
	IpfsGateway: "https://gateway.valist.io",
}

type Options struct {
	// EthereumRpc is the URL of the ethereum RPC.
	EthereumRpc string
	// IpfsHost is the URL of the IPFS host.
	IpfsHost string
	// IpfsGateway is the URL of the IPFS gateway.
	IpfsGateway string
	// PrivateKey is the key to use for signing transactions.
	PrivateKey *ecdsa.PrivateKey
}

type backend interface {
	bind.ContractBackend
	bind.DeployBackend
}

type Client struct {
	rpc             backend
	ipfs            coreiface.CoreAPI
	gateway         string
	private         *ecdsa.PrivateKey
	chainID         *big.Int
	registry        *registry_contract.Registry
	license         *license_contract.License
	registryAddress common.Address
	licenseAddress  common.Address
}

// NewClient returns a client using the given contract and storage providers.
func NewClient(ctx context.Context, opts *Options) (*Client, error) {
	if opts == nil {
		opts = DefaultOptions
	}
	ipfs, err := httpapi.NewURLApiWithClient(opts.IpfsHost, &http.Client{})
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
	registryAddress, ok := contracts.RegistryAddresses[chainID.String()]
	if !ok {
		return nil, fmt.Errorf("chain with id=%d not supported", chainID)
	}
	licenseAddress, ok := contracts.LicenseAddresses[chainID.String()]
	if !ok {
		return nil, fmt.Errorf("chain with id=%d not supported", chainID)
	}
	registry, err := registry_contract.NewRegistry(registryAddress, rpc)
	if err != nil {
		return nil, err
	}
	license, err := license_contract.NewLicense(licenseAddress, rpc)
	if err != nil {
		return nil, err
	}
	return &Client{
		rpc:             rpc,
		ipfs:            ipfs,
		gateway:         opts.IpfsGateway,
		private:         opts.PrivateKey,
		chainID:         chainID,
		registry:        registry,
		license:         license,
		registryAddress: registryAddress,
		licenseAddress:  licenseAddress,
	}, nil
}

// CreateAccount creates an account with the given name, metadata, and members.
func (c *Client) CreateAccount(ctx context.Context, name string, account *valist.AccountMeta, members []string) (*types.Transaction, error) {
	data, err := json.Marshal(account)
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
	return c.registry.CreateAccount(txopts, name, metaURI, addresses)
}

// CreateProject creates a project under the account with the given id, metadata, and members.
func (c *Client) CreateProject(ctx context.Context, accountID *big.Int, name string, project *valist.ProjectMeta, members []string) (*types.Transaction, error) {
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
	return c.registry.CreateProject(txopts, accountID, name, metaURI, addresses)
}

// CreateRelease creates a release under the project with the given name and metadata.
func (c *Client) CreateRelease(ctx context.Context, projectID *big.Int, name string, release *valist.ReleaseMeta) (*types.Transaction, error) {
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
	return c.registry.CreateRelease(txopts, projectID, name, metaURI)
}

// GetAccountMeta returns the metadata of the account with the given ID.
func (c *Client) GetAccountMeta(ctx context.Context, accountID *big.Int) (*valist.AccountMeta, error) {
	metaURI, err := c.registry.MetaByID(c.callopts(ctx), accountID)
	if err != nil {
		return nil, err
	}
	data, err := utils.Fetch(metaURI)
	if err != nil {
		return nil, err
	}
	var account valist.AccountMeta
	if err := json.Unmarshal(data, &account); err != nil {
		return nil, err
	}
	return &account, nil
}

// GetProjectMeta returns the metadata of the project with the given ID.
func (c *Client) GetProjectMeta(ctx context.Context, projectID *big.Int) (*valist.ProjectMeta, error) {
	metaURI, err := c.registry.MetaByID(c.callopts(ctx), projectID)
	if err != nil {
		return nil, err
	}
	data, err := utils.Fetch(metaURI)
	if err != nil {
		return nil, err
	}
	var project valist.ProjectMeta
	if err := json.Unmarshal(data, &project); err != nil {
		return nil, err
	}
	return &project, nil
}

// GetReleaseMeta returns the metadata of the release with the given ID.
func (c *Client) GetReleaseMeta(ctx context.Context, releaseID *big.Int) (*valist.ReleaseMeta, error) {
	metaURI, err := c.registry.MetaByID(c.callopts(ctx), releaseID)
	if err != nil {
		return nil, err
	}
	data, err := utils.Fetch(metaURI)
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
func (c *Client) GetLatestReleaseMeta(ctx context.Context, projectID *big.Int) (*valist.ReleaseMeta, error) {
	releaseID, err := c.GetLatestReleaseID(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return c.GetReleaseMeta(ctx, releaseID)
}

// AccountExists returns true if the account with the given ID exists.
func (c *Client) AccountExists(ctx context.Context, accountID *big.Int) (bool, error) {
	metaURI, err := c.registry.MetaByID(c.callopts(ctx), accountID)
	if err != nil {
		return false, err
	}
	return metaURI != "", nil
}

// ProjectExists returns true if the project with the given ID exists.
func (c *Client) ProjectExists(ctx context.Context, projectID *big.Int) (bool, error) {
	metaURI, err := c.registry.MetaByID(c.callopts(ctx), projectID)
	if err != nil {
		return false, err
	}
	return metaURI != "", nil
}

// ReleaseExists returns true if the release with the given ID exists.
func (c *Client) ReleaseExists(ctx context.Context, releaseID *big.Int) (bool, error) {
	metaURI, err := c.registry.MetaByID(c.callopts(ctx), releaseID)
	if err != nil {
		return false, err
	}
	return metaURI != "", nil
}

// AddAccountMember adds a member to the account. Requires the sender to be a member of the account.
func (c *Client) AddAccountMember(ctx context.Context, accountID *big.Int, address string) (*types.Transaction, error) {
	txopts, err := c.txopts(ctx)
	if err != nil {
		return nil, err
	}
	return c.registry.AddAccountMember(txopts, accountID, common.HexToAddress(address))
}

// RemoveAccountMember removes a member from the account. Requires the sender to be a member of the account.
func (c *Client) RemoveAccountMember(ctx context.Context, accountID *big.Int, address string) (*types.Transaction, error) {
	txopts, err := c.txopts(ctx)
	if err != nil {
		return nil, err
	}
	return c.registry.RemoveAccountMember(txopts, accountID, common.HexToAddress(address))
}

// AddProjectMemeber adds a member to the project. Requires the sender to be a member of the account.
func (c *Client) AddProjectMember(ctx context.Context, projectID *big.Int, address string) (*types.Transaction, error) {
	txopts, err := c.txopts(ctx)
	if err != nil {
		return nil, err
	}
	return c.registry.AddProjectMember(txopts, projectID, common.HexToAddress(address))
}

// RemoveProjectMember removes a member from the project. Requires the sender to be a member of the account.
func (c *Client) RemoveProjectMember(ctx context.Context, projectID *big.Int, address string) (*types.Transaction, error) {
	txopts, err := c.txopts(ctx)
	if err != nil {
		return nil, err
	}
	return c.registry.RemoveProjectMember(txopts, projectID, common.HexToAddress(address))
}

// SetAccountMetaURI sets the account metadata. Requires the sender to be a member of the account.
func (c *Client) SetAccountMetaURI(ctx context.Context, accountID *big.Int, account *valist.AccountMeta) (*types.Transaction, error) {
	data, err := json.Marshal(account)
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
	return c.registry.SetAccountMetaURI(txopts, accountID, metaURI)
}

// SetProjectMetaURI sets the project metadata content ID. Requires the sender to be a member of the account.
func (c *Client) SetProjectMetaURI(ctx context.Context, projectID *big.Int, project *valist.ProjectMeta) (*types.Transaction, error) {
	data, err := json.Marshal(project)
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
	return c.registry.SetProjectMetaURI(txopts, projectID, metaURI)
}

// ApproveRelease adds the senders address to the signers list.
func (c *Client) ApproveRelease(ctx context.Context, releaseID *big.Int) (*types.Transaction, error) {
	txopts, err := c.txopts(ctx)
	if err != nil {
		return nil, err
	}
	return c.registry.ApproveRelease(txopts, releaseID)
}

// RevokeRelease removes the senders address from the signers list.
func (c *Client) RevokeRelease(ctx context.Context, releaseID *big.Int) (*types.Transaction, error) {
	txopts, err := c.txopts(ctx)
	if err != nil {
		return nil, err
	}
	return c.registry.RevokeRelease(txopts, releaseID)
}

// SetProductLimit sets the limit on the total supply of products in the given project.
func (c *Client) SetProductLimit(ctx context.Context, projectID, limit *big.Int) (*types.Transaction, error) {
	txopts, err := c.txopts(ctx)
	if err != nil {
		return nil, err
	}
	return c.license.SetLimit(txopts, projectID, limit)
}

// SetProductRoyalty sets the royalty recipient and amount in basis points for the given project.
func (c *Client) SetProductRoyalty(ctx context.Context, projectID *big.Int, recipient string, amount *big.Int) (*types.Transaction, error) {
	txopts, err := c.txopts(ctx)
	if err != nil {
		return nil, err
	}
	return c.license.SetRoyalty(txopts, projectID, common.HexToAddress(recipient), amount)
}

// SetProductPrice sets the price of the product for the given project.
func (c *Client) SetProductPrice(ctx context.Context, projectID, price *big.Int) (*types.Transaction, error) {
	txopts, err := c.txopts(ctx)
	if err != nil {
		return nil, err
	}
	return c.license.SetPrice0(txopts, projectID, price)
}

// SetProductTokenPrice sets the token price of the product for the given project.
func (c *Client) SetProductTokenPrice(ctx context.Context, token string, projectID, price *big.Int) (*types.Transaction, error) {
	txopts, err := c.txopts(ctx)
	if err != nil {
		return nil, err
	}
	return c.license.SetPrice(txopts, common.HexToAddress(token), projectID, price)
}

// PurchaseProduct purchases the product for the given project on behalf of the recipient.
func (c *Client) PurchaseProduct(ctx context.Context, projectID *big.Int, recipient string) (*types.Transaction, error) {
	price, err := c.GetProductPrice(ctx, projectID)
	if err != nil {
		return nil, err
	}
	txopts, err := c.txopts(ctx)
	if err != nil {
		return nil, err
	}
	txopts.Value = price
	return c.license.Purchase(txopts, projectID, common.HexToAddress(recipient))
}

// PurchaseProductToken purchases the product for the given project on behalf of the recipient.
func (c *Client) PurchaseProductToken(ctx context.Context, token string, projectID *big.Int, recipient string) (*types.Transaction, error) {
	price, err := c.GetProductTokenPrice(ctx, token, projectID)
	if err != nil {
		return nil, err
	}
	erc20, err := erc20_contract.NewERC20(common.HexToAddress(token), c.rpc)
	if err != nil {
		return nil, err
	}
	txopts, err := c.txopts(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := erc20.Approve(txopts, c.licenseAddress, price)
	if err != nil {
		return nil, err
	}
	_, err = bind.WaitMined(ctx, c.rpc, tx)
	if err != nil {
		return nil, err
	}
	return c.license.Purchase0(txopts, common.HexToAddress(token), projectID, common.HexToAddress(recipient))
}

// GetProductPrice returns the price of the product for the given project.
func (c *Client) GetProductPrice(ctx context.Context, projectID *big.Int) (*big.Int, error) {
	return c.license.GetPrice0(c.callopts(ctx), projectID)
}

// GetProductTokenPrice returns the token price of the product for the given project.
func (c *Client) GetProductTokenPrice(ctx context.Context, token string, projectID *big.Int) (*big.Int, error) {
	return c.license.GetPrice(c.callopts(ctx), common.HexToAddress(token), projectID)
}

// GetProductBalance returns the current balance of the product for the given project.
func (c *Client) GetProductBalance(ctx context.Context, projectID *big.Int) (*big.Int, error) {
	return c.license.GetBalance(c.callopts(ctx), projectID)
}

// GetProductTokenBalance returns the current balance of the product for the given project.
func (c *Client) GetProductTokenBalance(ctx context.Context, token string, projectID *big.Int) (*big.Int, error) {
	return c.license.GetBalance0(c.callopts(ctx), common.HexToAddress(token), projectID)
}

// GetProductLimit returns the limit of supply for the product in the given project.
func (c *Client) GetProductLimit(ctx context.Context, projectID *big.Int) (*big.Int, error) {
	return c.license.GetLimit(c.callopts(ctx), projectID)
}

// GetProductSupply returns the total supply of the product in the given project.
func (c *Client) GetProductSupply(ctx context.Context, projectID *big.Int) (*big.Int, error) {
	return c.license.GetSupply(c.callopts(ctx), projectID)
}

// ProductRoyaltyInfo returns the royalty recipient and amount owed for product.
func (c *Client) ProductRoyaltyInfo(ctx context.Context, projectID, price *big.Int) (common.Address, *big.Int, error) {
	return c.license.RoyaltyInfo(c.callopts(ctx), projectID, price)
}

// ProductBalanceOf returns the product balance of the address for the given project.
func (c *Client) ProductBalanceOf(ctx context.Context, address string, projectID *big.Int) (*big.Int, error) {
	return c.license.BalanceOf(c.callopts(ctx), common.HexToAddress(address), projectID)
}

// GetAccountMembers returns a list of account members.
func (c *Client) GetAccountMembers(ctx context.Context, accountID *big.Int) ([]common.Address, error) {
	return c.registry.GetAccountMembers(c.callopts(ctx), accountID)
}

// GetProjectMembers returns a list of project members.
func (c *Client) GetProjectMembers(ctx context.Context, projectID *big.Int) ([]common.Address, error) {
	return c.registry.GetProjectMembers(c.callopts(ctx), projectID)
}

// GetReleaseSigners returns a list of release signers.
func (c *Client) GetReleaseSigners(ctx context.Context, releaseID *big.Int) ([]common.Address, error) {
	return c.registry.GetReleaseSigners(c.callopts(ctx), releaseID)
}

// GetLatestReleaseID returns the ID of the latest release for the given project.
func (c *Client) GetLatestReleaseID(ctx context.Context, projectID *big.Int) (*big.Int, error) {
	return c.registry.GetLatestReleaseID(c.callopts(ctx), projectID)
}

// GetPreviousReleaseID returns the ID of the previous release for the given release.
func (c *Client) GetPreviousReleaseID(ctx context.Context, releaseID *big.Int) (*big.Int, error) {
	return c.registry.GetPreviousReleaseID(c.callopts(ctx), releaseID)
}

// GetProjectAccountID returns the ID of the account for the given project.
func (c *Client) GetProjectAccountID(ctx context.Context, projectID *big.Int) (*big.Int, error) {
	return c.registry.GetProjectAccountID(c.callopts(ctx), projectID)
}

// GetReleaseProjectID returns the ID of the project for the given release.
func (c *Client) GetReleaseProjectID(ctx context.Context, releaseID *big.Int) (*big.Int, error) {
	return c.registry.GetReleaseProjectID(c.callopts(ctx), releaseID)
}

// IsAccountMember returns true if the given address is an account member.
func (c *Client) IsAccountMember(ctx context.Context, accountID *big.Int, address string) (bool, error) {
	return c.registry.IsAccountMember(c.callopts(ctx), accountID, common.HexToAddress(address))
}

// IsProjectMember returns true if the given address is a project member.
func (c *Client) IsProjectMember(ctx context.Context, projectID *big.Int, address string) (bool, error) {
	return c.registry.IsProjectMember(c.callopts(ctx), projectID, common.HexToAddress(address))
}

// IsReleaseSigner returns true if the given address is a release signer.
func (c *Client) IsReleaseSigner(ctx context.Context, releaseID *big.Int, address string) (bool, error) {
	return c.registry.IsReleaseSigner(c.callopts(ctx), releaseID, common.HexToAddress(address))
}

// GenerateID returns the ID of the object with the given parentID and name.
func (c *Client) GenerateID(parentID *big.Int, name string) *big.Int {
	return utils.GenerateID(parentID, name)
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
