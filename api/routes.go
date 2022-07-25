package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/trace"
)

type route struct {
	pattern string
	handler http.Handler
}

func (s *Server) parseApiRequest(url string) ApiRequest {
	uriParts := strings.Split(url, "/")[2:]
	var (
		id  uuid.UUID
		err error
	)
	if len(uriParts) > 1 {
		id, err = uuid.Parse(uriParts[2])
		if err != nil {
			panic(err)
		}
	}
	return ApiRequest{
		apiVersion:    uriParts[0],
		resource:      uriParts[1],
		sub_resources: strings.Split(uriParts[1], "_")[1:],
		id:            id,
	}
}

func (s *Server) HandleFunc(pattern string, handler http.Handler) {
	oltphandler := otelhttp.NewHandler(
		handler,
		pattern,
	)
	s.Mux.Handle(pattern, oltphandler)
}

func (s *Server) addRoutes() error {
	var routes = []route{
		newRoute("/api/v1/config", s.apiV1Config),
		newRoute("/api/v1/executor_local/", s.apiV1ExecutorLocal),
		newRoute("/api/v1/executor_kubernetesjob/", s.apiV1ExecutorKubernetesjob),
		newRoute("/", func(w http.ResponseWriter, r *http.Request) {
			http.NotFound(w, r)
		}),
	}
	for _, route := range routes {
		s.HandleFunc(route.pattern, route.handler)
	}
	ctx := context.Background()

	// Configure a new exporter using environment variables for sending data to Honeycomb over gRPC.
	exporter, err := otlptracegrpc.New(ctx)
	if err != nil {
		log.Fatalf("failed to initialize exporter: %v", err)
	}

	// Create a new tracer provider with a batch span processor and the otlp exporter.
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
	)

	// Handle shutdown errors in a sensible manner where possible
	defer func() { _ = tp.Shutdown(ctx) }()

	// Set the Tracer Provider global
	otel.SetTracerProvider(tp)

	// Register the trace context and baggage propagators so data is propagated across services/processes.
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)
	// Implement an HTTP Handler func to be instrumented
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World")
	})

	// Initialize HTTP handler instrumentation
	otelHandler := otelhttp.NewHandler(handler, "hello")
	s.Mux.Handle("/hello", otelHandler)
	return nil
}

func newRoute(pattern string, handler http.HandlerFunc) route {
	return route{pattern, handler}
}
