package main

import (
	"Diary/internal/app"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//var Counter = make(map[string]int)         // Счётчик обращений
//var Stormaps = make(map[string]string, 10) // Ежедневник

func main() {
	r := mux.NewRouter()
	app.SetHandlers(r)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
