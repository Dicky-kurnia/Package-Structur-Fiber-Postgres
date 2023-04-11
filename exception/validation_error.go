package exception

import "fmt"

type ValidationError struct {
	Message string
}

func (validationError ValidationError) Error() string {
	return validationError.Message
}

func NewValidationError(field, message string) ValidationError {
	return ValidationError{
		Message: fmt.Sprintf(`{"%s": "%s"}`, field, message),
	}
}
