package exception

const (
	ErrInternal ErrorType = iota
	ErrNotFound
	ErrInvalid
	ErrUnauthorized
	ErrForbidden
	ErrAlreadyExists
)

const (
	RepositoryError ErrorType = iota
	ValidatorError
	OtherError
)

func NewRepositoryError(msg string) AppError {
	return AppError{Type: RepositoryError, Message: msg}
}

func NewValidatorError(msg string) AppError {
	return AppError{Type: ValidatorError, Message: msg}
}

func NewOtherError(msg string) AppError {
	return AppError{Type: OtherError, Message: msg}
}
