package app

import (
	"net/http"
	"Diary/internal/config"
)


// Обработка на запись в ежедневник
func DiaryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	affairs := r.URL.Query().Get("affairs")
	date := r.URL.Query().Get("date")

	Record(date, affairs, config.Stormaps)
}