package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
	"github.com/shreyas1307/go-server/internal/handlers"
	"github.com/shreyas1307/go-server/internal/store"
)

type Server struct {
	mux *chi.Mux
	db  store.Store
}

func main() {
	ctx := context.Background()
	loglevel := "debug"

	err := setupLogger(ctx, loglevel)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
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

func setupLogger(ctx context.Context, loglevel string) error {
	l, err := zerolog.ParseLevel(loglevel)
	if err != nil {
		return fmt.Errorf("failed to set default log level: %w", err)
	}
	zerolog.SetGlobalLevel(l)
	return nil
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hello world \n")
	io.WriteString(w, "Hello world \n")
}

func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Health check \n")
	io.WriteString(w, "200 - OK \n")
}
