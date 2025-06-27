package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/shreyas1307/go-server/internal/handlers"
	"github.com/shreyas1307/go-server/internal/store"
)

type Server struct {
	mux *chi.Mux
	db  store.ServerStore
}

func main() {
	srv := Server{
		mux: chi.NewRouter(),
	}
	srv.mux.Use(middleware.Logger)
	srv.mux.Use(middleware.Recoverer)
	srv.mux.Get("/hello-world", HelloWorldHandler)

	srv.mux.Mount("/todo", (&handlers.TodoHandler{
		Store: srv.db,
	}).Router())

	srv.mux.Get("/healthz", HealthzHandler)

	log.Fatal(http.ListenAndServe(":8080", srv.mux))
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hello world \n")
	io.WriteString(w, "Hello world \n")
}

func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Health check \n")
	io.WriteString(w, "200 - OK \n")
}
