# Go Metaname

A Go library implementing the [Metaname](https://metaname.net.nz) API.

## Usage

Create a client with NewClient, passing it your account reference and API key, then use it to call the needed functions.

```go
client := NewClient(os.Getenv("ACCOUNT_REF"), os.Getenv("API_KEY"))
client.DeleteDnsRecord(ctx, "example.org", "1234")
```

### Implemented functions

These functions have been implemented:

* [CreateDnsRecord](https://metaname.net/api/1.1/doc#create_dns_record)
* [UpdateDnsRecord](https://metaname.net/api/1.1/doc#update_dns_record)
* [DeleteDnsRecord](https://metaname.net/api/1.1/doc#delete_dns_record)
* [DnsZone](https://metaname.net/api/1.1/doc#dns_zone)
* [ConfigureZone](https://metaname.net/api/1.1/doc#configure_zone)
