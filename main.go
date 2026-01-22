package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	log.Printf("server running on port :3000")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
