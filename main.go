package main

import (
	"log"
	"net/http"

	"go_templ/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.ListAllData)

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
