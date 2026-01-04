package main

import (
	"fmt"
	"net/http"
	"time"
)

func logRequest(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler(w, r)
		fmt.Printf(
			"%s %s %s\n",
			r.Method,
			r.URL.Path,
			time.Since(start),
		)
	}
}
