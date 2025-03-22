# New Relic Integration with Gin

This example demonstrates how to integrate New Relic with a Gin web application.

## Prerequisites

- Go 1.16 or later
- A New Relic account and API key

## Setup

1. Clone the repository and navigate to the `new_relic` directory:

```sh
git clone https://github.com/your-repo/gin-examples.git
cd gin-examples/new_relic
```

1. Install the dependencies:

```sh
go mod tidy
```

1. Set the environment variables for your New Relic application:

```sh
export NEW_RELIC_APP_NAME="YourAppName"
export NEW_RELIC_LICENSE_KEY="YourNewRelicLicenseKey"
```

## Running the Application

To run the application, use the following command:

```sh
go run main.go
```

The application will start a web server on `http://localhost:8080`. You can access it in your browser to see the "Hello World!" message.

## Code Overview

The main components of the application are:

- **Gin Router**: The web framework used to handle HTTP requests.
- **New Relic Application**: The New Relic agent used to monitor the application.

### main.go

```go
package main

import (
  "log"
  "net/http"

  "github.com/gin-gonic/gin"
  "github.com/newrelic/go-agent/v3/integrations/nrgin"
  newrelic "github.com/newrelic/go-agent/v3/newrelic"
)

func main() {
  router := gin.Default()

  app, err := newrelic.NewApplication(
    newrelic.ConfigAppName("MyApp"),
    newrelic.ConfigFromEnvironment(),
  )
  if err != nil {
    log.Fatalf("failed to make new_relic app: %v", err)
  }

  router.Use(nrgin.Middleware(app))
  router.GET("/", func(c *gin.Context) {
    c.String(http.StatusOK, "Hello World!\n")
  })
  router.Run()
}
```

## License

This project is licensed under the MIT License. See the [LICENSE](../LICENSE) file for details.

## Contributing

Contributions are welcome! Please see the [CONTRIBUTING](../CONTRIBUTING.md) file for more information.

## Acknowledgments

- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [New Relic Go Agent](https://github.com/newrelic/go-agent)
