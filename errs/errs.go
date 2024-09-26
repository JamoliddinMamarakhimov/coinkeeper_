package errs

import "errors"

var (
	ErrPermissionDenied            = errors.New("ErrPermissionDenied")
	ErrValidationFailed            = errors.New("ErrValidationFailed")
	ErrUsernameUniquenessFailed    = errors.New("ErrUsernameUniquenessFailed")
	ErrOperationNotFound           = errors.New("ErrOperationNotFound")
	ErrIncorrectUsernameOrPassword = errors.New("ErrIncorrectUsernameOrPassword")
	ErrRecordNotFound              = errors.New("ErrRecordNotFound")
	ErrUserNotFound                = errors.New("ErrUserNotFound")
	ErrSomethingWentWrong          = errors.New("ErrSomethingWentWrong")
)
