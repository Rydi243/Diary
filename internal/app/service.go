package app

import (
	"Diary/internal/app/middleware"
	"Diary/internal/config"
	"net/http"

	"github.com/gorilla/mux"
)

type Service struct {
	Port   string
	router *mux.Router
}

func NewServise() Service {
	tmp := Service{
		Port:   config.ServPort,
		router: mux.NewRouter(),
	}
	tmp.SetHandlers()
	return tmp
}

func (s Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s Service) SetHandlers() {
	s.router.Use(middleware.CounterMiddleware)
	s.router.Use(middleware.LogMiddlleware)
	s.router.HandleFunc("/diary", DiaryHandler)

	admin := s.router.PathPrefix("/admin").Subrouter()
	admin.Use(middleware.AuthCheckMiddleware)

	admin.HandleFunc("/allaffairs", HelloAdminFunc)
	admin.HandleFunc("/counter", CounterHandler)

}
