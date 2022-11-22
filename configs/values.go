package configs

import (
	"fmt"
	"os"
)

const (
	defaultHost              = "http://localhost:8080"
	organisationsAccountsUri = "/v1/organisation/accounts"
)

var OrganizationsAccountAddress = resolveAddress(defaultHost, organisationsAccountsUri)

func resolveAddress(host string, uri string) string {
	return fmt.Sprintf("%s%s", host, uri)
}

func init() {
	h := os.Getenv("F3_HOST")
	if len(h) > 0 {
		OrganizationsAccountAddress = resolveAddress(h, organisationsAccountsUri)
	}
}
