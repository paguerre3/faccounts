package pkg

import (
	"bytes"
	"encoding/json"
	"f3.com/accounts/configs"
	"f3.com/accounts/internal"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type accountHandlerImpl struct {
	client *http.Client
}

func NewAccountHandler() AccountHandler {
	return accountHandlerImpl{
		client: &http.Client{},
	}
}

func (accountHandlerImpl) Create(req AccountData) (resp *AccountData, err error) {
	r := request[AccountData]{
		Data: req,
	}
	bs, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	br := bytes.NewBuffer(bs)
	// add retry strategy for waiting until account api is healthy:
	retry := configs.HttpMaxRetries
	var httpResp *http.Response
	for retry > 0 {
		httpResp, err = http.Post(configs.OrganizationsAccountAddress,
			configs.ApplicationJson, br)
		if err != nil {
			retry--
			time.Sleep(time.Second * 1)
		} else {
			break
		}
	}
	return processResponse(httpResp, err, http.StatusCreated)
}

func (accountHandlerImpl) Fetch(id string) (*AccountData, error) {
	address := internal.ResolveAddress(configs.OrganizationsAccountAddress, id)
	retry := configs.HttpMaxRetries
	var httpResp *http.Response
	var err error
	// add retry strategy for waiting until account api is healthy:
	for retry > 0 {
		httpResp, err = http.Get(address)
		if err != nil {
			retry--
			time.Sleep(time.Second * 1)
		} else {
			break
		}
	}
	return processResponse(httpResp, err, http.StatusOK)
}

func (ah accountHandlerImpl) Delete(id string, version int64) error {
	address := internal.ResolveAddress(configs.OrganizationsAccountAddress, id)
	req, err := http.NewRequest("DELETE", address, nil)
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("version", strconv.FormatInt(version, 10))
	req.URL.RawQuery = q.Encode()
	// add retry strategy for waiting until account api is healthy:
	retry := configs.HttpMaxRetries
	var httpResp *http.Response
	for retry > 0 {
		httpResp, err = ah.client.Do(req)
		if err != nil {
			retry--
			time.Sleep(time.Second * 1)
		} else {
			break
		}
	}
	if httpResp != nil {
		defer httpResp.Body.Close()
	}
	return checkStatusError(httpResp, err, http.StatusNoContent)
}

func processResponse(httpResp *http.Response, httpErr error, expectedStatus int) (
	resp *AccountData, err error) {
	if httpResp != nil {
		// 1st of all ensure to close body if applicable:
		defer httpResp.Body.Close()
	}
	if err = checkStatusError(httpResp, httpErr, expectedStatus); err != nil {
		return nil, err
	}
	r := response[AccountData]{}
	if err = json.NewDecoder(httpResp.Body).Decode(&r); err != nil {
		return nil, err
	}
	resp = &r.Data
	return resp, err
}

func checkStatusError(httpResp *http.Response, httpErr error, expectedStatus int) (err error) {
	if httpErr != nil {
		err = httpErr
	}
	if err == nil && httpResp.StatusCode != expectedStatus {
		err = fmt.Errorf("invalid status %v expected %v",
			httpResp.StatusCode, expectedStatus)
	}
	return err
}
