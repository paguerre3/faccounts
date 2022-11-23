package configs

import (
	"f3.com/accounts/internal"
	"os"
)

const (
	defaultHostAddress            = "http://localhost:8080"
	organisationsAccountsResource = "v1/organisation/accounts"

	// ApplicationJson global exposed constant:
	ApplicationJson = "application/json"
)

// OrganizationsAccountAddress global exposed variable that might by overwritten if applicable:
var OrganizationsAccountAddress = internal.ResolveAddress(defaultHostAddress,
	organisationsAccountsResource)

func init() {
	// support overwrite of address from docker compose in case needed, e.g.
	// v2 version support or public address access:
	h := os.Getenv("ORG_ACCOUNT_ADDR")
	if len(h) > 0 {
		OrganizationsAccountAddress = h
	}
}
