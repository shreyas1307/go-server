package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Hello world")
		io.WriteString(w, "Hello world")
	})

	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Health check")
		io.WriteString(w, "200 - OK")
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}
