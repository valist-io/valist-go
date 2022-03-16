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
	"github.com/ipfs/go-ipfs/core/corehttp"
	coremock "github.com/ipfs/go-ipfs/core/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	valist "github.com/valist-io/valist-go"
	"github.com/valist-io/valist-go/contract"
	license_contract "github.com/valist-io/valist-go/contract/license"
	valist_contract "github.com/valist-io/valist-go/contract/valist"
	"github.com/valist-io/valist-go/storage"
)

var chainID = big.NewInt(1337)

// hard coded private key for deterministic deploys
const privateKey = "289c2857d4598e37fb9647507e47a309d6133539bf21a8b9cb6df88fd5232032"

func TestClient(t *testing.T) {
	priv, err := crypto.HexToECDSA(privateKey)
	require.NoError(t, err, "failed to generate private key")

	contract, err := setupEVM(priv)
	require.NoError(t, err, "failed to create evm contract")

	storage, err := setupIPFS()
	require.NoError(t, err, "failed to create ipfs storage")

	client := NewClient(contract, storage)
	address := crypto.PubkeyToAddress(priv.PublicKey)
	members := []string{address.Hex()}

	team := &valist.TeamMeta{
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

	license := &valist.LicenseMeta{
		Image:       "https://gateway.valist.io/ipfs/Qm789",
		Name:        "Valist Pro",
		Description: "Access pro tier features on Valist",
		ExternalURL: "https://valist.io/pro",
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tx, err := client.CreateTeam(ctx, "valist", team, address.Hex(), members)
	require.NoError(t, err, "failed to create team")
	tx.Wait(ctx)

	tx, err = client.CreateProject(ctx, "valist", "sdk", project, members)
	require.NoError(t, err, "failed to create project")
	tx.Wait(ctx)

	tx, err = client.CreateRelease(ctx, "valist", "sdk", "v0.5.0", release)
	require.NoError(t, err, "failed to create release")
	tx.Wait(ctx)

	tx, err = client.CreateLicense(ctx, "valist", "sdk", "pro", license, big.NewInt(1000))
	require.NoError(t, err, "failed to create license")
	tx.Wait(ctx)

	otherTeam, err := client.GetTeamMeta(ctx, "valist")
	require.NoError(t, err, "failed to get team")
	assert.Equal(t, team, otherTeam)

	otherProject, err := client.GetProjectMeta(ctx, "valist", "sdk")
	require.NoError(t, err, "failed to get project")
	assert.Equal(t, project, otherProject)

	otherRelease, err := client.GetReleaseMeta(ctx, "valist", "sdk", "v0.5.0")
	require.NoError(t, err, "failed to get release")
	assert.Equal(t, release, otherRelease)

	otherLicense, err := client.GetLicenseMeta(ctx, "valist", "sdk", "pro")
	require.NoError(t, err, "failed to get license")
	assert.Equal(t, license, otherLicense)

	beneficiary, err := client.Contract.GetTeamBeneficiary(ctx, "valist")
	require.NoError(t, err, "failed to get team beneficiary")
	assert.Equal(t, beneficiary, address.Hex())

	tx, err = client.Contract.SetTeamBeneficiary(ctx, "valist", "0x9399BB24DBB5C4b782C70c2969F58716Ebbd6a3b")
	require.NoError(t, err, "failed to set team beneficiary")
	tx.Wait(ctx)

	otherBeneficiary, err := client.Contract.GetTeamBeneficiary(ctx, "valist")
	require.NoError(t, err, "failed to get team beneficiary")
	assert.Equal(t, otherBeneficiary, "0x9399BB24DBB5C4b782C70c2969F58716Ebbd6a3b")

	licensePrice, err := client.Contract.GetLicensePrice(ctx, "valist", "sdk", "pro")
	require.NoError(t, err, "failed to get license price")
	assert.Equal(t, licensePrice, big.NewInt(1000))

	tx, err = client.Contract.MintLicense(ctx, "valist", "sdk", "pro", address.Hex())
	require.NoError(t, err, "failed to mint license")
	tx.Wait(ctx)
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

	// deploy the contracts
	address, _, _, err := valist_contract.DeployValist(txopts, backend, common.HexToAddress("0x0"))
	if err != nil {
		return nil, err
	}
	_, _, _, err = license_contract.DeployLicense(txopts, backend, address, common.HexToAddress("0x0"))
	if err != nil {
		return nil, err
	}

	// mine the transactions
	backend.Commit()

	return contract.NewEVM(backend, chainID, private)
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

	// start the http server
	go corehttp.ListenAndServe(node, "/ip4/127.0.0.1/tcp/9090", corehttp.GatewayOption(true, "/ipfs", "/ipns"))

	return storage.NewIPFS(api, "http://localhost:9090"), nil
}
