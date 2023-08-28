package errors

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBusinessError(t *testing.T) {
	t.Log("New business error should return a new business error")

	err := NewBusinessError("test message", "test_code")

	assert.Equal(t, "test message", err.Error())
	assert.Equal(t, "test message", err.Msg)
	assert.Equal(t, "test_code", err.Err)
	assert.False(t, err.Fatal)
}

func TestNewFatalError(t *testing.T) {
	t.Log("New fatal error should return a new fatal error")

	err := NewFatalError("test message")

	assert.Equal(t, "test message", err.Error())
	assert.Equal(t, "test message", err.Msg)
	assert.Equal(t, FatalErrorCode, err.Err)
	assert.True(t, err.Fatal)
}

func TestNewNotFoundError(t *testing.T) {
	t.Log("New not found error should return a new not found error")

	err := NewNotFoundError("test message")

	assert.Equal(t, "test message", err.Error())
	assert.Equal(t, "test message", err.Msg)
	assert.Equal(t, NotFoundErrorCode, err.Err)
	assert.False(t, err.Fatal)
}

func TestNewBusinessUnauthorizedError(t *testing.T) {
	t.Log("New business unauthorized error should return a new unauthorized")

	err := NewBusinessUnauthorizedError("test message")

	assert.Equal(t, "test message", err.Error())
	assert.Equal(t, "test message", err.Msg)
	assert.Equal(t, UnauthorizedErrorCode, err.Err)
	assert.False(t, err.Fatal)
}

func TestNewValidationErrorError(t *testing.T) {
	t.Log("New validation error should return a new validation error")

	err := NewValidationError("test message")

	assert.Equal(t, "test message", err.Error())
	assert.Equal(t, "test message", err.Msg)
	assert.Equal(t, ValidationErrorCode, err.Err)
	assert.False(t, err.Fatal)
}

func TestHandleFetcherErrorResponseBadRequest(t *testing.T) {
	t.Log("Handle fetcher error response should return a Business Error when a bad request response was received")

	response := APIError{
		Status:  400,
		Err:     "user_does_not_exist",
		Message: "User does not exist",
	}
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(response)
	responseData := buffer.Bytes()

	err := HandleFetcherErrorResponse(http.StatusBadRequest, responseData)

	assert.Equal(t, response.Message, err.Error())
	assert.Equal(t, response.Message, err.Msg)
	assert.Equal(t, response.Err, err.Err)
	assert.False(t, err.Fatal)
}

func TestHandleFetcherErrorResponseNotFound(t *testing.T) {
	t.Log("Handle fetcher error response should return a Business Error when a not found response was received")

	response := APIError{
		Status:  404,
		Err:     "not_found",
		Message: "User not found",
	}
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(response)
	responseData := buffer.Bytes()

	err := HandleFetcherErrorResponse(http.StatusNotFound, responseData)

	assert.Equal(t, response.Message, err.Error())
	assert.Equal(t, response.Message, err.Msg)
	assert.Equal(t, response.Err, err.Err)
	assert.False(t, err.Fatal)
}

func TestHandleFetcherErrorResponseUnauthorized(t *testing.T) {
	t.Log("Handle fetcher error response should return a Business Error when an unauthorized response was received")

	response := APIError{
		Status:  401,
		Err:     "unauthorized",
		Message: "User unauthorized",
	}
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(response)
	responseData := buffer.Bytes()

	err := HandleFetcherErrorResponse(http.StatusUnauthorized, responseData)

	assert.Equal(t, response.Message, err.Error())
	assert.Equal(t, response.Message, err.Msg)
	assert.Equal(t, response.Err, err.Err)
	assert.False(t, err.Fatal)
}

func TestHandleFetcherErrorResponseInternalServerError(t *testing.T) {
	t.Log("Handle fetcher error response should return a Business Error when an internal server error was received")

	response := APIError{
		Status:  500,
		Err:     "fatal_error",
		Message: "Unexpected error",
	}
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(response)
	responseData := buffer.Bytes()

	err := HandleFetcherErrorResponse(http.StatusInternalServerError, responseData)

	assert.Equal(t, response.Message, err.Error())
	assert.Equal(t, response.Message, err.Msg)
	assert.Equal(t, response.Err, err.Err)
	assert.True(t, err.Fatal)
}

func TestHandleFetcherErrorResponseNotValidResponseData(t *testing.T) {
	t.Log("Handle fetcher error response should return a generic error when not valid response data was received")

	response := "THIS IS NOT AN APIERROR!"
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(response)
	responseData := buffer.Bytes()

	err := HandleFetcherErrorResponse(http.StatusInternalServerError, responseData)

	assert.Equal(t, "unexpected error when try to decode the error response", err.Error())
	assert.Equal(t, "unexpected error when try to decode the error response", err.Msg)
	assert.Equal(t, "fatal_error", err.Err)
	assert.True(t, err.Fatal)
}
