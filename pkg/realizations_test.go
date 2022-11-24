package pkg

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	c  chan string
	vs []int64
)

func init() {
	c = make(chan string)
	vs = make([]int64, 0)
}

// Test_WhenCreateOrganisationAccountWithValidAttributes_ThenResultIsEnrichedWithSuccess core account creation
// success test:
func Test_WhenCreateOrganisationAccountWithValidAttributes_ThenResultIsEnrichedWithSuccess(t *testing.T) {
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
	go addAccountCreation(id, c)
	assert.NotNil(resp.Attributes)
	assert.Equal(n, resp.Attributes.Name[0])
	assert.Equal(id, resp.ID)
	assert.NotNil(resp.Version)
	vs = append(vs, *resp.Version)
}

func Test_WhenCreateOrganisationAccountHavingDuplicatedId_ThenViolationErrorOccurs(t *testing.T) {
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
	go addAccountCreation(id, c)
	assert.Equal(n, resp.Attributes.Name[0])
	vs = append(vs, *resp.Version)
	req, err = mockBasicAccountData(id, n)
	assert.Nil(err)
	assert.NotNil(req)
	_, err = ah.Create(*req)
	assert.NotNil(err)
	assert.Equal("invalid status 409 expected 201", err.Error())
}

func Test_WhenCreateOrganisationAccountInvalidDataOrMissing_ThenBadRequestErrorOccurs(t *testing.T) {
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
	assert := assert.New(t)
	ah := NewAccountHandler()
	id := uuid.New().String()
	n := "Lola Aguerre"
	req, err := mockBasicAccountData(id, n)
	_, err = ah.Create(*req)
	assert.Nil(err)
	go addAccountCreation(id, c)
	resp, err := ah.Fetch(id)
	assert.Nil(err)
	assert.NotNil(resp)
	assert.NotNil(resp.Attributes)
	assert.Equal(n, resp.Attributes.Name[0])
	assert.Equal(id, resp.ID)
	assert.NotNil(resp.Version)
	vs = append(vs, *resp.Version)
}

func Test_WhenFetchingOrganisationAccountHavingUnExistentResourceId_ThenNotFoundErrorOccurs(
	t *testing.T) {
	assert := assert.New(t)
	ah := NewAccountHandler()
	_, err := ah.Fetch("ad07e007-0007-0b7b-a0e7-0070ea0cc7da")
	assert.NotNil(err)
	assert.Equal("invalid status 404 expected 200", err.Error())
}

// Test_WhenDeletingOrganisationAccountHavingExistentIdAfterCreation_ThenResultIsNoContentWithSuccess core account delete
// success test:
func Test_WhenDeletingOrganisationAccountHavingExistentIdAfterCreation_ThenResultIsNoContentWithSuccess(
	t *testing.T) {
	assert := assert.New(t)
	ah := NewAccountHandler()
	id := uuid.New().String()
	n := "Milagros Aguerre"
	req, _ := mockBasicAccountData(id, n)
	assert.NotNil(req)
	resp, _ := ah.Create(*req)
	assert.NotNil(resp)
	assert.NotNil(resp.Version)
	err := ah.Delete(resp.ID, *resp.Version)
	assert.Nil(err)
}

func Test_WhenDeletingOrganisationAccountHavingUnExistentResourceId_ThenNotFoundErrorOccurs(
	t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	ah := NewAccountHandler()
	err := ah.Delete("ad07e007-0007-0b7b-a0e7-0070ea0cc7da", 0)
	assert.NotNil(err)
	assert.Equal("invalid status 404 expected 204", err.Error())
}

func Test_WhenDeletingOrganisationAccountHavingExistentIdWithInvalidVersionAfterCreation_ThenConflictErrorOccurs(
	t *testing.T) {
	assert := assert.New(t)
	ah := NewAccountHandler()
	id := uuid.New().String()
	n := "Sol Aguerre"
	req, _ := mockBasicAccountData(id, n)
	assert.NotNil(req)
	resp, _ := ah.Create(*req)
	assert.NotNil(resp)
	go addAccountCreation(id, c)
	assert.NotNil(resp)
	assert.NotNil(resp.Version)
	vs = append(vs, *resp.Version)
	err := ah.Delete(resp.ID, *resp.Version+1)
	assert.NotNil(err)
	assert.Equal("invalid status 409 expected 204", err.Error())
}

// Test_WhenDeletingAllPreviouslyCreatedAccounts_ThenResultIsNoContentWithSuccess clean up:
func Test_WhenDeletingAllPreviouslyCreatedAccounts_ThenResultIsNoContentWithSuccess(t *testing.T) {
	time.Sleep(time.Second * 1)
	if len(vs) > 0 {
		ah := NewAccountHandler()
		for _, v := range vs {
			// <-c is a blocking wait:
			go func(accountId string, ver int64, tt *testing.T) {
				err := ah.Delete(accountId, ver)
				assert.Nil(tt, err)
			}(<-c, v, t)
		}
	}
}

func addAccountCreation(id string, c chan string) {
	c <- id
}
