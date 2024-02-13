# pfsense-api-goclient

Go client library to call the pfsense API: https://github.com/jaredhendrickson13/pfsense-api. 

[![GoDoc](https://godoc.org/github.com/sjafferali/pfsense-api-goclient/v2?status.svg)](https://pkg.go.dev/github.com/sjafferali/pfsense-api-goclient/v2)
[![Go Report Card](https://goreportcard.com/badge/github.com/sjafferali/pfsense-api-goclient/v2)](https://goreportcard.com/report/github.com/sjafferali/pfsense-api-goclient/v2)
[![Unit](https://github.com/sjafferali/pfsense-api-goclient/actions/workflows/unit.yaml/badge.svg?branch=v2)](https://github.com/sjafferali/pfsense-api-goclient/actions?query=branch%3Av2)
[![golangci-lint](https://github.com/sjafferali/pfsense-api-goclient/actions/workflows/golang-ci-lint.yaml/badge.svg?branch=v2)](https://github.com/sjafferali/pfsense-api-goclient/actions?query=branch%3Av2)
[![govulncheck](https://github.com/sjafferali/pfsense-api-goclient/actions/workflows/govulncheck.yaml/badge.svg?branch=v2)](https://github.com/sjafferali/pfsense-api-goclient/actions?query=branch%3Av2)
[![Test Coverage](https://codecov.io/gh/sjafferali/pfsense-api-goclient/branch/main/graph/badge.svg)](https://codecov.io/gh/sjafferali/pfsense-api-goclient)
[![latest version](https://img.shields.io/github/tag/sjafferali/pfsense-api-goclient.svg)](https://github.com/sjafferali/pfsense-apfsense-api-goclient)

## Usage

### Supported Authentication Methods
- Local Authentication (Username/Password)
- JWT Authentication
- Token Authentication

### Example (Local Authentication)
```go
package main

import (
	"context"
	"fmt"

	"github.com/sjafferali/pfsense-api-goclient/pfsenseapi"
)

func main() {
	ctx := context.Background()
	client := pfsenseapi.NewClientWithLocalAuth(
		"https://192.168.10.1",
		"admin",
		"adminpassword",
	)

	leases, err := client.DHCP.ListLeases(ctx)
	if err != nil {
		panic(err)
	}

	for _, lease := range leases {
		fmt.Println(lease.Ip)
	}
}
```

## Contributing

PRs welcome.
