package configs

import (
	"f3.com/accounts/internal"
	"os"
)

const (
	defaultHostAddress            = "http://localhost:8080"
	organisationsAccountsResource = "v1/organisation/accounts"
	ApplicationJson               = "application/json"
)

var OrganizationsAccountAddress = internal.ResolveAddress(defaultHostAddress,
	organisationsAccountsResource)

func init() {
	h := os.Getenv("HOST_ADDR")
	if len(h) > 0 {
		OrganizationsAccountAddress = internal.ResolveAddress(h,
			organisationsAccountsResource)
	}
}
