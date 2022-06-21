package domain

type UnprocessableError struct {
	Message string
}

func (s UnprocessableError) Error() string {
	return s.Message
}

type ValidationError struct {
	Message string
}

func (v ValidationError) Error() string {
	return v.Message
}

type AccountServiceError struct {
	Message string
}

func (a AccountServiceError) Error() string {
	return a.Message
}

type InsufficientFundsError struct{}

func (i InsufficientFundsError) Error() string {
	return "user does not have enough funds for debit"
}

type NotFoundError struct {
	Message string
}

func (n NotFoundError) Error() string {
	return n.Message
}
