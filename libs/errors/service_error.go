package errors

import (
	"encoding/json"
	"net/http"
)

const (
	FatalErrorCode        = "fatal_error"
	NotFoundErrorCode     = "not_found"
	ValidationErrorCode   = "validation_error"
	UnauthorizedErrorCode = "unauthorized"
)

func (e *BusinessError) Error() string {
	return e.Msg
}

// NewBusinessError creates and initializes a BusinessError.
func NewBusinessError(msg string, err string) *BusinessError {
	return &BusinessError{
		Msg:   msg,
		Err:   err,
		Fatal: false,
	}
}

// NewFatalError creates and initializes a BusinessError with fatal mark
func NewFatalError(msg string) *BusinessError {
	return &BusinessError{
		Msg:   msg,
		Err:   FatalErrorCode,
		Fatal: true,
	}
}

// NewNotFoundError creates and initializes a not found BusinessError
func NewNotFoundError(msg string) *BusinessError {
	return &BusinessError{
		Msg:   msg,
		Err:   NotFoundErrorCode,
		Fatal: false,
	}
}

// NewValidationError creates and initializes a validation BusinessError
func NewValidationError(msg string) *BusinessError {
	return &BusinessError{
		Msg:   msg,
		Err:   ValidationErrorCode,
		Fatal: false,
	}
}

// NewBusinessUnauthorizedError creates and initializes an unauthorized BusinessError
func NewBusinessUnauthorizedError(msg string) *BusinessError {
	return &BusinessError{
		Msg:   msg,
		Err:   UnauthorizedErrorCode,
		Fatal: false,
	}
}

// HandleFetcherResponse handles errors from fetchers returning an BusinessError
func HandleFetcherErrorResponse(status int, response []byte) *BusinessError {
	var apiErr APIError

	if err := json.Unmarshal(response, &apiErr); err != nil {
		return NewFatalError("unexpected error when try to decode the error response")
	}

	switch apiErr.Status {
	case http.StatusBadRequest:
		return NewBusinessError(apiErr.Message, apiErr.Err)
	case http.StatusUnauthorized:
		return NewBusinessUnauthorizedError(apiErr.Message)
	case http.StatusNotFound:
		return NewNotFoundError(apiErr.Message)
	default:
		return NewFatalError(apiErr.Message)
	}
}
