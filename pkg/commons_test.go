package pkg

import (
	"testing"
)

func Test_MockingBasicAccountData_ThenMockIsEnriched(t *testing.T) {
	t.Parallel()
	id := "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"
	oid := "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c"
	n := "Camila Aguerre"
	an := n[:3]
	ad, err := mockBasicAccountData(id, n)
	if ad == nil || err != nil {
		t.Errorf("unexpected response accountData %v error %v", ad, err)
		t.Fail()
	}
	if n != ad.Attributes.Name[0] && an != ad.Attributes.AlternativeNames[0] {
		t.Errorf("unexpected names main %s alternative %s", n, an)
		t.Fail()
	}
	if id != ad.ID && oid != ad.OrganisationID {
		t.Errorf("unexpected ids account %s organisationId %s", id, oid)
		t.Fail()
	}
}
