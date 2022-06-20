package accounts

import (
	"encoding/json"
	"fmt"
	"github.com/leonardosaid/stc/accounts/pkg/stc-sdk/config"
	"github.com/leonardosaid/stc/accounts/pkg/stc-sdk/request"
	"github.com/leonardosaid/stc/accounts/pkg/stc-sdk/session/accounts/payload"
	"net/http"
)

type accountClient struct {
	requests       request.Requests
	accountAddress string
}

type Client interface {
	FindByCPF(accountID string) (payload.AccountResponse, error)
	FindBalanceByID(id string) (payload.BalanceResponse, error)
	UpdateBalanceByID(id string, data payload.UpdateBalanceRequest) error
}

func NewClient(r request.Requests, address string) Client {
	return &accountClient{
		requests:       r,
		accountAddress: address,
	}
}

func (a *accountClient) FindByCPF(cpf string) (payload.AccountResponse, error) {
	uri := fmt.Sprintf("%v%v/%v", a.accountAddress, config.URIAccountsResource, cpf)
	resp, err := a.requests.Request(nil, uri, http.MethodGet, http.StatusOK)
	return handleAccountResponse(err, resp)
}

func (a *accountClient) FindBalanceByID(id string) (payload.BalanceResponse, error) {
	uri := fmt.Sprintf(config.URIGetAccountBalance, id)
	resp, err := a.requests.Request(nil, fmt.Sprintf("%v%v", a.accountAddress, uri), http.MethodGet, http.StatusOK)
	return handleBalanceResponse(err, resp)
}

func (a *accountClient) UpdateBalanceByID(id string, data payload.UpdateBalanceRequest) error {
	uri := fmt.Sprintf("%v%v/%v", a.accountAddress, config.URIAccountsResource, id)
	_, err := a.requests.Request(data, uri, http.MethodPatch, http.StatusOK)
	return err
}

func handleAccountResponse(err error, resp []byte) (payload.AccountResponse, error) {
	if err != nil {
		return payload.AccountResponse{}, err
	}

	var response payload.AccountResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return payload.AccountResponse{}, err
	}

	return response, err
}

func handleBalanceResponse(err error, resp []byte) (payload.BalanceResponse, error) {
	if err != nil {
		return payload.BalanceResponse{}, err
	}

	var response payload.BalanceResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return payload.BalanceResponse{}, err
	}

	return response, err
}
