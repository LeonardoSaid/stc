package transfers

import (
	"time"

	"github.com/google/uuid"
	"github.com/leonardosaid/stc/accounts/internal/domain"
	"github.com/mitchellh/mapstructure"
)

type ResponseDTO struct {
	ID                   uuid.UUID `json:"id"`
	AccountOriginID      uuid.UUID `json:"account_origin_id"`
	AccountDestinationID uuid.UUID `json:"account_destination_id"`
	Amount               int64     `json:"amount"`
	CreatedAt            time.Time `json:"created_at"`
}

func ToTransferDTO(from *domain.Transfer) (to ResponseDTO, err error) {
	err = mapstructure.Decode(from, &to)
	to.CreatedAt = from.CreatedAt
	return
}
