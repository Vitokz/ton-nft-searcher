package tonclient

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tonkeeper/tongo"
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
	addr, err := tongo.ParseAccountID("0:89043793a728830e118536426b60efc820a19b602276d36677cbd629cef9f500")
	require.NoError(t, err)

	// bz, err := hex.DecodeString("89043793a728830e118536426b60efc820a19b602276d36677cbd629cef9f500")
	// require.NoError(t, err)
	// addr := address.NewAddressExt(0, uint(0), bz)

	fmt.Print(addr.ToHuman(true, false))
}
