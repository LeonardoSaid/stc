package accounts

import (
	"github.com/leonardosaid/stc/accounts/internal/domain"
	"github.com/mitchellh/mapstructure"
)

type CreateAccountDTO struct {
	Name   string `json:"name"`
	Cpf    string `json:"cpf"`
	Secret string `json:"secret"`
}

func (acc *CreateAccountDTO) ToEntity() (to domain.Account, err error) {
	err = mapstructure.Decode(acc, &to)
	if err != nil {
		return
	}
	return
}
