# Use OTEL in Gin

This example shows a minimum case to use otel in Gin.

To run the example:
```bash
go run main.go
```

In stander output like this:
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


