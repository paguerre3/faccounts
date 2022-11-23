package realizations

import (
	"f3.com/accounts/pckg"
	"testing"
)

func Test_WhenCreateOrganizationAccountWithValidAttributes_ThenResultIsEnriched(t *testing.T) {
	t.Parallel()
	ah := NewAccountHandler()
	ac := "Personal"
	r := pckg.Request[pckg.AccountAttributes]{
		Data: pckg.AccountAttributes{
			AccountClassification: &ac,
			AccountMatchingOptOut: nil,
			AccountNumber: "",
			AlternativeNames        []string `json:"alternative_names,omitempty"`
			BankID                  string   `json:"bank_id,omitempty"`
			BankIDCode              string   `json:"bank_id_code,omitempty"`
			BaseCurrency            string   `json:"base_currency,omitempty"`
			Bic                     string   `json:"bic,omitempty"`
			Country                 *string  `json:"country,omitempty"`
			Iban                    string   `json:"iban,omitempty"`
			JointAccount            *bool    `json:"joint_account,omitempty"`
			Name                    []string `json:"name,omitempty"`
			SecondaryIdentification string   `json:"secondary_identification,omitempty"`
			Status                  *string  `json:"status,omitempty"`
			Switched                *bool    `json:"switched,omitempty"`
		},
	}
	resp, err := ah.Create(r)
	if resp == nil || err != nil {
		t.Fail()
	}
}
