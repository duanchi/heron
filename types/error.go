package types

import "net/http"

type RuntimeError struct {
	Message string
	ErrorCode    int
}

func (this RuntimeError) Error() string {
	return this.Message
}

func (this RuntimeError) Code() int {
	return this.ErrorCode
}

func (this RuntimeError) NewWithCode (err error, code int) RuntimeError {
	return RuntimeError{
		Message:   err.Error(),
		ErrorCode: code,
	}
}

func (this RuntimeError) New (err error) RuntimeError {
	return RuntimeError{
		Message:   err.Error(),
		ErrorCode: http.StatusInternalServerError,
	}
}