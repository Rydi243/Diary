package app

import (
	"github.com/gorilla/mux"
	"Diary/internal/app/middleware"
)

func SetHandlers(r *mux.Router) {
	r.Use(middleware.CounterMiddleware)
	r.Use(middleware.LogMiddlleware)
	r.HandleFunc("/diary", DiaryHandler)

	admin := r.PathPrefix("/admin").Subrouter()
	admin.Use(middleware.AuthCheckMiddleware)

	admin.HandleFunc("/allaffairs", HelloAdminFunc)
	admin.HandleFunc("/counter", CounterHandler)

}