package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

var applicationName string = "demo"

func main() {
	otel.SetTracerProvider(sdktrace.NewTracerProvider(sdktrace.WithSampler(sdktrace.ParentBased(sdktrace.AlwaysSample()))))
	app := gin.Default()
	app.ContextWithFallback = true
	app.Use(otelgin.Middleware(applicationName))

	app.GET("/", func(ctx *gin.Context) {
		traceID, spanID, isSampled := GetTraceInfo(ctx)
		fmt.Printf("traceID: %v; spanID: %v; isSampled: %v\n", traceID, spanID, isSampled)
	})

	go func() {
		if err := app.Run(":8080"); err != nil {
			log.Fatalln(err)
		}
	}()

	http.Get("http://localhost:8080/")

}

func GetTraceInfo(ctx context.Context) (traceID string, spanID string, isSampled bool) {
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
