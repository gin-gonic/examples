# OTEL Integration with Gin

This project demonstrates how to integrate OpenTelemetry (OTEL) with the Gin web framework in Go. It provides a minimal example to get you started with tracing HTTP requests in a Gin application using OTEL.

## Prerequisites

- Go 1.23 or later
- Git

## Installation

### Clone the repository

```bash
git clone https://github.com/your-repo/otel-gin-example.git
cd otel-gin-example
```

### Install the dependencies

```bash
go mod tidy
```

## Running the Example

To run the example application, use the following command:

```bash
go run main.go
```

This will start the Gin server on `http://localhost:8080`.

## Example Output

When you run the example, you should see output similar to the following in the terminal:

```bash
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /                         --> main.main.func1 (4 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8080
traceID: 00477d7b56b757d0581328ef21d17271; spanID: 9d05e83c0c188a16; isSampled: true
[GIN] 2024/10/15 - 11:44:32 | 200 |      31.209µs |             ::1 | GET      "/"
```

## Project Structure

- `main.go`: The main application file that sets up the Gin server and OTEL middleware.
- `go.mod`: The Go module file that lists the project dependencies.
- `go.sum`: The Go checksum file for dependencies.

## Dependencies

This project uses the following dependencies:

- [Gin](https://github.com/gin-gonic/gin) v1.10.0: A web framework written in Go.
- [OpenTelemetry](https://github.com/open-telemetry/opentelemetry-go) v1.35.0: A set of APIs, libraries, agents, and instrumentation to provide observability.
- [OTEL Gin Middleware](https://github.com/open-telemetry/opentelemetry-go-contrib/tree/main/instrumentation/github.com/gin-gonic/gin/otelgin) v0.60.0: Middleware for integrating OTEL with Gin.

## Configuration

The example sets `ContextWithFallback = true` to use OTEL in Gin. This configuration is necessary for the example to work but may not be ideal for production use. Adjust the configuration as needed for your use case.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please see the [CONTRIBUTING](CONTRIBUTING.md) file for more information.

## Acknowledgments

- [Gin](https://github.com/gin-gonic/gin)
- [OpenTelemetry](https://github.com/open-telemetry/opentelemetry-go)
- [OTEL Gin Middleware](https://github.com/open-telemetry/opentelemetry-go-contrib/tree/main/instrumentation/github.com/gin-gonic/gin/otelgin)
