package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	newrelic "github.com/newrelic/go-agent/v3/newrelic"
)

func main() {
	router := gin.Default()

	// cfg := newrelic.NewConfig(os.Getenv("APP_NAME"), os.Getenv("NEW_RELIC_API_KEY"))
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("MyApp"),
		newrelic.ConfigFromEnvironment(),
		newrelic.ConfigDebugLogger(os.Stdout),
		newrelic.ConfigAppLogForwardingEnabled(true),
		newrelic.ConfigCodeLevelMetricsEnabled(true),
		newrelic.ConfigCodeLevelMetricsPathPrefixes("go-agent/v3"),
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
