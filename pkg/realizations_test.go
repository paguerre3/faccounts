package pkg

import (
	"testing"
)

func Test_WhenCreateOrganizationAccountWithValidAttributes_ThenResultIsEnriched(t *testing.T) {
	t.Parallel()
	ah := NewAccountHandler()
	id := "ad27e265-9605-4b4b-a0e5-3003ea9cc4db"
	n := "Camila Aguerre"
	req, err := mockBasicAccountData(id, n)
	if err != nil {
		t.Fail()
	}
	resp, err := ah.Create(*req)
	if resp == nil || err != nil {
		t.Errorf("unexpected create response %v error %v", resp, err)
		t.Fail()
	}
	if n != resp.Attributes.Name[0] {
		t.Errorf("unexpected name %s", n)
		t.Fail()
	}
	if id != resp.ID {
		t.Errorf("unexpected id %s", id)
		t.Fail()
	}
	if resp.Version == nil {
		t.Errorf("invalid version")
		t.Fail()
	}
}
