package configs

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_WhenInitConfigs_ThenOrganizationsAccountAddressIsSet(t *testing.T) {
	t.Parallel()
	// ea is expected address:
	ea := os.Getenv("ORG_ACCOUNT_ADDR")
	if len(ea) == 0 {
		ea = "http://localhost:8080/v1/organisation/accounts"
	}
	assert.Equal(t, OrganizationsAccountAddress, ea)
}
