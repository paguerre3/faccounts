package realizations

import (
	"bytes"
	"encoding/json"
	"f3.com/accounts/configs"
	"f3.com/accounts/internal"
	"f3.com/accounts/pckg"
	"fmt"
	"net/http"
)

type accountHandlerImpl struct{}

func NewAccountHandler() pckg.AccountHandler {
	return accountHandlerImpl{}
}

func (accountHandlerImpl) Create(req pckg.Request[pckg.AccountAttributes]) (
	resp *pckg.Response[pckg.AccountData], err error) {
	bs, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	br := bytes.NewBuffer(bs)
	httpResp, err := http.Post(configs.OrganizationsAccountAddress,
		configs.ApplicationJson, br)
	return processSingleResponse(httpResp, err, http.StatusCreated)
}

func (accountHandlerImpl) Fetch(id string) (resp *pckg.Response[pckg.AccountData], err error) {
	address := internal.ResolveAddress(configs.OrganizationsAccountAddress, id)
	httpResp, err := http.Get(address)
	return processSingleResponse(httpResp, err, http.StatusOK)
}

func (accountHandlerImpl) FetchAll(pageNumber *uint) (*pckg.ResponseComposition[pckg.AccountData], error) {
	return nil, nil
}

func (accountHandlerImpl) Delete(id string) error {
	return nil
}

func processSingleResponse(httpResp *http.Response, httpErr error, expectedStatus int) (
	resp *pckg.Response[pckg.AccountData], err error) {
	if httpResp != nil && httpResp.Body != nil {
		// 1st of all ensure to close body if applicable:
		defer httpResp.Body.Close()
	}
	if httpErr != nil {
		err = httpErr
	}
	if err == nil && httpResp.StatusCode != expectedStatus {
		err = fmt.Errorf("invalid status %b expected %b",
			httpResp.StatusCode, expectedStatus)
	}
	if err != nil {
		return nil, err
	}
	if err = json.NewDecoder(httpResp.Body).Decode(resp); err != nil {
		return nil, err
	}
	return resp, err
}
