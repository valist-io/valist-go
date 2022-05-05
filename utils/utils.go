package utils

import (
	"fmt"
	"io"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/crypto"
)

// Fetch returns the response of the given url.
func Fetch(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("request failed status=%d body=%s", res.StatusCode, body)
	}
	return body, nil
}

// GenerateID returns the ID of the object with the given parentID and name.
func GenerateID(parentID *big.Int, name string) *big.Int {
	// abi encode parent ID as uint256
	pack := make([]byte, 32)
	parentID.FillBytes(pack)
	// append name hash as abi encoded bytes32
	nameHash := crypto.Keccak256([]byte(name))
	pack = append(pack, nameHash...)
	// hash packed abi encoded values and convert to big int
	hash := crypto.Keccak256(pack)
	return big.NewInt(0).SetBytes(hash)
}
