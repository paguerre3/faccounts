package pkg

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_WhenCreateOrganizationAccountWithValidAttributes_ThenResultIsEnriched(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	ah := NewAccountHandler()
	id := uuid.New().String()
	n := "Camila Aguerre"
	req, err := mockBasicAccountData(id, n)
	if err != nil {
		t.Fail()
	}
	resp, err := ah.Create(*req)
	assert.Nil(err)
	assert.NotNil(resp)
	assert.NotNil(resp.Attributes)
	assert.Equal(n, resp.Attributes.Name[0])
	assert.Equal(id, resp.ID)
	assert.NotNil(resp.Version)
}
