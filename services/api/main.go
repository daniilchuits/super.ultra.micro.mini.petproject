package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/go-chi/chi/v5"
)

func main() {

	auth := "http://auth-service:8081"
	authLink, err := url.Parse(auth)
	if err != nil {
		log.Println("Auth link parsing error:", err)
		return
	}
	authProxy := httputil.NewSingleHostReverseProxy(authLink)

	r := chi.NewRouter()

	r.Handle("/auth/*",
		http.StripPrefix("/auth", authProxy),
	)

	log.Fatal(http.ListenAndServe(":8080", r))
}
