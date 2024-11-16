package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/getsentry/sentry-go"
	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/joho/godotenv"
)

func EnvVariable(key string) string {
  err := godotenv.Load(".env")

  if err != nil {
    log.Fatalf("Error loading .env file")
  }

  return os.Getenv(key)
}

func main() {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: EnvVariable("SENTRY_DSN"),
		EnableTracing: true,
		
    TracesSampleRate: 1.0,
    TracesSampler: sentry.TracesSampler(func(ctx sentry.SamplingContext) float64 {
    if ctx.Span.Name == "GET /health" {
    	return 0.0
    }
    	return 1.0
		}),
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}

	sentryHandler := sentryhttp.New(sentryhttp.Options{})

	http.HandleFunc("/test", sentryHandler.HandleFunc(func (w http.ResponseWriter, r *http.Request) {
		numbers := []int{1, 2, 3}
    fmt.Println(numbers[5])

		w.WriteHeader(500)
	}))

	http.HandleFunc("/test-2", sentryHandler.HandleFunc(func (w http.ResponseWriter, r *http.Request) {
		sentry.AddBreadcrumb(&sentry.Breadcrumb{Category:"test", Level: sentry.LevelInfo, Message: "test endpoint requested!"})

		numbers := []int{1, 2, 3}
    fmt.Println(numbers[5])

		w.WriteHeader(500)
	}))

	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}