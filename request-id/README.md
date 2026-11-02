# Request ID Middleware

This example demonstrates how to implement Request ID middleware in [Gin](https://github.com/gin-gonic/gin) to uniquely identify and track incoming HTTP requests.

Request IDs help correlate application logs, simplify debugging, and track requests across multiple services in distributed applications.

## Getting Started

### Install Dependencies

```bash
go mod tidy
```

### Run the Example

```bash
go run .
```

The server starts on port `8080`.

## Usage

### 1. Generate a Request ID

Send a request without an `X-Request-ID` header:

```bash
curl -i http://localhost:8080/hello
```

The middleware generates a UUID and includes it in the response header and JSON response.

Example response:

```http
HTTP/1.1 200 OK
X-Request-Id: 550e8400-e29b-41d4-a716-446655440000
Content-Type: application/json

{
  "message": "Hello from Gin!",
  "request_id": "550e8400-e29b-41d4-a716-446655440000"
}
```

### 2. Reuse an Existing Request ID

Send a request with an existing valid UUID:

```bash
curl -i \
  -H "X-Request-ID: 550e8400-e29b-41d4-a716-446655440000" \
  http://localhost:8080/hello
```

The middleware validates the supplied UUID and reuses it if valid.

If the header is missing or contains an invalid UUID, the middleware generates a new one.

### 3. Access Request ID in Handlers

The middleware stores the Request ID in the Gin context, allowing handlers to retrieve it:

```go
requestID, exists := c.Get("RequestID")
if exists {
    log.Printf("[%s] Processing request", requestID)
}
```

This allows application logs to be correlated using the same Request ID.

## Running Tests

Run the unit tests:

```bash
go test -v ./...
```

Run tests with the Go race detector:

```bash
go test -race ./...
```

## Why Use Request IDs?

Request IDs are useful for:

- **Debugging:** Identify logs associated with a specific request.
- **Log Correlation:** Group related log entries together.
- **Microservices:** Propagate request identifiers across service boundaries.
- **Production Support:** Help investigate failed API requests.
- **Observability:** Improve log correlation with monitoring and tracing systems.

## Important Considerations

- Validate incoming Request IDs before accepting them.
- Generate a new ID when the incoming value is missing or invalid.
- Avoid storing request-specific data in global variables.
- Propagate Request IDs to downstream services when appropriate.
- Establish a trust policy for incoming IDs, especially when requests originate from untrusted clients.
- Request IDs are for correlation and should not be used for authentication or authorization.