package middleware

import (
	"net/http"
	"Diary/internal/pkg/auth"
	"fmt"
)

// Аутентификация
func AuthCheckMiddleware(n http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := r.URL.Query().Get("log")
		pas := r.URL.Query().Get("pas")

		if !auth.AuthCheck(log, pas) {
			fmt.Fprintln(w, "Error!!")
			w.WriteHeader(http.StatusForbidden)
			return
		}

		n.ServeHTTP(w, r)
	})
}