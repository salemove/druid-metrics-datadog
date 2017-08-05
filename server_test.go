package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthz(t *testing.T) {
	server := Server{}

	handler := http.HandlerFunc(server.healthzHandler)
	responseWriter := httptest.NewRecorder()

	request, _ := http.NewRequest("GET", "http://example.com/healthz", nil)
	handler.ServeHTTP(responseWriter, request)

	response := responseWriter.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode, "should be successful request")
}
