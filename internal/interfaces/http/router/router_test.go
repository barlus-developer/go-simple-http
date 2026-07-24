package router

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	apphealth "github.com/barlus-developer/go-simple-http/internal/application/health"
	"github.com/barlus-developer/go-simple-http/internal/infrastructure/config"
	"github.com/barlus-developer/go-simple-http/internal/interfaces/http/handler"
	"go.uber.org/zap/zaptest"
)

func TestRootEndpointReturnsHealthStatus(t *testing.T) {
	engine := New(
		config.Config{App: config.AppConfig{Environment: "test"}},
		zaptest.NewLogger(t),
		handler.NewHealthHandler(apphealth.NewService()),
	)

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	engine.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status code %d, got %d", http.StatusOK, response.Code)
	}

	var body map[string]string
	if err := json.Unmarshal(response.Body.Bytes(), &body); err != nil {
		t.Fatalf("decode response body: %v", err)
	}
	if body["status"] != "ok" {
		t.Fatalf("expected body status ok, got %q", body["status"])
	}
	if body["message"] != "Hello, World!!!" {
		t.Fatalf("expected hello message, got %q", body["message"])
	}
}

func TestUnknownEndpointReturnsNotFound(t *testing.T) {
	engine := New(
		config.Config{App: config.AppConfig{Environment: "test"}},
		zaptest.NewLogger(t),
		handler.NewHealthHandler(apphealth.NewService()),
	)

	request := httptest.NewRequest(http.MethodGet, "/missing", nil)
	response := httptest.NewRecorder()

	engine.ServeHTTP(response, request)

	if response.Code != http.StatusNotFound {
		t.Fatalf("expected status code %d, got %d", http.StatusNotFound, response.Code)
	}
}
