package errors

import (
	"errors"
	"net/http"
	"strings"
)

const (
	BadRequestMessage          = "invalid request parameters"
	ResourceNotFoundMessage    = "resource not found"
	MethodNotAllowedMessage    = "method not allowed on the current resource"
	InternalServerErrorMessage = "internal Server Error"
	NotFoundErrorMessage       = "not found"
	UnathorizedErrorMessage    = "unauthorized"
)

// NewAPIError creates and initializes an APIError.
func NewAPIError(code int, message string, err string) *APIError {
	return &APIError{
		Status:  code,
		Message: message,
		Err:     err,
	}
}

// NewBadRequest creates an API Error for an invalid or malformed request.
func NewBadRequest(messages ...string) *APIError {
	message := BadRequestMessage
	if len(messages) > 0 {
		message = strings.Join(messages, " - ")
	}

	return NewAPIError(http.StatusBadRequest, message, "bad_request")
}

// NewResourceNotFound creates an API Error for an unexisting resource.
func NewResourceNotFound(messages ...string) *APIError {
	message := ResourceNotFoundMessage
	if len(messages) > 0 {
		message = strings.Join(messages, " - ")
	}

	return NewAPIError(http.StatusNotFound, message, "not_found")
}

// NewMethodNotAllowed creates an API Error for a forbidden verb on a resource.
func NewMethodNotAllowed(messages ...string) *APIError {
	message := MethodNotAllowedMessage
	if len(messages) > 0 {
		message = strings.Join(messages, " - ")
	}

	return NewAPIError(http.StatusMethodNotAllowed, message, "method_not_allowed")
}

// NewUnauthorizedError creates an API Error for an unauthorized access on a resource.
func NewUnauthorizedError(messages ...string) *APIError {
	message := UnathorizedErrorMessage
	if len(messages) > 0 {
		message = strings.Join(messages, " - ")
	}

	return NewAPIError(http.StatusUnauthorized, message, "unauthorized")
}

// NewInternalServerError creates an API Error for an unexpected condition.
func NewInternalServerError(messages ...string) *APIError {
	message := InternalServerErrorMessage
	if len(messages) > 0 {
		message = strings.Join(messages, " - ")
	}

	return NewAPIError(http.StatusInternalServerError, message, "internal_error")
}

// HandleBusinessError handles errors from services and use cases. Converts the errors to their REST equivalent
func HandleBusinessError(err error) *APIError {
	var bisErr *BusinessError
	switch {
	case errors.As(err, &bisErr):
		if bisErr.Err == NotFoundErrorCode {
			return NewResourceNotFound(bisErr.Msg)
		} else if bisErr.Err == UnauthorizedErrorCode {
			return NewUnauthorizedError(bisErr.Msg)
		} else if !bisErr.Fatal {
			return NewAPIError(http.StatusBadRequest, bisErr.Msg, bisErr.Err)
		}
		return NewAPIError(http.StatusInternalServerError, bisErr.Msg, bisErr.Err)
	default:
		return NewInternalServerError(err.Error())
	}
}
