package utils

type ErrorInterfaces interface {
	Code() int
}

type BaseError struct {
	code int
	message string
}

func (i *BaseError) Code() int {
	return i.code
}
func (i *BaseError) Error() string {
	return i.message
}

func NewError(code int,message string) *BaseError {
	return &BaseError{
		code: code,
		message: message,
	}
}