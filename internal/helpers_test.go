package internal

import (
	"testing"
)

func Test_WhenResolveAddressHavingValidAttributes_ThenResultIsUrl(t *testing.T) {
	t.Parallel()
	a := ResolveAddress("http://localhost:8080", "v1/organisation/accounts")
	if a != "http://localhost:8080/v1/organisation/accounts" {
		t.Fail()
	}
	a = ResolveAddress("http://localhost:8080/v1/organisation/accounts", "007")
	if a != "http://localhost:8080/v1/organisation/accounts/007" {
		t.Fail()
	}
}

func Test_WhenResolveAddressHavingInvalidAttributes_ThenResultIsIncorrect(t *testing.T) {
	t.Parallel()
	a := ResolveAddress("", "")
	if a != "/" {
		t.Fail()
	}
}
