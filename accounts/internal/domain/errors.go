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

type NotFoundError struct {
	Message string
}

func (n NotFoundError) Error() string {
	return n.Message
}
