package errors

import "fmt"

type AppError struct {
	Code      int               `json:"code"`
	Message   string            `json:"message"`
	Metadata  map[string]string `json:"metadata"`
	Temporary bool              `json:"temporary"`
	err       error
}

func (a *AppError) Error() string {
	return fmt.Sprintf("[%d] %s", a.Code, a.Message)
}

func (a *AppError) Unwrap() error {
	return a.err
}

func Wrap(err error) *AppError {
	return &AppError{
		Code:      -1,
		Message:   err.Error(),
		Metadata:  map[string]string{},
		Temporary: false,
		err:       err,
	}
}

func WrapWithMessage(err error, msg string) *AppError {
	return &AppError{
		Code:      -1,
		Message:   msg,
		Metadata:  map[string]string{},
		Temporary: false,
		err:       err,
	}
}
