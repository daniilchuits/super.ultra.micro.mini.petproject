package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {

	r := chi.NewRouter()

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("worker ok"))
	}) // endpoints for worker

	log.Fatal(http.ListenAndServe(":8082", r)) // zachekinil, works ok, do handlers
}
