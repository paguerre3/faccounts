package realizations

import (
	"bytes"
	"encoding/json"
	"f3.com/accounts/configs"
	"f3.com/accounts/internal"
	"f3.com/accounts/pckg"
	"fmt"
	"log"
	"net/http"
)

type accountHandlerImpl struct{}

func NewAccountHandler() pckg.AccountHandler {
	return accountHandlerImpl{}
}

func (accountHandlerImpl) Create(req pckg.AccountData) (resp *pckg.AccountData, err error) {
	bs, err := json.Marshal(req)
	if err != nil {
		log.Printf("Error while parsing request account %v", err)
		return nil, err
	}
	br := bytes.NewBuffer(bs)
	httpResp, err := http.Post(configs.OrganizationsAccountAddress,
		configs.ApplicationJson, br)
	return processResponse(httpResp, err, http.StatusCreated)
}

func (accountHandlerImpl) Fetch(id string) (resp *pckg.AccountData, err error) {
	address := internal.ResolveAddress(configs.OrganizationsAccountAddress, id)
	httpResp, err := http.Get(address)
	return processResponse(httpResp, err, http.StatusOK)
}

func (accountHandlerImpl) FetchAll() ([]pckg.AccountData, error) {
	return nil, nil
}

func (accountHandlerImpl) Delete(id string) error {
	return nil
}

func processResponse(httpResp *http.Response, httpErr error, expectedStatus int) (resp *pckg.AccountData, err error) {
	if httpResp != nil && httpResp.Body != nil {
		// 1st of all ensure to close body if applicable:
		defer httpResp.Body.Close()
	}
	if httpErr == nil && httpResp.StatusCode != expectedStatus {
		err = fmt.Errorf("invalid status %b expected %b",
			httpResp.StatusCode, expectedStatus)
	}
	if httpErr != nil {
		log.Printf("Error while processing account action %v", httpErr)
		return nil, httpErr
	}
	if err = json.NewDecoder(httpResp.Body).Decode(resp); err != nil {
		log.Printf("Error while decoding account response %v", err)
		return nil, err
	}
	return resp, err
}
