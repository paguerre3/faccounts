package pkg

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test_WhenCreatingOrganisationAccountHavingValidAttributes_ThenResultIsEnrichedWithSuccess core account creation
// success test:
func Test_WhenCreatingOrganisationAccountHavingValidAttributes_ThenResultIsEnrichedWithSuccess(t *testing.T) {
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

func Test_WhenCreatingOrganisationAccountHavingDuplicatedId_ThenViolationErrorOccurs(t *testing.T) {
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

func Test_WhenCreatingOrganisationAccountHavingInvalidDataOrMissing_ThenBadRequestErrorOccurs(t *testing.T) {
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
	assert.NotNil(resp)
	assert.NotNil(resp.Version)
	err := ah.Delete(resp.ID, *resp.Version+1)
	assert.NotNil(err)
	assert.Equal("invalid status 409 expected 204", err.Error())
}

// Test_WhenDeletingCreatedAccountsInParallel_ThenResultIsNoContentWithSuccess create and delete in parallel.
func Test_WhenDeletingCreatedAccountsInParallel_ThenResultIsNoContentWithSuccess(t *testing.T) {
	ids := make(chan string)
	defer close(ids)
	laps := 77
	ah := NewAccountHandler()
	// no need of using wg as <-c during consumption is a blocking wait:
	for i := 0; i <= laps; i++ {
		// wg add delta
		// anonymous func:
		go func(ahh *AccountHandler, cc chan string, tt *testing.T) {
			// wg defer done
			cc <- addAccountCreation(ahh, tt)
		}(&ah, ids, t)
	}
	// wg wait
	for j := 0; j <= laps; j++ {
		//  <-c is a blocking wait, i.e. the same as doing "for" over the chan:
		go addAccountDeletion(<-ids, &ah, t)
	}
}

func addAccountCreation(h *AccountHandler, tt *testing.T) string {
	id := uuid.New().String()
	n := "Camila Milagros Aguerre"
	req, _ := mockBasicAccountData(id, n)
	resp, _ := (*h).Create(*req)
	if resp == nil {
		tt.Errorf("response is nil for request %v", *req)
	}
	return id
}

func addAccountDeletion(id string, h *AccountHandler, tt *testing.T) {
	// recently created accounts have 0 version always:
	err := (*h).Delete(id, 0)
	if err != nil {
		tt.Errorf("delete of id %s failed with error %v", id, err)
	}
}
