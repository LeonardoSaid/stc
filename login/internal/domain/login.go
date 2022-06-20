package domain

import "github.com/golang-jwt/jwt"

type LoginToken struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

type LoginCredentials struct {
	CPF    string `json:"cpf"`
	Secret string `json:"secret"`
}
