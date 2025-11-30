// upload-limit-bytes/main_test.go
// Unit & integration tests for Gin file upload with http.MaxBytesReader limit

package main

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

// Use setupRouter and uploadHandler from the main package.
// (Removed local setupRouter to avoid redeclaration; tests use main's setupRouter().)

// Helper: create multipart body with given size
func createMultipartBody(fieldName string, size int) (contentType string, bodyBytes []byte) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	fw, _ := w.CreateFormFile(fieldName, "test.bin")
	_, _ = io.CopyN(fw, bytes.NewReader(make([]byte, size)), int64(size)) // Write 'size' bytes
	_ = w.Close()
	return w.FormDataContentType(), b.Bytes()
}

func TestUploadWithinLimit(t *testing.T) {
	router := setupRouter()
	// Prepare file just under max size
	// Reduce file size so the total request (including multipart headers/overhead) is safely below the limit.
	contentType, body := createMultipartBody("file", MaxUploadSize-10*1024)
	req, _ := http.NewRequest("POST", "/upload", bytes.NewReader(body))
	req.Header.Set("Content-Type", contentType)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	require.Equal(t, http.StatusOK, resp.Code)
	var result map[string]string
	err := json.Unmarshal(resp.Body.Bytes(), &result)
	require.NoError(t, err)
	require.Equal(t, "upload successful", result["message"])
}

func TestUploadOverLimit(t *testing.T) {
	router := setupRouter()
	contentType, body := createMultipartBody("file", MaxUploadSize+100)
	req, _ := http.NewRequest("POST", "/upload", bytes.NewReader(body))
	req.Header.Set("Content-Type", contentType)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	require.Equal(t, http.StatusRequestEntityTooLarge, resp.Code)
	var result map[string]string
	err := json.Unmarshal(resp.Body.Bytes(), &result)
	require.NoError(t, err)
	require.Contains(t, result["error"], "file too large")
}

func TestUploadMissingFile(t *testing.T) {
	router := setupRouter()
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	_ = w.Close()
	req, _ := http.NewRequest("POST", "/upload", &b)
	req.Header.Set("Content-Type", w.FormDataContentType())
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	require.Equal(t, http.StatusBadRequest, resp.Code)
	var result map[string]string
	err := json.Unmarshal(resp.Body.Bytes(), &result)
	require.NoError(t, err)
	require.Contains(t, result["error"], "file form required")
}
