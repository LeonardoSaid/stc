package usecases

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/leonardosaid/stc/accounts/internal/domain"
	"github.com/leonardosaid/stc/accounts/internal/infra/repositories"
	"github.com/leonardosaid/stc/accounts/pkg/crypt"
)

type AccountUseCase interface {
	List(context.Context) ([]domain.Account, error)
	Create(context.Context, *domain.Account) error
	FindBalanceByID(context.Context, uuid.UUID) (int64, error)
	FindByCPF(context.Context, string) (*domain.Account, error)
	UpdateBalanceByID(context.Context, *domain.Account) error
}

type AccountUseCaseImpl struct {
	Repository repositories.AccountRepository
}

func NewAccountUseCaseImpl(r repositories.AccountRepository) AccountUseCase {
	return &AccountUseCaseImpl{r}
}

func (a *AccountUseCaseImpl) List(ctx context.Context) ([]domain.Account, error) {
	accs, err := a.Repository.List(ctx)
	if err != nil {
		return nil, &domain.UnprocessableError{Message: err.Error()}
	}

	if accs != nil {
		return accs, nil
	}

	return nil, &domain.NotFoundError{}
}

func (a *AccountUseCaseImpl) Create(ctx context.Context, acc *domain.Account) error {
	err := acc.Validate()
	if err != nil {
		return &domain.ValidationError{Message: err.Error()}
	}

	acc.Secret, err = crypt.HashSecret(acc.Secret)
	if err != nil {
		return &domain.UnprocessableError{Message: err.Error()}
	}

	err = a.Repository.Create(ctx, acc)
	if err != nil {
		return &domain.UnprocessableError{Message: err.Error()}
	}

	return nil
}

func (a *AccountUseCaseImpl) FindBalanceByID(ctx context.Context, id uuid.UUID) (int64, error) {
	acc, err := a.Repository.FindByID(ctx, id)
	if err != nil {
		message := err.Error()
		if strings.Contains(message, "no rows in result set") {
			return 0, &domain.NotFoundError{Message: message}
		}
		return 0, &domain.UnprocessableError{Message: message}
	}

	return acc.Balance, nil
}

func (a *AccountUseCaseImpl) FindByCPF(ctx context.Context, cpf string) (*domain.Account, error) {
	acc, err := a.Repository.FindByCPF(ctx, cpf)
	if err != nil {
		message := err.Error()
		if strings.Contains(message, "sql: no rows in result set") {
			return nil, &domain.NotFoundError{Message: message}
		}
		return nil, &domain.UnprocessableError{Message: message}
	}

	return acc, nil
}

func (a *AccountUseCaseImpl) UpdateBalanceByID(ctx context.Context, acc *domain.Account) error {
	err := a.Repository.UpdateBalanceByID(ctx, acc)
	if err != nil {
		message := err.Error()
		if strings.Contains(message, "sql: no rows in result set") {
			return &domain.NotFoundError{Message: message}
		}
		return &domain.UnprocessableError{Message: message}
	}

	return nil
}
