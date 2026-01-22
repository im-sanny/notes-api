package main

import (
	"log"
	"net/http"
	"notes/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /note", handlers.Get)
	mux.HandleFunc("GET /note/{id}", handlers.GetById)
	mux.HandleFunc("POST /note", handlers.Post)
	mux.HandleFunc("PUT /note/{id}", handlers.Post)
	mux.HandleFunc("PATCH /note/{id}", handlers.Patch)
	mux.HandleFunc("DELETE /note/{id}", handlers.Delete)

	log.Printf("server running on port :3000")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
