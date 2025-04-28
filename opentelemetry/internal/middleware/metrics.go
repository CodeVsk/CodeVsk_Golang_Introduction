package middleware

import (
	"net/http"
	"os"

	"github.com/CodeVsk/CodeVsk_Golang_Introduction/opentelemetry/pkg/otel"
	"github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

var httpRequestCounter metric.Int64Counter
var httpActiveConnections metric.Int64Counter
var httpRequestDuration metric.Float64Histogram

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

	httpRequestDuration, err = otel.Meter.Float64Histogram(
		"http_request_duration_seconds",
		metric.WithDescription("Measures the duration of HTTP requests"),
	)
	if err != nil {
		return err
	}

	return nil
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func HTTPRequestCounter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timer := prometheus.NewTimer(prometheus.ObserverFunc(func(v float64) {
			httpRequestDuration.Record(r.Context(), v, metric.WithAttributes(
				attribute.String("method", r.Method),
				attribute.String("path", r.URL.Path),
				attribute.String("service", os.Getenv("SERVICE_NAME")),
			))
		}))
		defer timer.ObserveDuration()

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

		rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(rw, r)

		attrs := []attribute.KeyValue{
			attribute.String("method", r.Method),
			attribute.String("path", r.URL.Path),
			attribute.Int("status_code", rw.statusCode),
			attribute.String("service", os.Getenv("SERVICE_NAME")),
		}

		httpRequestCounter.Add(r.Context(), 1, metric.WithAttributes(attrs...))
	})
}
