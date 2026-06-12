# Gin Example: Restrict File Upload Size with http.MaxBytesReader

This example demonstrates how to use [Gin](https://github.com/gin-gonic/gin) and Go's `http.MaxBytesReader` to strictly limit uploaded file size and return custom error messages when the size limit is exceeded.

## Features

- REST API - POST `/upload` endpoint for file upload.
- Server-side enforcement of maximum file size (default: **1MB**).
- Immediate error response (`413 Request Entity Too Large` with custom JSON) when file is too large.
- Well-commented code for clarity.
- Unit + integration tests verifying successful uploads, oversized file rejection, and edge cases.

## How It Works

- The core implementation wraps the incoming request body using Go's [`http.MaxBytesReader`](https://pkg.go.dev/net/http#MaxBytesReader), which limits how many bytes the server will read.
- If a client uploads a file exceeding this limit, parsing fails and Gin responds with a **custom error message**.
- See [`main.go`](./main.go) for details and comments.

## API Usage

### Start the server

```bash
go run main.go
```

Server listens on port **8080**.

### Upload a file

**cURL valid file (within 1MB):**

```bash
curl -F "file=@smallfile.bin" http://localhost:8080/upload
# Response: {"message":"upload successful"}
```

**cURL oversized file (e.g., 2MB):**

```bash
curl -F "file=@bigfile.bin" http://localhost:8080/upload
# Response: {"error":"file too large (max: 1048576 bytes)"}
```

### Error cases

- Missing file: `{"error":"file form required"}` (status 400)  
- File too large: `{"error":"file too large (max: 1048576 bytes)"}` (status 413)

## Testing

This example includes unit/integration tests in [`main_test.go`](./main_test.go):

```bash
go test
```

- **TestUploadWithinLimit**: uploads a file under the limit, expects success.
- **TestUploadOverLimit**: uploads a file over the limit, expects a custom error.
- **TestUploadMissingFile**: no file uploaded, expects validation error.

## Code Reference

- [`main.go`](./main.go): Gin server setup, upload handler with file size enforcement.
- [`main_test.go`](./main_test.go): Contains all unit and integration tests.

## Modify the Limit

Change `MaxUploadSize` constant in [`main.go`](./main.go) and [`main_test.go`](./main_test.go) to test other limits.

---

**License:** MIT  
**Author:** Gin Example Contributors
