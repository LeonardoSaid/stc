package usecases

import (
	"context"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/leonardosaid/stc/accounts/internal/domain"
	"github.com/leonardosaid/stc/accounts/mocks"
	mocks2 "github.com/leonardosaid/stc/accounts/mocks/pkg"
	"github.com/leonardosaid/stc/accounts/pkg/stc-sdk/session/accounts/payload"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTransferUseCase(t *testing.T) {
	ctx := context.TODO()

	t.Run("create transfer successfully", func(t *testing.T) {
		clientsMock := new(mocks.ServiceClients)
		accMock := new(mocks2.Client)
		repoMock := new(mocks.TransferRepository)

		transfer := &domain.Transfer{
			AccountDestinationID: uuid.New(),
			Amount:               10,
		}

		token := &jwt.Token{
			Claims: &domain.LoginToken{ID: uuid.New().String()},
		}

		clientsMock.On("GetAccountServiceClient").Return(accMock)
		accMock.On("FindBalanceByID", mock.Anything).Return(payload.BalanceResponse{Balance: 1000}, nil)
		accMock.On("UpdateBalanceByID", mock.Anything, mock.Anything).Return(nil)
		repoMock.On("Create", mock.Anything, mock.Anything).Return(nil)

		uc := NewTransferUseCaseImpl(clientsMock, repoMock)
		err := uc.Create(ctx, transfer, token)

		assert.NoError(t, err)
		clientsMock.AssertNumberOfCalls(t, "GetAccountServiceClient", 4)
		accMock.AssertNumberOfCalls(t, "FindBalanceByID", 2)
		accMock.AssertNumberOfCalls(t, "UpdateBalanceByID", 2)
		repoMock.AssertNumberOfCalls(t, "Create", 1)
	})

	t.Run("create transfer with validation error", func(t *testing.T) {
		clientsMock := new(mocks.ServiceClients)
		accMock := new(mocks2.Client)
		repoMock := new(mocks.TransferRepository)

		transfer := &domain.Transfer{
			AccountDestinationID: uuid.New(),
			Amount:               -1,
		}

		token := &jwt.Token{
			Claims: &domain.LoginToken{ID: uuid.New().String()},
		}

		uc := NewTransferUseCaseImpl(clientsMock, repoMock)
		err := uc.Create(ctx, transfer, token)

		assert.Error(t, err)
		clientsMock.AssertNumberOfCalls(t, "GetAccountServiceClient", 0)
		accMock.AssertNumberOfCalls(t, "FindBalanceByID", 0)
		accMock.AssertNumberOfCalls(t, "UpdateBalanceByID", 0)
		repoMock.AssertNumberOfCalls(t, "Create", 0)
	})

	t.Run("create transfer with account service error", func(t *testing.T) {
		clientsMock := new(mocks.ServiceClients)
		accMock := new(mocks2.Client)
		repoMock := new(mocks.TransferRepository)

		transfer := &domain.Transfer{
			AccountDestinationID: uuid.New(),
			Amount:               10,
		}

		token := &jwt.Token{
			Claims: &domain.LoginToken{ID: uuid.New().String()},
		}

		clientsMock.On("GetAccountServiceClient").Return(accMock)
		accMock.On("FindBalanceByID", mock.Anything).Return(payload.BalanceResponse{}, &domain.AccountServiceError{})

		uc := NewTransferUseCaseImpl(clientsMock, repoMock)
		err := uc.Create(ctx, transfer, token)

		assert.Error(t, err)
		clientsMock.AssertNumberOfCalls(t, "GetAccountServiceClient", 1)
		accMock.AssertNumberOfCalls(t, "FindBalanceByID", 1)
		accMock.AssertNumberOfCalls(t, "UpdateBalanceByID", 0)
		repoMock.AssertNumberOfCalls(t, "Create", 0)
	})
}
