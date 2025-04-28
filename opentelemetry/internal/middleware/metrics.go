package middleware

import (
	"net/http"
	"os"

	"github.com/CodeVsk/CodeVsk_Golang_Introduction/opentelemetry/pkg/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

var httpRequestCounter metric.Int64Counter
var httpActiveConnections metric.Int64Counter

func InitMetrics() error {
	var err error

	httpRequestCounter, err = otel.Meter.Int64Counter(
		"http_request_total",
		metric.WithDescription("Counts total HTTP requests"),
	)
	if err != nil {
		return err
	}

	httpActiveConnections, err = otel.Meter.Int64Counter(
		"http_request_total_active",
		metric.WithDescription("Counts total active HTTP requests"),
	)
	if err != nil {
		return err
	}

	return nil
}

func HTTPRequestCounter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		httpActiveConnections.Add(r.Context(), 1, metric.WithAttributes(
			attribute.String("method", r.Method),
			attribute.String("path", r.URL.Path),
			attribute.String("service", os.Getenv("SERVICE_NAME")),
		))

		defer func() {
			httpActiveConnections.Add(r.Context(), -1, metric.WithAttributes(
				attribute.String("method", r.Method),
				attribute.String("path", r.URL.Path),
				attribute.String("service", os.Getenv("SERVICE_NAME")),
			))
		}()

		next.ServeHTTP(w, r)

		attrs := []attribute.KeyValue{
			attribute.String("method", r.Method),
			attribute.String("path", r.URL.Path),
			attribute.String("service", os.Getenv("SERVICE_NAME")),
		}

		httpRequestCounter.Add(r.Context(), 1, metric.WithAttributes(attrs...))
	})
}
