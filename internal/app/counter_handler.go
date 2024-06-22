package app

import (
	"fmt"
	"net/http"
	"Diary/internal/config"
)

func CounterHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, Counter)
	for key, value := range config.Counter {
		fmt.Fprintf(w, "%s : %d\n", key, value)
	}
}
