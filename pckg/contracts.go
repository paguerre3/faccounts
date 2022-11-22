package pckg

type AccountHandler interface {
	Create(req AccountData) (resp *AccountData, err error)
	Fetch(id string) (*AccountData, error)
	FetchAll() ([]AccountData, error)
	Delete(id string) error
}
