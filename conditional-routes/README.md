# Conditional Routes

This example demonstrates how to enable or disable routes using command-line flags and environment variables.

## Run

```bash
go run main.go
```

Enable admin routes:

```bash
go run main.go --admin
```

Enable metrics endpoint:

```bash
ENABLE_METRICS=true go run main.go
```