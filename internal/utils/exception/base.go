package exception

type ErrorType int

type AppError struct {
	Type    ErrorType
	Message string
}

func (e AppError) Error() string {
	return e.Message
}
