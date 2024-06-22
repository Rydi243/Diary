package main

import (
	"Diary/internal/app"
	"fmt"
	"log"
	"net/http"
)

func main() {
	//r := mux.NewRouter()
	//app.SetHandlers(r)

	service := app.NewServise()

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", service); err != nil {
		log.Fatal(err)
	}
}
