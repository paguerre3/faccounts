package configs

import "testing"

func Test_WhenInitWithoutHavingHostAddress_ThenOrganizationsAccountsUrlStartsWithLocalhost(t *testing.T) {
	t.Parallel()
	if OrganizationsAccountAddress != "http://localhost:8080/v1/organisation/accounts" {
		t.Fail()
	}
}
