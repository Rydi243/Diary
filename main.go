package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var Counter = make(map[string]int)         // Счётчик обращений
var stormaps = make(map[string]string, 10) // Ежедневник

// Запись в ежедневник
func Record(date string, affairs string, stormaps map[string]string) {
	stormaps[date] = affairs
	fmt.Println(time.Now(), "Новая запись в ежедневнике =>", date, affairs)
}

// Печать ежедневника
func PrintStoremaps(w http.ResponseWriter, stormaps map[string]string) {
	for key, value := range stormaps {
		fmt.Fprintf(w, "%s : %v\n", key, value)
		//fmt.Printf("%s : %v\n", key, value)
	}
}

// Обработка на запись в ежедневник
func diaryHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/diary" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	affairs := r.FormValue("affairs")
	date := r.FormValue("date")

	Record(date, affairs, stormaps)
}

// Счётчик запросов
func CounterMiddleware(callBack func(w http.ResponseWriter, r *http.Request), Counter map[string]int) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		Counter[r.URL.String()]++
		fmt.Println(Counter)
		callBack(w, r)
	}
}

// Аутентификация
func AuthCheck(login string, password string) bool {
	return login == "admin" && password == "admin"
}

// Аутентификация
func AuthCheckMiddleware(n http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := r.URL.Query().Get("log")
		pas := r.URL.Query().Get("pas")

		if !AuthCheck(log, pas) {
			fmt.Fprintln(w, "Error!!")
			w.WriteHeader(http.StatusForbidden)
			return
		}

		n.ServeHTTP(w, r)
	})
}

// Приветствие админа и печать ежедневника всего
func HelloAdminFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello, admin!")
	PrintStoremaps(w, stormaps)

}

// Печать счётчика обращений
func CounterHandler(w http.ResponseWriter, Counter map[string]int) {
	fmt.Fprintln(w, Counter)
}

func setHandlers(r *mux.Router) {

	r.HandleFunc("/diary", CounterMiddleware(diaryHandler, Counter))

	admin := r.PathPrefix("/admin").Subrouter()
	admin.Use(AuthCheckMiddleware)

	admin.HandleFunc("/allaffairs", CounterMiddleware(HelloAdminFunc, Counter))
	admin.HandleFunc("/counter", func(w http.ResponseWriter, r *http.Request) {
		CounterHandler(w, Counter)
	})
}

func main() {
	r := mux.NewRouter()
	setHandlers(r)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
