package exception

type BaseError struct {
	ErrorMsg string
	Code  int
}

func NewBaseError(error string, code int) BaseError {
	return BaseError{ErrorMsg: error, Code: code}
}

// implement ini supaya memnuhi kontrak error golang
func (e BaseError) Error() string {
	return e.ErrorMsg
}