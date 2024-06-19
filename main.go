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
	//fmt.Println(time.Now(), "Новая запись в ежедневнике =>", date, affairs)
}

// Печать ежедневника
func PrintStoremaps(w http.ResponseWriter, stormaps map[string]string) {
	for key, value := range stormaps {
		fmt.Fprintf(w, "%s : %v\n", key, value)
	}
}

// Обработка на запись в ежедневник
func diaryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	affairs := r.URL.Query().Get("affairs")
	date := r.URL.Query().Get("date")

	Record(date, affairs, stormaps)
}

// Счётчик запросов
func CounterMiddleware(n http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Counter[r.URL.String()]++
		n.ServeHTTP(w, r)
	})
}

func CounterHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, Counter)
	for key, value := range Counter {
		fmt.Fprintf(w, "%s : %d\n", key, value)
	}
}

func LogMiddlleware(n http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(time.Now())
		n.ServeHTTP(w, r)
	})

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

// Навешать мидлвару через Use - сделал CounterMiddleware/ Переписать мидлвару под сигнатуру  (дз) - сделал CounterMiddleware!!! //  логирование сделать как мидлвару!!!!
func setHandlers(r *mux.Router) {
	r.Use(CounterMiddleware)
	r.Use(LogMiddlleware)
	r.HandleFunc("/diary", diaryHandler)

	admin := r.PathPrefix("/admin").Subrouter()
	admin.Use(AuthCheckMiddleware)

	admin.HandleFunc("/allaffairs", HelloAdminFunc)
	admin.HandleFunc("/counter", CounterHandler)

}

func main() {
	r := mux.NewRouter()
	setHandlers(r)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
