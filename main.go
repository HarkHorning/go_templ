package main

import (
	"log"
	"main/handlers"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handlers.ListAllData)
	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	component := components.Base()
	// 	component.Render(context.Background(), os.Stdout)
	// 	fmt.Println(component)

	// })

	// log.Println("Listening on port 8080")
	// log.Fatal(http.ListenAndServe(":8080", mux))

	// fmt.Println("Started http server on port: 8080")
	// http.ListenAndServe(":8080", nil)

}
