package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleRoot(t *testing.T) {
	t.Run("root path", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()

		handleRoot(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
		}

		expected := []byte("Go REST Mastery")
		if !bytes.Equal(expected, w.Body.Bytes()) {
			t.Errorf("Expected body %q, got %q", expected, w.Body.Bytes())
		}
	})

	t.Run("non-root path returns 404", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/something", nil)
		w := httptest.NewRecorder()

		handleRoot(w, req)

		if w.Code != http.StatusNotFound {
			t.Errorf("Expected status %d, got %d", http.StatusNotFound, w.Code)
		}
		// body is set by http.NotFound, it's HTML, we don't check it
	})
}

func TestHandleHello(t *testing.T) {
	req := httptest.NewRequest("GET", "/hello", nil)
	w := httptest.NewRecorder()

	handleHello(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	expected := []byte("Hello, HTTP!")
	if !bytes.Equal(expected, w.Body.Bytes()) {
		t.Errorf("Expected body %q, got %q", expected, w.Body.Bytes())
	}
}

func TestHandleEchoBody(t *testing.T) {
	t.Run("successful echo", func(t *testing.T) {
		requestBody := []byte("test echo body")
		req := httptest.NewRequest("POST", "/echo", bytes.NewReader(requestBody))
		w := httptest.NewRecorder()

		handleEchoBody(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
		}

		if !bytes.Equal(requestBody, w.Body.Bytes()) {
			t.Errorf("Expected body %q, got %q", requestBody, w.Body.Bytes())
		}
	})

	t.Run("method not allowed", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/echo", nil)
		w := httptest.NewRecorder()

		handleEchoBody(w, req)

		if w.Code != http.StatusMethodNotAllowed {
			t.Errorf("Expected status %d, got %d", http.StatusMethodNotAllowed, w.Code)
		}
	})

	t.Run("empty body", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/echo", bytes.NewReader([]byte{}))
		w := httptest.NewRecorder()

		handleEchoBody(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
		}

		if len(w.Body.Bytes()) != 0 {
			t.Errorf("Expected empty body, got %q", w.Body.Bytes())
		}
	})
}

func TestHandleAnyMethod(t *testing.T) {
	methods := []string{"GET", "POST", "PUT", "DELETE", "PATCH"}
	for _, method := range methods {
		t.Run("method "+method, func(t *testing.T) {
			req := httptest.NewRequest(method, "/method", nil)
			w := httptest.NewRecorder()

			handleAnyMethod(w, req)

			if w.Code != http.StatusOK {
				t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
			}

			var resp Response
			if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			if resp.Message != "Hello, JSON!" {
				t.Errorf("Expected message 'Hello, JSON!', got %q", resp.Message)
			}
			if resp.Status != 200 {
				t.Errorf("Expected status 200, got %d", resp.Status)
			}
		})
	}
}

func TestHandleHeaders(t *testing.T) {
	t.Run("successful headers", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/headers", nil)
		req.Header.Set("X-Custom", "test123")
		req.Header.Set("User-Agent", "test-agent")
		w := httptest.NewRecorder()

		handleHeaders(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
		}

		var resp HeadersResponse
		if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
			t.Fatalf("Failed to unmarshal response: %v", err)
		}

		if resp.Headers["X-Custom"] != "test123" {
			t.Errorf("Expected X-Custom header 'test123', got %q", resp.Headers["X-Custom"])
		}
		if resp.Headers["User-Agent"] != "test-agent" {
			t.Errorf("Expected User-Agent 'test-agent', got %q", resp.Headers["User-Agent"])
		}
	})

	t.Run("method not allowed", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/headers", nil)
		w := httptest.NewRecorder()

		handleHeaders(w, req)

		if w.Code != http.StatusMethodNotAllowed {
			t.Errorf("Expected status %d, got %d", http.StatusMethodNotAllowed, w.Code)
		}
	})
}

