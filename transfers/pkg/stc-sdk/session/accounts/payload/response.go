package payload

import "time"

type AccountResponse struct {
	ID        string
	Name      string
	CPF       string
	Secret    string
	Balance   int64
	CreatedAt time.Time
}

type BalanceResponse struct {
	Balance int64
}
