package usecases

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/leonardosaid/stc/accounts/internal/domain"
	"github.com/leonardosaid/stc/accounts/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAccountUseCase(t *testing.T) {
	ctx := context.TODO()

	t.Run("create account successfully", func(t *testing.T) {
		os.Setenv("JWT_TOKEN_SECRET", "secret")
		repoMock := new(mocks.AccountRepository)

		repoMock.On("Create", mock.Anything, mock.Anything).Return(nil)

		data := &domain.Account{
			CPF:    "00000000191",
			Secret: "secret",
		}

		uc := NewAccountUseCaseImpl(repoMock)
		err := uc.Create(ctx, data)

		assert.NoError(t, err)
		repoMock.AssertNumberOfCalls(t, "Create", 1)
	})

	t.Run("create account with validation error", func(t *testing.T) {
		os.Setenv("JWT_TOKEN_SECRET", "secret")
		repoMock := new(mocks.AccountRepository)

		data := &domain.Account{
			CPF:    "invalid",
			Secret: "secret",
		}

		uc := NewAccountUseCaseImpl(repoMock)
		err := uc.Create(ctx, data)

		assert.Error(t, err)
		repoMock.AssertNumberOfCalls(t, "Create", 0)
	})

	t.Run("find balance by id with not found error", func(t *testing.T) {
		os.Setenv("JWT_TOKEN_SECRET", "secret")
		repoMock := new(mocks.AccountRepository)
		accID, _ := uuid.NewUUID()

		repoMock.On("FindByID", mock.Anything, mock.Anything).Return(nil, errors.New(""))

		uc := NewAccountUseCaseImpl(repoMock)
		_, err := uc.FindBalanceByID(ctx, accID)

		assert.Error(t, err)
		repoMock.AssertNumberOfCalls(t, "FindByID", 1)
	})
}
