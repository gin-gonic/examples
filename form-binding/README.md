# Form Binding Example with Gin

This example demonstrates how to use Gin to bind and validate form data in a Go web application. The sample creates a simple booking API that accepts form submissions, validates the fields, and parses date values with custom formats.

## Features

- Binds form data from client requests into Go structs.
- Validates form fields using Gin's binding and validation tags.
- Handles date fields with custom format parsing.
- Responds with JSON status and parsed data.

## Prerequisites

- [Go](https://golang.org/dl/) 1.13+ installed
- [Gin](https://github.com/gin-gonic/gin) web framework

Install Gin with:

```bash
go get -u github.com/gin-gonic/gin
```

## How It Works

The [`main.go`](main.go) file defines a `Booking` struct with form tags and validators:

```go
type Booking struct {
    Name     string     `form:"name" binding:"required"`
    CheckIn  *time.Time `form:"check_in" time_format:"2006-01-02" binding:"required"`
    CheckOut *time.Time `form:"check_out" time_format:"2006-01-02"`
}
```

- `form:"..."`: Maps incoming form field names to struct fields.
- `binding:"required"`: Marks a field as mandatory. If missing, binding fails and an error is returned.
- `time_format:"..."`: Tells Gin how to parse date strings into [`time.Time`](https://pkg.go.dev/time#Time).

The `/book` endpoint receives POST requests, binds form data into a Booking struct, and responds with JSON.

## Running the Example

1. Change into the `form-binding` directory:

   ```bash
   cd form-binding
   ```

2. Run the application:

   ```bash
   go run main.go
   ```

3. The server starts on [http://localhost:8080](http://localhost:8080).

## Example Usage

### Request

Send a POST request to `/book` with form data:

```bash
curl -X POST http://localhost:8080/book \
  -d "name=John Doe" \
  -d "check_in=2025-06-21" \
  -d "check_out=2025-06-25"
```

### Successful Response

```json
{
  "name": "John Doe",
  "check_in": "2025-06-21T00:00:00Z",
  "check_out": "2025-06-25T00:00:00Z"
}
```

### Validation Error Example

If a required field (e.g., `name` or `check_in`) is missing:

```json
{
  "error": "Key: 'Booking.Name' Error:Field validation for 'Name' failed on the 'required' tag"
}
```

## Key Points

- The `ShouldBind` method automatically performs validation based on struct tags.
- If validation fails, the handler returns HTTP 400 with an error message.
- Dates are parsed according to the `time_format` specified in struct tags (`YYYY-MM-DD`).

## References

- [Gin Documentation: Binding and Validation](https://gin-gonic.com/docs/examples/binding-and-validation/)
- [Go Documentation: time.Time](https://pkg.go.dev/time#Time)
