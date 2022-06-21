package repositories

import (
	"context"

	"github.com/google/uuid"

	"github.com/leonardosaid/stc/accounts/internal/domain"
	"github.com/uptrace/bun"
)

type TransferRepository interface {
	ListByAccountID(context.Context, uuid.UUID) ([]domain.Transfer, error)
	Create(context.Context, *domain.Transfer) error
}

type TransferRepositoryImpl struct {
	DB *bun.DB
}

func NewTransferRepositoryImpl(db *bun.DB) TransferRepository {
	return &TransferRepositoryImpl{db}
}

func (a *TransferRepositoryImpl) ListByAccountID(ctx context.Context, id uuid.UUID) ([]domain.Transfer, error) {
	rows := []domain.Transfer{}
	err := a.DB.NewSelect().Model(&rows).Where("account_origin_id = ?", id.String()).Scan(ctx)
	return rows, err
}

func (a *TransferRepositoryImpl) Create(ctx context.Context, i *domain.Transfer) (err error) {
	_, err = a.DB.NewInsert().Model(i).Returning("*").Exec(ctx)
	return
}
