// upload-limit-bytes/main.go
// Example: Restrict file upload size using Gin + http.MaxBytesReader
//
// This example shows how to use Gin together with http.MaxBytesReader to safely and *strictly*
// limit the maximum size of uploaded files. It returns a custom error message if the file is too large.

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	MaxUploadSize = 1 << 20 // 1MB (can set to other sizes; used for tests and demos)
)

func uploadHandler(c *gin.Context) {
	// Wrap the body reader so only MaxUploadSize bytes are allowed
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MaxUploadSize)

	// Parse multipart form (limit maxMemory for demonstration, not for size restriction)
	if err := c.Request.ParseMultipartForm(MaxUploadSize); err != nil {
		// Check for *http.MaxBytesError for upload size limit exceeded (portable way)
		if _, ok := err.(*http.MaxBytesError); ok {
			c.JSON(http.StatusRequestEntityTooLarge, gin.H{
				"error": fmt.Sprintf("file too large (max: %d bytes)", MaxUploadSize),
			})
			return
		}
		// Other form errors
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get uploaded file; parameter name is "file"
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file form required"})
		return
	}

	defer func() { _ = file.Close() }()
	// For demonstration we ignore saving the file, but could copy to disk or buffer
	// _, err = io.Copy(io.Discard, file) // Uncomment if you want to read all bytes

	c.JSON(http.StatusOK, gin.H{
		"message": "upload successful",
	})
}

// setupRouter creates and returns the Gin router
func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/upload", uploadHandler)
	return r
}

func main() {
	r := setupRouter()
	// Run Gin server on :8080
	_ = r.Run(":8080")
}
