# CORS Middleware Example

This example demonstrates how to enable **Cross-Origin Resource Sharing (CORS)** in a Gin web server.  
It also provides basic RESTful endpoints (`GET`, `POST`, `PUT`, `DELETE`) to verify that cross-origin requests are properly handled.

---

## Features
- Custom CORS middleware allowing all origins (`*`)
- RESTful routes for `/ping` and `/data/:id`
- JSON responses with timestamps
- Handles `OPTIONS` preflight requests
- Includes logging for incoming requests

---

## Run

```bash
go run main.go
