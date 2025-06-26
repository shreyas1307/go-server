package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	mux *chi.Mux
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Hello world")
		io.WriteString(w, "Hello world")
	})

	r.Post("/todo", TodoHandler)

	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Health check")
		io.WriteString(w, "200 - OK")
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}

func TodoHandler(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Error reading request body")
	}
	fmt.Printf("Received todo: %s\n", b)
	io.WriteString(w, "Received todo")
}
