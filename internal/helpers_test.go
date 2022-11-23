package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_WhenResolveAddressHavingValidAttributes_ThenResultIsUrl(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	a := ResolveAddress("http://localhost:8080", "v1/organisation/accounts")
	assert.Equal("http://localhost:8080/v1/organisation/accounts", a)
	a = ResolveAddress("http://localhost:8080/v1/organisation/accounts", "007")
	assert.Equal("http://localhost:8080/v1/organisation/accounts/007", a)
}

func Test_WhenResolveAddressHavingInvalidAttributes_ThenResultIsIncorrect(t *testing.T) {
	t.Parallel()
	a := ResolveAddress("", "")
	assert.Equal(t, "/", a)
}
