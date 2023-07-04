// Package metaname provides a client for the Metaname API.
package metaname

import (
	"context"
	"github.com/AdamSLevy/jsonrpc2/v14"
)

type IClient interface {
	Request(context context.Context, host string, method string, params interface{}, result interface{}) error
}

type Client struct {
	Client           IClient
	Host             string
	AccountReference string
	APIKey           string
}

type ResourceRecord struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Aux  int    `json:"aux,omitempty"`
	Ttl  int    `json:"ttl"`
	Data string `json:"data"`
}

func NewClient(accountReference string, apiKey string) *Client {
	return &Client{
		Client:           &jsonrpc2.Client{},
		Host:             "https://metaname.net/api/1.1",
		AccountReference: accountReference,
		APIKey:           apiKey,
	}
}

func (c *Client) CreateDnsRecord(ctx context.Context, domainName string, record ResourceRecord) (string, error) {
	params := []interface{}{c.AccountReference, c.APIKey, domainName, record}
	var result string
	err := c.Client.Request(ctx, c.Host, "create_dns_record", params, result)
	return result, err
}

func (c *Client) UpdateDnsRecord(ctx context.Context, domainName string, reference string, record ResourceRecord) error {
	params := []interface{}{c.AccountReference, c.APIKey, domainName, reference, record}
	err := c.Client.Request(ctx, c.Host, "update_dns_record", params, nil)
	return err
}

func (c *Client) DeleteDnsRecord(ctx context.Context, domainName string, reference string) error {
	params := []interface{}{c.AccountReference, c.APIKey, domainName, reference}
	err := c.Client.Request(ctx, c.Host, "delete_dns_record", params, nil)
	return err
}

func (c *Client) DnsZone(ctx context.Context, domainName string) ([]ResourceRecord, error) {
	params := []interface{}{c.AccountReference, c.APIKey, domainName}
	var result []ResourceRecord
	err := c.Client.Request(ctx, c.Host, "dns_zone", params, result)
	return result, err
}

func (c *Client) ConfigureZone(ctx context.Context, zoneName string) error {
	params := []interface{}{c.AccountReference, c.APIKey, zoneName, []ResourceRecord{}, nil}
	err := c.Client.Request(ctx, c.Host, "configure_zone", params, nil)
	return err
}
