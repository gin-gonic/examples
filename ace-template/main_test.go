package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAceTemplate(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequestWithContext(context.Background(), "GET", "/ace-example", nil)
	router.ServeHTTP(w, req)

	expected := "<!DOCTYPE html><html><body><h1>Hello Ace & Gin</h1></body></html>"

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expected, w.Body.String())
}
