package exception

type NotFoundError struct {
	Error string
	Code  int
}

func NewNotFoundError(error string, code int) NotFoundError {
	return NotFoundError{Error: error, Code: code}
}