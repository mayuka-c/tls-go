# TLS-GO
TLS Practice using Go

## Self Signed certificate

### Usage
```bash
go run cmd/tls/main.go --help
go run cmd/tls/main.go create help
```

### Testing

Note: The CA and server certificate configuration is read from tls.yaml

```bash
# Generate CA certificate
go run cmd/tls/main.go create ca

# Generate Server certificate
go run cmd/tls/main.go create cert --ca-cert="ca.crt" --ca-key="ca.key" --name="go-demo-tls.localtest.me"

# Run the testing server
go run cmd/test-server/main.go

# Test using cURL
curl https://go-demo-tls.localtest.me:8001 --cacert ca.crt
```

## mTLS

### Testing

```bash
# Generate Client certificate
# Make sure the name used below is changed in tls.yaml (in name (under certs) and commonName)
go run cmd/tls/main.go create cert --ca-cert="ca.crt" --ca-key="ca.key" --name="go-demo-client.localtest.me" --cert-out="client.crt" --key-out="client.key"

# Run mTLS Server
go run cmd/mtls-server/main.go

# Testing
curl https://go-demo-tls.localtest.me:9001 --cacert ca.crt --key client.key --cert client.crt

# Run mTLS Client
go run cmd/mtls-client/main.go
```