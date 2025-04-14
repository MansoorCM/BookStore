package middleware

import (
	"log"
	"net/http"
)

func LoggingMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("request received %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("request handled successfully.")
	})
}
