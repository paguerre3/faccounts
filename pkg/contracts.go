package pkg

type AccountHandler interface {
	Create(req AccountData) (resp *AccountData, err error)
	Fetch(id string) (*AccountData, error)
	FetchAll(pageNumber *uint, link *Link) ([]AccountData, error)
	Delete(id string) error
}
