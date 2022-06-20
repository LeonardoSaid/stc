package login

import (
	"github.com/leonardosaid/stc/accounts/internal/domain"
	"github.com/mitchellh/mapstructure"
)

type LoginDTO struct {
	CPF    string `json:"cpf"`
	Secret string `json:"secret"`
}

func (l *LoginDTO) ToEntity() (to domain.LoginCredentials, err error) {
	err = mapstructure.Decode(l, &to)
	if err != nil {
		return
	}
	return
}
