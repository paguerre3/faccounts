package pkg

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test_WhenCreateOrganisationAccountWithValidAttributes_ThenResultIsEnrichedWithSuccess core account creation
// success test:
func Test_WhenCreateOrganisationAccountWithValidAttributes_ThenResultIsEnrichedWithSuccess(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	ah := NewAccountHandler()
	id := uuid.New().String()
	n := "Camila Aguerre"
	req, err := mockBasicAccountData(id, n)
	assert.Nil(err)
	assert.NotNil(req)
	resp, err := ah.Create(*req)
	assert.Nil(err)
	assert.NotNil(resp)
	assert.NotNil(resp.Attributes)
	assert.Equal(n, resp.Attributes.Name[0])
	assert.Equal(id, resp.ID)
	assert.NotNil(resp.Version)
}

func Test_WhenCreateOrganisationAccountHavingDuplicatedId_ThenViolationErrorOccurs(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	ah := NewAccountHandler()
	id := uuid.New().String()
	n := "Malena Aguerre"
	req, err := mockBasicAccountData(id, n)
	assert.Nil(err)
	assert.NotNil(req)
	resp, err := ah.Create(*req)
	assert.Nil(err)
	assert.NotNil(resp)
	assert.Equal(n, resp.Attributes.Name[0])
	req, err = mockBasicAccountData(id, n)
	assert.Nil(err)
	assert.NotNil(req)
	_, err = ah.Create(*req)
	assert.NotNil(err)
	assert.Equal("invalid status 409 expected 201", err.Error())
}

func Test_WhenCreateOrganisationAccountInvalidDataOrMissing_ThenBadRequestErrorOccurs(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	ah := NewAccountHandler()
	id := ""
	n := "Aguerre"
	req, err := mockBasicAccountData(id, n)
	assert.Nil(err)
	assert.NotNil(req)
	_, err = ah.Create(*req)
	assert.NotNil(err)
	assert.Equal("invalid status 400 expected 201", err.Error())
}

// Test_WhenFetchingOrganisationAccountHavingExistentIdAfterCreation_ThenResultIsEnrichedWithSuccess core account fetch
// success test:
func Test_WhenFetchingOrganisationAccountHavingExistentIdAfterCreation_ThenResultIsEnrichedWithSuccess(
	t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	ah := NewAccountHandler()
	id := uuid.New().String()
	n := "Lola Aguerre"
	req, err := mockBasicAccountData(id, n)
	_, err = ah.Create(*req)
	assert.Nil(err)
	resp, err := ah.Fetch(id)
	assert.Nil(err)
	assert.NotNil(resp)
	assert.NotNil(resp.Attributes)
	assert.Equal(n, resp.Attributes.Name[0])
	assert.Equal(id, resp.ID)
	assert.NotNil(resp.Version)
}

func Test_WhenFetchingOrganisationAccountHavingUnExistentResourceId_ThenNotFoundErrorOccurs(
	t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	ah := NewAccountHandler()
	_, err := ah.Fetch("ad07e007-0007-0b7b-a0e7-0070ea0cc7da")
	assert.NotNil(err)
	assert.Equal("invalid status 404 expected 200", err.Error())
}
