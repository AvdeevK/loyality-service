package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func run(r chi.Router) error {
	return http.ListenAndServe(":8080", r)
}

func main() {
	srv := chi.NewRouter()
	if err := run(srv); err != nil {
		log.Fatal(err)
	}
}
