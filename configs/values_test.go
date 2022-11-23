package configs

import (
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
	if OrganizationsAccountAddress != ea {
		// fail is included inside error:
		t.Errorf("unexpected account address %s", OrganizationsAccountAddress)
		t.Fail()
	}
}
