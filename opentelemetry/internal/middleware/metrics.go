package middleware

import (
	"net/http"
	"os"

	"github.com/CodeVsk/CodeVsk_Golang_Introduction/opentelemetry/pkg/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

var httpRequestCounter metric.Int64Counter

func InitMetrics() error {
	var err error

	httpRequestCounter, err = otel.Meter.Int64Counter(
		"http_request_total_batata",
		metric.WithDescription("Counts total HTTP requests"),
	)
	if err != nil {
		return err
	}

	return nil
}

func HTTPRequestCounter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		next.ServeHTTP(w, r)

		attrs := []attribute.KeyValue{
			attribute.String("method", r.Method),
			attribute.String("path", r.URL.Path),
			//attribute.Int("status_code", r.Response.StatusCode),
			attribute.String("service", os.Getenv("SERVICE_NAME")),
		}

		httpRequestCounter.Add(r.Context(), 1, metric.WithAttributes(attrs...))
	})
}
