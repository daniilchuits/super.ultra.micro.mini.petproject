package main

import (
	jwtmiddlewear "api/jwt_middlewear"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/go-chi/chi/v5"
)

func main() {

	secretKey := os.Getenv("JWT_SECRET")

	auth := "http://auth-service:8081"
	worker := "http://worker:8082"

	authLink, err := url.Parse(auth)
	if err != nil {
		log.Println("Auth link parsing error:", err)
		return
	}
	authProxy := httputil.NewSingleHostReverseProxy(authLink)

	workerLink, err := url.Parse(worker)
	if err != nil {
		log.Println("Worker link parsing error:", err)
		return
	}
	workerProxy := httputil.NewSingleHostReverseProxy(workerLink)
	newSecretKey := jwtmiddlewear.NewSecretKey([]byte(secretKey))
	workerProxyJWT := newSecretKey.JWTCheck(workerProxy)

	r := chi.NewRouter()

	r.Handle("/auth/*",
		http.StripPrefix("/auth", authProxy),
	)
	r.Handle("/work/*",
		http.StripPrefix("/work", workerProxyJWT),
	)

	log.Fatal(http.ListenAndServe(":8080", r))
}
