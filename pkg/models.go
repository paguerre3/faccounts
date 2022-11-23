package pkg

// request generic template for building a "single" request:
type request[T any] struct {
	Data T `json:"data"`
}

// response generic template for building a single response:
type response[T any] struct {
	Data  T      `json:"data"`
	Links *links `json:"links,omitempty"`
}

// responseComposition generic template for building a multiple responses normally paginated:
type responseComposition[T any] struct {
	Data  []T    `json:"data"`
	Links *links `json:"links,omitempty"`
}

type links struct {
	Self  string  `json:"self"`
	First *string `json:"first,omitempty"`
	Last  *string `json:"last,omitempty"`
}

// link enum used by fetch all retrieval:
type link string

const (
	First link = "first"
	Last  link = "last"
)

// AccountData represents an account in the form3 org section.
// See https://api-docs.form3.tech/api.html#organisation-accounts for
// more information about fields.
type AccountData struct {
	Attributes     *AccountAttributes `json:"attributes,omitempty"`
	ID             string             `json:"id,omitempty"`
	OrganisationID string             `json:"organisation_id,omitempty"`
	Type           string             `json:"type,omitempty"`
	Version        *int64             `json:"version,omitempty"`
}

type AccountAttributes struct {
	AccountClassification   *string  `json:"account_classification,omitempty"`
	AccountMatchingOptOut   *bool    `json:"account_matching_opt_out,omitempty"`
	AccountNumber           string   `json:"account_number,omitempty"`
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
}
