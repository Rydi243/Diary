package app

import (
	"Diary/internal/config"
	"fmt"
	"net/http"
)

func CounterHandler(w http.ResponseWriter, r *http.Request) {
	for key, value := range config.Counter {
		fmt.Fprintf(w, "%s : %d\n", key, value)
	}
}
