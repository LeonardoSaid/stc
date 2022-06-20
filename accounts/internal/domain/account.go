package domain

import (
	"regexp"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Account struct {
	bun.BaseModel `bun:"table:account,alias:a"`
	ID            uuid.UUID `bun:"id,pk,nullzero"`
	Name          string
	CPF           string `bun:",unique"`
	Secret        string
	Balance       int64
	CreatedAt     time.Time `bun:",nullzero,default:current_timestamp"`
}

func (a Account) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.CPF, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{11}$"))),
	)
}
