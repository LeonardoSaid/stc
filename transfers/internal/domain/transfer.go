package domain

import (
	"time"

	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Transfer struct {
	bun.BaseModel        `bun:"table:transfer,alias:a"`
	ID                   uuid.UUID `bun:"id,pk,nullzero"`
	AccountOriginID      uuid.UUID
	AccountDestinationID uuid.UUID
	Amount               int64
	CreatedAt            time.Time `bun:",nullzero,default:current_timestamp"`
}

func (t Transfer) Validate() error {
	return validation.ValidateStruct(&t,
		validation.Field(&t.AccountOriginID, validation.Required, is.UUID),
		validation.Field(&t.AccountDestinationID, validation.Required, is.UUID),
		validation.Field(&t.Amount, validation.Required, validation.Min(1)),
	)
}
