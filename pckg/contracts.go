package pckg

type AccountHandler interface {
	Create(req Request[AccountAttributes]) (resp *Response[AccountData], err error)
	Fetch(id string) (*Response[AccountData], error)
	FetchAll(pageNumber *uint) (*ResponseComposition[AccountData], error)
	Delete(id string) error
}
