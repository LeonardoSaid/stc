package error

import "fmt"

type ResponseError struct {
	Address   string
	Status    int
	ErrorBody string
}

func (r *ResponseError) Error() string {
	err := fmt.Sprintf("Error on request. url: %s. status: %d. content: %s", r.Address, r.Status, r.ErrorBody)
	return err
}
