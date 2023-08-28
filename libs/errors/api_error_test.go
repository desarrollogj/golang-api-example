package errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBadRequest(t *testing.T) {
	t.Log("New bad request should return a new bad request error")

	err := NewBadRequest("test message")

	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Equal(t, "test message", err.Message)
	assert.Equal(t, "bad_request", err.Err)
}

func TestNewMethodNotAllowedError(t *testing.T) {
	t.Log("TestNewMethodNotAllowedError should return a new method not allowed error")

	err := NewMethodNotAllowed("some error")

	assert.Equal(t, http.StatusMethodNotAllowed, err.Status)
	assert.Equal(t, "some error", err.Message)
	assert.Equal(t, "method_not_allowed", err.Err)
}

func TestNewResourceNotFoundError(t *testing.T) {
	t.Log("TestNewResourceNotFoundError should return a new resource not found error")

	err := NewResourceNotFound("some error")

	assert.Equal(t, http.StatusNotFound, err.Status)
	assert.Equal(t, "some error", err.Message)
	assert.Equal(t, "not_found", err.Err)
}

func TestNewInternalServerError(t *testing.T) {
	t.Log("NewInternalServerError should return a new internal server error")

	err := NewInternalServerError("some error")

	assert.Equal(t, http.StatusInternalServerError, err.Status)
	assert.Equal(t, "some error", err.Message)
	assert.Equal(t, "internal_error", err.Err)
}

func TestNewUnauthorizedError(t *testing.T) {
	t.Log("NewUnauthorizedError should return an unauthorized error")

	err := NewUnauthorizedError("some error")

	assert.Equal(t, http.StatusUnauthorized, err.Status)
	assert.Equal(t, "some error", err.Message)
	assert.Equal(t, "unauthorized", err.Err)
}

func TestHandleBusinessErrorWithResourceNotFoundError(t *testing.T) {
	t.Log("NewResourceNotFound should be get when a business NotFoundError is passed by parameters")

	notFoundErr := NewNotFoundError("Not found")

	apiErr := HandleBusinessError(notFoundErr)

	assert.Equal(t, http.StatusNotFound, apiErr.Status)
	assert.Equal(t, "Not found", apiErr.Message)
	assert.Equal(t, "not_found", apiErr.Err)
}

func TestHandleBusinessErrorWithNonFatalBusinessError(t *testing.T) {
	t.Log("Bad Request Api error should be get when a non fatal BusinessError is passed by parameters")

	notFatalErr := NewBusinessError("Not fatal error", "not_fatal_error")

	apiErr := HandleBusinessError(notFatalErr)

	assert.Equal(t, http.StatusBadRequest, apiErr.Status)
	assert.Equal(t, "Not fatal error", apiErr.Message)
	assert.Equal(t, "not_fatal_error", apiErr.Err)
}

func TestHandleBusinessErrorWithFatalBusinessError(t *testing.T) {
	t.Log("ServerError Api error should be get when a fatal BusinessError is passed by parameters")

	fatalErr := NewFatalError("fatal error")

	apiErr := HandleBusinessError(fatalErr)

	assert.Equal(t, http.StatusInternalServerError, apiErr.Status)
	assert.Equal(t, "fatal error", apiErr.Message)
	assert.Equal(t, "fatal_error", apiErr.Err)
}

func TestHandleBusinessErrorWithOtherError(t *testing.T) {
	t.Log("ServerError Api error should be get when a generic error is passed by parameters")

	genericErr := errors.New("generic error")

	apiErr := HandleBusinessError(genericErr)

	assert.Equal(t, http.StatusInternalServerError, apiErr.Status)
	assert.Equal(t, "generic error", apiErr.Message)
	assert.Equal(t, "internal_error", apiErr.Err)
}

func TestHandleBusinessErrorWithUnauthorizedError(t *testing.T) {
	t.Log("Unauthorized Api error should be get when a UnauthorizedError is passed by parameters")

	genericErr := NewBusinessUnauthorizedError("unauthorized resource")

	apiErr := HandleBusinessError(genericErr)

	assert.Equal(t, http.StatusUnauthorized, apiErr.Status)
	assert.Equal(t, "unauthorized resource", apiErr.Message)
	assert.Equal(t, "unauthorized", apiErr.Err)
}
