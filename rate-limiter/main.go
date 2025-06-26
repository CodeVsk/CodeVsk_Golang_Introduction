package main

import (
	"net/http"

	"golang.org/x/time/rate"
)

func rateLimit(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		limiter := rate.NewLimiter(1, 50)
		if limiter.Allow() {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)

			return
		}

		h.ServeHTTP(w, r)
	})
}

func main() {
	http.Handle("/", rateLimit(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { 
		w.Write([]byte("Capybara!"))
	})))

	http.ListenAndServe(":3030", nil)
}