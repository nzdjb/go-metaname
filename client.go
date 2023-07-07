// Package metaname provides a client for the Metaname API.
package metaname

import (
	"context"
	"encoding/json"

	"github.com/AdamSLevy/jsonrpc2/v14"
)

type IJsonRpc2Client interface {
	Request(context context.Context, host string, method string, params interface{}, result interface{}) error
}

type MetanameClient struct {
	RpcClient        IJsonRpc2Client
	Host             string
	AccountReference string
	APIKey           string
}

type ResourceRecord struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Aux  *int   `json:"aux"`
	Ttl  int    `json:"ttl"`
	Data string `json:"data"`
}

func NewMetanameClient(accountReference string, apiKey string) *MetanameClient {
	return &MetanameClient{
		RpcClient:        &jsonrpc2.Client{},
		Host:             "https://metaname.net/api/1.1",
		AccountReference: accountReference,
		APIKey:           apiKey,
	}
}

func (c *MetanameClient) CreateDnsRecord(ctx context.Context, domainName string, record ResourceRecord) (string, error) {
	params := []interface{}{c.AccountReference, c.APIKey, domainName, record}
	var result string
	err := c.RpcClient.Request(ctx, c.Host, "create_dns_record", params, result)
	return result, err
}

func (c *MetanameClient) UpdateDnsRecord(ctx context.Context, domainName string, reference string, record ResourceRecord) error {
	params := []interface{}{c.AccountReference, c.APIKey, domainName, reference, record}
	err := c.RpcClient.Request(ctx, c.Host, "update_dns_record", params, nil)
	return ignoreNullResultError(err)
}

func (c *MetanameClient) DeleteDnsRecord(ctx context.Context, domainName string, reference string) error {
	params := []interface{}{c.AccountReference, c.APIKey, domainName, reference}
	err := c.RpcClient.Request(ctx, c.Host, "delete_dns_record", params, nil)
	return ignoreNullResultError(err)
}

func (c *MetanameClient) DnsZone(ctx context.Context, domainName string) ([]ResourceRecord, error) {
	params := []interface{}{c.AccountReference, c.APIKey, domainName}
	var result []ResourceRecord
	err := c.RpcClient.Request(ctx, c.Host, "dns_zone", params, result)
	return result, err
}

func (c *MetanameClient) ConfigureZone(ctx context.Context, zoneName string, records []ResourceRecord, options interface{}) error {
	params := []interface{}{c.AccountReference, c.APIKey, zoneName, records, options}
	err := c.RpcClient.Request(ctx, c.Host, "configure_zone", params, nil)
	return ignoreNullResultError(err)
}

// Workaround until https://github.com/AdamSLevy/jsonrpc2/issues/11 is fixed.
type nullSafeResponse struct {
	Result interface{} `json:"result"`
}

func ignoreNullResultError(err error) error {
	if unexerr, ok := err.(jsonrpc2.ErrorUnexpectedHTTPResponse); ok {
		var res nullSafeResponse
		unmerr := json.Unmarshal(unexerr.Body, &res)
		if unmerr != nil {
			return err
		} else if res.Result == nil {
			return nil
		}
	}
	return err
}
