package otel

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	trace_tracer "go.opentelemetry.io/otel/trace"
)

var Tracer trace_tracer.Tracer

type OtelTracer struct {
	Exporter trace.SpanExporter
	Name     string
}

func NewTracer(name string) *OtelTracer {
	return &OtelTracer{
		Name: name,
	}
}

func (o *OtelTracer) UseJaeger(url string) *OtelTracer {
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		panic(err)
	}

	o.Exporter = exporter

	return o
}

func (o *OtelTracer) UseHttp(url string) *OtelTracer {
	exporter, err := otlptracehttp.New(context.Background(),
		otlptracehttp.WithEndpointURL(url),
		otlptracehttp.WithInsecure())
	if err != nil {
		panic(err)
	}

	o.Exporter = exporter

	return o
}

func (o *OtelTracer) UseStdout(url string) *OtelTracer {
	exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		panic(err)
	}

	o.Exporter = exporter

	return o
}

func (o *OtelTracer) Init() (*trace.TracerProvider, error) {
	tracerProvider := trace.NewTracerProvider(
		trace.WithBatcher(o.Exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(o.Name),
		)),
	)

	otel.SetTracerProvider(tracerProvider)

	Tracer = tracerProvider.Tracer(o.Name)

	return tracerProvider, nil
}
