package client

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ipfs/go-ipfs/core/coreapi"
	coremock "github.com/ipfs/go-ipfs/core/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	valist "github.com/valist-io/valist-go"
	"github.com/valist-io/valist-go/contract"
	valist_contract "github.com/valist-io/valist-go/contract/valist"
	"github.com/valist-io/valist-go/storage"
)

var chainID = big.NewInt(1337)

func TestClient(t *testing.T) {
	priv, err := crypto.GenerateKey()
	require.NoError(t, err, "failed to generate private key")

	contract, err := setupEVM(priv)
	require.NoError(t, err, "failed to create evm contract")

	storage, err := setupIPFS()
	require.NoError(t, err, "failed to create ipfs storage")

	// set the signer for transactions
	contract.Signer(priv, chainID)
	client := NewClient(contract, storage)

	address := crypto.PubkeyToAddress(priv.PublicKey)
	members := []string{address.Hex()}

	team := &valist.Team{
		Name:        "valist",
		Description: "Web3 digital distribution",
		Homepage:    "https://valist.io",
	}

	project := &valist.Project{
		Name:        "sdk",
		Description: "Valist Typescript SDK",
		Homepage:    "https://docs.valist.io",
		Repository:  "https://github.com/valist-io/valist-js",
	}

	artifact := valist.Artifact{
		Provider: "Qm123",
	}

	release := &valist.Release{
		Name:    "sdk@v0.5.0",
		Version: "v0.5.0",
		Artifacts: map[string]valist.Artifact{
			"package.json": artifact,
		},
	}

	err = client.CreateTeam(context.Background(), "valist", team, members)
	require.NoError(t, err, "failed to create team")

	err = client.CreateProject(context.Background(), "valist", "sdk", project, members)
	require.NoError(t, err, "failed to create project")

	err = client.CreateRelease(context.Background(), "valist", "sdk", "v0.5.0", release)
	require.NoError(t, err, "failed to create release")

	otherTeam, err := client.GetTeam(context.Background(), "valist")
	require.NoError(t, err, "failed to get team")

	assert.Equal(t, team.Name, otherTeam.Name)
	assert.Equal(t, team.Description, otherTeam.Description)
	assert.Equal(t, team.Homepage, otherTeam.Homepage)

	otherProject, err := client.GetProject(context.Background(), "valist", "sdk")
	require.NoError(t, err, "failed to get project")

	assert.Equal(t, project.Name, otherProject.Name)
	assert.Equal(t, project.Description, otherProject.Description)
	assert.Equal(t, project.Homepage, otherProject.Homepage)
	assert.Equal(t, project.Repository, otherProject.Repository)

	otherRelease, err := client.GetRelease(context.Background(), "valist", "sdk", "v0.5.0")
	require.NoError(t, err, "failed to get release")

	assert.Equal(t, release.Name, otherRelease.Name)
	assert.Equal(t, release.Version, otherRelease.Version)
	assert.Equal(t, len(release.Artifacts), len(otherRelease.Artifacts))

	otherArtifact, ok := otherRelease.Artifacts["package.json"]
	require.True(t, ok)

	assert.Equal(t, artifact.Provider, otherArtifact.Provider)
}

func setupEVM(private *ecdsa.PrivateKey) (*contract.EVM, error) {
	// fund the account
	account := crypto.PubkeyToAddress(private.PublicKey)
	genesis := make(core.GenesisAlloc)
	genesis[account] = core.GenesisAccount{
		Balance: big.NewInt(9223372036854775807),
	}

	// simulate an ethereum blockchain
	backend := backends.NewSimulatedBackend(genesis, uint64(8000000))

	// sign the transaction with the account
	txopts, err := bind.NewKeyedTransactorWithChainID(private, chainID)
	if err != nil {
		return nil, err
	}

	// deploy the contract
	address, _, _, err := valist_contract.DeployValist(txopts, backend, common.HexToAddress("0x0"))
	if err != nil {
		return nil, err
	}

	// mine the transactions
	backend.Commit()

	return contract.NewEVM(address, backend)
}

func setupIPFS() (*storage.IPFS, error) {
	node, err := coremock.NewMockNode()
	if err != nil {
		return nil, err
	}
	api, err := coreapi.NewCoreAPI(node)
	if err != nil {
		return nil, err
	}
	return storage.NewIPFS(api), nil
}
