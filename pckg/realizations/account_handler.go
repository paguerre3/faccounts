package realizations

import "f3.com/accounts/pckg"

type accountHandlerImpl struct{}

func NewAccountHandler() pckg.AccountHandler {
	return accountHandlerImpl{}
}

func (accountHandlerImpl) Create(req pckg.AccountData) (resp *pckg.AccountData, err error) {
	return nil, nil
}

func (accountHandlerImpl) Fetch(id string) (*pckg.AccountData, error) {
	return nil, nil
}

func (accountHandlerImpl) FetchAll() ([]pckg.AccountData, error) {
	return nil, nil
}

func (accountHandlerImpl) Delete(id string) error {
	return nil
}
