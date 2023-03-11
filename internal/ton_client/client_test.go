package tonclient

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConstructURL(t *testing.T) {
	baseURL := "https://base.ru"

	parsedURL, err := url.Parse(baseURL)
	require.NoError(t, err)

	t.Log(parsedURL.RawQuery)

	client := Client{
		baseURL: parsedURL,
	}

	path := "/test/path"
	queryURL := client.constructQueryURL(path)

	require.Equal(t, baseURL+path, queryURL)
}

func TestName(t *testing.T) {
	var empty interface{}

	fmt.Print(empty == nil)
}
