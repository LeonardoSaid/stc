package domain

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/golang-jwt/jwt"
)

type LoginToken struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

type LoginCredentials struct {
	CPF    string `json:"cpf"`
	Secret string `json:"secret"`
}

func (l LoginCredentials) Validate() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.CPF, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{11}$"))),
	)
}
