package otel

import (
	"time"

	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"

	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"

	metric_meter "go.opentelemetry.io/otel/metric"
)

var Meter metric_meter.Meter

var PeriodicReaderDefault = 3

type OtelMeter struct {
	Reader         metric.Reader
	Name           string
	PeriodicReader int
}

func NewMeter(name string) *OtelMeter {
	return &OtelMeter{
		Name:           name,
		PeriodicReader: PeriodicReaderDefault,
	}
}

func (o *OtelMeter) UseStdout() *OtelMeter {
	exporter, err := stdoutmetric.New()
	if err != nil {
		panic(err)
	}

	o.Reader = metric.NewPeriodicReader(exporter, metric.WithInterval(time.Duration(o.PeriodicReader)*time.Second))

	return o
}

func (o *OtelMeter) UsePrometheus() *OtelMeter {
	exporter, err := prometheus.New()
	if err != nil {
		panic(err)
	}

	o.Reader = exporter

	return o
}

func (o *OtelMeter) SetPeriodicReader(seconds int) *OtelMeter {
	o.PeriodicReader = seconds

	return o
}

func (o *OtelMeter) Init() (*metric.MeterProvider, *OtelMeter, error) {
	meterProvider := metric.NewMeterProvider(
		metric.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(o.Name),
			semconv.ServiceVersion("0.1.0"),
		)), metric.WithReader(o.Reader))

	Meter = meterProvider.Meter(o.Name)

	return meterProvider, o, nil
}
