package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var stormaps = make(map[string]string, 10)

func Record(date string, affairs string, stormaps map[string]string) {
	stormaps[date] = affairs
	fmt.Println("Новая запись в ежедневнике =>", date, affairs)
}

func PrintStoremaps(w http.ResponseWriter, stormaps map[string]string) {
	for key, value := range stormaps {
		fmt.Fprintf(w, "%s : %v\n", key, value)
		fmt.Printf("%s : %v\n", key, value)
	}
}

func diaryHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/diary" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
	}

	affairs := r.FormValue("affairs")
	date := r.FormValue("date")

	Record(date, affairs, stormaps)
}

func affairsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/allaffairs" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
	}
	PrintStoremaps(w, stormaps)
}

func setHandlers(r *mux.Router) {
	r.HandleFunc("/diary", diaryHandler)
	r.HandleFunc("/allaffairs", affairsHandler)
}

func main() {
	r := mux.NewRouter()
	setHandlers(r)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
