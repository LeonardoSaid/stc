package domain

type UnprocessableError struct {
	Message string
}

func (s UnprocessableError) Error() string {
	return s.Message
}

type NotFoundError struct {
	Message string
}

func (n NotFoundError) Error() string {
	return n.Message
}

type InvalidCredentialsError struct{}

func (i InvalidCredentialsError) Error() string {
	return "invalid credentials"
}

type ValidationError struct {
	Message string
}

func (v ValidationError) Error() string {
	return v.Message
}
