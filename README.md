# pfsense-api-goclient

Go client library to call the pfsense API: https://github.com/jaredhendrickson13/pfsense-api. 

[![GoDoc](https://godoc.org/github.com/sjafferali/pfsense-api-goclient/v2?status.svg)](https://pkg.go.dev/github.com/sjafferali/pfsense-api-goclient/v2)
[![Go Report Card](https://goreportcard.com/badge/github.com/sjafferali/pfsense-api-goclient/v2)](https://goreportcard.com/report/github.com/sjafferali/pfsense-api-goclient/v2)
[![Unit](https://github.com/sjafferali/pfsense-api-goclient/actions/workflows/unit.yaml/badge.svg?branch=v2)](https://github.com/sjafferali/pfsense-api-goclient/actions?query=branch%3Av2)
[![golangci-lint](https://github.com/sjafferali/pfsense-api-goclient/actions/workflows/golang-ci-lint.yaml/badge.svg?branch=v2)](https://github.com/sjafferali/pfsense-api-goclient/actions?query=branch%3Av2)
[![govulncheck](https://github.com/sjafferali/pfsense-api-goclient/actions/workflows/govulncheck.yaml/badge.svg?branch=v2)](https://github.com/sjafferali/pfsense-api-goclient/actions?query=branch%3Av2)
[![Test Coverage](https://codecov.io/gh/sjafferali/pfsense-api-goclient/branch/v2/graph/badge.svg)](https://codecov.io/gh/sjafferali/pfsense-api-goclient)
[![latest version](https://img.shields.io/github/tag/sjafferali/pfsense-api-goclient.svg)](https://github.com/sjafferali/pfsense-apfsense-api-goclient)

## Usage

### Supported APIs

This client supports the following pfSense API endpoints:

- **Firewall** - Manage firewall rules and aliases
- **WireGuard VPN** - Manage WireGuard tunnels and peers
- **Interface Management** - Create, read, update, delete interfaces, VLANs, bridges, and groups
- **User Management** - Manage users and user groups

### Supported Authentication Methods
- Local Authentication (Username/Password)
- JWT Authentication
- Token Authentication

### Example

```go
package main

import (
	"context"
	"fmt"
	
	"github.com/sjafferali/pfsense-api-goclient/v2/pfsenseapi"
)

func main() {
	ctx := context.Background()
	
	// Create client with local authentication
	client := pfsenseapi.NewClientWithLocalAuth(
		"https://192.168.10.1",
		"admin",
		"adminpassword",
	)

	// List firewall rules
	rules, err := client.Firewall.ListFirewallRules(ctx)
	if err != nil {
		panic(err)
	}
	
	for _, rule := range rules {
		descrValue, _ := rule.Descr.Get()
		fmt.Printf("Rule ID %d: %s\n", rule.Id, descrValue)
	}
}
```

## Testing

### Running Unit Tests

Run all unit tests:

```bash
go test ./pfsenseapi/
```

### Running E2E Tests

E2E tests validate functionality against a real pfSense server and are isolated using build tags.

#### Setup

1. Create a configuration file `e2e_config.json` in the project root:

```json
{
  "url": "https://your-pfsense-host.example.com",
  "username": "your-username",
  "password": "your-password"
}
```

**Note:** This file is gitignored to prevent credential leaks.

2. Alternatively, set a custom config path:

```bash
export E2E_CONFIG_FILE=/path/to/your/e2e_config.json
```

#### Run E2E Tests

Run all E2E tests:

```bash
go test -tags=e2e -v ./pfsenseapi/
```

## Contributing

PRs welcome.
