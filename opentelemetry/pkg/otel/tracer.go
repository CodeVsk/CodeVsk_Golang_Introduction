package otel

import (
	"context"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

var Tracer = otel.Tracer(os.Getenv("TRACER_SERVICE_NAME"))

type OpenTelemetry struct {
	Exporter trace.SpanExporter
	Trace    *trace.TracerProvider
	Name     string
}

func NewTracer(name string) *OpenTelemetry {
	return &OpenTelemetry{
		Name: name,
	}
}

func (o *OpenTelemetry) UseJaeger(url string) *OpenTelemetry {
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		panic(err)
	}

	o.Exporter = exporter

	return o
}

func (o *OpenTelemetry) UseHttp(url string) *OpenTelemetry {
	exporter, err := otlptracehttp.New(context.Background(),
		otlptracehttp.WithEndpointURL(url),
		otlptracehttp.WithInsecure())
	if err != nil {
		panic(err)
	}

	o.Exporter = exporter

	return o
}

func (o *OpenTelemetry) UseStdout(url string) *OpenTelemetry {
	exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		panic(err)
	}

	o.Exporter = exporter

	return o
}

func (o *OpenTelemetry) Init() (*trace.TracerProvider, error) {
	tracerProvider := trace.NewTracerProvider(
		trace.WithBatcher(o.Exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(o.Name),
		)),
	)

	otel.SetTracerProvider(tracerProvider)

	return tracerProvider, nil
}
