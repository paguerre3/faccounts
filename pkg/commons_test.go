package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MockingBasicAccountData_ThenMockIsEnriched(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	id := "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"
	oid := "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c"
	n := "Camila Aguerre"
	an := n[:3]
	ad, err := mockBasicAccountData(id, n)
	assert.Nil(err)
	assert.NotNil(ad)
	assert.Equal(n, ad.Attributes.Name[0])
	assert.Equal(an, ad.Attributes.AlternativeNames[0])
	assert.Equal(id, ad.ID)
	assert.Equal(oid, ad.OrganisationID)
}
