package metaname

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

type FauxClient struct {
	requestCalled bool
	hostUsed      string
	methodUsed    string
	paramsUsed    interface{}
}

func (fc *FauxClient) Request(context context.Context, host string, method string, params interface{}, result interface{}) error {
	fc.requestCalled = true
	fc.hostUsed = host
	fc.methodUsed = method
	fc.paramsUsed = params
	return nil
}

func TestNewClient(t *testing.T) {
	c := NewMetanameClient("username", "apikey")
	assert.Equal(t, "https://metaname.net/api/1.1", c.Host)
	assert.Equal(t, "username", c.AccountReference)
	assert.Equal(t, "apikey", c.APIKey)
}

func TestDnsZone(t *testing.T) {
	c := &MetanameClient{
		RpcClient:        &FauxClient{},
		Host:             "abc",
		AccountReference: "def",
		APIKey:           "ghi",
	}
	_, err := c.DnsZone(context.TODO(), "testzone.nz")
	if err != nil {
		panic(err)
	}
	assert.True(t, c.RpcClient.(*FauxClient).requestCalled)
	assert.Equal(t, "abc", c.RpcClient.(*FauxClient).hostUsed)
	assert.Equal(t, "dns_zone", c.RpcClient.(*FauxClient).methodUsed)
	assert.Equal(t, []interface{}{"def", "ghi", "testzone.nz"}, c.RpcClient.(*FauxClient).paramsUsed)
}
func TestConfigureZone(t *testing.T) {
	c := &MetanameClient{
		RpcClient:        &FauxClient{},
		Host:             "abc",
		AccountReference: "def",
		APIKey:           "ghi",
	}
	err := c.ConfigureZone(context.TODO(), "testzone.nz", []ResourceRecord{}, nil)
	if err != nil {
		panic(err)
	}
	assert.True(t, c.RpcClient.(*FauxClient).requestCalled)
	assert.Equal(t, "abc", c.RpcClient.(*FauxClient).hostUsed)
	assert.Equal(t, "configure_zone", c.RpcClient.(*FauxClient).methodUsed)
	assert.Equal(t, []interface{}{"def", "ghi", "testzone.nz", []ResourceRecord{}, nil}, c.RpcClient.(*FauxClient).paramsUsed)
}

func TestCreateDnsRecord(t *testing.T) {
	c := &MetanameClient{
		RpcClient:        &FauxClient{},
		Host:             "abc",
		AccountReference: "def",
		APIKey:           "ghi",
	}
	record := ResourceRecord{
		Name: "testrecord",
		Type: "A",
		Aux:  nil,
		Data: "127.0.0.1",
		Ttl:  300,
	}
	_, err := c.CreateDnsRecord(context.TODO(), "testzone.nz", record)
	if err != nil {
		panic(err)
	}
	assert.True(t, c.RpcClient.(*FauxClient).requestCalled)
	assert.Equal(t, "abc", c.RpcClient.(*FauxClient).hostUsed)
	assert.Equal(t, "create_dns_record", c.RpcClient.(*FauxClient).methodUsed)
	assert.Equal(t, []interface{}{"def", "ghi", "testzone.nz", record}, c.RpcClient.(*FauxClient).paramsUsed)
}

func TestCreateMXDnsRecord(t *testing.T) {
	c := &MetanameClient{
		RpcClient:        &FauxClient{},
		Host:             "abc",
		AccountReference: "def",
		APIKey:           "ghi",
	}
	x := int(30)
	record := ResourceRecord{
		Name: "testrecord",
		Type: "MX",
		Aux:  &x,
		Data: "127.0.0.1",
		Ttl:  300,
	}
	_, err := c.CreateDnsRecord(context.TODO(), "testzone.nz", record)
	if err != nil {
		panic(err)
	}
	assert.True(t, c.RpcClient.(*FauxClient).requestCalled)
	assert.Equal(t, "abc", c.RpcClient.(*FauxClient).hostUsed)
	assert.Equal(t, "create_dns_record", c.RpcClient.(*FauxClient).methodUsed)
	assert.Equal(t, []interface{}{"def", "ghi", "testzone.nz", record}, c.RpcClient.(*FauxClient).paramsUsed)
}

func TestUpdateDnsRecord(t *testing.T) {
	c := &MetanameClient{
		RpcClient:        &FauxClient{},
		Host:             "abc",
		AccountReference: "def",
		APIKey:           "ghi",
	}
	record := ResourceRecord{
		Name: "testrecord",
		Type: "A",
		Data: "127.0.0.1",
		Ttl:  300,
	}
	err := c.UpdateDnsRecord(context.TODO(), "testzone.nz", "1234", record)
	if err != nil {
		panic(err)
	}
	assert.True(t, c.RpcClient.(*FauxClient).requestCalled)
	assert.Equal(t, "abc", c.RpcClient.(*FauxClient).hostUsed)
	assert.Equal(t, "update_dns_record", c.RpcClient.(*FauxClient).methodUsed)
	assert.Equal(t, []interface{}{"def", "ghi", "testzone.nz", "1234", record}, c.RpcClient.(*FauxClient).paramsUsed)
}
func TestDeleteDnsRecord(t *testing.T) {
	c := &MetanameClient{
		RpcClient:        &FauxClient{},
		Host:             "abc",
		AccountReference: "def",
		APIKey:           "ghi",
	}
	err := c.DeleteDnsRecord(context.TODO(), "testzone.nz", "1234")
	if err != nil {
		panic(err)
	}
	assert.True(t, c.RpcClient.(*FauxClient).requestCalled)
	assert.Equal(t, "abc", c.RpcClient.(*FauxClient).hostUsed)
	assert.Equal(t, "delete_dns_record", c.RpcClient.(*FauxClient).methodUsed)
	assert.Equal(t, []interface{}{"def", "ghi", "testzone.nz", "1234"}, c.RpcClient.(*FauxClient).paramsUsed)
}
