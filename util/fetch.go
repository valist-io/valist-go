package util

import (
	"fmt"
	"io"
	"net/http"
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
