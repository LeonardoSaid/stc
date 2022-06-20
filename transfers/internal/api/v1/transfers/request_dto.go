package transfers

import (
	"github.com/leonardosaid/stc/accounts/internal/config"
	"github.com/leonardosaid/stc/accounts/internal/domain"
)

type CreateTransferDTO struct {
	AccountDestinationID string `json:"account_destination_id"`
	Amount               int64  `json:"amount"`
}

type ListTransfersDTO struct {
	LoginToken string `header:"Authorization"`
}

func (t *CreateTransferDTO) ToEntity() (to domain.Transfer, err error) {
	err = config.CustomDecode(t, &to)
	if err != nil {
		return
	}
	return
}
