package util

import (
	"runtime"
)

// adapted from here: https://dev.to/tigorlazuardi/go-creating-custom-error-wrapper-and-do-proper-error-equality-check-11k7
var (
	CodeClientError         = 1001
	CodeNotFoundError       = 1002
	CodeConflictError       = 1003
	CodeServerError         = 1004
	CodeClientUnauthorized  = 1005
	CodeClientForbidden     = 1006
	CodeUnprocessableEntity = 1007

	CodeCacheMiss = 1008
)

var (
	MsgServerError               = "server error"
	MsgClientBadFormattedRequest = "bad format request"
	MsgClientNotFoundRequest     = "resource not found"
	MsgClientUnauthorized        = "unauthorized"
	MsgClientForbidden           = "forbidden"
	MsgValidationError           = "validation error"
	MsgConflictRequest           = "resource conflict"
)

type ErrorWrapper struct {
	Message    string      `json:"message"` // human readable error
	Validation interface{} `json:"-"`       //
	Code       int         `json:"-"`       // code
	Err        error       `json:"-"`       // original error
	Filename   string      `json:"-"`
	LineNumber int         `json:"-"`
}

func (w *ErrorWrapper) Error() string {
	// guard against panics
	if w.Err != nil {
		return w.Err.Error()
	}
	return w.Message
}

func NewErrorWrapper(code int, msg string, validation interface{}, err error) *ErrorWrapper {
	// getting previous call stack file and line info
	_, filename, line, _ := runtime.Caller(1)
	return &ErrorWrapper{
		Code:       code,
		Message:    msg,
		Err:        err,
		Validation: validation,
		Filename:   filename,
		LineNumber: line,
	}
}

func NewNotFoundError(message string) *ErrorWrapper {
	_, filename, line, _ := runtime.Caller(1)
	return &ErrorWrapper{
		Code:       CodeNotFoundError,
		Message:    message,
		Err:        nil,
		Validation: nil,
		Filename:   filename,
		LineNumber: line,
	}
}

func NewUnAuthorizedError() *ErrorWrapper {
	_, filename, line, _ := runtime.Caller(1)
	return &ErrorWrapper{
		Code:       CodeClientUnauthorized,
		Message:    MsgClientUnauthorized,
		Err:        nil,
		Validation: nil,
		Filename:   filename,
		LineNumber: line,
	}
}

func NewForbiddenError() *ErrorWrapper {
	_, filename, line, _ := runtime.Caller(1)
	return &ErrorWrapper{
		Code:       CodeClientForbidden,
		Message:    MsgClientForbidden,
		Err:        nil,
		Validation: nil,
		Filename:   filename,
		LineNumber: line,
	}
}

func NewBadRequestError(message string) *ErrorWrapper {
	_, filename, line, _ := runtime.Caller(1)
	return &ErrorWrapper{
		Code:       CodeClientError,
		Message:    message,
		Err:        nil,
		Validation: nil,
		Filename:   filename,
		LineNumber: line,
	}
}

func NewConflictError(message string) *ErrorWrapper {
	_, filename, line, _ := runtime.Caller(1)
	return &ErrorWrapper{
		Code:       CodeConflictError,
		Message:    message,
		Err:        nil,
		Validation: nil,
		Filename:   filename,
		LineNumber: line,
	}
}

func NewUnprocessibleEntityError(validation interface{}) *ErrorWrapper {
	_, filename, line, _ := runtime.Caller(1)
	return &ErrorWrapper{
		Code:       CodeUnprocessableEntity,
		Message:    MsgValidationError,
		Err:        nil,
		Validation: validation,
		Filename:   filename,
		LineNumber: line,
	}
}
