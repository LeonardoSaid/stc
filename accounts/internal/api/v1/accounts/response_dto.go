package accounts

import (
	"time"

	"github.com/google/uuid"
	"github.com/leonardosaid/stc/accounts/internal/domain"
	"github.com/mitchellh/mapstructure"
)

type ResponseDTO struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Cpf       string    `json:"cpf"`
	Secret    string    `json:"secret"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

type BalanceResponseDTO struct {
	Balance int64 `json:"balance"`
}

func ToAccountDTO(from *domain.Account) (to ResponseDTO, err error) {
	err = mapstructure.Decode(from, &to)
	to.CreatedAt = from.CreatedAt
	return
}

func ToBalanceDTO(amount int64) (to BalanceResponseDTO) {
	return BalanceResponseDTO{
		Balance: amount,
	}
}
