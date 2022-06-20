package usecases

import (
	"context"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/leonardosaid/stc/accounts/internal/clients"
	"github.com/leonardosaid/stc/accounts/internal/domain"
	"github.com/leonardosaid/stc/accounts/internal/infra/repositories"
	"github.com/leonardosaid/stc/accounts/pkg/stc-sdk/session/accounts/payload"
)

type TransferUseCase interface {
	ListByAccountID(context.Context, *jwt.Token) ([]domain.Transfer, error)
	Create(context.Context, *domain.Transfer, *jwt.Token) error
}

type TransferUseCaseImpl struct {
	Clients    clients.ServiceClients
	Repository repositories.TransferRepository
}

func NewTransferUseCaseImpl(c clients.ServiceClients, r repositories.TransferRepository) TransferUseCase {
	return &TransferUseCaseImpl{c, r}
}

func (t *TransferUseCaseImpl) ListByAccountID(ctx context.Context, token *jwt.Token) ([]domain.Transfer, error) {
	// check if token is valid
	claims := token.Claims.(*domain.LoginToken)

	id, err := uuid.Parse(claims.ID)
	if err != nil {
		return nil, err
	}

	return t.Repository.ListByAccountID(ctx, id)
}

func (t *TransferUseCaseImpl) Create(ctx context.Context, tr *domain.Transfer, token *jwt.Token) error {
	// check if token is valid
	claims := token.Claims.(*domain.LoginToken)

	id, err := uuid.Parse(claims.ID)
	if err != nil {
		return err
	}

	tr.AccountOriginID = id

	debtorBalance, err := t.Clients.GetAccountServiceClient().FindBalanceByID(id.String())
	if err != nil {
		return err
	}

	creditorBalance, err := t.Clients.GetAccountServiceClient().FindBalanceByID(tr.AccountDestinationID.String())
	if err != nil {
		return err
	}

	newDebtorBalance := debtorBalance.Balance - tr.Amount
	if newDebtorBalance < 0 {
		// error
	}

	err = t.updateAccountBalance(id, newDebtorBalance)
	if err != nil {
		return err
	}

	newCreditorBalance := creditorBalance.Balance + tr.Amount
	err = t.updateAccountBalance(tr.AccountDestinationID, newCreditorBalance)
	if err != nil {
		return err
	}

	return t.Repository.Create(ctx, tr)
}

func (t *TransferUseCaseImpl) updateAccountBalance(id uuid.UUID, balance int64) error {
	requestData := payload.UpdateBalanceRequest{
		Balance: balance,
	}
	return t.Clients.GetAccountServiceClient().UpdateBalanceByID(id.String(), requestData)
}
