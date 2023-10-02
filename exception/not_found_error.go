package exception

type NotFoundError struct {
	Error string
}

func NewNotFoundError(error string) *NotFoundError {
	return &NotFoundError{Error: error}
}

func PanicNotFoundIfError(err error, message string) {
	if err != nil {
		panic(NewNotFoundError(message))
	}
}
