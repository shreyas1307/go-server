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
	// Setting up the context and logger
	ctx := context.Background()
	loglevel := os.Getenv("LOG_LEVEL")

	if err := setupLogger(ctx, loglevel); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Creatinng a new server instance
	srv := Server{
		mux: chi.NewRouter(),
	}

	// Initialise the in-memory store
	var memorydb store.Store = store.NewInMemoryStore()

	// Initialising the TodoHandler with the in-memory store
	todohandler, err := handlers.NewTodoHandler(memorydb)
	if err != nil {
		fmt.Printf("Error creating TodoHandler: %v\n", err)
		os.Exit(1)
	}

	// Initialising the server middleware and routes
	srv.mux.Use(middleware.Logger)
	srv.mux.Use(middleware.Recoverer)
	srv.mux.Get("/hello-world", HelloWorldHandler)

	srv.mux.Mount("/todo", todohandler.Router())

	srv.mux.Get("/healthz", HealthzHandler)

	// Serving the application
	log.Fatal(http.ListenAndServe(":8080", srv.mux))
}

// setupLogger initialises the global logger with a specified log level.
//
// It returns an error if the log level is invalid.
func setupLogger(_ context.Context, loglevel string) error {
	l, err := zerolog.ParseLevel(loglevel)
	if err != nil {
		return fmt.Errorf("failed to set default log level: %w", err)
	}
	zerolog.SetGlobalLevel(l)
	return nil
}

// HelloWorldHandler is a simple http handler that reponds with "Hello world".
func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hello world \n")
	io.WriteString(w, "Hello world \n")
}

// HealthzHandler is a health check endpoint that responds with the server's current health status.
func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Health check \n")
	io.WriteString(w, "200 - OK \n")
}
