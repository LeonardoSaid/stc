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
	claims := token.Claims.(*domain.LoginToken)

	id, err := uuid.Parse(claims.ID)
	if err != nil {
		return nil, &domain.UnprocessableError{Message: err.Error()}
	}

	trs, err := t.Repository.ListByAccountID(ctx, id)
	if err != nil {
		return nil, &domain.UnprocessableError{Message: err.Error()}
	}

	if len(trs) > 0 {
		return trs, nil
	}

	return nil, &domain.NotFoundError{}
}

func (t *TransferUseCaseImpl) Create(ctx context.Context, tr *domain.Transfer, token *jwt.Token) error {
	err := tr.Validate()
	if err != nil {
		return &domain.ValidationError{Message: err.Error()}
	}

	claims := token.Claims.(*domain.LoginToken)
	id, err := uuid.Parse(claims.ID)
	if err != nil {
		return &domain.UnprocessableError{Message: err.Error()}
	}

	tr.AccountOriginID = id

	debtorBalance, err := t.Clients.GetAccountServiceClient().FindBalanceByID(id.String())
	if err != nil {
		return &domain.AccountServiceError{Message: err.Error()}
	}

	creditorBalance, err := t.Clients.GetAccountServiceClient().FindBalanceByID(tr.AccountDestinationID.String())
	if err != nil {
		return &domain.AccountServiceError{Message: err.Error()}
	}

	newDebtorBalance := debtorBalance.Balance - tr.Amount
	if newDebtorBalance < 0 {
		return &domain.InsufficientFundsError{}
	}

	err = t.updateAccountBalance(id, newDebtorBalance)
	if err != nil {
		return &domain.AccountServiceError{Message: err.Error()}
	}

	newCreditorBalance := creditorBalance.Balance + tr.Amount
	err = t.updateAccountBalance(tr.AccountDestinationID, newCreditorBalance)
	if err != nil {
		return &domain.AccountServiceError{Message: err.Error()}
	}

	err = t.Repository.Create(ctx, tr)
	if err != nil {
		return &domain.UnprocessableError{Message: err.Error()}
	}
	return nil
}

func (t *TransferUseCaseImpl) updateAccountBalance(id uuid.UUID, balance int64) error {
	requestData := payload.UpdateBalanceRequest{
		Balance: balance,
	}
	err := t.Clients.GetAccountServiceClient().UpdateBalanceByID(id.String(), requestData)
	if err != nil {
		return &domain.AccountServiceError{Message: err.Error()}
	}

	return nil
}
