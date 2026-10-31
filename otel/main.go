package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

var applicationName = "demo"

func main() {
	otel.SetTracerProvider(
		sdktrace.NewTracerProvider(
			sdktrace.WithSampler(sdktrace.ParentBased(sdktrace.AlwaysSample())),
		),
	)
	app := gin.Default()
	app.ContextWithFallback = true
	app.Use(otelgin.Middleware(applicationName))

	app.GET("/", func(ctx *gin.Context) {
		traceID, spanID, isSampled := GetTraceInfo(ctx)
		log.Printf("traceID: %v; spanID: %v; isSampled: %v\n", traceID, spanID, isSampled)
	})

	go func() {
		if err := app.Run(":8080"); err != nil {
			log.Fatalln(err)
		}
	}()

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		"http://localhost:8080/",
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}

func GetTraceInfo(ctx context.Context) (traceID, spanID string, isSampled bool) {
	spanCtx := trace.SpanContextFromContext(ctx)

	if spanCtx.HasTraceID() {
		traceID = spanCtx.TraceID().String()
	}
	if spanCtx.HasSpanID() {
		spanID = spanCtx.SpanID().String()
	}

	isSampled = spanCtx.IsSampled()

	return traceID, spanID, isSampled
}
