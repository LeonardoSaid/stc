package usecases

import (
	"context"
	"errors"
	"github.com/leonardosaid/stc/accounts/internal/domain"
	"github.com/leonardosaid/stc/accounts/mocks"
	mocks2 "github.com/leonardosaid/stc/accounts/mocks/pkg"
	"github.com/leonardosaid/stc/accounts/pkg/crypt"
	"github.com/leonardosaid/stc/accounts/pkg/stc-sdk/session/accounts/payload"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"os"
	"testing"
)

func TestLogin(t *testing.T) {
	ctx := context.TODO()

	password, _ := crypt.HashSecret("secret")

	t.Run("login successfully", func(t *testing.T) {
		os.Setenv("JWT_TOKEN_SECRET", "secret")
		accMock := new(mocks.ServiceClients)
		clientMock := new(mocks2.Client)

		accMock.On("GetAccountServiceClient").Return(clientMock)
		clientMock.On("FindByCPF", mock.Anything).Return(payload.AccountResponse{Secret: password}, nil)

		uc := NewLoginUseCaseImpl(accMock)
		_, err := uc.Login(ctx, &domain.LoginCredentials{
			CPF:    "00000000191",
			Secret: "secret",
		})

		assert.NoError(t, err)
		accMock.AssertNumberOfCalls(t, "GetAccountServiceClient", 1)
		clientMock.AssertNumberOfCalls(t, "FindByCPF", 1)
	})

	t.Run("error searching for account", func(t *testing.T) {
		os.Setenv("JWT_TOKEN_SECRET", "secret")
		accMock := new(mocks.ServiceClients)
		clientMock := new(mocks2.Client)

		accMock.On("GetAccountServiceClient").Return(clientMock)
		clientMock.On("FindByCPF", mock.Anything).Return(payload.AccountResponse{}, errors.New(""))

		uc := NewLoginUseCaseImpl(accMock)
		_, err := uc.Login(ctx, &domain.LoginCredentials{
			CPF:    "00000000191",
			Secret: "secret",
		})

		assert.Error(t, err)
		accMock.AssertNumberOfCalls(t, "GetAccountServiceClient", 1)
		clientMock.AssertNumberOfCalls(t, "FindByCPF", 1)
	})
}
