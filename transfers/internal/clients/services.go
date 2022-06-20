package clients

import (
	"github.com/leonardosaid/stc/accounts/pkg/stc-sdk/request"
	"github.com/leonardosaid/stc/accounts/pkg/stc-sdk/session/accounts"
	"net/http"
	"os"
)

type ServiceClients interface {
	GetAccountServiceClient() accounts.Client
}

type serviceClientsImpl struct {
	AccountServiceClient accounts.Client
}

func NewServiceClientsImpl() ServiceClients {
	return &serviceClientsImpl{
		AccountServiceClient: accounts.NewClient(request.NewRequests(&http.Client{}), os.Getenv("ACCOUNTS_ADDRESS")),
	}
}

func (s *serviceClientsImpl) GetAccountServiceClient() accounts.Client {
	return s.AccountServiceClient
}
