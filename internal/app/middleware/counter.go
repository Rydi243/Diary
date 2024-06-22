package middleware

import (
	"net/http"
	"Diary/internal/config"
)


// Счётчик запросов
func CounterMiddleware(n http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		config.Counter[r.URL.String()]++
		n.ServeHTTP(w, r)
	})
}