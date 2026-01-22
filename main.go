package main

import (
	"log"
	"net/http"
	"notes/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /note", handlers.Get)

	log.Printf("server running on port :3000")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