func TestHandleStatus(t *testing.T) {
	t.Run("valid status 200", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/status/200", nil)
		w := httptest.NewRecorder()

		handleStatus(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
		}

		var resp Response
		if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
			t.Fatalf("Failed to unmarshal response: %v", err)
		}
		if resp.Message != "Status 200 - OK" {
			t.Errorf("Expected message 'Status 200 - OK', got %q", resp.Message)
		}
		if resp.Status != 200 {
			t.Errorf("Expected status 200, got %d", resp.Status)
		}
	})

	t.Run("valid status 404", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/status/404", nil)
		w := httptest.NewRecorder()

		handleStatus(w, req)

		if w.Code != http.StatusNotFound {
			t.Errorf("Expected status %d, got %d", http.StatusNotFound, w.Code)
		}

		var resp Response
		if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
			t.Fatalf("Failed to unmarshal response: %v", err)
		}
		if resp.Message != "Status 404 - Not Found" {
			t.Errorf("Expected message 'Status 404 - Not Found', got %q", resp.Message)
		}
	})

	t.Run("valid status 500", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/status/500", nil)
		w := httptest.NewRecorder()

		handleStatus(w, req)

		if w.Code != http.StatusInternalServerError {
			t.Errorf("Expected status %d, got %d", http.StatusInternalServerError, w.Code)
		}
	})

	t.Run("invalid status code (non-numeric)", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/status/abc", nil)
		w := httptest.NewRecorder()

		handleStatus(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
		}
	})

	t.Run("invalid status code (out of range)", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/status/999", nil)
		w := httptest.NewRecorder()

		handleStatus(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
		}
	})

	t.Run("method not allowed", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/status/200", nil)
		w := httptest.NewRecorder()

		handleStatus(w, req)

		if w.Code != http.StatusMethodNotAllowed {
			t.Errorf("Expected status %d, got %d", http.StatusMethodNotAllowed, w.Code)
		}
	})
}

func TestHandleSearch(t *testing.T) {
	t.Run("with q and page", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/search?q=golang&page=2", nil)
		w := httptest.NewRecorder()

		handleSearch(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
		}

		var resp QueryResponse
		if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
			t.Fatalf("Failed to unmarshal response: %v", err)
		}
		if resp.Query != "golang" {
			t.Errorf("Expected query 'golang', got %q", resp.Query)
		}
		if resp.Page != "2" {
			t.Errorf("Expected page '2', got %q", resp.Page)
		}
	})

	t.Run("with only q", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/search?q=rust", nil)
		w := httptest.NewRecorder()

		handleSearch(w, req)

		var resp QueryResponse
		if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
			t.Fatalf("Failed to unmarshal response: %v", err)
		}
		if resp.Query != "rust" {
			t.Errorf("Expected query 'rust', got %q", resp.Query)
		}
		if resp.Page != "" {
			t.Errorf("Expected page empty, got %q", resp.Page)
		}
	})

	t.Run("with no params", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/search", nil)
		w := httptest.NewRecorder()

		handleSearch(w, req)

		var resp QueryResponse
		if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
			t.Fatalf("Failed to unmarshal response: %v", err)
		}
		if resp.Query != "" || resp.Page != "" {
			t.Errorf("Expected empty fields, got query=%q page=%q", resp.Query, resp.Page)
		}
	})

	t.Run("method not allowed", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/search", nil)
		w := httptest.NewRecorder()

		handleSearch(w, req)

		if w.Code != http.StatusMethodNotAllowed {
			t.Errorf("Expected status %d, got %d", http.StatusMethodNotAllowed, w.Code)
		}
	})
}

func TestHandleNotFound(t *testing.T) {
	t.Run("unknown path", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/unknown/path", nil)
		w := httptest.NewRecorder()

		handleNotFound(w, req)

		if w.Code != http.StatusNotFound {
			t.Errorf("Expected status %d, got %d", http.StatusNotFound, w.Code)
		}

		var resp ErrorResponse
		if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
			t.Fatalf("Failed to unmarshal response: %v", err)
		}
		if resp.Error != "not found" {
			t.Errorf("Expected error 'not found', got %q", resp.Error)
		}
		if resp.Path != "/unknown/path" {
			t.Errorf("Expected path '/unknown/path', got %q", resp.Path)
		}
	})

	t.Run("any method", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/anything", nil)
		w := httptest.NewRecorder()

		handleNotFound(w, req)

		if w.Code != http.StatusNotFound {
			t.Errorf("Expected status %d, got %d", http.StatusNotFound, w.Code)
		}
	})
}
