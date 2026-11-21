package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"api-gateway/internal/server"
)

func TestHealthCheck(t *testing.T) {
	t.Run("returns the server status and status code = 200", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/check", nil)
		response := httptest.NewRecorder()

		server.CheckHealth(response, request)

		if response.Code != http.StatusOK {
			t.Errorf("got status %d, want %d", response.Code, http.StatusOK)
		}

		if got, want := response.Body.String(), "OK"; got != want {
			t.Errorf("got body %q, want %q", got, want)
		}
	})
}
