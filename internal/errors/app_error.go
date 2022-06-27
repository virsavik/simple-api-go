package errors

type AppError struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

func NewAppError(message string, code string) AppError {
	return AppError{
		Message: message,
		Code:    code,
	}
}

func (e AppError) Error() string {
	return e.Message
}
