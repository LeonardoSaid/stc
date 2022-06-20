package domain

import "time"

type Account struct {
	ID        string
	Name      string
	CPF       string
	Secret    string
	Balance   int64
	CreatedAt time.Time
}
