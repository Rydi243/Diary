package middleware

import (
	"net/http"
	"fmt"
	"time"
)


func LogMiddlleware(n http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[%s] [Запрос: %s] [url:%s]\n", time.Now().Format("02.01.2006 15:04:01"), r.Method, r.URL)
		n.ServeHTTP(w, r)
	})

}