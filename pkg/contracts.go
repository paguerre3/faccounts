package pkg

type AccountHandler interface {
	// Create organisation account
	Create(req AccountData) (resp *AccountData, err error)
	// Fetch organisation account by accountId
	Fetch(id string) (*AccountData, error)
	// Delete organisation account by accountId and accountVersion
	Delete(id string, version int64) error
}
