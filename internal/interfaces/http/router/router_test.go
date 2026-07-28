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
		config.Config{App: config.AppConfig{Debug: true}},
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
	if body["message"] == "" {
		t.Fatalf("expected a non-empty message")
	}
}

func TestUnknownEndpointReturnsNotFound(t *testing.T) {
	engine := New(
		config.Config{App: config.AppConfig{Debug: true}},
		zaptest.NewLogger(t),
		handler.NewHealthHandler(apphealth.NewService()),
	)

	request := httptest.NewRequest(http.MethodGet, "/missing", nil)
	response := httptest.NewRecorder()

	engine.ServeHTTP(response, request)

	if response.Code != http.StatusNotFound {
		t.Fatalf("expected status code %d, got %d", http.StatusNotFound, response.Code)
	}

	if response.Header().Get("Content-Type") == "" {
		t.Fatalf("expected a Content-Type header to be set")
	}

	var body map[string]string
	if err := json.Unmarshal(response.Body.Bytes(), &body); err != nil {
		t.Fatalf("decode response body: %v", err)
	}
	if body["status"] != "error" {
		t.Fatalf("expected body status error, got %q", body["status"])
	}
	if body["message"] == "" {
		t.Fatalf("expected a message describing the missing route")
	}
}

func TestUnsupportedMethodReturnsMethodNotAllowed(t *testing.T) {
	engine := New(
		config.Config{App: config.AppConfig{Debug: true}},
		zaptest.NewLogger(t),
		handler.NewHealthHandler(apphealth.NewService()),
	)

	request := httptest.NewRequest(http.MethodPost, "/", nil)
	response := httptest.NewRecorder()

	engine.ServeHTTP(response, request)

	if response.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status code %d, got %d", http.StatusMethodNotAllowed, response.Code)
	}

	var body map[string]string
	if err := json.Unmarshal(response.Body.Bytes(), &body); err != nil {
		t.Fatalf("decode response body: %v", err)
	}
	if body["status"] != "error" {
		t.Fatalf("expected body status error, got %q", body["status"])
	}
	if body["message"] == "" {
		t.Fatalf("expected a message describing the disallowed method")
	}
}
