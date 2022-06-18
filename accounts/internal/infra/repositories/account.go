package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/leonardosaid/stc/accounts/internal/domain"
	"github.com/uptrace/bun"
)

type AccountRepository interface {
	List(context.Context) ([]domain.Account, error)
	Create(context.Context, *domain.Account) error
	FindByID(context.Context, uuid.UUID) (*domain.Account, error)
}

type AccountRepositoryImpl struct {
	DB *bun.DB
}

func NewAccountRepositoryImpl(db *bun.DB) AccountRepository {
	return &AccountRepositoryImpl{db}
}

func (a *AccountRepositoryImpl) List(ctx context.Context) ([]domain.Account, error) {
	rows := []domain.Account{}
	err := a.DB.NewSelect().Model(&rows).Scan(ctx)
	return rows, err
}

func (a *AccountRepositoryImpl) Create(ctx context.Context, i *domain.Account) (err error) {
	_, err = a.DB.NewInsert().Model(i).Returning("*").Exec(ctx)
	return
}

func (a *AccountRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (*domain.Account, error) {
	row := &domain.Account{}
	err := a.DB.NewSelect().Model(row).Where("id = ?", id.String()).Scan(ctx)
	return row, err
}
