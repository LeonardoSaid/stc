package usecases

import (
	"context"
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
	return a.Repository.List(ctx)
}

func (a *AccountUseCaseImpl) Create(ctx context.Context, acc *domain.Account) error {
	err := acc.Validate()
	if err != nil {
		return err
	}

	acc.Secret, err = crypt.HashSecret(acc.Secret)
	if err != nil {
		return err
	}

	return a.Repository.Create(ctx, acc)
}

func (a *AccountUseCaseImpl) FindBalanceByID(ctx context.Context, id uuid.UUID) (int64, error) {
	acc, err := a.Repository.FindByID(ctx, id)
	if err != nil {
		return 0, err
	}

	return acc.Balance, nil
}

func (a *AccountUseCaseImpl) FindByCPF(ctx context.Context, cpf string) (*domain.Account, error) {
	return a.Repository.FindByCPF(ctx, cpf)
}

func (a *AccountUseCaseImpl) UpdateBalanceByID(ctx context.Context, acc *domain.Account) error {
	return a.Repository.UpdateBalanceByID(ctx, acc)
}
