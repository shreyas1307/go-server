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
	srv := Server{
		mux: chi.NewRouter(),
	}
	srv.mux.Use(middleware.Logger)
	srv.mux.Use(middleware.Recoverer)
	srv.mux.Get("/hello-world", HelloWorldHandler)

	srv.mux.Post("/todo", TodoHandler)

	srv.mux.Get("/healthz", HealthzHandler)

	log.Fatal(http.ListenAndServe(":8080", srv.mux))
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

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hello world")
	io.WriteString(w, "Hello world")
}

func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Health check")
	io.WriteString(w, "200 - OK")
}
