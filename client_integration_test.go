//go:build integration

package metaname

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testAccountReference string = os.Getenv("METANAME_ACCOUNT_REFERENCE")
var testAccountAPIKey string = os.Getenv("METANAME_ACCOUNT_API_KEY")

func TestIntegrationConfigureZone(t *testing.T) {
	c := NewMetanameClient(testAccountReference, testAccountAPIKey)
	c.Host = "https://test.metaname.net/api/1.1"
	err := c.ConfigureZone(context.TODO(), "testzone.nz")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, "testzone.nz", "testzone.nz")
}

func TestIntegrationDnsZone(t *testing.T) {
	c := NewMetanameClient(testAccountReference, testAccountAPIKey)
	c.Host = "https://test.metaname.net/api/1.1"
	_, err := c.DnsZone(context.TODO(), "testzone.nz")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, "testzone.nz", "testzone.nz")
}
