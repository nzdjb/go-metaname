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
	aux := int(30)
	err := c.ConfigureZone(context.TODO(), "testzone.nz", []ResourceRecord{{
		Name: "george",
		Type: "CNAME",
		Aux:  nil,
		Ttl:  600,
		Data: "example.org",
	},
		{
			Name: "mail",
			Type: "MX",
			Aux:  &aux,
			Ttl:  600,
			Data: "example.org",
		},
	}, nil)
	if err != nil {
		panic(err)
	}
}

func TestIntegrationDnsZone(t *testing.T) {
	c := NewMetanameClient(testAccountReference, testAccountAPIKey)
	c.Host = "https://test.metaname.net/api/1.1"
	result, err := c.DnsZone(context.TODO(), "testzone.nz")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, []ResourceRecord(nil), result)
}
