package usecases

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/leonardosaid/stc/accounts/internal/clients"
	"github.com/leonardosaid/stc/accounts/internal/domain"
	"github.com/leonardosaid/stc/accounts/pkg/crypt"
	"os"
)

type LoginUseCase interface {
	Login(context.Context, *domain.LoginCredentials) (string, error)
}

type LoginUseCaseImpl struct {
	Clients     clients.ServiceClients
	TokenSecret []byte
}

func NewLoginUseCaseImpl(c clients.ServiceClients) LoginUseCase {
	return &LoginUseCaseImpl{c, []byte(os.Getenv("JWT_TOKEN_SECRET"))}
}

func (l *LoginUseCaseImpl) Login(ctx context.Context, payload *domain.LoginCredentials) (string, error) {
	acc, err := l.Clients.GetAccountServiceClient().FindByCPF(payload.CPF)
	if err != nil {
		return "", err
	}

	err = crypt.CompareHash(acc.Secret, payload.Secret)
	if err != nil {
		fmt.Println("Invalid password !!!!!!!!!!!!!!!!!!!!!!!!!!")
		return "", err
	}

	claims := &domain.LoginToken{
		ID: acc.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString(l.TokenSecret)
	if err != nil {
		return "", err
	}

	return t, nil
}
