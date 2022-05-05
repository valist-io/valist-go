package client

import (
	"context"
	"math/big"
	"net/http"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	httpapi "github.com/ipfs/go-ipfs-http-client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	valist "github.com/valist-io/valist-go"
	license_contract "github.com/valist-io/valist-go/contracts/license"
	registry_contract "github.com/valist-io/valist-go/contracts/registry"
)

var chainID = big.NewInt(1337)

func TestClient(t *testing.T) {
	// generate deterministic key
	private, err := crypto.HexToECDSA("289c2857d4598e37fb9647507e47a309d6133539bf21a8b9cb6df88fd5232032")
	require.NoError(t, err, "failed to generate private key")

	// fund the account
	address := crypto.PubkeyToAddress(private.PublicKey)
	genesis := make(core.GenesisAlloc)
	genesis[address] = core.GenesisAccount{
		Balance: big.NewInt(9223372036854775807),
	}

	// simulate an ethereum blockchain
	backend := backends.NewSimulatedBackend(genesis, uint64(8000000))

	// sign the transaction with the account
	txopts, err := bind.NewKeyedTransactorWithChainID(private, chainID)
	require.NoError(t, err, "failed create keyed transactor")

	// deploy the contracts
	registryAddress, _, registryContract, err := registry_contract.DeployRegistry(txopts, backend, common.HexToAddress("0x0"))
	require.NoError(t, err, "failed deploy valist contract")

	licenseAddress, _, licenseContract, err := license_contract.DeployLicense(txopts, backend, registryAddress)
	require.NoError(t, err, "failed deploy license contract")

	// mine the transactions
	backend.Commit()

	ipfs, err := httpapi.NewURLApiWithClient("https://pin.valist.io", &http.Client{})
	require.NoError(t, err, "failed to create IPFS http client")

	client := &Client{
		rpc:             backend,
		ipfs:            ipfs,
		gateway:         "https://gateway.valist.io",
		chainID:         chainID,
		private:         private,
		registry:        registryContract,
		license:         licenseContract,
		registryAddress: registryAddress,
		licenseAddress:  licenseAddress,
	}

	account := &valist.AccountMeta{
		Image:       "https://gateway.valist.io/ipfs/Qm456",
		Name:        "valist",
		Description: "Web3 digital distribution",
		ExternalURL: "https://valist.io",
	}

	project := &valist.ProjectMeta{
		Image:       "https://gateway.valist.io/ipfs/Qm456",
		Name:        "sdk",
		Description: "Valist Typescript SDK",
		ExternalURL: "https://github.com/valist-io/valist-js",
	}

	release := &valist.ReleaseMeta{
		Image:       "https://gateway.valist.io/ipfs/Qm456",
		Name:        "sdk@v0.5.0",
		Description: "Release v0.5.0",
		ExternalURL: "https://gateway.valist.io/ipfs/Qm123",
	}

	accountID := client.GenerateID(chainID, "valist")
	projectID := client.GenerateID(accountID, "sdk")
	releaseID := client.GenerateID(projectID, "v0.5.0")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_, err = client.CreateAccount(ctx, "valist", account, []string{address.Hex()})
	require.NoError(t, err, "failed to create account")
	backend.Commit()

	_, err = client.CreateProject(ctx, accountID, "sdk", project, []string{})
	require.NoError(t, err, "failed to create project")
	backend.Commit()

	_, err = client.CreateRelease(ctx, projectID, "v0.5.0", release)
	require.NoError(t, err, "failed to create release")
	backend.Commit()

	otherAccount, err := client.GetAccountMeta(ctx, accountID)
	require.NoError(t, err, "failed to get account")
	assert.Equal(t, account, otherAccount)

	otherProject, err := client.GetProjectMeta(ctx, projectID)
	require.NoError(t, err, "failed to get project")
	assert.Equal(t, project, otherProject)

	otherRelease, err := client.GetReleaseMeta(ctx, releaseID)
	require.NoError(t, err, "failed to get release")
	assert.Equal(t, release, otherRelease)
}
