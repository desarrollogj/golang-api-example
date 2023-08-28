package handler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthSuccess(t *testing.T) {
	t.Log("Successfully response for health check")

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)

	r := testRouter()
	r.GET("/health", Health)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	bodyBytes, _ := io.ReadAll(w.Body)

	assert.Equal(t, "{\"status\":\"OK\",\"environment\":\"LOCAL\",\"app\":\"UNKNOWN\",\"version\":\"UNKNOWN\"}", string(bodyBytes))
}
