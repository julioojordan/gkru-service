package exception

type InternalServerError struct {
	Error string
	Code  int
}

func NewInternalServerError() InternalServerError {
	return InternalServerError{Error: "internal server error", Code: 500}
}