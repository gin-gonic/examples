package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func setupRouter() *gin.Engine {
	router := gin.New()
	router.Use(RequestID())

	router.GET("/test", func(c *gin.Context) {
		requestID, exists := c.Get("RequestID")
		if !exists {
			c.Status(http.StatusInternalServerError)
			return
		}

		c.String(http.StatusOK, requestID.(string))
	})

	return router
}

func TestRequestID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("generates request ID when header is missing", func(t *testing.T) {
		router := setupRouter()

		req := httptest.NewRequestWithContext(
			context.Background(), http.MethodGet, "/test", nil,
		)
		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		requestID := recorder.Header().Get(RequestIDHeader)

		if recorder.Code != http.StatusOK {
			t.Fatalf("expected status %d, got %d", http.StatusOK, recorder.Code)
		}

		if _, err := uuid.Parse(requestID); err != nil {
			t.Fatalf("expected a valid UUID, got %q: %v", requestID, err)
		}

		if recorder.Body.String() != requestID {
			t.Errorf("expected body request ID %q, got %q", requestID, recorder.Body.String())
		}
	})

	t.Run("reuses existing valid request ID", func(t *testing.T) {
		router := setupRouter()

		validID := "550e8400-e29b-41d4-a716-446655440000"

		req := httptest.NewRequestWithContext(
			context.Background(), http.MethodGet, "/test", nil,
		)
		req.Header.Set(RequestIDHeader, validID)

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		if recorder.Code != http.StatusOK {
			t.Fatalf("expected status %d, got %d", http.StatusOK, recorder.Code)
		}

		if got := recorder.Header().Get(RequestIDHeader); got != validID {
			t.Errorf("expected request ID %q, got %q", validID, got)
		}

		if got := recorder.Body.String(); got != validID {
			t.Errorf("expected body request ID %q, got %q", validID, got)
		}
	})

	t.Run("generates new ID for invalid incoming ID", func(t *testing.T) {
		router := setupRouter()

		req := httptest.NewRequestWithContext(
			context.Background(), http.MethodGet, "/test", nil,
		)
		req.Header.Set(RequestIDHeader, "invalid id with spaces")

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		requestID := recorder.Header().Get(RequestIDHeader)

		if recorder.Code != http.StatusOK {
			t.Fatalf("expected status %d, got %d", http.StatusOK, recorder.Code)
		}

		if _, err := uuid.Parse(requestID); err != nil {
			t.Fatalf("expected a newly generated valid UUID, got %q: %v", requestID, err)
		}

		if recorder.Body.String() != requestID {
			t.Errorf("expected body request ID %q, got %q", requestID, recorder.Body.String())
		}
	})

	t.Run("generates unique IDs for concurrent requests", func(t *testing.T) {
		router := setupRouter()

		const requestCount = 100

		requestIDs := make(chan string, requestCount)
		var wg sync.WaitGroup

		for range requestCount {
			wg.Go(func() {
				req := httptest.NewRequestWithContext(
					context.Background(), http.MethodGet, "/test", nil,
				)
				recorder := httptest.NewRecorder()

				router.ServeHTTP(recorder, req)

				if recorder.Code != http.StatusOK {
					t.Errorf("expected status %d, got %d", http.StatusOK, recorder.Code)
					return
				}

				requestIDs <- recorder.Header().Get(RequestIDHeader)
			})
		}

		wg.Wait()
		close(requestIDs)

		seen := make(map[string]struct{})

		for requestID := range requestIDs {
			if _, err := uuid.Parse(requestID); err != nil {
				t.Errorf("expected a valid UUID, got %q: %v", requestID, err)
				continue
			}

			if _, exists := seen[requestID]; exists {
				t.Errorf("duplicate request ID generated: %s", requestID)
			}

			seen[requestID] = struct{}{}
		}

		if len(seen) != requestCount {
			t.Errorf("expected %d unique IDs, got %d", requestCount, len(seen))
		}
	})
}
