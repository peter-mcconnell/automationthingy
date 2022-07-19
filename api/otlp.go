package api

import (
	"errors"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/trace"
)

func (s *Server) OltpInitialize() error {
	url := os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	if url == "" {
		return errors.New("OTEL_EXPORTER_OTLP_ENDPOINT must be set")
	}
	s.logger.Info("initializing oltp configuration")
	exporter, err := otlptracegrpc.New(s.ctx)
	if err != nil {
		s.logger.Fatalf("failed to initialize oltp exporter: %v", err)
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
	)

	defer func() { _ = tp.Shutdown(s.ctx) }()

	otel.SetTracerProvider(tp)

	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)

	return nil
}
