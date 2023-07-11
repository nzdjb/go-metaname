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
	assert.NoError(t, err)
}

func TestIntegrationDnsZone(t *testing.T) {
	c := NewMetanameClient(testAccountReference, testAccountAPIKey)
	c.Host = "https://test.metaname.net/api/1.1"
	result, err := c.DnsZone(context.TODO(), "testzone.nz")
	assert.NoError(t, err)
	noRefResult := []ResourceRecord{}
	for _, r := range result {
		assert.NotNil(t, r.Reference)
		assert.NotEqual(t, "", *r.Reference)
		r.Reference = nil
		noRefResult = append(noRefResult, r)
	}
	aux := int(30)
	assert.ElementsMatch(t, []ResourceRecord{{
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
	}, noRefResult)
}

func TestIntegrationCreateDnsRecord(t *testing.T) {
	c := NewMetanameClient(testAccountReference, testAccountAPIKey)
	c.Host = "https://test.metaname.net/api/1.1"
	ref, err := c.CreateDnsRecord(context.TODO(), "testzone.nz", ResourceRecord{
		Name: "bill",
		Type: "A",
		Aux:  nil,
		Ttl:  600,
		Data: "127.0.0.1",
	})
	assert.NoError(t, err)
	assert.NotEqual(t, "", ref)

	result, err := c.DnsZone(context.TODO(), "testzone.nz")
	noRefResult := []ResourceRecord{}
	for _, r := range result {
		assert.NotNil(t, r.Reference)
		assert.NotEqual(t, "", *r.Reference)
		r.Reference = nil
		noRefResult = append(noRefResult, r)
	}
	aux := int(30)
	assert.ElementsMatch(t, []ResourceRecord{{
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
		{
			Name: "bill",
			Type: "A",
			Aux:  nil,
			Ttl:  600,
			Data: "127.0.0.1",
		},
	}, noRefResult)
}

func TestIntegrationUpdateDnsRecord(t *testing.T) {
	c := NewMetanameClient(testAccountReference, testAccountAPIKey)
	c.Host = "https://test.metaname.net/api/1.1"
	result, err := c.DnsZone(context.TODO(), "testzone.nz")
	if err != nil {
		panic(err)
	}
	for _, r := range result {
		ref := *r.Reference
		r.Reference = nil
		r.Ttl = 300
		err := c.UpdateDnsRecord(context.TODO(), "testzone.nz", ref, r)
		assert.NoError(t, err)
	}

	result, err = c.DnsZone(context.TODO(), "testzone.nz")
	noRefResult := []ResourceRecord{}
	for _, r := range result {
		assert.NotNil(t, r.Reference)
		assert.NotEqual(t, "", *r.Reference)
		r.Reference = nil
		noRefResult = append(noRefResult, r)
	}
	aux := int(30)
	assert.ElementsMatch(t, []ResourceRecord{{
		Name: "george",
		Type: "CNAME",
		Aux:  nil,
		Ttl:  300,
		Data: "example.org",
	},
		{
			Name: "mail",
			Type: "MX",
			Aux:  &aux,
			Ttl:  300,
			Data: "example.org",
		},
		{
			Name: "bill",
			Type: "A",
			Aux:  nil,
			Ttl:  300,
			Data: "127.0.0.1",
		},
	}, noRefResult)
}

func TestIntegrationDeleteDnsRecord(t *testing.T) {
	c := NewMetanameClient(testAccountReference, testAccountAPIKey)
	c.Host = "https://test.metaname.net/api/1.1"
	result, err := c.DnsZone(context.TODO(), "testzone.nz")
	if err != nil {
		panic(err)
	}
	for _, r := range result {
		err = c.DeleteDnsRecord(context.TODO(), "testzone.nz", *r.Reference)
		assert.NoError(t, err)
	}
	result, err = c.DnsZone(context.TODO(), "testzone.nz")
	if err != nil {
		panic(err)
	}
	assert.Empty(t, result)
}
